package globals

import (
	pb "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/plogger-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var Logger *plogger.Logger
var KerberosPassword = "password" // * Default password
var AuthenticationServer *nex.Server
var SecureServer *nex.Server
var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pb.AccountClient
var GRPCAccountCommonMetadata metadata.MD
