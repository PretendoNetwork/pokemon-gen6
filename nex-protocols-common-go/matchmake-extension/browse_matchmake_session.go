package matchmake_extension

import (
	"math"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
)

func (commonProtocol *CommonProtocol) browseMatchmakeSession(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	searchCriterias := []*match_making_types.MatchmakeSessionSearchCriteria{searchCriteria}

	sessions := common_globals.FindSessionsByMatchmakeSessionSearchCriterias(connection.PID(), searchCriterias, commonProtocol.GameSpecificMatchmakeSessionSearchCriteriaChecks)

	// TODO - Is this right?
	if resultRange.Offset.Value != math.MaxUint32 {
		if len(sessions) < int(resultRange.Offset.Value) {
			return nil, nex.NewError(nex.ResultCodes.Core.InvalidIndex, "change_error")
		}

		sessions = sessions[resultRange.Offset.Value:]
	}


	if len(sessions) > int(resultRange.Length.Value) {
		sessions = sessions[:resultRange.Length.Value]
	}

	lstGathering := types.NewList[*types.AnyDataHolder]()
	lstGathering.Type = types.NewAnyDataHolder()

	for _, session := range sessions {
		matchmakeSessionDataHolder := types.NewAnyDataHolder()
		matchmakeSessionDataHolder.TypeName = types.NewString("MatchmakeSession")
		matchmakeSessionDataHolder.ObjectData = session.GameMatchmakeSession.Copy()

		lstGathering.Append(matchmakeSessionDataHolder)
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstGathering.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = matchmake_extension.ProtocolID
	rmcResponse.MethodID = matchmake_extension.MethodBrowseMatchmakeSession
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterBrowseMatchmakeSession != nil {
		go commonProtocol.OnAfterBrowseMatchmakeSession(packet, searchCriteria, resultRange)
	}

	return rmcResponse, nil
}
