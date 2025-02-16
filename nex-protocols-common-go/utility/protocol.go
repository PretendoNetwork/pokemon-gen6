package utility

import (
	"github.com/PretendoNetwork/nex-go/v2"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
)

type CommonProtocol struct {
	endpoint                  nex.EndpointInterface
	protocol                  utility.Interface
	GenerateNEXUniqueID       func() uint64
	OnAfterAcquireNexUniqueID func(packet nex.PacketInterface)
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol utility.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		endpoint: protocol.Endpoint(),
		protocol: protocol,
	}

	protocol.SetHandlerAcquireNexUniqueID(commonProtocol.acquireNexUniqueID)

	return commonProtocol
}
