package pfcpType

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const (
	Mask8 = 1<<8 - 1
	Mask7 = 1<<7 - 1
	Mask6 = 1<<6 - 1
	Mask5 = 1<<5 - 1
	Mask4 = 1<<4 - 1
	Mask3 = 1<<3 - 1
	Mask2 = 1<<2 - 1
	Mask1 = 1<<1 - 1
)

const (
	BitMask8 = 1 << 7
	BitMask7 = 1 << 6
	BitMask6 = 1 << 5
	BitMask5 = 1 << 4
	BitMask4 = 1 << 3
	BitMask3 = 1 << 2
	BitMask2 = 1 << 1
	BitMask1 = 1
)

var BASE_DATE_NTP_ERA0 time.Time = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)

func btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func utob(u uint8) bool {
	return u != 0
}

func fqdnToRfc1035(fqdn string, isDnn bool) ([]byte, error) {
	var rfc1035RR []byte
	domainSegments := strings.Split(fqdn, ".")

	typeName := "fqdn"
	maxLabelLen := 63
	maxTotalLen := 255
	if isDnn {
		typeName = "DNN"
		// In RFC 1035 max length of label is 63, but in TS 23.003 including length octet
		maxLabelLen = 62
		// In RFC 1035 max length of FQDN is 255, but DNN in TS 23.003 is 100
		maxTotalLen = 100
	}

	for _, segment := range domainSegments {
		if len(segment) > maxLabelLen {
			return nil, fmt.Errorf("%s limit the label to %d octets or less", typeName, maxLabelLen)
		}
		rfc1035RR = append(rfc1035RR, uint8(len(segment)))
		rfc1035RR = append(rfc1035RR, segment...)
	}

	if len(rfc1035RR) > maxTotalLen {
		return nil, fmt.Errorf("%s should less then %d octet", typeName, maxTotalLen)
	}
	return rfc1035RR, nil
}

func rfc1035tofqdn(rfc1035RR []byte) string {
	rfc1035Reader := bytes.NewBuffer(rfc1035RR)
	fqdn := ""

	for {
		// length of label
		if labelLen, err := rfc1035Reader.ReadByte(); err != nil {
			break
		} else {
			fqdn += string(rfc1035Reader.Next(int(labelLen))) + "."
		}
	}

	return fqdn[0 : len(fqdn)-1]
}
