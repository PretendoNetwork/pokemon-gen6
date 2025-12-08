package globals_rmc

func NewSubscription() ProtocolInfo {
	ss := ProtocolInfo{}
	ss.protocolName = "Subscription"
	ss.methodTable = map[uint32]string{
		1:  "CreateMySubscriptionData",
		2:  "UpdateMySubscriptionData",
		3:  "ClearMySubscriptionData",
		4:  "AddTargets",
		5:  "DeleteTargets",
		6:  "ClearTargets",
		7:  "GetFriendSubscriptionData",
		8:  "GetTargetSubscriptionData",
		9:  "GetActivePlayerSubscriptionData",
		10: "GetSubscriptionData",
		11: "ReplaceTargetAndGetSubscriptionData",
		12: "SetPrivacyLevel",
		13: "GetPrivacyLevel",
	}

	return ss
}
