package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go"
	"encoding/hex"
	"fmt"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	"github.com/PretendoNetwork/nex-protocols-go/subscription"
)

func GetSubscriptionData(err error, client *nex.Client, callID uint32, pids []uint32) {
	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	for _, pid := range pids {
		fmt.Println(pid)
		content := globals.Timeline[pid]
		rmcResponseStream.WriteUInt32LE(1)
		for i := 0; i < len(content); i++ {
			rmcResponseStream.WriteUInt8(content[i])
		}
	}

	rmcResponseBody := rmcResponseStream.Bytes()
	_ = rmcResponseBody
	fmt.Println(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCResponse(subscription.ProtocolID, callID)
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

	globals.SecureServer.Send(responsePacket)
}
