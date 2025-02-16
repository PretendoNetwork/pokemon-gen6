// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePrepareGetReplayParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStorePrepareGetReplayParam struct {
	types.Structure
	ReplayID  *types.PrimitiveU64
	ExtraData *types.List[*types.String]
}

// WriteTo writes the DataStorePrepareGetReplayParam to the given writable
func (dspgrp *DataStorePrepareGetReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspgrp.ReplayID.WriteTo(contentWritable)
	dspgrp.ExtraData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspgrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePrepareGetReplayParam from the given readable
func (dspgrp *DataStorePrepareGetReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspgrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam header. %s", err.Error())
	}

	err = dspgrp.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ReplayID. %s", err.Error())
	}

	err = dspgrp.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetReplayParam.ExtraData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetReplayParam
func (dspgrp *DataStorePrepareGetReplayParam) Copy() types.RVType {
	copied := NewDataStorePrepareGetReplayParam()

	copied.StructureVersion = dspgrp.StructureVersion
	copied.ReplayID = dspgrp.ReplayID.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dspgrp.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePrepareGetReplayParam contains the same data as the current DataStorePrepareGetReplayParam
func (dspgrp *DataStorePrepareGetReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetReplayParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetReplayParam)

	if dspgrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspgrp.ReplayID.Equals(other.ReplayID) {
		return false
	}

	return dspgrp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePrepareGetReplayParam
func (dspgrp *DataStorePrepareGetReplayParam) String() string {
	return dspgrp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePrepareGetReplayParam using the provided indentation level
func (dspgrp *DataStorePrepareGetReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dspgrp.ReplayID))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dspgrp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetReplayParam returns a new DataStorePrepareGetReplayParam
func NewDataStorePrepareGetReplayParam() *DataStorePrepareGetReplayParam {
	dspgrp := &DataStorePrepareGetReplayParam{
		ReplayID:  types.NewPrimitiveU64(0),
		ExtraData: types.NewList[*types.String](),
	}

	dspgrp.ExtraData.Type = types.NewString("")

	return dspgrp
}
