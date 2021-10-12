package pfcpType

import (
	"bytes"
	"errors"
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

func fqdnToRfc1035(fqdn string) ([]byte, error) {
	var rfc1035RR []byte
	domainSegments := strings.Split(fqdn, ".")

	for _, segment := range domainSegments {
		if len(segment) > 63 {
			return nil, errors.New("fqdn limit the label to 63 octets or less")
		}
		rfc1035RR = append(rfc1035RR, uint8(len(segment)))
		rfc1035RR = append(rfc1035RR, segment...)
	}

	if len(rfc1035RR) > 255 {
		return nil, errors.New("fqdn should less then 255 octet")
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
