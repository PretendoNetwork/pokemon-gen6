package nattraversal

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"

	//"encoding/hex"
)

func (commonProtocol *CommonProtocol) getRelaySignatureKey(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	relayMode := types.NewPrimitiveS32(0xff)        // * Unknown, 1 appears to use direct connection?
	currentUTCTime := types.NewDateTime(0).Now() // * Current time for the relay server, UTC
	address := types.NewString("")               // * Relay server address. We don't have one, so for now this is empty.
	port := types.NewPrimitiveU16(0)             // * Relay server port. We don't have one, so for now this is empty.
	relayAddressType := types.NewPrimitiveS32(0xff) // * Unknown, 1 appears to use direct connection?
	gameServerID := types.NewPrimitiveU32(0)     // * Game Server ID. I don't know if this is checked (it doesn't appear to be though).

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	relayMode.WriteTo(rmcResponseStream)
	currentUTCTime.WriteTo(rmcResponseStream)
	address.WriteTo(rmcResponseStream)
	port.WriteTo(rmcResponseStream)
	relayAddressType.WriteTo(rmcResponseStream)
	gameServerID.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()
	//rmcResponseBody, _ = hex.DecodeString("01000000ef748f991f00000021006c312d30303035356430312e722e6e2e6170702e6e696e74656e646f2e6e657400000001000000015d0500")

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = nat_traversal.ProtocolID
	rmcResponse.MethodID = nat_traversal.MethodGetRelaySignatureKey
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetRelaySignatureKey != nil {
		go commonProtocol.OnAfterGetRelaySignatureKey(packet)
	}

	return rmcResponse, nil
}
