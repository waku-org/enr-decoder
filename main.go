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
	"github.com/waku-org/go-waku/waku/v2/protocol/filter"
	"github.com/waku-org/go-waku/waku/v2/protocol/lightpush"
	"github.com/waku-org/go-waku/waku/v2/protocol/relay"
	"github.com/waku-org/go-waku/waku/v2/protocol/store"
)

func main() {
	var enrStr = flag.String("enr", "", "enr to be decoded")

	flag.Parse()

	node, err := enode.Parse(enode.ValidSchemes, *enrStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded ENR:")
	peerID, multiaddrs, err := wenr.Multiaddress(node)
	if err != nil {
		panic(err)
	}

	fmt.Println("seq:", node.Record().Seq())
	fmt.Println("signature:", "0x"+hex.EncodeToString(node.Record().Signature()))
	fmt.Println("peer-id: ", peerID)
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
			fmt.Println("tcp-port:", binary.BigEndian.Uint16(port))
		} else {
			fmt.Println("tcp-port:: field has no value")
		}
	}
	uport, err := GetValue(node.Record(), "udp")
	if err != nil {
		panic(err)
	} else {
		if len(uport) > 0 {
			fmt.Println("udp-port:", binary.BigEndian.Uint16(uport))
		} else {
			fmt.Println("udp-port: field has no value")
		}
	}

	shards, err := wenr.RelaySharding(node.Record())
	if err != nil {
		panic(err)
	}
	if shards != nil {
		fmt.Println("cluster-id: ", shards.ClusterID)
		fmt.Println("shards: ", shards.ShardIDs)
	} else {
		fmt.Println("cluster-id:", "not available")
		fmt.Println("shards:", "not available")
	}

	DecodeWaku2ENRField(node.Record())

	fmt.Println("multiaddresses:")
	for _, maddr := range multiaddrs {
		fmt.Println(maddr)
	}

}

func DecodeWaku2ENRField(record *enr.Record) {
	//Decoding Waku2 field
	var enrField wenr.WakuEnrBitfield
	var protosSupported []string
	if err := record.Load(enr.WithEntry("waku2", &enrField)); err != nil {
		if enr.IsNotFound(err) {
			fmt.Println("waku2:", "field contains no value")
		} else {
			panic(err)
		}
	}

	if enrField&relay.WakuRelayENRField != 0 {
		protosSupported = append(protosSupported, string(relay.WakuRelayID_v200))
	}
	if enrField&filter.FilterSubscribeENRField != 0 {
		protosSupported = append(protosSupported, string(filter.FilterSubscribeID_v20beta1))
	}
	if enrField&lightpush.LightPushENRField != 0 {
		protosSupported = append(protosSupported, string(lightpush.LightPushID_v20beta1))
	}
	if enrField&store.StoreENRField != 0 {
		protosSupported = append(protosSupported, string(store.StoreID_v20beta4))
	}
	fmt.Println("Wakuv2 Protocols Supported:")
	for _, proto := range protosSupported {
		fmt.Println(proto)
	}
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
