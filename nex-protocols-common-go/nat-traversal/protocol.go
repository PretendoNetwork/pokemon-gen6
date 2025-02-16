package nattraversal

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
)

type CommonProtocol struct {
	endpoint                              nex.EndpointInterface
	protocol                              nat_traversal.Interface
	OnAfterRequestProbeInitiationExt      func(packet nex.PacketInterface, targetList *types.List[*types.String], stationToProbe *types.String)
	OnAfterReportNATProperties            func(packet nex.PacketInterface, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32)
	OnAfterReportNATTraversalResult       func(packet nex.PacketInterface, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32)
	OnAfterGetRelaySignatureKey           func(packet nex.PacketInterface)
	OnAfterReportNATTraversalResultDetail func(packet nex.PacketInterface, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32)
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol nat_traversal.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		endpoint: protocol.Endpoint(),
		protocol: protocol,
	}

	protocol.SetHandlerRequestProbeInitiationExt(commonProtocol.requestProbeInitiationExt)
	protocol.SetHandlerReportNATProperties(commonProtocol.reportNATProperties)
	protocol.SetHandlerReportNATTraversalResult(commonProtocol.reportNATTraversalResult)
	protocol.SetHandlerGetRelaySignatureKey(commonProtocol.getRelaySignatureKey)
	protocol.SetHandlerReportNATTraversalResultDetail(commonProtocol.reportNATTraversalResultDetail)

	return commonProtocol
}
