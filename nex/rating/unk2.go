package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	rating "github.com/PretendoNetwork/nex-protocols-go/v2/rating"
)

func Unk2(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = rating.ProtocolID
	rmcResponse.MethodID = rating.MethodUnk2
	rmcResponse.CallID = callID

	return rmcResponse, nil
}

