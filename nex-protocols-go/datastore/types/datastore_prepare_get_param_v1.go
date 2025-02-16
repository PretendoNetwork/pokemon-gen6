// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePrepareGetParamV1 is a type within the DataStore protocol
type DataStorePrepareGetParamV1 struct {
	types.Structure
	DataID *types.PrimitiveU32
	LockID *types.PrimitiveU32
}

// WriteTo writes the DataStorePrepareGetParamV1 to the given writable
func (dspgpv *DataStorePrepareGetParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspgpv.DataID.WriteTo(contentWritable)
	dspgpv.LockID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspgpv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePrepareGetParamV1 from the given readable
func (dspgpv *DataStorePrepareGetParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspgpv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1 header. %s", err.Error())
	}

	err = dspgpv.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.DataID. %s", err.Error())
	}

	err = dspgpv.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.LockID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetParamV1
func (dspgpv *DataStorePrepareGetParamV1) Copy() types.RVType {
	copied := NewDataStorePrepareGetParamV1()

	copied.StructureVersion = dspgpv.StructureVersion
	copied.DataID = dspgpv.DataID.Copy().(*types.PrimitiveU32)
	copied.LockID = dspgpv.LockID.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given DataStorePrepareGetParamV1 contains the same data as the current DataStorePrepareGetParamV1
func (dspgpv *DataStorePrepareGetParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetParamV1); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetParamV1)

	if dspgpv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspgpv.DataID.Equals(other.DataID) {
		return false
	}

	return dspgpv.LockID.Equals(other.LockID)
}

// String returns the string representation of the DataStorePrepareGetParamV1
func (dspgpv *DataStorePrepareGetParamV1) String() string {
	return dspgpv.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePrepareGetParamV1 using the provided indentation level
func (dspgpv *DataStorePrepareGetParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetParamV1{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dspgpv.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dspgpv.LockID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParamV1 returns a new DataStorePrepareGetParamV1
func NewDataStorePrepareGetParamV1() *DataStorePrepareGetParamV1 {
	dspgpv := &DataStorePrepareGetParamV1{
		DataID: types.NewPrimitiveU32(0),
		LockID: types.NewPrimitiveU32(0),
	}

	return dspgpv
}
