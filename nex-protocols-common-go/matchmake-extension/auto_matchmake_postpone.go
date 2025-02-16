package matchmake_extension

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
)

func (commonProtocol *CommonProtocol) autoMatchmakePostpone(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message *types.String) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.CleanupSearchMatchmakeSession == nil {
		common_globals.Logger.Warning("MatchmakeExtension::AutoMatchmake_Postpone missing CleanupSearchMatchmakeSession!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	// * A client may disconnect from a session without leaving reliably,
	// * so let's make sure the client is removed from the session
	common_globals.RemoveConnectionFromAllSessions(connection)

	var matchmakeSession *match_making_types.MatchmakeSession
	anyGatheringDataType := anyGathering.TypeName

	if anyGatheringDataType.Value == "MatchmakeSession" {
		matchmakeSession = anyGathering.ObjectData.(*match_making_types.MatchmakeSession)
	} else {
		common_globals.Logger.Critical("Non-MatchmakeSession DataType?!")
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	searchMatchmakeSession := matchmakeSession.Copy().(*match_making_types.MatchmakeSession)
	commonProtocol.CleanupSearchMatchmakeSession(searchMatchmakeSession)
	sessionIndex := common_globals.FindSessionByMatchmakeSession(connection.PID(), searchMatchmakeSession)
	var session *common_globals.CommonMatchmakeSession

	if sessionIndex == 0 {
		var errCode *nex.Error
		session, errCode = common_globals.CreateSessionByMatchmakeSession(matchmakeSession, searchMatchmakeSession, connection.PID())
		if err != nil {
			common_globals.Logger.Error(errCode.Error())
			return nil, errCode
		}
	} else {
		session = common_globals.Sessions[sessionIndex]
	}

	errCode := common_globals.AddPlayersToSession(session, []uint32{connection.ID}, connection, message.Value)
	if errCode != nil {
		common_globals.Logger.Error(errCode.Error())
		return nil, errCode
	}

	matchmakeDataHolder := types.NewAnyDataHolder()

	matchmakeDataHolder.TypeName = types.NewString("MatchmakeSession")
	matchmakeDataHolder.ObjectData = session.GameMatchmakeSession.Copy()

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	matchmakeDataHolder.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmake_extension.ProtocolID
	rmcResponse.MethodID = matchmake_extension.MethodAutoMatchmakePostpone
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterAutoMatchmakePostpone != nil {
		go commonProtocol.OnAfterAutoMatchmakePostpone(packet, anyGathering, message)
	}

	return rmcResponse, nil
}
