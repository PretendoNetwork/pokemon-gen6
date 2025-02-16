package nattraversal

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	"fmt"
)

func (commonProtocol *CommonProtocol) requestProbeInitiationExt(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)
	server := endpoint.Server

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = nat_traversal.ProtocolID
	rmcResponse.MethodID = nat_traversal.MethodRequestProbeInitiationExt
	rmcResponse.CallID = callID

	rmcRequestStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	//lol nice hack thanks pokemon
	stationToProbeStation := types.NewStationURL(stationToProbe.Value)
	_, exists := stationToProbeStation.StreamID()
	if(!exists){
		stationToProbeStation.SetNATFiltering(0)
		stationToProbeStation.SetNATMapping(0)
		stationToProbeStation.SetNATPMPSupport(false)
		stationToProbeStation.SetStreamID(15)
		stationToProbeStation.SetUPnPSupport(false)
		//stationToProbe = types.NewString(stationToProbeStation.EncodeToString())
	}
	stationToProbe.WriteTo(rmcRequestStream)
	fmt.Println("stationToProbe")
	fmt.Println(stationToProbe)
	fmt.Println(connection.PID())

	rmcRequestBody := rmcRequestStream.Bytes()

	rmcRequest := nex.NewRMCRequest(endpoint)
	rmcRequest.ProtocolID = nat_traversal.ProtocolID
	rmcRequest.CallID = 0xFFFF0000 + callID
	rmcRequest.MethodID = nat_traversal.MethodInitiateProbe
	rmcRequest.Parameters = rmcRequestBody

	rmcRequestBytes := rmcRequest.Bytes()

	for _, target := range targetList.Slice() {
		targetStation := types.NewStationURL(target.Value)
		fmt.Println("target.Value")
		fmt.Println(target.Value)

		if connectionID, ok := targetStation.RVConnectionID(); ok {
			target := endpoint.FindConnectionByID(connectionID)
			if target == nil {
				common_globals.Logger.Warning("Client not found")
				continue
			}

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
			messagePacket.SetPayload(rmcRequestBytes)

			server.Send(messagePacket)
		}
	}

	if commonProtocol.OnAfterRequestProbeInitiationExt != nil {
		go commonProtocol.OnAfterRequestProbeInitiationExt(packet, targetList, stationToProbe)
	}

	return rmcResponse, nil
}
