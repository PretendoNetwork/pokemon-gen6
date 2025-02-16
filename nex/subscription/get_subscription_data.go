package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"encoding/hex"
	"fmt"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
)

func GetSubscriptionData(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	for _, pid := range pids.Slice() {
		fmt.Println(pid)
		content := globals.Timeline[pid.Value]
		types.NewPrimitiveU32(1).WriteTo(rmcResponseStream)
		types.NewPrimitiveU32(pid.Value).WriteTo(rmcResponseStream)
		for i := 0; i < len(content); i++ {
			types.NewPrimitiveU8(content[i]).WriteTo(rmcResponseStream)
			//rmcResponseStream.WriteUInt8(content[i])
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

	/*rmcResponse := nex.NewRMCResponse(subscription.ProtocolID, callID)
	rmcResponse.SetSuccess(subscription.MethodGetSubscriptionData, rmcResponseBody)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.SecureServer.Send(responsePacket)*/
}
