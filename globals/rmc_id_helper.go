package globals

import (
	globals_rmc "github.com/PretendoNetwork/pokemon-gen6/globals/rmc"
)

func GetProtocolByID(protocolId uint16) globals_rmc.ProtocolInfo {
	switch protocolId {
	case 3:
		return globals_rmc.NewNATTraversal()
	case 10:
		return globals_rmc.NewTicketGranting()
	case 11:
		return globals_rmc.NewSecureConnection()
	case 21:
		return globals_rmc.NewMatchMaking()
	case 27:
		return globals_rmc.NewMessageDelivery()
	case 50:
		return globals_rmc.NewMatchMakingExtension()
	case 109:
		return globals_rmc.NewMatchmakeExtension()
	case 112:
		if SecureEndpoint.LibraryVersions().Main.Major <= 1 {
			return globals_rmc.NewRankingLegacyOld()
		} else if SecureEndpoint.LibraryVersions().Main.Major == 2 {
			return globals_rmc.NewRankingLegacy()
		}

		return globals_rmc.NewRanking()
	case 117:
		return globals_rmc.NewSubscription()
	case 118:
		return globals_rmc.NewRating()
	}

	return globals_rmc.NewProtocolInfo()
}
