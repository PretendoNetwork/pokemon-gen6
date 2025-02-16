// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePrepareGetParam is a type within the DataStore protocol
type DataStorePrepareGetParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	LockID            *types.PrimitiveU32
	PersistenceTarget *DataStorePersistenceTarget
	AccessPassword    *types.PrimitiveU64
	ExtraData         *types.List[*types.String] // * NEX v3.5.0
}

// WriteTo writes the DataStorePrepareGetParam to the given writable
func (dspgp *DataStorePrepareGetParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.DataStore

	contentWritable := writable.CopyNew()

	dspgp.DataID.WriteTo(contentWritable)
	dspgp.LockID.WriteTo(contentWritable)
	dspgp.PersistenceTarget.WriteTo(contentWritable)
	dspgp.AccessPassword.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.5.0") {
		dspgp.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dspgp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePrepareGetParam from the given readable
func (dspgp *DataStorePrepareGetParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.DataStore

	var err error

	err = dspgp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam header. %s", err.Error())
	}

	err = dspgp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.DataID. %s", err.Error())
	}

	err = dspgp.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.LockID. %s", err.Error())
	}

	err = dspgp.PersistenceTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.PersistenceTarget. %s", err.Error())
	}

	err = dspgp.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.AccessPassword. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = dspgp.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareGetParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetParam
func (dspgp *DataStorePrepareGetParam) Copy() types.RVType {
	copied := NewDataStorePrepareGetParam()

	copied.StructureVersion = dspgp.StructureVersion
	copied.DataID = dspgp.DataID.Copy().(*types.PrimitiveU64)
	copied.LockID = dspgp.LockID.Copy().(*types.PrimitiveU32)
	copied.PersistenceTarget = dspgp.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.AccessPassword = dspgp.AccessPassword.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dspgp.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the given DataStorePrepareGetParam contains the same data as the current DataStorePrepareGetParam
func (dspgp *DataStorePrepareGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetParam)

	if dspgp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspgp.DataID.Equals(other.DataID) {
		return false
	}

	if !dspgp.LockID.Equals(other.LockID) {
		return false
	}

	if !dspgp.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if !dspgp.AccessPassword.Equals(other.AccessPassword) {
		return false
	}

	return dspgp.ExtraData.Equals(other.ExtraData)
}

// String returns the string representation of the DataStorePrepareGetParam
func (dspgp *DataStorePrepareGetParam) String() string {
	return dspgp.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePrepareGetParam using the provided indentation level
func (dspgp *DataStorePrepareGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dspgp.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dspgp.LockID))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dspgp.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dspgp.AccessPassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, dspgp.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParam returns a new DataStorePrepareGetParam
func NewDataStorePrepareGetParam() *DataStorePrepareGetParam {
	dspgp := &DataStorePrepareGetParam{
		DataID:            types.NewPrimitiveU64(0),
		LockID:            types.NewPrimitiveU32(0),
		PersistenceTarget: NewDataStorePersistenceTarget(),
		AccessPassword:    types.NewPrimitiveU64(0),
		ExtraData:         types.NewList[*types.String](),
	}

	dspgp.ExtraData.Type = types.NewString("")

	return dspgp
}
