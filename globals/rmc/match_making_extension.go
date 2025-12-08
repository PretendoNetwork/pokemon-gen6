package globals_rmc

func NewMatchMakingExtension() ProtocolInfo {
	mme := ProtocolInfo{}
	mme.protocolName = "Match Making Extension"
	mme.methodTable = map[uint32]string{
		1: "EndParticipation",
		2: "GetParticipants",
		3: "GetDetailedParticipants",
		4: "GetParticipantsURLs",
		5: "GetGatheringRelations",
		6: "DeleteFromDeletions",
	}

	return mme
}
