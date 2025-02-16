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

func DeliverMessage(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server
	
	/*recipientPid := fmt.Sprintf("%.8x",(oUserMessage.(*nexproto.TextMessage).UserMessage.M_messageRecipient.M_principalID))
	recipientPidString := recipientPid[6:8] + recipientPid[4:6] + recipientPid[2:4] + recipientPid[0:2]
	senderPid := fmt.Sprintf("%.8x",(client.PID()))
	senderPidString := senderPid[6:8] + senderPid[4:6] + senderPid[2:4] + senderPid[0:2]*/

	oUserMessage.ObjectData.(*messaging_types.TextMessage).UserMessage.StrSender = types.NewString(strconv.Itoa(int(client.PID().Value())))
	oUserMessage.ObjectData.(*messaging_types.TextMessage).UserMessage.PIDSender = client.PID()
	//hexData, _ := hex.DecodeString("0C00546578744D65737361676500560000005200000000000000"+recipientPidString+"0100000000000000"+senderPidString+"D01952961F00000000000000010000001400537562736372697074696F6E4D616E61676572000B00313837363834313330350009005570646174696E6700")

	messageStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())
	oUserMessage.WriteTo(messageStream)

	messageRequest := nex.NewRMCRequest(endpoint)
	messageRequest.ProtocolID = message_delivery.ProtocolID
	messageRequest.CallID = 0xffffffff
	messageRequest.MethodID = message_delivery.MethodDeliverMessage
	messageRequest.Parameters = messageStream.Bytes()

	messageRequestBytes := messageRequest.Bytes()

	fmt.Println(hex.EncodeToString(messageRequest.Parameters))

	//fmt.Println(uint32(oUserMessage.(*nexproto.TextMessage).UserMessage.M_messageRecipient.M_principalID))
	fmt.Println(uint32(oUserMessage.ObjectData.(*messaging_types.TextMessage).UserMessage.UIIDRecipient.Value))
	target := endpoint.FindConnectionByPID(uint64(oUserMessage.ObjectData.(*messaging_types.TextMessage).UserMessage.UIIDRecipient.Value))

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
