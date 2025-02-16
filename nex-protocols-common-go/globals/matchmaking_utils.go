package common_globals

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	notifications "github.com/PretendoNetwork/nex-protocols-go/v2/notifications"
	notifications_types "github.com/PretendoNetwork/nex-protocols-go/v2/notifications/types"
	"golang.org/x/exp/slices"
)

// GetAvailableGatheringID returns a gathering ID which doesn't belong to any session
// Returns 0 if no IDs are available (math.MaxUint32 has been reached)
func GetAvailableGatheringID() uint32 {
	return CurrentGatheringID.Next()
}

// FindOtherConnectionID searches a connection ID on the session that isn't the given one
// Returns 0 if no connection ID could be found
func FindOtherConnectionID(excludedConnectionID uint32, gatheringID uint32) uint32 {
	var otherConnectionID uint32 = 0
	if session, ok := Sessions[gatheringID]; ok {
		session.ConnectionIDs.Each(func(_ int, connectionID uint32) bool {
			if connectionID != excludedConnectionID {
				otherConnectionID = connectionID
				return true
			}

			return false
		})
	}

	return otherConnectionID
}

// RemoveConnectionIDFromSession removes a PRUDP connection from the session
func RemoveConnectionIDFromSession(id uint32, gathering uint32) {
	Sessions[gathering].ConnectionIDs.DeleteAll(id)

	if Sessions[gathering].ConnectionIDs.Size() == 0 {
		delete(Sessions, gathering)
	} else {
		// Update the participation count with the new connection ID count
		Sessions[gathering].GameMatchmakeSession.ParticipationCount.Value = uint32(Sessions[gathering].ConnectionIDs.Size())
	}
}

// FindConnectionSession searches for session the given connection ID is connected to
func FindConnectionSession(id uint32) uint32 {
	for gatheringID := range Sessions {
		if Sessions[gatheringID].ConnectionIDs.Has(id) {
			return gatheringID
		}
	}

	return 0
}

// RemoveConnectionFromAllSessions removes a connection from every session
func RemoveConnectionFromAllSessions(connection *nex.PRUDPConnection) {
	// * Keep checking until no session is found
	for gid := FindConnectionSession(connection.ID); gid != 0; {
		session := Sessions[gid]
		lenParticipants := session.ConnectionIDs.Size()

		RemoveConnectionIDFromSession(connection.ID, gid)

		if lenParticipants <= 1 {
			gid = FindConnectionSession(connection.ID)
			continue
		}

		ownerPID := session.GameMatchmakeSession.Gathering.OwnerPID

		if ownerPID.Equals(connection.PID()) {
			// * This flag tells the server to change the matchmake session owner if they disconnect
			// * If the flag is not set, delete the session
			// * More info: https://nintendo-wiki.pretendo.network/docs/nex/protocols/match-making/types#flags
			if session.GameMatchmakeSession.Gathering.Flags.PAND(match_making.GatheringFlags.DisconnectChangeOwner) == 0 {
				delete(Sessions, gid)
			} else {
				ChangeSessionOwner(connection, gid, true)
			}
		} else {
			endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)
			server := endpoint.Server

			category := notifications.NotificationCategories.Participation
			subtype := notifications.NotificationSubTypes.Participation.Disconnected

			oEvent := notifications_types.NewNotificationEvent()
			oEvent.PIDSource = connection.PID()
			oEvent.Type = types.NewPrimitiveU32(notifications.BuildNotificationType(category, subtype))
			oEvent.Param1 = types.NewPrimitiveU32(gid)
			oEvent.Param2 = types.NewPrimitiveU32(connection.PID().LegacyValue()) // TODO - This assumes a legacy client. This won't work on the Switch

			stream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

			oEvent.WriteTo(stream)

			rmcRequest := nex.NewRMCRequest(endpoint)
			rmcRequest.ProtocolID = notifications.ProtocolID
			rmcRequest.CallID = CurrentMatchmakingCallID.Next()
			rmcRequest.MethodID = notifications.MethodProcessNotificationEvent
			rmcRequest.Parameters = stream.Bytes()

			rmcRequestBytes := rmcRequest.Bytes()

			target := endpoint.FindConnectionByPID(ownerPID.Value())
			if target == nil {
				Logger.Warning("Target connection not found")
				gid = FindConnectionSession(connection.ID)
				continue
			}

			var messagePacket nex.PRUDPPacketInterface

			if connection.DefaultPRUDPVersion == 0 {
				messagePacket, _ = nex.NewPRUDPPacketV0(server, target, nil)
			} else {
				messagePacket, _ = nex.NewPRUDPPacketV1(server, target, nil)
			}

			messagePacket.SetType(constants.DataPacket)
			messagePacket.AddFlag(constants.PacketFlagNeedsAck)
			messagePacket.AddFlag(constants.PacketFlagReliable)
			messagePacket.SetSourceVirtualPortStreamType(connection.StreamType)
			messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
			messagePacket.SetDestinationVirtualPortStreamType(connection.StreamType)
			messagePacket.SetDestinationVirtualPortStreamID(connection.StreamID)
			messagePacket.SetPayload(rmcRequestBytes)

			server.Send(messagePacket)
		}

		gid = FindConnectionSession(connection.ID)
	}
}

// CreateSessionByMatchmakeSession creates a gathering from a MatchmakeSession
func CreateSessionByMatchmakeSession(matchmakeSession *match_making_types.MatchmakeSession, searchMatchmakeSession *match_making_types.MatchmakeSession, hostPID *types.PID) (*CommonMatchmakeSession, *nex.Error) {
	sessionIndex := GetAvailableGatheringID()
	if sessionIndex == 0 {
		sessionIndex = GetAvailableGatheringID() // * Skip to index 1
	}

	session := CommonMatchmakeSession{
		SearchMatchmakeSession: searchMatchmakeSession,
		GameMatchmakeSession:   matchmakeSession,
		ConnectionIDs:          nex.NewMutexSlice[uint32](),
	}

	session.GameMatchmakeSession.Gathering.ID = types.NewPrimitiveU32(sessionIndex)
	session.GameMatchmakeSession.Gathering.OwnerPID = hostPID
	session.GameMatchmakeSession.Gathering.HostPID = hostPID

	session.GameMatchmakeSession.StartedTime = types.NewDateTime(0).Now()
	session.GameMatchmakeSession.SessionKey = types.NewBuffer(make([]byte, 32))

	rand.Read(session.GameMatchmakeSession.SessionKey.Value)

	SR := types.NewVariant()
	SR.TypeID = types.NewPrimitiveU8(3)
	SR.Type = types.NewPrimitiveBool(true)

	GIR := types.NewVariant()
	GIR.TypeID = types.NewPrimitiveU8(1)
	GIR.Type = types.NewPrimitiveS64(3)

	session.GameMatchmakeSession.MatchmakeParam.Params.Set(types.NewString("@SR"), SR)
	session.GameMatchmakeSession.MatchmakeParam.Params.Set(types.NewString("@GIR"), GIR)

	Sessions[sessionIndex] = &session

	return Sessions[sessionIndex], nil
}

// FindSessionByMatchmakeSession finds a gathering that matches with a MatchmakeSession
func FindSessionByMatchmakeSession(pid *types.PID, searchMatchmakeSession *match_making_types.MatchmakeSession) uint32 {
	// * This portion finds any sessions that match the search session
	// * It does not care about anything beyond that, such as if the match is already full
	// * This is handled below
	candidateSessionIndexes := make([]uint32, 0, len(Sessions))
	for index, session := range Sessions {
		if session.SearchMatchmakeSession.Equals(searchMatchmakeSession) {
			candidateSessionIndexes = append(candidateSessionIndexes, index)
		}
	}

	// TODO - This whole section assumes legacy clients. None of it will work on the Switch
	var friendList []uint32
	for _, sessionIndex := range candidateSessionIndexes {
		sessionToCheck := Sessions[sessionIndex]
		if sessionToCheck.ConnectionIDs.Size() >= int(sessionToCheck.GameMatchmakeSession.MaximumParticipants.Value) {
			continue
		}

		if !sessionToCheck.GameMatchmakeSession.OpenParticipation.Value {
			continue
		}

		// * If the session only allows friends, check if the owner is in the friend list of the PID
		// TODO - Is this a flag or a constant?
		if sessionToCheck.GameMatchmakeSession.ParticipationPolicy.Value == 98 {
			if GetUserFriendPIDsHandler == nil {
				Logger.Warning("Missing GetUserFriendPIDsHandler!")
				continue
			}

			if len(friendList) == 0 {
				friendList = GetUserFriendPIDsHandler(pid.LegacyValue()) // TODO - This grpc method needs to support the Switch
			}

			if !slices.Contains(friendList, sessionToCheck.GameMatchmakeSession.OwnerPID.LegacyValue()) {
				continue
			}
		}

		return sessionIndex // * Found a match
	}

	return 0
}

// FindSessionsByMatchmakeSessionSearchCriterias finds a gathering that matches with the given search criteria
func FindSessionsByMatchmakeSessionSearchCriterias(pid *types.PID, searchCriterias []*match_making_types.MatchmakeSessionSearchCriteria, gameSpecificChecks func(searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, matchmakeSession *match_making_types.MatchmakeSession) bool) []*CommonMatchmakeSession {
	candidateSessions := make([]*CommonMatchmakeSession, 0, len(Sessions))

	// TODO - This whole section assumes legacy clients. None of it will work on the Switch
	var friendList []uint32
	for _, session := range Sessions {
		for _, criteria := range searchCriterias {
			// * Check things like game specific attributes
			if gameSpecificChecks != nil {
				if !gameSpecificChecks(criteria, session.GameMatchmakeSession) {
					continue
				}
			} else {
				if !compareAttributesSearchCriteria(session.GameMatchmakeSession.Attributes.Slice(), criteria.Attribs.Slice()) {
					continue
				}
			}

			if !compareSearchCriteria(session.GameMatchmakeSession.MaximumParticipants.Value, criteria.MaxParticipants.Value) {
				continue
			}

			if !compareSearchCriteria(session.GameMatchmakeSession.MinimumParticipants.Value, criteria.MinParticipants.Value) {
				continue
			}

			if !compareSearchCriteria(session.GameMatchmakeSession.MatchmakeSystemType.Value, criteria.MatchmakeSystemType.Value) {
				continue
			}

			if !compareSearchCriteria(session.GameMatchmakeSession.GameMode.Value, criteria.GameMode.Value) {
				continue
			}

			if session.ConnectionIDs.Size() >= int(session.GameMatchmakeSession.MaximumParticipants.Value) {
				continue
			}

			if !session.GameMatchmakeSession.OpenParticipation.Value {
				continue
			}

			// * If the session only allows friends, check if the owner is in the friend list of the PID
			// TODO - Is this a flag or a constant?
			if session.GameMatchmakeSession.ParticipationPolicy.Value == 98 {
				if GetUserFriendPIDsHandler == nil {
					Logger.Warning("Missing GetUserFriendPIDsHandler!")
					continue
				}

				if len(friendList) == 0 {
					friendList = GetUserFriendPIDsHandler(pid.LegacyValue()) // TODO - Support the Switch
				}

				if !slices.Contains(friendList, session.GameMatchmakeSession.OwnerPID.LegacyValue()) {
					continue
				}
			}

			candidateSessions = append(candidateSessions, session)

			// We don't have to compare with other search criterias
			break
		}
	}

	return candidateSessions
}

func compareAttributesSearchCriteria(original []*types.PrimitiveU32, search []*types.String) bool {
	if len(original) != len(search) {
		return false
	}

	for index, originalAttribute := range original {
		searchAttribute := search[index]

		if !compareSearchCriteria(originalAttribute.Value, searchAttribute.Value) {
			return false
		}
	}

	return true
}

func compareSearchCriteria[T ~uint16 | ~uint32](original T, search string) bool {
	if search == "" { // * Accept any value
		return true
	}

	before, after, found := strings.Cut(search, ",")
	if found {
		min, err := strconv.ParseUint(before, 10, 64)
		if err != nil {
			return false
		}

		max, err := strconv.ParseUint(after, 10, 64)
		if err != nil {
			return false
		}

		return min <= uint64(original) && max >= uint64(original)
	} else {
		searchNum, err := strconv.ParseUint(before, 10, 64)
		if err != nil {
			return false
		}

		return searchNum == uint64(original)
	}
}

// AddPlayersToSession updates the given sessions state to include the provided connection IDs
// Returns a NEX error code if failed
func AddPlayersToSession(session *CommonMatchmakeSession, connectionIDs []uint32, initiatingConnection *nex.PRUDPConnection, joinMessage string) *nex.Error {
	if (session.ConnectionIDs.Size() + len(connectionIDs)) > int(session.GameMatchmakeSession.Gathering.MaximumParticipants.Value) {
		return nex.NewError(nex.ResultCodes.RendezVous.SessionFull, fmt.Sprintf("Gathering %d is full", session.GameMatchmakeSession.Gathering.ID))
	}

	for _, connectedID := range connectionIDs {
		if session.ConnectionIDs.Has(connectedID) {
			return nex.NewError(nex.ResultCodes.RendezVous.AlreadyParticipatedGathering, fmt.Sprintf("Connection ID %d is already in gathering %d", connectedID, session.GameMatchmakeSession.Gathering.ID))
		}

		session.ConnectionIDs.Add(connectedID)

		// Update the participation count with the new connection ID count
		session.GameMatchmakeSession.ParticipationCount.Value = uint32(session.ConnectionIDs.Size())
	}

	endpoint := initiatingConnection.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server

	//session.ConnectionIDs.Each(func(_ int, connectionID uint32) bool {
		target := endpoint.FindConnectionByPID(session.GameMatchmakeSession.OwnerPID.Value())
		if target == nil {
			// TODO - Error here?
			Logger.Warning("Player not found")
			return nil
		}

		notificationCategory := notifications.NotificationCategories.Participation
		notificationSubtype := notifications.NotificationSubTypes.Participation.NewParticipant

		oEvent := notifications_types.NewNotificationEvent()
		oEvent.PIDSource = initiatingConnection.PID()
		oEvent.Type = types.NewPrimitiveU32(notifications.BuildNotificationType(notificationCategory, notificationSubtype))
		oEvent.Param1 = session.GameMatchmakeSession.ID.Copy().(*types.PrimitiveU32)
		oEvent.Param2 = types.NewPrimitiveU32(initiatingConnection.PID().LegacyValue()) // TODO - This assumes a legacy client. Will not work on the Switch
		oEvent.StrParam = types.NewString(joinMessage)
		//oEvent.Param3 = types.NewPrimitiveU32(uint32(len(connectionIDs)))

		notificationStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

		oEvent.WriteTo(notificationStream)

		notificationRequest := nex.NewRMCRequest(endpoint)
		notificationRequest.ProtocolID = notifications.ProtocolID
		notificationRequest.CallID = CurrentMatchmakingCallID.Next()
		notificationRequest.MethodID = notifications.MethodProcessNotificationEvent
		notificationRequest.Parameters = notificationStream.Bytes()

		notificationRequestBytes := notificationRequest.Bytes()

		var messagePacket nex.PRUDPPacketInterface

		if target.DefaultPRUDPVersion == 0 {
			messagePacket, _ = nex.NewPRUDPPacketV0(server, target, nil)
		} else {
			messagePacket, _ = nex.NewPRUDPPacketV1(server, target, nil)
		}

		messagePacket.SetType(constants.DataPacket)
		messagePacket.AddFlag(constants.PacketFlagNeedsAck)
		messagePacket.AddFlag(constants.PacketFlagReliable)
		messagePacket.SetSourceVirtualPortStreamType(target.StreamType)
		messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
		messagePacket.SetDestinationVirtualPortStreamType(target.StreamType)
		messagePacket.SetDestinationVirtualPortStreamID(target.StreamID)
		messagePacket.SetPayload(notificationRequestBytes)

		server.Send(messagePacket)

		//return false
	//})

	// * This appears to be correct. Tri-Force Heroes uses 3.9.0,
	// * and has issues if these notifications are sent.
	// * Minecraft, however, requires these to be sent
	// TODO - Check other games both pre and post 3.10.0 and validate
	/*if server.LibraryVersions.MatchMaking.GreaterOrEqual("3.10.0")* / {
		session.ConnectionIDs.Each(func(_ int, connectionID uint32) bool {
			target := endpoint.FindConnectionByID(connectionID)
			if target == nil {
				// TODO - Error here?
				Logger.Warning("Player not found")
				return false
			}

			notificationCategory := notifications.NotificationCategories.Participation
			notificationSubtype := notifications.NotificationSubTypes.Participation.NewParticipant

			oEvent := notifications_types.NewNotificationEvent()
			oEvent.PIDSource = initiatingConnection.PID()
			oEvent.Type = types.NewPrimitiveU32(notifications.BuildNotificationType(notificationCategory, notificationSubtype))
			oEvent.Param1 = session.GameMatchmakeSession.ID.Copy().(*types.PrimitiveU32)
			oEvent.Param2 = types.NewPrimitiveU32(target.PID().LegacyValue()) // TODO - This assumes a legacy client. Will not work on the Switch
			oEvent.StrParam = types.NewString(joinMessage)
			oEvent.Param3 = types.NewPrimitiveU32(uint32(len(connectionIDs)))

			notificationStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

			oEvent.WriteTo(notificationStream)

			notificationRequest := nex.NewRMCRequest(endpoint)
			notificationRequest.ProtocolID = notifications.ProtocolID
			notificationRequest.CallID = CurrentMatchmakingCallID.Next()
			notificationRequest.MethodID = notifications.MethodProcessNotificationEvent
			notificationRequest.Parameters = notificationStream.Bytes()

			notificationRequestBytes := notificationRequest.Bytes()

			var messagePacket nex.PRUDPPacketInterface

			if target.DefaultPRUDPVersion == 0 {
				messagePacket, _ = nex.NewPRUDPPacketV0(server, initiatingConnection, nil)
			} else {
				messagePacket, _ = nex.NewPRUDPPacketV1(server, initiatingConnection, nil)
			}

			messagePacket.SetType(constants.DataPacket)
			messagePacket.AddFlag(constants.PacketFlagNeedsAck)
			messagePacket.AddFlag(constants.PacketFlagReliable)
			messagePacket.SetSourceVirtualPortStreamType(target.StreamType)
			messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
			messagePacket.SetDestinationVirtualPortStreamType(target.StreamType)
			messagePacket.SetDestinationVirtualPortStreamID(target.StreamID)
			messagePacket.SetPayload(notificationRequestBytes)

			server.Send(messagePacket)

			return false
		})

		notificationCategory := notifications.NotificationCategories.Participation
		notificationSubtype := notifications.NotificationSubTypes.Participation.NewParticipant

		oEvent := notifications_types.NewNotificationEvent()
		oEvent.PIDSource = initiatingConnection.PID()
		oEvent.Type = types.NewPrimitiveU32(notifications.BuildNotificationType(notificationCategory, notificationSubtype))
		oEvent.Param1 = session.GameMatchmakeSession.ID
		oEvent.Param2 = types.NewPrimitiveU32(initiatingConnection.PID().LegacyValue()) // TODO - This assumes a legacy client. Will not work on the Switch
		oEvent.StrParam = types.NewString(joinMessage)
		oEvent.Param3 = types.NewPrimitiveU32(uint32(len(connectionIDs)))

		notificationStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

		oEvent.WriteTo(notificationStream)

		notificationRequest := nex.NewRMCRequest(endpoint)
		notificationRequest.ProtocolID = notifications.ProtocolID
		notificationRequest.CallID = CurrentMatchmakingCallID.Next()
		notificationRequest.MethodID = notifications.MethodProcessNotificationEvent
		notificationRequest.Parameters = notificationStream.Bytes()

		notificationRequestBytes := notificationRequest.Bytes()

		var messagePacket nex.PRUDPPacketInterface

		if initiatingConnection.DefaultPRUDPVersion == 0 {
			messagePacket, _ = nex.NewPRUDPPacketV0(server, initiatingConnection, nil)
		} else {
			messagePacket, _ = nex.NewPRUDPPacketV1(server, initiatingConnection, nil)
		}

		messagePacket.SetType(constants.DataPacket)
		messagePacket.AddFlag(constants.PacketFlagNeedsAck)
		messagePacket.AddFlag(constants.PacketFlagReliable)
		messagePacket.SetSourceVirtualPortStreamType(initiatingConnection.StreamType)
		messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
		messagePacket.SetDestinationVirtualPortStreamType(initiatingConnection.StreamType)
		messagePacket.SetDestinationVirtualPortStreamID(initiatingConnection.StreamID)
		messagePacket.SetPayload(notificationRequestBytes)

		server.Send(messagePacket)

		target := endpoint.FindConnectionByPID(session.GameMatchmakeSession.Gathering.OwnerPID.Value())
		if target == nil {
			// TODO - Error here?
			Logger.Warning("Player not found")
			return nil
		}

		if target.DefaultPRUDPVersion == 0 {
			messagePacket, _ = nex.NewPRUDPPacketV0(server, target, nil)
		} else {
			messagePacket, _ = nex.NewPRUDPPacketV1(server, target, nil)
		}

		messagePacket.SetType(constants.DataPacket)
		messagePacket.AddFlag(constants.PacketFlagNeedsAck)
		messagePacket.AddFlag(constants.PacketFlagReliable)
		messagePacket.SetSourceVirtualPortStreamType(target.StreamType)
		messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
		messagePacket.SetDestinationVirtualPortStreamType(target.StreamType)
		messagePacket.SetDestinationVirtualPortStreamID(target.StreamID)
		messagePacket.SetPayload(notificationRequestBytes)

		server.Send(messagePacket)
	}*/

	return nil
}

// ChangeSessionOwner changes the session owner to a different connection
func ChangeSessionOwner(currentOwner *nex.PRUDPConnection, gathering uint32, isLeaving bool) {
	endpoint := currentOwner.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server
	session := Sessions[gathering]

	var newOwner *nex.PRUDPConnection

	newOwnerConnectionID := FindOtherConnectionID(currentOwner.ID, gathering)
	if newOwnerConnectionID != 0 {
		newOwner = endpoint.FindConnectionByID(newOwnerConnectionID)
		if newOwner == nil {
			Logger.Warning("Other connection not found")
			return
		}

		// If the current owner is the host and they are leaving, change it by the new owner
		if session.GameMatchmakeSession.Gathering.HostPID.Equals(currentOwner.PID()) && isLeaving {
			session.GameMatchmakeSession.Gathering.HostPID = newOwner.PID()
		}
		session.GameMatchmakeSession.Gathering.OwnerPID = newOwner.PID()
	} else {
		return
	}

	category := notifications.NotificationCategories.OwnershipChanged
	subtype := notifications.NotificationSubTypes.OwnershipChanged.None

	oEvent := notifications_types.NewNotificationEvent()
	oEvent.PIDSource = currentOwner.PID()
	oEvent.Type = types.NewPrimitiveU32(notifications.BuildNotificationType(category, subtype))
	oEvent.Param1 = types.NewPrimitiveU32(gathering)
	oEvent.Param2 = types.NewPrimitiveU32(newOwner.PID().LegacyValue()) // TODO - This assumes a legacy client. Will not work on the Switch

	// TODO - StrParam doesn't have this value on some servers
	// * https://github.com/kinnay/NintendoClients/issues/101
	// * unixTime := time.Now()
	// * oEvent.StrParam = strconv.FormatInt(unixTime.UnixMicro(), 10)

	stream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	oEvent.WriteTo(stream)

	rmcRequest := nex.NewRMCRequest(endpoint)
	rmcRequest.ProtocolID = notifications.ProtocolID
	rmcRequest.CallID = CurrentMatchmakingCallID.Next()
	rmcRequest.MethodID = notifications.MethodProcessNotificationEvent
	rmcRequest.Parameters = stream.Bytes()

	rmcRequestBytes := rmcRequest.Bytes()

	session.ConnectionIDs.Each(func(_ int, connectionID uint32) bool {
		target := endpoint.FindConnectionByID(connectionID)
		if target == nil {
			Logger.Warning("Connection not found")
			return false
		}

		var messagePacket nex.PRUDPPacketInterface

		if target.DefaultPRUDPVersion == 0 {
			messagePacket, _ = nex.NewPRUDPPacketV0(server, target, nil)
		} else {
			messagePacket, _ = nex.NewPRUDPPacketV1(server, target, nil)
		}

		messagePacket.SetType(constants.DataPacket)
		messagePacket.AddFlag(constants.PacketFlagNeedsAck)
		messagePacket.AddFlag(constants.PacketFlagReliable)
		messagePacket.SetSourceVirtualPortStreamType(target.StreamType)
		messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
		messagePacket.SetDestinationVirtualPortStreamType(target.StreamType)
		messagePacket.SetDestinationVirtualPortStreamID(target.StreamID)
		messagePacket.SetPayload(rmcRequestBytes)

		server.Send(messagePacket)

		return false
	})
}
