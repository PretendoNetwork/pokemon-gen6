package matchmaking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	_ "github.com/PretendoNetwork/nex-protocols-go/v2"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
)

type CommonProtocol struct {
	endpoint                   *nex.PRUDPEndPoint
	protocol                   match_making.Interface
	OnAfterUnregisterGathering func(packet nex.PacketInterface, idGathering *types.PrimitiveU32)
	OnAfterFindBySingleID      func(packet nex.PacketInterface, id *types.PrimitiveU32)
	OnAfterUpdateSessionURL    func(packet nex.PacketInterface, idGathering *types.PrimitiveU32, strURL *types.String)
	OnAfterUpdateSessionHostV1 func(packet nex.PacketInterface, gid *types.PrimitiveU32)
	OnAfterGetSessionURLs      func(packet nex.PacketInterface, gid *types.PrimitiveU32)
	OnAfterUpdateSessionHost   func(packet nex.PacketInterface, gid *types.PrimitiveU32, isMigrateOwner *types.PrimitiveBool)
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol match_making.Interface) *CommonProtocol {
	endpoint := protocol.Endpoint().(*nex.PRUDPEndPoint)

	commonProtocol := &CommonProtocol{
		endpoint: endpoint,
		protocol: protocol,
	}

	common_globals.Sessions = make(map[uint32]*common_globals.CommonMatchmakeSession)

	protocol.SetHandlerUnregisterGathering(commonProtocol.unregisterGathering)
	protocol.SetHandlerFindBySingleID(commonProtocol.findBySingleID)
	protocol.SetHandlerUpdateSessionURL(commonProtocol.updateSessionURL)
	protocol.SetHandlerUpdateSessionHostV1(commonProtocol.updateSessionHostV1)
	protocol.SetHandlerGetSessionURLs(commonProtocol.getSessionURLs)
	protocol.SetHandlerUpdateSessionHost(commonProtocol.updateSessionHost)

	endpoint.OnConnectionEnded(func(connection *nex.PRUDPConnection) {
		common_globals.RemoveConnectionFromAllSessions(connection)
	})

	return commonProtocol
}
