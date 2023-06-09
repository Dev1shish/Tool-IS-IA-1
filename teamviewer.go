package packets

import (
	"encoding/binary"
	"fmt"
)

const TeamViewerPort = 5938

var teamViewerCommands = map[uint8]string{
	10:  "CMD_IDENTIFY",
	11:  "CMD_REQUESTCONNECT",
	13:  "CMD_DISCONNECT",
	14:  "CMD_VNCDISCONNECT",
	15:  "CMD_TVCONNECTIONFAILED",
	16:  "CMD_PING",
	17:  "CMD_PINGOK",
	18:  "CMD_MASTERCOMMAND",
	19:  "CMD_MASTERRESPONSE",
	20:  "CMD_CHANGECONNECTION",
	21:  "CMD_NOPARTNERCONNECT",
	22:  "CMD_CONNECTTOWAITINGTHREAD",
	23:  "CMD_SESSIONMODE",
	24:  "CMD_REQUESTROUTINGSESSION",
	25:  "CMD_TIMEOUT",
	26:  "CMD_JAVACONNECT",
	27:  "CMD_KEEPALIVEBEEP",
	28:  "CMD_REQUESTKEEPALIVE",
	29:  "CMD_MASTERCOMMAND_ENCRYPTED",
	30:  "CMD_MASTERRESPONSE_ENCRYPTED",
	31:  "CMD_REQUESTRECONNECT",
	32:  "CMD_RECONNECTTOWAITINGTHREAD",
	33:  "CMD_STARTLOGGING",
	34:  "CMD_SERVERAVAILABLE",
	35:  "CMD_KEEPALIVEREQUEST",
	36:  "CMD_OK",
	37:  "CMD_FAILED",
	38:  "CMD_PING_PERFORMANCE",
	39:  "CMD_PING_PERFORMANCE_RESPONSE",
	40:  "CMD_REQUESTKEEPALIVE2",
	41:  "CMD_DISCONNECT_SWITCHEDTOUDP",
	42:  "CMD_SENDMODE_UDP",
	43:  "CMD_KEEPALIVEREQUEST_ANSWER",
	44:  "CMD_ROUTE_CMD_TO_CLIENT",
	45:  "CMD_NEW_MASTERLOGIN",
	46:  "CMD_BUDDY",
	47:  "CMD_ACCEPTROUTINGSESSION",
	48:  "CMD_NEW_MASTERLOGIN_ANSWER",
	49:  "CMD_BUDDY_ENCRYPTED",
	50:  "CMD_REQUEST_ROUTE_BUDDY",
	51:  "CMD_CONTACT_OTHER_MASTER",
	52:  "CMD_REQUEST_ROUTE_ENCRYPTED",
	53:  "CMD_ENDSESSION",
	54:  "CMD_SESSIONID",
	55:  "CMD_RECONNECT_TO_SESSION",
	56:  "CMD_RECONNECT_TO_SESSION_ANSWER",
	57:  "CMD_MEETING_CONTROL",
	58:  "CMD_CARRIER_SWITCH",
	59:  "CMD_MEETING_AUTHENTICATION",
	60:  "CMD_ROUTERCMD",
	61:  "CMD_PARTNERRECONNECT",
	62:  "CMD_CONGRESTION_CONTROL",
	63:  "CMD_ACK",
	70:  "CMD_UDPREQUESTCONNECT",
	71:  "CMD_UDPPING",
	72:  "CMD_UDPREQUESTCONNECT_VPN",
	90:  "CMD_DATA",
	91:  "CMD_DATA2",
	92:  "CMD_DATA_ENCRYPTED",
	93:  "CMD_REQUESTENCRYPTION",
	94:  "CMD_CONFIRMENCRYPTION",
	95:  "CMD_ENCRYPTIONREQUESTFAILED",
	96:  "CMD_REQUESTNOENCRYPTION",
	97:  "CMD_UDPFLOWCONTROL",
	98:  "CMD_DATA3",
	99:  "CMD_DATA3_ENCRYPTED",
	100: "CMD_DATA3_RESENDPACKETS",
	101: "CMD_DATA3_ACKPACKETS",
	102: "CMD_AUTH_CHALLENGE",
	103: "CMD_AUTH_RESPONSE",
	104: "CMD_AUTH_RESULT",
	105: "CMD_RIP_MESSAGES",
	106: "CMD_DATA4",
	107: "CMD_DATASTREAM",
	108: "CMD_UDPHEARTBEAT",
	109: "CMD_DATA_DIRECTED",
	110: "CMD_UDP_RESENDPACKETS",
	111: "CMD_UDP_ACKPACKETS",
	112: "CMD_UDP_PROTECTEDCOMMAND",
	113: "CMD_FLUSHSENDBUFFER",
}

type TeamViewerPacket struct {
	Magic       uint16
	Version     string
	CommandCode uint8
	Command     string
}

func ParseTeamViewer(data []byte) *TeamViewerPacket {
	if len(data) < 3 {
		return nil
	}

	magic := binary.BigEndian.Uint16(data[0:2])
	cmdCode := uint8(data[3])
	version := "1.x"
	cmd := "CMD_?"

	if magic == 0x1724 {
		version = "1.x"
	} else if magic == 0x1130 {
		version = "2.x"
	} else {
		version = fmt.Sprintf("?.? (magic=0x%x)", magic)
	}

	if val, found := teamViewerCommands[cmdCode]; !found {
		return nil
	} else {
		cmd = val
	}

	return &TeamViewerPacket{
		Magic:       magic,
		Version:     version,
		CommandCode: cmdCode,
		Command:     cmd,
	}
}
