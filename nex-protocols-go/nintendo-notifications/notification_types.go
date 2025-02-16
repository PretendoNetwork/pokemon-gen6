// Package protocol implements the Nintendo Notfications protocol
package protocol

type notificationTypes struct {
	FriendPresenceUpdated3DS               uint32
	FriendFavoriteGameUpdated3DS           uint32
	FriendCommentUpdated3DS                uint32
	FriendMiiChanged3DS                    uint32
	FriendshipCompleted3DS                 uint32
	FriendOffline                          uint32
	FriendMiiChanged                       uint32
	Unknown1MiiRelated                     uint32
	FriendPreferencesChanged               uint32
	FriendStartedTitle                     uint32
	Unknown2FriendRequestRelated           uint32
	FriendRemoved                          uint32 // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestCanceled                  uint32 // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestReceived                  uint32
	Unknown3FriendRequestRelated           uint32
	Unknown4BlacklistRelated               uint32
	FriendRequestAccepted                  uint32
	Unknown5BlacklistRelated               uint32
	Unknown6BlacklistRelated               uint32
	FriendStatusMessageChanged             uint32
	Unknown7                               uint32
	Unknown8FriendshipRelated              uint32
	Unknown9PersistentNotificationsRelated uint32
}

// NotificationTypes is an enum of all the types a notification can be in the NintendoNotifications protocol
var NotificationTypes = notificationTypes{
	FriendPresenceUpdated3DS:               1,
	FriendFavoriteGameUpdated3DS:           2,
	FriendCommentUpdated3DS:                3,
	FriendMiiChanged3DS:                    5,
	FriendshipCompleted3DS:                 7,
	FriendOffline:                          10,
	FriendMiiChanged:                       21,
	Unknown1MiiRelated:                     22,
	FriendPreferencesChanged:               23,
	FriendStartedTitle:                     24,
	Unknown2FriendRequestRelated:           25,
	FriendRemoved:                          26, // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestCanceled:                  26, // * These 2 are the same event. Split them into 2 for cleaner naming
	FriendRequestReceived:                  27,
	Unknown3FriendRequestRelated:           28,
	Unknown4BlacklistRelated:               29,
	FriendRequestAccepted:                  30,
	Unknown5BlacklistRelated:               31,
	Unknown6BlacklistRelated:               32,
	FriendStatusMessageChanged:             33,
	Unknown7:                               34,
	Unknown8FriendshipRelated:              35,
	Unknown9PersistentNotificationsRelated: 36,
}
