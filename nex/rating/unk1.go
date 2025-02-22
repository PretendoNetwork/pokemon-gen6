package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	rating "github.com/PretendoNetwork/nex-protocols-go/v2/rating"
)

func Unk1(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	types.UInt32(2259334670).WriteTo(rmcResponseStream)
	types.UInt32(0).WriteTo(rmcResponseStream)
	types.NewString("e00d70431ddca6ae").WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = rating.ProtocolID
	rmcResponse.MethodID = rating.MethodUnk1
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

