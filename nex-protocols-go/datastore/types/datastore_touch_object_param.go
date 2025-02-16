// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreTouchObjectParam is a type within the DataStore protocol
type DataStoreTouchObjectParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	LockID         *types.PrimitiveU32
	AccessPassword *types.PrimitiveU64
}

// WriteTo writes the DataStoreTouchObjectParam to the given writable
func (dstop *DataStoreTouchObjectParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dstop.DataID.WriteTo(contentWritable)
	dstop.LockID.WriteTo(contentWritable)
	dstop.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dstop.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreTouchObjectParam from the given readable
func (dstop *DataStoreTouchObjectParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dstop.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam header. %s", err.Error())
	}

	err = dstop.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.DataID. %s", err.Error())
	}

	err = dstop.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.LockID. %s", err.Error())
	}

	err = dstop.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreTouchObjectParam
func (dstop *DataStoreTouchObjectParam) Copy() types.RVType {
	copied := NewDataStoreTouchObjectParam()

	copied.StructureVersion = dstop.StructureVersion
	copied.DataID = dstop.DataID.Copy().(*types.PrimitiveU64)
	copied.LockID = dstop.LockID.Copy().(*types.PrimitiveU32)
	copied.AccessPassword = dstop.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreTouchObjectParam contains the same data as the current DataStoreTouchObjectParam
func (dstop *DataStoreTouchObjectParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreTouchObjectParam); !ok {
		return false
	}

	other := o.(*DataStoreTouchObjectParam)

	if dstop.StructureVersion != other.StructureVersion {
		return false
	}

	if !dstop.DataID.Equals(other.DataID) {
		return false
	}

	if !dstop.LockID.Equals(other.LockID) {
		return false
	}

	return dstop.AccessPassword.Equals(other.AccessPassword)
}

// String returns the string representation of the DataStoreTouchObjectParam
func (dstop *DataStoreTouchObjectParam) String() string {
	return dstop.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreTouchObjectParam using the provided indentation level
func (dstop *DataStoreTouchObjectParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreTouchObjectParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dstop.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dstop.LockID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dstop.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() *DataStoreTouchObjectParam {
	dstop := &DataStoreTouchObjectParam{
		DataID:         types.NewPrimitiveU64(0),
		LockID:         types.NewPrimitiveU32(0),
		AccessPassword: types.NewPrimitiveU64(0),
	}

	return dstop
}
