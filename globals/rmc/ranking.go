package globals_rmc

func NewRanking() ProtocolInfo {
	tg := ProtocolInfo{}
	tg.protocolName = "Ranking (Modern)"
	tg.methodTable = map[uint32]string{
		1:  "UploadScore",
		2:  "DeleteScore",
		3:  "DeleteAllScores",
		4:  "UploadCommonData",
		5:  "DeleteCommonData",
		6:  "GetCommonData",
		7:  "ChangeAttributes",
		8:  "ChangeAllAttributes",
		9:  "GetRanking",
		10: "GetApproxOrder",
		11: "GetStats",
		12: "GetRankingByPIDList",
		13: "GetRankingByUniqueIdList",
		14: "GetCachedTopXRanking",
		15: "GetCachedTopXRankings",
	}

	return tg
}
