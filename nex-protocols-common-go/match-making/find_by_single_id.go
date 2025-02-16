package matchmaking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
)

func (commonProtocol *CommonProtocol) findBySingleID(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	session, ok := common_globals.Sessions[id.Value]
	if !ok {
		return nil, nex.NewError(nex.ResultCodes.RendezVous.SessionVoid, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	bResult := types.NewPrimitiveBool(true)
	pGathering := types.NewAnyDataHolder()

	pGathering.TypeName = types.NewString("MatchmakeSession")
	pGathering.ObjectData = session.GameMatchmakeSession.Copy()

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	bResult.WriteTo(rmcResponseStream)
	pGathering.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = match_making.ProtocolID
	rmcResponse.MethodID = match_making.MethodFindBySingleID
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterFindBySingleID != nil {
		go commonProtocol.OnAfterFindBySingleID(packet, id)
	}

	return rmcResponse, nil
}
