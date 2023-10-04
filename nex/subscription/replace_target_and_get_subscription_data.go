package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	"github.com/PretendoNetwork/nex-protocols-go/subscription"
)

func ReplaceTargetAndGetSubscriptionData(err error, client *nex.Client, callID uint32) {
	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteUInt32LE(0)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCResponse(subscription.ProtocolID, callID)
	rmcResponse.SetSuccess(subscription.MethodReplaceTargetAndGetSubscriptionData, rmcResponseBody)

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
