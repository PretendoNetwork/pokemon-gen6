package globals_rmc

func NewMessageDelivery() ProtocolInfo {
	md := ProtocolInfo{}
	md.protocolName = "Message Delivery"
	md.methodTable = map[uint32]string{
		1: "DeliverMessage",
		2: "DeliverMessageMultiTarget",
	}

	return md
}
