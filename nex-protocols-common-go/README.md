# NEX Protocols Common Go
## NEX protocols used by many games with premade handlers and a high level API

[![GoDoc](https://godoc.org/github.com/PretendoNetwork/nex-protocols-common-go/v2?status.svg)](https://godoc.org/github.com/PretendoNetwork/nex-protocols-common-go/v2)

### Other NEX libraries
[nex-go](https://github.com/PretendoNetwork/nex-go) - Barebones NEX/PRUDP server implementation

[nex-protocols-go](https://github.com/PretendoNetwork/nex-protocols-go) - NEX protocol definitions

### Install

`go get github.com/PretendoNetwork/nex-protocols-common-go/v2`

### Usage

`nex-protocols-common-go` provides a higher level API than the [NEX Protocols Go module](https://github.com/PretendoNetwork/nex-protocols-go). This module handles many of the more common protcols and methods used shared by many servers. Instead of working directly with the NEX server, this module exposes an API for defining helper functions to provide the module with the data it needs to run

### Example, friends (Wii U) authentication server
### For a complete example, see the complete [Friends Server](https://github.com/PretendoNetwork/friends), and other game servers

```go
package main

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/nex-go/v2"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
	common_ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/v2/ticket-granting"
)

var nexServer *nex.PRUDPServer

func main() {
	nexServer := nex.NewPRUDPServer()

	endpoint := nex.NewPRUDPEndPoint(1)
	endpoint.ServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", "password"))
	endpoint.AccountDetailsByPID = accountDetailsByPID
	endpoint.AccountDetailsByUsername = accountDetailsByUsername

	endpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("==Friends - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("==================")
	})

	nexServer.SetFragmentSize(962)
	nexServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(1, 1, 0))
	nexServer.SessionKeyLength = 16
	nexServer.AccessKey = "ridfebb9"

	ticketGrantingProtocol := ticket_granting.NewProtocol(endpoint)
	endpoint.RegisterServiceProtocol(ticketGrantingProtocol)
	commonTicketGrantingProtocol := common_ticket_granting.NewCommonProtocol(ticketGrantingProtocol)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.Scheme = "prudps"
	secureStationURL.Fields.Set("address", os.Getenv("SECURE_SERVER_LOCATION"))
	secureStationURL.Fields.Set("port", os.Getenv("SECURE_SERVER_PORT"))
	secureStationURL.Fields.Set("CID", "1")
	secureStationURL.Fields.Set("PID", "2")
	secureStationURL.Fields.Set("sid", "1")
	secureStationURL.Fields.Set("stream", "10")
	secureStationURL.Fields.Set("type", "2")

	commonTicketGrantingProtocol.SecureStationURL = secureStationURL
	commonTicketGrantingProtocol.BuildName = "Pretendo Friends Auth"

	nexServer.Listen(60000)
}
```
