package globals_rmc

func NewRankingLegacy() ProtocolInfo {
	tg := ProtocolInfo{}
	tg.protocolName = "Ranking (Legacy, 2.0)"
	tg.methodTable = map[uint32]string{
		1:  "UploadScore",
		2:  "UploadScores",
		3:  "DeleteScore",
		4:  "DeleteAllScore",
		5:  "UploadCommonData",
		6:  "DeleteCommonData",
		7:  "UnknownMethod0x7",
		8:  "UnknownMethod0x8",
		9:  "UnknownMethod0x9",
		10: "GetTopScore",
		11: "GetCommonData",
		12: "UnknownMethod0xC",
		13: "UnknownMethod0xD",
		14: "GetScore",
		15: "GetSelfScore",
		16: "GetTotal",
		17: "UploadScoreWithLimit",
		18: "UploadScoresWithLimit",
		19: "UnknownMethod0x13",
	}

	return tg
}
