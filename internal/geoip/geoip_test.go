package geoip

import (
	"encoding/binary"
	"net"
	"testing"
)

func ipNum(s string) uint32 {
	return binary.BigEndian.Uint32(net.ParseIP(s).To4())
}

func TestGeoLookup(t *testing.T) {
	db := NewDB([]Range{{Start: ipNum("1.2.3.0"), End: ipNum("1.2.3.255"), CC: "US"}})
	cc, err := db.Lookup("1.2.3.10")
	if err != nil || cc != "US" {
		t.Fatal(err, cc)
	}
	cc, err = db.Lookup("8.8.8.8")
	if err != nil || cc != "ZZ" {
		t.Fatal(err, cc)
	}
}
