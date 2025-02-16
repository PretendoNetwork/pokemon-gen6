// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreRateObjectParam is a type within the DataStore protocol
type DataStoreRateObjectParam struct {
	types.Structure
	RatingValue    *types.PrimitiveS32
	AccessPassword *types.PrimitiveU64
}

// WriteTo writes the DataStoreRateObjectParam to the given writable
func (dsrop *DataStoreRateObjectParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrop.RatingValue.WriteTo(contentWritable)
	dsrop.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrop.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreRateObjectParam from the given readable
func (dsrop *DataStoreRateObjectParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrop.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam header. %s", err.Error())
	}

	err = dsrop.RatingValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.RatingValue. %s", err.Error())
	}

	err = dsrop.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateObjectParam
func (dsrop *DataStoreRateObjectParam) Copy() types.RVType {
	copied := NewDataStoreRateObjectParam()

	copied.StructureVersion = dsrop.StructureVersion
	copied.RatingValue = dsrop.RatingValue.Copy().(*types.PrimitiveS32)
	copied.AccessPassword = dsrop.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given DataStoreRateObjectParam contains the same data as the current DataStoreRateObjectParam
func (dsrop *DataStoreRateObjectParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreRateObjectParam); !ok {
		return false
	}

	other := o.(*DataStoreRateObjectParam)

	if dsrop.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrop.RatingValue.Equals(other.RatingValue) {
		return false
	}

	return dsrop.AccessPassword.Equals(other.AccessPassword)
}

// String returns the string representation of the DataStoreRateObjectParam
func (dsrop *DataStoreRateObjectParam) String() string {
	return dsrop.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreRateObjectParam using the provided indentation level
func (dsrop *DataStoreRateObjectParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreRateObjectParam{\n")
	b.WriteString(fmt.Sprintf("%sRatingValue: %s,\n", indentationValues, dsrop.RatingValue))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dsrop.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreRateObjectParam returns a new DataStoreRateObjectParam
func NewDataStoreRateObjectParam() *DataStoreRateObjectParam {
	dsrop := &DataStoreRateObjectParam{
		RatingValue:    types.NewPrimitiveS32(0),
		AccessPassword: types.NewPrimitiveU64(0),
	}

	return dsrop
}
