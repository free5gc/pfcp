package pfcp

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/free5gc/pfcp/logger"
)

type TransactionType uint8

type TxTable struct {
	m sync.Map // map[uint32]*Transaction
}

func (t *TxTable) Store(sequenceNumber uint32, tx *Transaction) {
	t.m.Store(sequenceNumber, tx)
}

func (t *TxTable) Load(sequenceNumber uint32) (*Transaction, bool) {
	tx, ok := t.m.Load(sequenceNumber)
	if ok {
		return tx.(*Transaction), ok
	}
	return nil, false
}

func (t *TxTable) LoadOrStore(sequenceNumber uint32, storeTx *Transaction) (*Transaction, bool) {
	tx, loaded := t.m.LoadOrStore(sequenceNumber, storeTx)
	return tx.(*Transaction), loaded
}

func (t *TxTable) Delete(sequenceNumber uint32) {
	t.m.Delete(sequenceNumber)
}

const (
	SendingRequest TransactionType = iota
	SendingResponse
)

const (
	NumOfResend                 = 3
	ResendRequestTimeOutPeriod  = 3
	ResendResponseTimeOutPeriod = 15
)

// Transaction - represent the transaction state of pfcp message
type Transaction struct {
	SendMsg        []byte
	SequenceNumber uint32
	MessageType    MessageType
	TxType         TransactionType
	EventChannel   chan ReceiveEvent
	Conn           *net.UDPConn
	DestAddr       *net.UDPAddr
	ConsumerAddr   string
}

// NewTransaction - create pfcp transaction object
func NewTransaction(pfcpMSG *Message, binaryMSG []byte, Conn *net.UDPConn, DestAddr *net.UDPAddr) (tx *Transaction) {
	tx = &Transaction{
		SendMsg:        binaryMSG,
		SequenceNumber: pfcpMSG.Header.SequenceNumber,
		MessageType:    pfcpMSG.Header.MessageType,
		EventChannel:   make(chan ReceiveEvent),
		Conn:           Conn,
		DestAddr:       DestAddr,
	}

	if pfcpMSG.IsRequest() {
		tx.TxType = SendingRequest
		tx.ConsumerAddr = Conn.LocalAddr().String()
	} else if pfcpMSG.IsResponse() {
		tx.TxType = SendingResponse
		tx.ConsumerAddr = DestAddr.String()
	}

	logger.PFCPLog.Tracef("New Transaction SEQ[%d] DestAddr[%s]", tx.SequenceNumber, DestAddr.String())
	return
}

func (tx *Transaction) StartSendingRequest() (*ReceiveEvent, error) {
	if tx.TxType != SendingRequest {
		return nil, errors.New("this transaction is not for sending request")
	}

	logger.PFCPLog.Tracef("Start Request Transaction [%d]", tx.SequenceNumber)

	for iter := 0; iter < NumOfResend; iter++ {
		_, err := tx.Conn.WriteToUDP(tx.SendMsg, tx.DestAddr)
		if err != nil {
			return nil, fmt.Errorf("Request Transaction [%d]: %s", tx.SequenceNumber, err)
		}
		logger.PFCPLog.Tracef("Request Transaction [%d]: Sent a PFCP request packet", tx.SequenceNumber)

		select {
		case event := <-tx.EventChannel:
			if event.Type == ReceiveEventTypeValidResponse {
				logger.PFCPLog.Tracef("Request Transaction [%d]: receive valid response", tx.SequenceNumber)
				return &event, nil
			}
		case <-time.After(ResendRequestTimeOutPeriod * time.Second):
			logger.PFCPLog.Tracef("Request Transaction [%d]: timeout expire", tx.SequenceNumber)
			continue
		}
	}
	return nil, fmt.Errorf("Request Transaction [%d]: retry-out", tx.SequenceNumber)
}

func (tx *Transaction) StartSendingResponse() error {
	if tx.TxType != SendingResponse {
		return errors.New("this transaction is not for sending response")
	}

	logger.PFCPLog.Tracef("Start Response Transaction [%d]", tx.SequenceNumber)

	for {
		_, err := tx.Conn.WriteToUDP(tx.SendMsg, tx.DestAddr)
		if err != nil {
			return fmt.Errorf("Response Transaction [%d]: sending error", tx.SequenceNumber)
		}

		select {
		case event := <-tx.EventChannel:
			if event.Type == ReceiveEventTypeResendRequest {
				logger.PFCPLog.Tracef("Response Transaction [%d]: receive resend request", tx.SequenceNumber)
				logger.PFCPLog.Tracef("Response Transaction [%d]: Resend packet", tx.SequenceNumber)
				continue
			} else {
				logger.PFCPLog.Warnf("Response Transaction [%d]: receive invalid request", tx.SequenceNumber)
			}
		case <-time.After(ResendResponseTimeOutPeriod * time.Second):
			logger.PFCPLog.Tracef("Response Transaction [%d]: timeout expire", tx.SequenceNumber)
			return nil
		}
	}
}
