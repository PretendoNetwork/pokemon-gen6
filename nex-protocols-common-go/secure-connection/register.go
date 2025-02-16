package secureconnection

import (
	"net"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"

	"fmt"
)

func (commonProtocol *CommonProtocol) register(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint()

	// * vecMyURLs may contain multiple StationURLs. Search them all
	var localStation *types.StationURL
	var publicStation *types.StationURL

	for _, stationURL := range vecMyURLs.Slice() {
		natf, ok := stationURL.NATFiltering()
		if !ok {
			continue
		}

		natm, ok := stationURL.NATMapping()
		if !ok {
			continue
		}

		// * Station reports itself as being non-public (local)
		if localStation == nil && !stationURL.IsPublic() {
			localStation = stationURL.Copy().(*types.StationURL)
		}

		// * Still did not find the station, trying heuristics
		if localStation == nil && natf == constants.UnknownNATFiltering && natm == constants.UnknownNATMapping {
			localStation = stationURL.Copy().(*types.StationURL)
		}

		if publicStation == nil && stationURL.IsPublic() {
			publicStation = stationURL.Copy().(*types.StationURL)
		}
	}

	if localStation == nil {
		common_globals.Logger.Error("Failed to find local station")
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	if publicStation == nil {
		publicStation = localStation.Copy().(*types.StationURL)

		var address string
		var port uint16

		// * We have to duplicate this because Go automatically breaks on switch statements
		switch clientAddress := connection.Address().(type) {
		case *net.UDPAddr:
			address = clientAddress.IP.String()
			port = uint16(clientAddress.Port)
		case *net.TCPAddr:
			address = clientAddress.IP.String()
			port = uint16(clientAddress.Port)
		}

		publicStation.SetAddress(address)
		publicStation.SetPortNumber(port)
		publicStation.SetNATFiltering(constants.UnknownNATFiltering)
		publicStation.SetNATMapping(constants.UnknownNATMapping)
		publicStation.SetType(uint8(constants.StationURLFlagPublic) | uint8(constants.StationURLFlagBehindNAT))
	}

	/*localStation.SetPrincipalID(connection.PID())
	publicStation.SetPrincipalID(connection.PID())

	localStation.SetRVConnectionID(connection.ID)
	publicStation.SetRVConnectionID(connection.ID)*/

	connection.StationURLs.Append(localStation)
	connection.StationURLs.Append(publicStation)

	retval := types.NewQResultSuccess(nex.ResultCodes.Core.Unknown)
	pidConnectionID := types.NewPrimitiveU32(connection.ID)
	urlPublic := types.NewString(publicStation.EncodeToString())
	fmt.Println(publicStation)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	retval.WriteTo(rmcResponseStream)
	pidConnectionID.WriteTo(rmcResponseStream)
	urlPublic.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = secure_connection.ProtocolID
	rmcResponse.MethodID = secure_connection.MethodRegister
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterRegister != nil {
		go commonProtocol.OnAfterRegister(packet, vecMyURLs)
	}

	return rmcResponse, nil
}
