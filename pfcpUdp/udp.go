package pfcpUdp

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/free5gc/pfcp"
	"github.com/free5gc/pfcp/logger"
)

const (
	PFCP_PORT        = 8805
	PFCP_MAX_UDP_LEN = 2048
)

var ErrReceivedResentRequest = errors.New("received a request that is re-sent")

type PfcpServer struct {
	Addr string
	Conn *net.UDPConn
	// Consumer Table
	// Map Consumer IP to its tx table
	ConsumerTable ConsumerTable
}

type ConsumerTable struct {
	m sync.Map // map[string]pfcp.TxTable
}

func (t *ConsumerTable) Load(consumerAddr string) (*pfcp.TxTable, bool) {
	txTable, ok := t.m.Load(consumerAddr)
	if ok {
		return txTable.(*pfcp.TxTable), ok
	}
	return nil, false
}

func (t *ConsumerTable) Store(consumerAddr string, txTable *pfcp.TxTable) {
	t.m.Store(consumerAddr, txTable)
}

func (t *ConsumerTable) Delete(consumerAddr string) {
	t.m.Delete(consumerAddr)
}

func NewPfcpServer(addr string) *PfcpServer {
	server := PfcpServer{Addr: addr}
	return &server
}

func (pfcpServer *PfcpServer) Listen() error {
	var serverIp net.IP
	if pfcpServer.Addr == "" {
		serverIp = net.IPv4zero
	} else {
		serverIp = net.ParseIP(pfcpServer.Addr)
	}

	addr := &net.UDPAddr{
		IP:   serverIp,
		Port: PFCP_PORT,
	}

	conn, err := net.ListenUDP("udp", addr)
	pfcpServer.Conn = conn
	return err
}

func (pfcpServer *PfcpServer) ReadFrom() (*Message, error) {

	buf := make([]byte, PFCP_MAX_UDP_LEN)
	n, addr, err := pfcpServer.Conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}

	pfcpMsg := &pfcp.Message{}
	msg := NewMessage(addr, pfcpMsg)

	err = pfcpMsg.Unmarshal(buf[:n])
	if err != nil {
		return msg, err
	}

	if pfcpMsg.IsRequest() {
		//Todo: Implement SendingResponse type of reliable delivery
		tx, err := pfcpServer.FindTransaction(pfcpMsg, addr)
		if err != nil {
			return msg, err
		}
		if tx != nil {
			// tx != nil => Already Replied => Resend Request
			tx.EventChannel <- pfcp.ReceiveEvent{
				Type:       pfcp.ReceiveEventTypeResendRequest,
				RemoteAddr: addr,
				RcvMsg:     pfcpMsg,
			}
			return msg, ErrReceivedResentRequest
		} else {
			// tx == nil => New Request
			return msg, nil
		}
	} else if pfcpMsg.IsResponse() {
		tx, err := pfcpServer.FindTransaction(pfcpMsg, pfcpServer.Conn.LocalAddr().(*net.UDPAddr))
		if err != nil {
			return msg, err
		}

		tx.EventChannel <- pfcp.ReceiveEvent{
			Type:       pfcp.ReceiveEventTypeValidResponse,
			RemoteAddr: addr,
			RcvMsg:     pfcpMsg,
		}
	}

	return msg, nil
}

func (pfcpServer *PfcpServer) WriteRequestTo(reqMsg *pfcp.Message, addr *net.UDPAddr) (resMsg *Message, err error) {
	if !reqMsg.IsRequest() {
		return nil, errors.New("not a request message")
	}

	buf, err := reqMsg.Marshal()
	if err != nil {
		return nil, err
	}

	tx := pfcp.NewTransaction(reqMsg, buf, pfcpServer.Conn, addr)

	err = pfcpServer.PutTransaction(tx)
	if err != nil {
		return nil, err
	}

	return pfcpServer.StartReqTxLifeCycle(tx)
}

func (pfcpServer *PfcpServer) WriteResponseTo(resMsg *pfcp.Message, addr *net.UDPAddr) {
	if !resMsg.IsResponse() {
		logger.PFCPLog.Warn("not a response message")
		return
	}

	buf, err := resMsg.Marshal()
	if err != nil {
		logger.PFCPLog.Warnf("marshal error: %+v", err)
		return
	}

	tx := pfcp.NewTransaction(resMsg, buf, pfcpServer.Conn, addr)

	err = pfcpServer.PutTransaction(tx)
	if err != nil {
		logger.PFCPLog.Warnf("PutTransaction error: %+v", err)
		return
	}

	go pfcpServer.StartResTxLifeCycle(tx)
}

func (pfcpServer *PfcpServer) Close() error {
	return pfcpServer.Conn.Close()
}

func (pfcpServer *PfcpServer) PutTransaction(tx *pfcp.Transaction) (err error) {
	logger.PFCPLog.Traceln("In PutTransaction")

	consumerAddr := tx.ConsumerAddr
	if _, exist := pfcpServer.ConsumerTable.Load(consumerAddr); !exist {
		pfcpServer.ConsumerTable.Store(consumerAddr, &pfcp.TxTable{})
	}

	txTable, _ := pfcpServer.ConsumerTable.Load(consumerAddr)
	if _, exist := txTable.Load(tx.SequenceNumber); !exist {
		txTable.Store(tx.SequenceNumber, tx)
	} else {
		logger.PFCPLog.Warnln("In PutTransaction")
		logger.PFCPLog.Warnln("Consumer Addr: ", consumerAddr)
		logger.PFCPLog.Warnln("Sequence number ", tx.SequenceNumber, " already exist!")
		err = fmt.Errorf("Insert tx error: duplicate sequence number %d", tx.SequenceNumber)
	}

	logger.PFCPLog.Traceln("End PutTransaction")
	return
}

func (pfcpServer *PfcpServer) RemoveTransaction(tx *pfcp.Transaction) (err error) {
	logger.PFCPLog.Traceln("In RemoveTransaction")
	consumerAddr := tx.ConsumerAddr
	txTable, _ := pfcpServer.ConsumerTable.Load(consumerAddr)

	if txTmp, exist := txTable.Load(tx.SequenceNumber); exist {
		tx = txTmp
		if tx.TxType == pfcp.SendingRequest {
			logger.PFCPLog.Infof("Remove Request Transaction [%d]\n", tx.SequenceNumber)
		} else if tx.TxType == pfcp.SendingResponse {
			logger.PFCPLog.Infof("Remove Request Transaction [%d]\n", tx.SequenceNumber)
		}

		txTable.Delete(tx.SequenceNumber)
	} else {
		logger.PFCPLog.Warnln("In RemoveTransaction")
		logger.PFCPLog.Warnln("Consumer IP: ", consumerAddr)
		logger.PFCPLog.Warnln("Sequence number ", tx.SequenceNumber, " doesn't exist!")
		err = fmt.Errorf("Remove tx error: transaction [%d] doesn't exist", tx.SequenceNumber)
	}

	logger.PFCPLog.Traceln("End RemoveTransaction")
	return
}

func (pfcpServer *PfcpServer) StartReqTxLifeCycle(tx *pfcp.Transaction) (resMsg *Message, err error) {
	defer func() {
		//End Transaction
		err := pfcpServer.RemoveTransaction(tx)
		if err != nil {
			logger.PFCPLog.Warnf("RemoveTransaction error: %+v", err)
		}
	}()

	//Start Transaction
	event, err := tx.StartSendingRequest()
	if err != nil {
		return nil, err
	}
	return NewMessage(event.RemoteAddr, event.RcvMsg), nil
}

// StartResTxLifeCycle does not return an error because if an error occurs, a resend request will be sent
func (pfcpServer *PfcpServer) StartResTxLifeCycle(tx *pfcp.Transaction) {
	//Start Transaction
	err := tx.StartSendingResponse()
	if err != nil {
		logger.PFCPLog.Warnf("SendingResponse error: %+v", err)
		return
	}
	//End Transaction
	err = pfcpServer.RemoveTransaction(tx)
	if err != nil {
		logger.PFCPLog.Warnf("RemoveTransaction error: %+v", err)
	}
}

func (pfcpServer *PfcpServer) FindTransaction(msg *pfcp.Message, addr *net.UDPAddr) (*pfcp.Transaction, error) {
	var tx *pfcp.Transaction

	logger.PFCPLog.Traceln("In FindTransaction")
	consumerAddr := addr.String()

	if msg.IsResponse() {
		if _, exist := pfcpServer.ConsumerTable.Load(consumerAddr); !exist {
			logger.PFCPLog.Warnln("In FindTransaction")
			logger.PFCPLog.Warnf("Can't find txTable from consumer addr: [%s]", consumerAddr)
			return nil, fmt.Errorf("FindTransaction Error: txTable not found")
		}

		txTable, _ := pfcpServer.ConsumerTable.Load(consumerAddr)
		seqNum := msg.Header.SequenceNumber

		if _, exist := txTable.Load(seqNum); !exist {
			logger.PFCPLog.Warnln("In FindTransaction")
			logger.PFCPLog.Warnln("Consumer Addr: ", consumerAddr)
			logger.PFCPLog.Warnf("Can't find tx [%d] from txTable: ", seqNum)
			return nil, fmt.Errorf("FindTransaction Error: sequence number [%d] not found", seqNum)
		}

		tx, _ = txTable.Load(seqNum)
	} else if msg.IsRequest() {
		if _, exist := pfcpServer.ConsumerTable.Load(consumerAddr); !exist {
			return nil, nil
		}

		txTable, _ := pfcpServer.ConsumerTable.Load(consumerAddr)
		seqNum := msg.Header.SequenceNumber

		if _, exist := txTable.Load(seqNum); !exist {
			return nil, nil
		}

		tx, _ = txTable.Load(seqNum)
	}
	logger.PFCPLog.Traceln("End FindTransaction")
	return tx, nil
}

// Send a PFCP message and close UDP connection
func SendPfcpMessage(msg pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {
	var conn net.Conn
	if connTmp, err := net.DialUDP("udp", srcAddr, dstAddr); err != nil {
		return err
	} else {
		conn = connTmp
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.PFCPLog.Warnf("Connection close error: %v", err)
		}
	}()

	buf, err := msg.Marshal()
	if err != nil {
		return err
	}

	/*TODO: check if all bytes of buf are sent*/
	_, err = conn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

// Receive a PFCP message and close UDP connection
func ReceivePfcpMessage(msg *pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {
	var conn net.Conn
	if connTmp, err := net.DialUDP("udp", srcAddr, dstAddr); err != nil {
		return err
	} else {
		conn = connTmp
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.PFCPLog.Warnf("Connection close error: %v", err)
		}
	}()

	buf := make([]byte, PFCP_MAX_UDP_LEN)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

	err = msg.Unmarshal(buf[:n])
	if err != nil {
		return err
	}

	return nil
}
