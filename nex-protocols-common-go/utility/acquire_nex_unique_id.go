package utility

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	utility "github.com/PretendoNetwork/nex-protocols-go/v2/utility"
)

func (commonProtocol *CommonProtocol) acquireNexUniqueID(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	if commonProtocol.GenerateNEXUniqueID == nil {
		common_globals.Logger.Warning("Utility::AcquireNexUniqueID missing GenerateNEXUniqueID!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	pNexUniqueID := types.NewPrimitiveU64(commonProtocol.GenerateNEXUniqueID())

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pNexUniqueID.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = utility.ProtocolID
	rmcResponse.MethodID = utility.MethodAcquireNexUniqueID
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterAcquireNexUniqueID != nil {
		go commonProtocol.OnAfterAcquireNexUniqueID(packet)
	}

	return rmcResponse, nil
}
