package globals_rmc

func NewRankingLegacyOld() ProtocolInfo {
	tg := ProtocolInfo{}
	tg.protocolName = "Ranking (Legacy, Pre-2.0)"
	tg.methodTable = map[uint32]string{
		1:  "UploadScore",
		2:  "DeleteScore",
		3:  "DeleteAllScore",
		4:  "UploadCommonData",
		5:  "DeleteCommonData",
		6:  "UnknownMethod0x7",
		7:  "UnknownMethod0x8",
		8:  "GetTopScore",
		9:  "GetCommonData",
		10: "UnknownMethod0xC",
		11: "UnknownMethod0xD",
		12: "GetScore",
		13: "GetSelfScore",
		14: "GetTotal",
	}

	return tg
}
