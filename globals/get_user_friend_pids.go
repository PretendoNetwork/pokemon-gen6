package globals

import (
	"context"

	pbfriends "github.com/PretendoNetwork/grpc-go/friends"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func GetUserFriendPIDs(pid uint32) []uint32 {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCFriendsCommonMetadata)

	response, err := GRPCFriendsClient.GetUserFriendPIDs(ctx, &pbfriends.GetUserFriendPIDsRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return make([]uint32, 0)
	}

	return response.Pids
}

func GetOnlineFriendPIDs(pid uint32) []uint64 {
	friendPids := GetUserFriendPIDs(pid)

	onlineFriendPids := make([]uint64, 0)

	for _, friendPid := range friendPids {
		SecureEndpoint.Connections.Each(func(key string, value *nex.PRUDPConnection) bool {
			if friendPid == uint32(value.PID()) {
				onlineFriendPids = append(onlineFriendPids, uint64(friendPid))
			}

			return false
		})
	}

	return onlineFriendPids
}
