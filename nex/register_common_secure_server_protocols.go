package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
	matchmaking "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking-ext"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/matchmake-extension"
	nattraversal "github.com/PretendoNetwork/nex-protocols-common-go/nat-traversal"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
	matchmaking.NewCommonMatchMakingProtocol(globals.SecureServer)
	matchmakingext.NewCommonMatchMakingExtProtocol(globals.SecureServer)
	matchmakeextension.NewCommonMatchmakeExtensionProtocol(globals.SecureServer)
	nattraversal.NewCommonNATTraversalProtocol(globals.SecureServer)
}
