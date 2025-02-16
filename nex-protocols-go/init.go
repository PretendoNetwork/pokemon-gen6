// Package nexproto provides all the main NEX protocols.
//
// Each folder contains a different package for that specific protocol,
// with all their types and methods needed to parse and build packets with RMC payloads
package nexproto

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/account-management/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	ticket_granting_types "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting/types"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

func init() {
	types.RegisterDataHolderType("NintendoCreateAccountData", account_management_types.NewNintendoCreateAccountData())
	types.RegisterDataHolderType("AccountExtraInfo", account_management_types.NewAccountExtraInfo())
	types.RegisterDataHolderType("NintendoLoginData", ticket_granting_types.NewNintendoLoginData())
	types.RegisterDataHolderType("AuthenticationInfo", ticket_granting_types.NewAuthenticationInfo())
	types.RegisterDataHolderType("Gathering", match_making_types.NewGathering())
	types.RegisterDataHolderType("MatchmakeSession", match_making_types.NewMatchmakeSession())
	types.RegisterDataHolderType("TextMessage", messaging_types.NewTextMessage())
}
