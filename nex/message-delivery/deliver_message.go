package nex_message_delivery

import (
	"strconv"
	"time"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func DeliverMessage(err error, packet nex.PacketInterface, callID uint32, oUserMessage types.DataHolder) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server

	var idRecipient uint64

	switch oUserMessage.Object.ObjectID() {
	case types.NewString("BinaryMessage"):
		binaryMessage := oUserMessage.Object.(messaging_types.BinaryMessage)

		binaryMessage.StrSender = types.NewString(strconv.Itoa(int(client.PID())))
		binaryMessage.PIDSender = client.PID()
		binaryMessage.Receptiontime.FromTimestamp(time.Now().UTC())
		idRecipient = uint64(binaryMessage.IDRecipient)

		oUserMessage.Object = binaryMessage
	case types.NewString("TextMessage"):
		textMessage := oUserMessage.Object.(messaging_types.TextMessage)

		textMessage.StrSender = types.NewString(strconv.Itoa(int(client.PID())))
		textMessage.PIDSender = client.PID()
		textMessage.Receptiontime.FromTimestamp(time.Now().UTC())
		idRecipient = uint64(textMessage.IDRecipient)

		oUserMessage.Object = textMessage
	default:
		idRecipient = 0
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "Invalid type for deliver message")
	}

	messageStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())
	oUserMessage.WriteTo(messageStream)

	messageRequest := nex.NewRMCRequest(endpoint)
	messageRequest.ProtocolID = message_delivery.ProtocolID
	messageRequest.CallID = callID
	messageRequest.MethodID = message_delivery.MethodDeliverMessage
	messageRequest.Parameters = messageStream.Bytes()

	messageRequestBytes := messageRequest.Bytes()

	target := endpoint.FindConnectionByPID(idRecipient)

	var messagePacket nex.PRUDPPacketInterface

	if target.DefaultPRUDPVersion == 0 {
		messagePacket, _ = nex.NewPRUDPPacketV0(server, target, nil)
	} else {
		messagePacket, _ = nex.NewPRUDPPacketV1(server, target, nil)
	}

	messagePacket.SetType(constants.DataPacket)
	messagePacket.AddFlag(constants.PacketFlagNeedsAck)
	messagePacket.AddFlag(constants.PacketFlagReliable)
	messagePacket.SetSourceVirtualPortStreamType(target.StreamType)
	messagePacket.SetSourceVirtualPortStreamID(endpoint.StreamID)
	messagePacket.SetDestinationVirtualPortStreamType(target.StreamType)
	messagePacket.SetDestinationVirtualPortStreamID(target.StreamID)
	messagePacket.SetPayload(messageRequestBytes)

	server.Send(messagePacket)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, nil)
	rmcResponse.ProtocolID = message_delivery.ProtocolID
	rmcResponse.MethodID = message_delivery.MethodDeliverMessage
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
