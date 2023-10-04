package nex_subscription

import (
	"encoding/hex"
	"fmt"
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	"github.com/PretendoNetwork/nex-protocols-go/subscription"
)

func GetActivePlayerSubscriptionData(err error, client *nex.Client, callID uint32) {
	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteUInt32LE(uint32(len(globals.Timeline)))
	for clientPID := range globals.Timeline {
		for j := 0; j < len(globals.Timeline[clientPID]); j++ {
			rmcResponseStream.WriteUInt8(globals.Timeline[clientPID][j])
		}
	}

	rmcResponseBody := rmcResponseStream.Bytes()
	fmt.Println(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCResponse(subscription.ProtocolID, callID)
	rmcResponse.SetSuccess(subscription.MethodGetActivePlayerSubscriptionData, rmcResponseBody)

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
