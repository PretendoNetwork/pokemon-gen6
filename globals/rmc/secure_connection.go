package globals_rmc

func NewSecureConnection() ProtocolInfo {
	sc := ProtocolInfo{}
	sc.protocolName = "Secure Connection"
	sc.methodTable = map[uint32]string{
		1: "Register",
		2: "RequestConnectionData",
		3: "RequestUrls",
		4: "RegisterEx",
		5: "TestConnectivity",
		6: "UpdateURLs",
		7: "ReplaceURL",
		8: "SendRequest",
	}

	return sc
}
