package globals_rmc

func NewTicketGranting() ProtocolInfo {
	tg := ProtocolInfo{}
	tg.protocolName = "Ticket Granting"
	tg.methodTable = map[uint32]string{
		1: "Login/ValidateAndRequestTicket",
		2: "LoginEx/ValidateAndRequestTicketWithCustomData",
		3: "RequestTicket",
		4: "GetPID",
		5: "GetName",
		6: "LoginWithContext/ValidateAndRequestTicketWithParam",
	}

	return tg
}
