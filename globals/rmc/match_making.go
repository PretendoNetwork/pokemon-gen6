package globals_rmc

func NewMatchMaking() ProtocolInfo {
	mm := ProtocolInfo{}
	mm.protocolName = "Match Making"
	mm.methodTable = map[uint32]string{
		1:  "RegisterGathering",
		2:  "UnregisterGathering",
		3:  "UnregisterGatherings",
		4:  "UpdateGathering",
		5:  "Invite",
		6:  "AcceptInvitation",
		7:  "DeclineInvitation",
		8:  "CancelInvitation",
		9:  "GetInvitationsSent",
		10: "GetInvitationsReceived",
		11: "Participate",
		12: "CancelParticipation",
		13: "GetParticipants",
		14: "AddParticipants",
		15: "GetDetailedParticipants",
		16: "GetParticipantsURLs",
		17: "FindByType",
		18: "FindByDescription",
		19: "FindByDescriptionRegex",
		20: "FindByID",
		21: "FindBySingleID",
		22: "FindByOwner",
		23: "FindByParticipants",
		24: "FindInvitations",
		25: "FindBySQLQuery",
		26: "LaunchSession",
		27: "UpdateSessionURL",
		28: "GetSessionURL",
		29: "GetState",
		30: "SetState",
		31: "ReportStats",
		32: "GetStats",
		33: "DeleteGathering",
		34: "GetPendingDeletions",
		35: "DeleteFromDeletions",
		36: "MigrateGatheringOwnershipV1",
		37: "FindByDescriptionLike",
		38: "RegisterLocalURL",
		39: "RegisterLocalURLs",
		40: "UpdateSessionHostV1",
		41: "GetSessionURLs",
		42: "UpdateSessionHost",
		43: "UpdateGatheringOwnership",
		44: "MigrateGatheringOwnership",
	}

	return mm
}
