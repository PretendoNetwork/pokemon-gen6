package nattraversal

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
)

func (commonProtocol *CommonProtocol) reportNATTraversalResultDetail(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = nat_traversal.ProtocolID
	rmcResponse.MethodID = nat_traversal.MethodReportNATTraversalResultDetail
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterReportNATTraversalResultDetail != nil {
		go commonProtocol.OnAfterReportNATTraversalResultDetail(packet, cid, result, detail, rtt)
	}

	return rmcResponse, nil
}
