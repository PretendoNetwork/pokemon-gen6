// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemNotice is a type within the ServiceItem protocol
type ServiceItemNotice struct {
	types.Structure
	ScheduleID   *types.PrimitiveU64
	ScheduleType *types.PrimitiveU32
	ParamInt     *types.PrimitiveS32
	ParamString  *types.String
	ParamBinary  *types.QBuffer
	TimeBegin    *types.DateTime
	TimeEnd      *types.DateTime
}

// WriteTo writes the ServiceItemNotice to the given writable
func (sin *ServiceItemNotice) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sin.ScheduleID.WriteTo(contentWritable)
	sin.ScheduleType.WriteTo(contentWritable)
	sin.ParamInt.WriteTo(contentWritable)
	sin.ParamString.WriteTo(contentWritable)
	sin.ParamBinary.WriteTo(contentWritable)
	sin.TimeBegin.WriteTo(contentWritable)
	sin.TimeEnd.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sin.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemNotice from the given readable
func (sin *ServiceItemNotice) ExtractFrom(readable types.Readable) error {
	var err error

	err = sin.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice header. %s", err.Error())
	}

	err = sin.ScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleID. %s", err.Error())
	}

	err = sin.ScheduleType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleType. %s", err.Error())
	}

	err = sin.ParamInt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamInt. %s", err.Error())
	}

	err = sin.ParamString.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamString. %s", err.Error())
	}

	err = sin.ParamBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamBinary. %s", err.Error())
	}

	err = sin.TimeBegin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeBegin. %s", err.Error())
	}

	err = sin.TimeEnd.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeEnd. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemNotice
func (sin *ServiceItemNotice) Copy() types.RVType {
	copied := NewServiceItemNotice()

	copied.StructureVersion = sin.StructureVersion
	copied.ScheduleID = sin.ScheduleID.Copy().(*types.PrimitiveU64)
	copied.ScheduleType = sin.ScheduleType.Copy().(*types.PrimitiveU32)
	copied.ParamInt = sin.ParamInt.Copy().(*types.PrimitiveS32)
	copied.ParamString = sin.ParamString.Copy().(*types.String)
	copied.ParamBinary = sin.ParamBinary.Copy().(*types.QBuffer)
	copied.TimeBegin = sin.TimeBegin.Copy().(*types.DateTime)
	copied.TimeEnd = sin.TimeEnd.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given ServiceItemNotice contains the same data as the current ServiceItemNotice
func (sin *ServiceItemNotice) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemNotice); !ok {
		return false
	}

	other := o.(*ServiceItemNotice)

	if sin.StructureVersion != other.StructureVersion {
		return false
	}

	if !sin.ScheduleID.Equals(other.ScheduleID) {
		return false
	}

	if !sin.ScheduleType.Equals(other.ScheduleType) {
		return false
	}

	if !sin.ParamInt.Equals(other.ParamInt) {
		return false
	}

	if !sin.ParamString.Equals(other.ParamString) {
		return false
	}

	if !sin.ParamBinary.Equals(other.ParamBinary) {
		return false
	}

	if !sin.TimeBegin.Equals(other.TimeBegin) {
		return false
	}

	return sin.TimeEnd.Equals(other.TimeEnd)
}

// String returns the string representation of the ServiceItemNotice
func (sin *ServiceItemNotice) String() string {
	return sin.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemNotice using the provided indentation level
func (sin *ServiceItemNotice) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemNotice{\n")
	b.WriteString(fmt.Sprintf("%sScheduleID: %s,\n", indentationValues, sin.ScheduleID))
	b.WriteString(fmt.Sprintf("%sScheduleType: %s,\n", indentationValues, sin.ScheduleType))
	b.WriteString(fmt.Sprintf("%sParamInt: %s,\n", indentationValues, sin.ParamInt))
	b.WriteString(fmt.Sprintf("%sParamString: %s,\n", indentationValues, sin.ParamString))
	b.WriteString(fmt.Sprintf("%sParamBinary: %s,\n", indentationValues, sin.ParamBinary))
	b.WriteString(fmt.Sprintf("%sTimeBegin: %s,\n", indentationValues, sin.TimeBegin.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTimeEnd: %s,\n", indentationValues, sin.TimeEnd.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemNotice returns a new ServiceItemNotice
func NewServiceItemNotice() *ServiceItemNotice {
	sin := &ServiceItemNotice{
		ScheduleID:   types.NewPrimitiveU64(0),
		ScheduleType: types.NewPrimitiveU32(0),
		ParamInt:     types.NewPrimitiveS32(0),
		ParamString:  types.NewString(""),
		ParamBinary:  types.NewQBuffer(nil),
		TimeBegin:    types.NewDateTime(0),
		TimeEnd:      types.NewDateTime(0),
	}

	return sin
}
