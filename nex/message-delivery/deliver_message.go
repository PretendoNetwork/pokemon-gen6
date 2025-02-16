package nex_message_delivery

import (
	"encoding/hex"
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	message_delivery "github.com/PretendoNetwork/nex-protocols-go/v2/message-delivery"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
	"fmt"
	"strconv"
)

func DeliverMessage(err error, packet nex.PacketInterface, callID uint32, oUserMessage types.DataHolder) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server

	oUserMessage.Object.(*messaging_types.TextMessage).UserMessage.StrSender = types.NewString(strconv.Itoa(int(client.PID())))
	oUserMessage.Object.(*messaging_types.TextMessage).UserMessage.PIDSender = client.PID()

	messageStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())
	oUserMessage.WriteTo(messageStream)

	messageRequest := nex.NewRMCRequest(endpoint)
	messageRequest.ProtocolID = message_delivery.ProtocolID
	messageRequest.CallID = 0xffffffff
	messageRequest.MethodID = message_delivery.MethodDeliverMessage
	messageRequest.Parameters = messageStream.Bytes()

	messageRequestBytes := messageRequest.Bytes()

	fmt.Println(hex.EncodeToString(messageRequest.Parameters))

	fmt.Println(uint32(oUserMessage.Object.(*messaging_types.TextMessage).UserMessage.IDRecipient))
	target := endpoint.FindConnectionByPID(uint64(oUserMessage.Object.(*messaging_types.TextMessage).UserMessage.IDRecipient))

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

