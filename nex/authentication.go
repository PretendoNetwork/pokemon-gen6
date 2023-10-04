package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(3)
	globals.AuthenticationServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 3,
		Patch: 0,
	})
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("876138df")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Pokemon X/Y/OR/AS (Gen 6) - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_POKEGEN6_AUTHENTICATION_SERVER_PORT")))
}
