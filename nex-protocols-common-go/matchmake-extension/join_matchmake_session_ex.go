package matchmake_extension

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
)

func (commonProtocol *CommonProtocol) joinMatchmakeSessionEx(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage *types.String, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	session, ok := common_globals.Sessions[gid.Value]
	if !ok {
		return nil, nex.NewError(nex.ResultCodes.RendezVous.SessionVoid, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server

	// TODO - More checks here
	errCode := common_globals.AddPlayersToSession(session, []uint32{connection.ID}, connection, strMessage.Value)
	if errCode != nil {
		common_globals.Logger.Error(errCode.Error())
		return nil, errCode
	}

	joinedMatchmakeSession := session.GameMatchmakeSession

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	if server.LibraryVersions.MatchMaking.GreaterOrEqual("3.0.0") {
		joinedMatchmakeSession.SessionKey.WriteTo(rmcResponseStream)
	}

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmake_extension.ProtocolID
	rmcResponse.MethodID = matchmake_extension.MethodJoinMatchmakeSessionEx
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterJoinMatchmakeSessionEx != nil {
		go commonProtocol.OnAfterJoinMatchmakeSessionEx(packet, gid, strMessage, dontCareMyBlockList, participationCount)
	}

	return rmcResponse, nil
}
