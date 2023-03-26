package packets

import (
	"strconv"

	"github.com/evilsocket/islazy/str"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	NBNSPort        = 137
	NBNSMinRespSize = 73
)

var (
	// NBNS hostname resolution request buffer.
	NBNSRequest = []byte{
		0x82, 0x28, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x20, 0x43, 0x4B, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41,
		0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x0,
		0x0, 0x21, 0x0, 0x1,
	}
)

func NBNSGetMeta(pkt gopacket.Packet) map[string]string {
	if ludp := pkt.Layer(layers.LayerTypeUDP); ludp != nil {
		if udp := ludp.(*layers.UDP); udp != nil && udp.SrcPort == NBNSPort && len(udp.Payload) >= NBNSMinRespSize {
			hostname := str.Trim(string(udp.Payload[57:72]))
			if strconv.IsPrint(rune(hostname[0])) {
				return map[string]string{
					"nbns:hostname": hostname,
				}
			}
		}
	}
	return nil
}
