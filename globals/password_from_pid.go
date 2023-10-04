package globals

import (
	"context"

	pb "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid uint32) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCAccountCommonMetadata)

	response, err := GRPCAccountClient.GetNEXPassword(ctx, &pb.GetNEXPasswordRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return "", nex.Errors.RendezVous.InvalidUsername
	}

	return response.Password, 0
}
