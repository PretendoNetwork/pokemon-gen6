package nex_subscription

import (
	"encoding/hex"
	"fmt"
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
)

func GetActivePlayerSubscriptionData(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	types.UInt32(uint32(len(globals.Timeline))).WriteTo(rmcResponseStream)
	for clientPID := range globals.Timeline {
		types.UInt32(clientPID).WriteTo(rmcResponseStream)
		for j := 0; j < len(globals.Timeline[clientPID]); j++ {
			types.UInt8(globals.Timeline[clientPID][j]).WriteTo(rmcResponseStream)
		}
	}

	rmcResponseBody := rmcResponseStream.Bytes()
	fmt.Println(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetActivePlayerSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

