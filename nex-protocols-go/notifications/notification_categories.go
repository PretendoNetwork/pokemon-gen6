// Package protocol implements the Notifications protocol
package protocol

type notificationCategories struct {
	Participation                uint32
	OwnershipChanged             uint32
	RequestJoinGathering         uint32 // * This is what these mean in WiiU Chat, unclear if this is the real use
	EndGathering                 uint32 // * This is what these mean in WiiU Chat, unclear if this is the real use
	GatheringUnregistered        uint32
	HostChanged                  uint32
	ServiceItemRequestCompleted  uint32
	MatchmakeRefereeRoundStarted uint32
	SystemPasswordChanged        uint32
	SystemPasswordCleared        uint32
	SwitchGathering              uint32
}

// NotificationCategories is a list of all the categories a notification can be in the Notifications protocol
//
// Not all of these are categories, some are stand-alone types
// This is a design choice made by NEX, not us
var NotificationCategories = notificationCategories{
	Participation:                3,
	OwnershipChanged:             4,
	RequestJoinGathering:         101, // * This is what these mean in WiiU Chat, unclear if this is the real use
	EndGathering:                 102, // * This is what these mean in WiiU Chat, unclear if this is the real use
	GatheringUnregistered:        109,
	HostChanged:                  110,
	ServiceItemRequestCompleted:  115,
	MatchmakeRefereeRoundStarted: 116,
	SystemPasswordChanged:        120,
	SystemPasswordCleared:        121,
	SwitchGathering:              122,
}
