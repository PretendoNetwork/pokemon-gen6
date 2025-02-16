package matchmake_extension

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
)

func (commonProtocol *CommonProtocol) autoMatchmakeWithParamPostpone(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	// * A client may disconnect from a session without leaving reliably,
	// * so let's make sure the client is removed from the session
	common_globals.RemoveConnectionFromAllSessions(connection)

	matchmakeSession := autoMatchmakeParam.SourceMatchmakeSession

	sessions := common_globals.FindSessionsByMatchmakeSessionSearchCriterias(connection.PID(), autoMatchmakeParam.LstSearchCriteria.Slice(), commonProtocol.GameSpecificMatchmakeSessionSearchCriteriaChecks)
	var session *common_globals.CommonMatchmakeSession

	if len(sessions) == 0 {
		var errCode *nex.Error
		session, errCode = common_globals.CreateSessionByMatchmakeSession(matchmakeSession, nil, connection.PID())
		if errCode != nil {
			common_globals.Logger.Error(errCode.Error())
			return nil, errCode
		}
	} else {
		session = sessions[0]
	}

	errCode := common_globals.AddPlayersToSession(session, []uint32{connection.ID}, connection, "")
	if errCode != nil {
		common_globals.Logger.Error(errCode.Error())
		return nil, errCode
	}

	matchmakeDataHolder := types.NewAnyDataHolder()

	matchmakeDataHolder.TypeName = types.NewString("MatchmakeSession")
	matchmakeDataHolder.ObjectData = session.GameMatchmakeSession.Copy()

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	session.GameMatchmakeSession.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmake_extension.ProtocolID
	rmcResponse.MethodID = matchmake_extension.MethodAutoMatchmakeWithParamPostpone
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterAutoMatchmakeWithParamPostpone != nil {
		go commonProtocol.OnAfterAutoMatchmakeWithParamPostpone(packet, autoMatchmakeParam)
	}

	return rmcResponse, nil
}
