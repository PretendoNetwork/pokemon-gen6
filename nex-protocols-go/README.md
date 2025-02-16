# NEX Protocols Go
## NEX servers with protocol support in Go

[![GoDoc](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go/v2?status.svg)](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go/v2)

### Other NEX libraries
[nex-go](https://github.com/PretendoNetwork/nex-go/v2) - Barebones NEX/PRUDP server implementation

[nex-protocols-common-go](https://github.com/PretendoNetwork/nex-protocols-common-go) - NEX protocols used by many games with premade handlers and a high level API

### Install

`go get github.com/PretendoNetwork/nex-protocols-go/v2`

### Usage

`nex-protocols-go` provides a higher level API than the [NEX Go module](https://github.com/PretendoNetwork/nex-go/v2) to the underlying PRUDP server by providing a set of NEX protocols. This module only provides access to the lower level raw RMC method calls, however, and all method handlers must be defined in full manually. For a higher level API, see the [common NEX method handlers module](https://github.com/PretendoNetwork/nex-protocols-common-go)

### Example, friends (Wii U) authentication server
### For a complete example, see the complete [Friends Server](https://github.com/PretendoNetwork/friends), and other game servers

```go
package main

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
)

var nexServer *nex.PRUDPServer

func main() {
	nexServer := nex.NewPRUDPServer()

	endpoint := nex.NewPRUDPEndPoint(1)
	endpoint.ServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", "password"))
	endpoint.AccountDetailsByPID = accountDetailsByPID
	endpoint.AccountDetailsByUsername = accountDetailsByUsername

	nexServer.BindPRUDPEndPoint(endpoint)
	nexServer.SetFragmentSize(962)
	nexServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(1, 1, 0))
	nexServer.SessionKeyLength = 16
	nexServer.AccessKey = "ridfebb9"

	ticketGrantingProtocol := ticket_granting.NewProtocol(endpoint)

	// Handle Login RMC method
	ticketGrantingProtocol.Login = login

	// Handle RequestTicket RMC method
	ticketGrantingProtocol.RequestTicket = requestTicket

	// Register the protocol on the endpoint
	endpoint.RegisterServiceProtocol(ticketGrantingProtocol)

	nexServer.Listen(60000)
}
```
