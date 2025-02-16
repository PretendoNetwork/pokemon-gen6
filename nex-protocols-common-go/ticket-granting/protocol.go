package ticket_granting

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	_ "github.com/PretendoNetwork/nex-protocols-go/v2"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
)

type CommonProtocol struct {
	protocol                   ticket_granting.Interface
	SecureStationURL           *types.StationURL
	SpecialProtocols           []*types.PrimitiveU8
	StationURLSpecialProtocols *types.StationURL
	BuildName                  *types.String
	allowInsecureLoginMethod   bool
	SessionKeyLength           int // TODO - Use server SessionKeyLength?
	SecureServerAccount        *nex.Account
	OnAfterLogin               func(packet nex.PacketInterface, strUserName *types.String)
	OnAfterLoginEx             func(packet nex.PacketInterface, strUserName *types.String, oExtraData *types.AnyDataHolder)
	OnAfterRequestTicket       func(packet nex.PacketInterface, idSource *types.PID, idTarget *types.PID)
}

func (commonProtocol *CommonProtocol) DisableInsecureLogin() {
	commonProtocol.allowInsecureLoginMethod = false
}

func (commonProtocol *CommonProtocol) EnableInsecureLogin() {
	common_globals.Logger.Warning("INSECURE LOGIN HAS BEEN ENABLED. THIS ALLOWS THE USE OF CUSTOM CLIENTS TO BYPASS THE ACCOUNT SERVER AND CONNECT DIRECTLY TO THIS GAME SERVER, EVADING BANS! USE WITH CAUTION!")
	commonProtocol.allowInsecureLoginMethod = true
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol ticket_granting.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		protocol:                   protocol,
		SecureStationURL:           types.NewStationURL("prudp:/"),
		SpecialProtocols:           make([]*types.PrimitiveU8, 0),
		StationURLSpecialProtocols: types.NewStationURL(""),
		BuildName:                  types.NewString(""),
		allowInsecureLoginMethod:   false,
		SessionKeyLength:           32,
	}

	protocol.SetHandlerLogin(commonProtocol.login)
	protocol.SetHandlerLoginEx(commonProtocol.loginEx)
	protocol.SetHandlerRequestTicket(commonProtocol.requestTicket)

	commonProtocol.DisableInsecureLogin() // * Disable insecure login by default

	return commonProtocol
}
