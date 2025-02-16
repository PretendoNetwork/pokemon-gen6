package matchmake_extension

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	"golang.org/x/exp/slices"
)

func (commonProtocol *CommonProtocol) getSimplePlayingSession(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	// * Does nothing if element is not present in the List
	listPID.Remove(connection.PID())

	if includeLoginUser.Value && !listPID.Contains(connection.PID()) {
		listPID.Append(connection.PID().Copy().(*types.PID))
	}

	simplePlayingSessions := make(map[string]*match_making_types.SimplePlayingSession)

	for gatheringID, session := range common_globals.Sessions {
		for _, pid := range listPID.Slice() {
			key := fmt.Sprintf("%d-%d", gatheringID, pid.Value())
			if simplePlayingSessions[key] == nil {
				connectedPIDs := make([]uint64, 0)
				session.ConnectionIDs.Each(func(_ int, connectionID uint32) bool {
					player := endpoint.FindConnectionByID(connectionID)
					if player == nil {
						common_globals.Logger.Warning("Player not found")
						return false
					}

					connectedPIDs = append(connectedPIDs, player.PID().Value())
					return false
				})

				if slices.Contains(connectedPIDs, pid.Value()) {
					attribute0, err := session.GameMatchmakeSession.Attributes.Get(0)
					if err != nil {
						common_globals.Logger.Error(err.Error())
						return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
					}

					simplePlayingSessions[key] = match_making_types.NewSimplePlayingSession()
					simplePlayingSessions[key].PrincipalID = pid.Copy().(*types.PID)
					simplePlayingSessions[key].GatheringID = types.NewPrimitiveU32(gatheringID)
					simplePlayingSessions[key].GameMode = session.GameMatchmakeSession.GameMode.Copy().(*types.PrimitiveU32)
					simplePlayingSessions[key].Attribute0 = attribute0.Copy().(*types.PrimitiveU32)
				}
			}
		}
	}

	lstSimplePlayingSession := types.NewList[*match_making_types.SimplePlayingSession]()

	for _, simplePlayingSession := range simplePlayingSessions {
		lstSimplePlayingSession.Append(simplePlayingSession)
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstSimplePlayingSession.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmake_extension.ProtocolID
	rmcResponse.MethodID = matchmake_extension.MethodGetSimplePlayingSession
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetSimplePlayingSession != nil {
		go commonProtocol.OnAfterGetSimplePlayingSession(packet, listPID, includeLoginUser)
	}

	return rmcResponse, nil
}
