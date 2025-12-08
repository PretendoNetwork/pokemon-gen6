package globals_rmc

type ProtocolInfo struct {
	methodTable  map[uint32]string
	protocolName string
}

func NewProtocolInfo() ProtocolInfo {
	pi := ProtocolInfo{}
	pi.protocolName = "Unknown"
	pi.methodTable = map[uint32]string{}

	return pi
}

func (protocol ProtocolInfo) Protocol() string {
	return protocol.protocolName
}

func (protocol ProtocolInfo) GetMethodByID(methodId uint32) string {
	method, exists := protocol.methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
