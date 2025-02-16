// Package globals implements variables and functions used by all protocol packages
package globals

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
)

// RespondError sends the client a given error code
func RespondError(packet nex.PacketInterface, protocolID uint16, err error) {
	sender := packet.Sender()
	request := packet.RMCMessage()
	errorCode := nex.ResultCodes.Core.Unknown

	if err, ok := err.(*nex.Error); ok {
		errorCode = err.ResultCode
		err.Packet = packet

		packet.Sender().Endpoint().EmitError(err)
	}

	rmcResponse := nex.NewRMCError(sender.Endpoint(), errorCode)
	rmcResponse.ProtocolID = request.ProtocolID
	rmcResponse.CallID = request.CallID

	var responsePacket nex.PacketInterface

	switch packet := packet.(type) {
	case nex.PRUDPPacketInterface:
		rmcResponseBytes := rmcResponse.Bytes()

		// * Go won't type assert responsePacket in the version check below,
		// * so to avoid a bunch of assertions just create a temp variable
		var prudpPacket nex.PRUDPPacketInterface

		endpoint := sender.(*nex.PRUDPConnection).Endpoint()
		server := endpoint.(*nex.PRUDPEndPoint).Server
		if packet.Version() == 1 {
			prudpPacket, _ = nex.NewPRUDPPacketV1(server, sender.(*nex.PRUDPConnection), nil)
		} else {
			prudpPacket, _ = nex.NewPRUDPPacketV0(server, sender.(*nex.PRUDPConnection), nil)
		}

		prudpPacket.SetType(constants.DataPacket)

		if packet.HasFlag(constants.PacketFlagReliable) {
			prudpPacket.AddFlag(constants.PacketFlagReliable)
		}

		prudpPacket.AddFlag(constants.PacketFlagNeedsAck)
		prudpPacket.SetSourceVirtualPortStreamType(packet.DestinationVirtualPortStreamType())
		prudpPacket.SetSourceVirtualPortStreamID(packet.DestinationVirtualPortStreamID())
		prudpPacket.SetDestinationVirtualPortStreamType(packet.SourceVirtualPortStreamType())
		prudpPacket.SetDestinationVirtualPortStreamID(packet.SourceVirtualPortStreamID())

		responsePacket = prudpPacket
		responsePacket.SetPayload(rmcResponseBytes)
	case *nex.HPPPacket:
		// * We reuse the same packet from input and replace
		// * the RMC message so that it can be delivered back
		responsePacket = packet
		responsePacket.SetRMCMessage(rmcResponse)
	}

	sender.Endpoint().Send(responsePacket)
}
