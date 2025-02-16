// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreReportCourseParam is a type within the DataStore protocol
type DataStoreReportCourseParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	MiiName        *types.String
	ReportCategory *types.PrimitiveU8
	ReportReason   *types.String
}

// WriteTo writes the DataStoreReportCourseParam to the given writable
func (dsrcp *DataStoreReportCourseParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dsrcp.DataID.WriteTo(contentWritable)
	dsrcp.MiiName.WriteTo(contentWritable)
	dsrcp.ReportCategory.WriteTo(contentWritable)
	dsrcp.ReportReason.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dsrcp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreReportCourseParam from the given readable
func (dsrcp *DataStoreReportCourseParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dsrcp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam header. %s", err.Error())
	}

	err = dsrcp.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.DataID. %s", err.Error())
	}

	err = dsrcp.MiiName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.MiiName. %s", err.Error())
	}

	err = dsrcp.ReportCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportCategory. %s", err.Error())
	}

	err = dsrcp.ReportReason.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreReportCourseParam.ReportReason. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreReportCourseParam
func (dsrcp *DataStoreReportCourseParam) Copy() types.RVType {
	copied := NewDataStoreReportCourseParam()

	copied.StructureVersion = dsrcp.StructureVersion
	copied.DataID = dsrcp.DataID.Copy().(*types.PrimitiveU64)
	copied.MiiName = dsrcp.MiiName.Copy().(*types.String)
	copied.ReportCategory = dsrcp.ReportCategory.Copy().(*types.PrimitiveU8)
	copied.ReportReason = dsrcp.ReportReason.Copy().(*types.String)

	return copied
}

// Equals checks if the given DataStoreReportCourseParam contains the same data as the current DataStoreReportCourseParam
func (dsrcp *DataStoreReportCourseParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreReportCourseParam); !ok {
		return false
	}

	other := o.(*DataStoreReportCourseParam)

	if dsrcp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dsrcp.DataID.Equals(other.DataID) {
		return false
	}

	if !dsrcp.MiiName.Equals(other.MiiName) {
		return false
	}

	if !dsrcp.ReportCategory.Equals(other.ReportCategory) {
		return false
	}

	return dsrcp.ReportReason.Equals(other.ReportReason)
}

// String returns the string representation of the DataStoreReportCourseParam
func (dsrcp *DataStoreReportCourseParam) String() string {
	return dsrcp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreReportCourseParam using the provided indentation level
func (dsrcp *DataStoreReportCourseParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreReportCourseParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dsrcp.DataID))
	b.WriteString(fmt.Sprintf("%sMiiName: %s,\n", indentationValues, dsrcp.MiiName))
	b.WriteString(fmt.Sprintf("%sReportCategory: %s,\n", indentationValues, dsrcp.ReportCategory))
	b.WriteString(fmt.Sprintf("%sReportReason: %s,\n", indentationValues, dsrcp.ReportReason))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreReportCourseParam returns a new DataStoreReportCourseParam
func NewDataStoreReportCourseParam() *DataStoreReportCourseParam {
	dsrcp := &DataStoreReportCourseParam{
		DataID:         types.NewPrimitiveU64(0),
		MiiName:        types.NewString(""),
		ReportCategory: types.NewPrimitiveU8(0),
		ReportReason:   types.NewString(""),
	}

	return dsrcp
}
