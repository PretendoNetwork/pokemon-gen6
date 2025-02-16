// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-mario-maker/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetMetasWithCourseRecord(packet nex.PacketInterface) {
	if protocol.GetMetasWithCourseRecord == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::GetMetasWithCourseRecord not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	params := types.NewList[*datastore_super_mario_maker_types.DataStoreGetCourseRecordParam]()
	params.Type = datastore_super_mario_maker_types.NewDataStoreGetCourseRecordParam()
	metaParam := datastore_types.NewDataStoreGetMetaParam()

	var err error

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = metaParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMetasWithCourseRecord(fmt.Errorf("Failed to read metaParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMetasWithCourseRecord(nil, packet, callID, params, metaParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
