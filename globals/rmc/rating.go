package globals_rmc

func NewRating() ProtocolInfo {
	rt := ProtocolInfo{}
	rt.protocolName = "Rating"
	rt.methodTable = map[uint32]string{
		1: "Unk0x1",
		2: "Unk0x2",
		3: "ReportRatingStats",
		4: "GetRanking",
		5: "DeleteScore",
		6: "Unk0x6",
		7: "UploadCommonData",
		8: "GetCommonData",
		9: "Unk0x9",
	}

	return rt
}
