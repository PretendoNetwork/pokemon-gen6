package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"encoding/hex"
	"fmt"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
)

func GetSubscriptionData(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	for _, pid := range pids {
		fmt.Println(pid)
		content := globals.Timeline[uint32(pid)]
		types.UInt32(1).WriteTo(rmcResponseStream)
		types.UInt32(uint32(pid)).WriteTo(rmcResponseStream)
		for i := 0; i < len(content); i++ {
			types.UInt8(content[i]).WriteTo(rmcResponseStream)
		}
	}

	rmcResponseBody := rmcResponseStream.Bytes()
	_ = rmcResponseBody
	fmt.Println(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

