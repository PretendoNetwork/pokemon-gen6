// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStorePreparePostParamV1 is a type within the DataStore protocol
type DataStorePreparePostParamV1 struct {
	types.Structure
	Size             *types.PrimitiveU32
	Name             *types.String
	DataType         *types.PrimitiveU16
	MetaBinary       *types.QBuffer
	Permission       *DataStorePermission
	DelPermission    *DataStorePermission
	Flag             *types.PrimitiveU32
	Period           *types.PrimitiveU16
	ReferDataID      *types.PrimitiveU32
	Tags             *types.List[*types.String]
	RatingInitParams *types.List[*DataStoreRatingInitParamWithSlot]
}

// WriteTo writes the DataStorePreparePostParamV1 to the given writable
func (dspppv *DataStorePreparePostParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dspppv.Size.WriteTo(contentWritable)
	dspppv.Name.WriteTo(contentWritable)
	dspppv.DataType.WriteTo(contentWritable)
	dspppv.MetaBinary.WriteTo(contentWritable)
	dspppv.Permission.WriteTo(contentWritable)
	dspppv.DelPermission.WriteTo(contentWritable)
	dspppv.Flag.WriteTo(contentWritable)
	dspppv.Period.WriteTo(contentWritable)
	dspppv.ReferDataID.WriteTo(contentWritable)
	dspppv.Tags.WriteTo(contentWritable)
	dspppv.RatingInitParams.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dspppv.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePreparePostParamV1 from the given readable
func (dspppv *DataStorePreparePostParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	err = dspppv.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1 header. %s", err.Error())
	}

	err = dspppv.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Size. %s", err.Error())
	}

	err = dspppv.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Name. %s", err.Error())
	}

	err = dspppv.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DataType. %s", err.Error())
	}

	err = dspppv.MetaBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.MetaBinary. %s", err.Error())
	}

	err = dspppv.Permission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Permission. %s", err.Error())
	}

	err = dspppv.DelPermission.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.DelPermission. %s", err.Error())
	}

	err = dspppv.Flag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Flag. %s", err.Error())
	}

	err = dspppv.Period.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Period. %s", err.Error())
	}

	err = dspppv.ReferDataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.ReferDataID. %s", err.Error())
	}

	err = dspppv.Tags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.Tags. %s", err.Error())
	}

	err = dspppv.RatingInitParams.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParamV1.RatingInitParams. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostParamV1
func (dspppv *DataStorePreparePostParamV1) Copy() types.RVType {
	copied := NewDataStorePreparePostParamV1()

	copied.StructureVersion = dspppv.StructureVersion
	copied.Size = dspppv.Size.Copy().(*types.PrimitiveU32)
	copied.Name = dspppv.Name.Copy().(*types.String)
	copied.DataType = dspppv.DataType.Copy().(*types.PrimitiveU16)
	copied.MetaBinary = dspppv.MetaBinary.Copy().(*types.QBuffer)
	copied.Permission = dspppv.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dspppv.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dspppv.Flag.Copy().(*types.PrimitiveU32)
	copied.Period = dspppv.Period.Copy().(*types.PrimitiveU16)
	copied.ReferDataID = dspppv.ReferDataID.Copy().(*types.PrimitiveU32)
	copied.Tags = dspppv.Tags.Copy().(*types.List[*types.String])
	copied.RatingInitParams = dspppv.RatingInitParams.Copy().(*types.List[*DataStoreRatingInitParamWithSlot])

	return copied
}

// Equals checks if the given DataStorePreparePostParamV1 contains the same data as the current DataStorePreparePostParamV1
func (dspppv *DataStorePreparePostParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePreparePostParamV1); !ok {
		return false
	}

	other := o.(*DataStorePreparePostParamV1)

	if dspppv.StructureVersion != other.StructureVersion {
		return false
	}

	if !dspppv.Size.Equals(other.Size) {
		return false
	}

	if !dspppv.Name.Equals(other.Name) {
		return false
	}

	if !dspppv.DataType.Equals(other.DataType) {
		return false
	}

	if !dspppv.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dspppv.Permission.Equals(other.Permission) {
		return false
	}

	if !dspppv.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dspppv.Flag.Equals(other.Flag) {
		return false
	}

	if !dspppv.Period.Equals(other.Period) {
		return false
	}

	if !dspppv.ReferDataID.Equals(other.ReferDataID) {
		return false
	}

	if !dspppv.Tags.Equals(other.Tags) {
		return false
	}

	return dspppv.RatingInitParams.Equals(other.RatingInitParams)
}

// String returns the string representation of the DataStorePreparePostParamV1
func (dspppv *DataStorePreparePostParamV1) String() string {
	return dspppv.FormatToString(0)
}

// FormatToString pretty-prints the DataStorePreparePostParamV1 using the provided indentation level
func (dspppv *DataStorePreparePostParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParamV1{\n")
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dspppv.Size))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, dspppv.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %s,\n", indentationValues, dspppv.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %s,\n", indentationValues, dspppv.MetaBinary))
	b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dspppv.Permission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dspppv.DelPermission.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sFlag: %s,\n", indentationValues, dspppv.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %s,\n", indentationValues, dspppv.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %s,\n", indentationValues, dspppv.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %s,\n", indentationValues, dspppv.Tags))
	b.WriteString(fmt.Sprintf("%sRatingInitParams: %s,\n", indentationValues, dspppv.RatingInitParams))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostParamV1 returns a new DataStorePreparePostParamV1
func NewDataStorePreparePostParamV1() *DataStorePreparePostParamV1 {
	dspppv := &DataStorePreparePostParamV1{
		Size:             types.NewPrimitiveU32(0),
		Name:             types.NewString(""),
		DataType:         types.NewPrimitiveU16(0),
		MetaBinary:       types.NewQBuffer(nil),
		Permission:       NewDataStorePermission(),
		DelPermission:    NewDataStorePermission(),
		Flag:             types.NewPrimitiveU32(0),
		Period:           types.NewPrimitiveU16(0),
		ReferDataID:      types.NewPrimitiveU32(0),
		Tags:             types.NewList[*types.String](),
		RatingInitParams: types.NewList[*DataStoreRatingInitParamWithSlot](),
	}

	dspppv.Tags.Type = types.NewString("")
	dspppv.RatingInitParams.Type = NewDataStoreRatingInitParamWithSlot()

	return dspppv
}
