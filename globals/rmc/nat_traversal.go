package globals_rmc

func NewNATTraversal() ProtocolInfo {
	nt := ProtocolInfo{}
	nt.protocolName = "NAT Traversal"
	nt.methodTable = map[uint32]string{
		1: "RequestProbeInitiation",
		2: "InitiateProbe",
		3: "RequestProbeInitiationExt",
		4: "ReportNATTraversalResult",
		5: "ReportNATProperties",
		6: "GetRelaySignatureKey",
		7: "ReportNATTraversalDetail",
	}

	return nt
}
