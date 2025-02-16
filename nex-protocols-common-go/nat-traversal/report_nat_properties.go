package nattraversal

import (
	"github.com/PretendoNetwork/nex-go/v2"
	//"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
)

func (commonProtocol *CommonProtocol) reportNATProperties(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	for _, station := range connection.StationURLs.Slice() {
		if !station.IsPublic() {
			//station.SetNATMapping(constants.NATMappingProperties(natmapping.Value))
			//station.SetNATFiltering(constants.NATFilteringProperties(natfiltering.Value))
		}

		station.SetRVConnectionID(connection.ID)
		station.SetPrincipalID(connection.PID())
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = nat_traversal.ProtocolID
	rmcResponse.MethodID = nat_traversal.MethodReportNATProperties
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterReportNATProperties != nil {
		go commonProtocol.OnAfterReportNATProperties(packet, natmapping, natfiltering, rtt)
	}

	return rmcResponse, nil
}
