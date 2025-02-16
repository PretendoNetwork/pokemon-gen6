package ticket_granting

import (
	"crypto/rand"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
)

func generateTicket(source, target *nex.Account, sessionKeyLength int, endpoint *nex.PRUDPEndPoint) ([]byte, *nex.Error) {
	if source == nil || target == nil {
		return []byte{}, nex.NewError(nex.ResultCodes.Authentication.Unknown, "Source or target account is nil")
	}

	sourceKey := nex.DeriveKerberosKey(source.PID, []byte(source.Password))
	targetKey := nex.DeriveKerberosKey(target.PID, []byte(target.Password))
	sessionKey := make([]byte, sessionKeyLength)

	_, err := rand.Read(sessionKey)
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return []byte{}, nex.NewError(nex.ResultCodes.Authentication.Unknown, "Failed to generate session key")
	}

	ticketInternalData := nex.NewKerberosTicketInternalData(endpoint.Server)
	serverTime := types.NewDateTime(0).Now()

	ticketInternalData.Issued = serverTime
	ticketInternalData.SourcePID = source.PID
	ticketInternalData.SessionKey = sessionKey

	encryptedTicketInternalData, err := ticketInternalData.Encrypt(targetKey, nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings()))
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return []byte{}, nex.NewError(nex.ResultCodes.Authentication.Unknown, "Failed to encrypt Ticket Internal Data")
	}

	ticket := nex.NewKerberosTicket()
	ticket.SessionKey = sessionKey
	ticket.TargetPID = target.PID
	ticket.InternalData = types.NewBuffer(encryptedTicketInternalData)

	encryptedTicket, err := ticket.Encrypt(sourceKey, nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings()))
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return []byte{}, nex.NewError(nex.ResultCodes.Authentication.Unknown, "Failed to encrypt Ticket")
	}

	return encryptedTicket, nil
}
