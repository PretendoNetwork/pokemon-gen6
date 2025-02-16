package secureconnection

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
)

func (commonProtocol *CommonProtocol) replaceURL(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint()

	connection.StationURLs.Each(func(i int, station *types.StationURL) bool {
		currentStationAddress, currentStationAddressOk := station.Address()
		currentStationPort, currentStationPortOk := station.PortNumber()
		oldStationAddress, oldStationAddressOk := target.Address()
		oldStationPort, oldStationPortOk := target.PortNumber()

		if currentStationAddressOk && currentStationPortOk && oldStationAddressOk && oldStationPortOk {
			if currentStationAddress == oldStationAddress && currentStationPort == oldStationPort {
				// * This fixes Minecraft, but is obviously incorrect
				// TODO - What are we really meant to do here?
				newStation := url.Copy().(*types.StationURL)

				newStation.SetPrincipalID(connection.PID())

				connection.StationURLs.SetIndex(i, newStation)
			}
		}

		return false
	})

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = secure_connection.ProtocolID
	rmcResponse.MethodID = secure_connection.MethodReplaceURL
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterReplaceURL != nil {
		go commonProtocol.OnAfterReplaceURL(packet, target, url)
	}

	return rmcResponse, nil
}
