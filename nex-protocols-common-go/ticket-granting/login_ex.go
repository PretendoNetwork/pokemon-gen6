package ticket_granting

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
)

func (commonProtocol *CommonProtocol) loginEx(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	// TODO - VALIDATE oExtraData!

	connection := packet.Sender().(*nex.PRUDPConnection)
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	sourceAccount, errorCode := endpoint.AccountDetailsByUsername(strUserName.Value)
	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.RendezVous.InvalidUsername {
		// * Some other error happened
		return nil, errorCode
	}

	targetAccount, errorCode := endpoint.AccountDetailsByUsername(commonProtocol.SecureServerAccount.Username)
	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.RendezVous.InvalidUsername {
		// * Some other error happened
		return nil, errorCode
	}

	encryptedTicket, errorCode := generateTicket(sourceAccount, targetAccount, commonProtocol.SessionKeyLength, endpoint)

	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.RendezVous.InvalidUsername {
		// * Some other error happened
		return nil, errorCode
	}

	var retval *types.QResult
	pidPrincipal := types.NewPID(0)
	pbufResponse := types.NewBuffer([]byte{})
	pConnectionData := types.NewRVConnectionData()
	strReturnMsg := types.NewString("")

	// * From the wiki:
	// *
	// * "If the username does not exist, the %retval% field is set to
	// * RendezVous::InvalidUsername and the other fields are left blank."
	if errorCode != nil && errorCode.ResultCode == nex.ResultCodes.RendezVous.InvalidUsername {
		retval = types.NewQResultError(errorCode.ResultCode)
	} else {
		retval = types.NewQResultSuccess(nex.ResultCodes.Core.Unknown)
		pidPrincipal = sourceAccount.PID
		pbufResponse = types.NewBuffer(encryptedTicket)
		strReturnMsg = commonProtocol.BuildName.Copy().(*types.String)

		specialProtocols := types.NewList[*types.PrimitiveU8]()

		specialProtocols.Type = types.NewPrimitiveU8(0)
		specialProtocols.SetFromData(commonProtocol.SpecialProtocols)

		pConnectionData.StationURL = commonProtocol.SecureStationURL
		pConnectionData.SpecialProtocols = specialProtocols
		pConnectionData.StationURLSpecialProtocols = commonProtocol.StationURLSpecialProtocols
		pConnectionData.Time = types.NewDateTime(0).Now()

		if endpoint.LibraryVersions().Main.GreaterOrEqual("v3.5.0") {
			pConnectionData.StructureVersion = 1
		}
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	retval.WriteTo(rmcResponseStream)
	pidPrincipal.WriteTo(rmcResponseStream)
	pbufResponse.WriteTo(rmcResponseStream)
	pConnectionData.WriteTo(rmcResponseStream)
	strReturnMsg.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = ticket_granting.ProtocolID
	rmcResponse.MethodID = ticket_granting.MethodLoginEx
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterLoginEx != nil {
		go commonProtocol.OnAfterLoginEx(packet, strUserName, oExtraData)
	}

	return rmcResponse, nil
}
