package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(3)
	globals.SecureServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 3,
		Patch: 0,
	})
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("876138df")

	globals.Timeline = make(map[uint32][]uint8)

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Pokemon X/Y/OR/AS (Gen 6) - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_POKEGEN6_SECURE_SERVER_PORT")))
}
