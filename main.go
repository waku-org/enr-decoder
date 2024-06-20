package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"net/netip"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	wenr "github.com/waku-org/go-waku/waku/v2/protocol/enr"
)

func main() {
	var enr = flag.String("enr", "", "enr to be decoded")

	flag.Parse()

	node, err := enode.Parse(enode.ValidSchemes, *enr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded ENR:")
	peerID, multiaddrs, err := wenr.Multiaddress(node)
	if err != nil {
		panic(err)
	}

	fmt.Println("===============================================================================")
	fmt.Println("seq:", node.Record().Seq())
	fmt.Println("signature:", "0x"+hex.EncodeToString(node.Record().Signature()))
	fmt.Println("PeerID: ", peerID)
	ip, err := GetValue(node.Record(), "ip")
	if err != nil {
		panic(err)
	} else {
		if len(ip) > 0 {
			ipaddr, ok := netip.AddrFromSlice(ip)
			if ok {
				fmt.Println("ipv4:", ipaddr)
			}
		} else {
			fmt.Println("ipv4: field has no value")
		}
	}
	port, err := GetValue(node.Record(), "tcp")
	if err != nil {
		panic(err)
	} else {
		if len(port) > 0 {
			fmt.Println("tcp:", binary.BigEndian.Uint16(port))
		} else {
			fmt.Println("ipv4: field has no value")
		}
	}
	ReadAndPrintValue(node.Record(), "waku2")
	ReadAndPrintValue(node.Record(), "rs")
	ReadAndPrintValue(node.Record(), "rsv")

	fmt.Println("Multiaddresses:")
	for _, maddr := range multiaddrs {
		fmt.Println(maddr)
	}
	fmt.Println("===============================================================================")

}

func GetValue(record *enr.Record, name string) ([]byte, error) {
	var field []byte
	if err := record.Load(enr.WithEntry(name, &field)); err != nil {
		if enr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return field, nil
}

func ReadAndPrintValue(record *enr.Record, name string) {
	var field []byte
	if err := record.Load(enr.WithEntry(name, &field)); err != nil {
		if enr.IsNotFound(err) {
			fmt.Println(name, ":", "field contains no value")
			return
		}
		panic(err)
	}
	fmt.Println(name, ":", "0x"+hex.EncodeToString(field))
}
