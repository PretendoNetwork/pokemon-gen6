// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemEvent is a type within the ServiceItem protocol
type ServiceItemEvent struct {
	types.Structure
	EventID           *types.PrimitiveU64
	ParamInt          *types.PrimitiveS32
	ParamString       *types.String
	ParamBinary       *types.QBuffer
	PresentTicketType *types.PrimitiveU32
	PresentTicketNum  *types.PrimitiveU32
	TimeBegin         *types.DateTime
	TimeEnd           *types.DateTime
}

// WriteTo writes the ServiceItemEvent to the given writable
func (sie *ServiceItemEvent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sie.EventID.WriteTo(contentWritable)
	sie.ParamInt.WriteTo(contentWritable)
	sie.ParamString.WriteTo(contentWritable)
	sie.ParamBinary.WriteTo(contentWritable)
	sie.PresentTicketType.WriteTo(contentWritable)
	sie.PresentTicketNum.WriteTo(contentWritable)
	sie.TimeBegin.WriteTo(contentWritable)
	sie.TimeEnd.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sie.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemEvent from the given readable
func (sie *ServiceItemEvent) ExtractFrom(readable types.Readable) error {
	var err error

	err = sie.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent header. %s", err.Error())
	}

	err = sie.EventID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.EventID. %s", err.Error())
	}

	err = sie.ParamInt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamInt. %s", err.Error())
	}

	err = sie.ParamString.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamString. %s", err.Error())
	}

	err = sie.ParamBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamBinary. %s", err.Error())
	}

	err = sie.PresentTicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketType. %s", err.Error())
	}

	err = sie.PresentTicketNum.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketNum. %s", err.Error())
	}

	err = sie.TimeBegin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeBegin. %s", err.Error())
	}

	err = sie.TimeEnd.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeEnd. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemEvent
func (sie *ServiceItemEvent) Copy() types.RVType {
	copied := NewServiceItemEvent()

	copied.StructureVersion = sie.StructureVersion
	copied.EventID = sie.EventID.Copy().(*types.PrimitiveU64)
	copied.ParamInt = sie.ParamInt.Copy().(*types.PrimitiveS32)
	copied.ParamString = sie.ParamString.Copy().(*types.String)
	copied.ParamBinary = sie.ParamBinary.Copy().(*types.QBuffer)
	copied.PresentTicketType = sie.PresentTicketType.Copy().(*types.PrimitiveU32)
	copied.PresentTicketNum = sie.PresentTicketNum.Copy().(*types.PrimitiveU32)
	copied.TimeBegin = sie.TimeBegin.Copy().(*types.DateTime)
	copied.TimeEnd = sie.TimeEnd.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given ServiceItemEvent contains the same data as the current ServiceItemEvent
func (sie *ServiceItemEvent) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemEvent); !ok {
		return false
	}

	other := o.(*ServiceItemEvent)

	if sie.StructureVersion != other.StructureVersion {
		return false
	}

	if !sie.EventID.Equals(other.EventID) {
		return false
	}

	if !sie.ParamInt.Equals(other.ParamInt) {
		return false
	}

	if !sie.ParamString.Equals(other.ParamString) {
		return false
	}

	if !sie.ParamBinary.Equals(other.ParamBinary) {
		return false
	}

	if !sie.PresentTicketType.Equals(other.PresentTicketType) {
		return false
	}

	if !sie.PresentTicketNum.Equals(other.PresentTicketNum) {
		return false
	}

	if !sie.TimeBegin.Equals(other.TimeBegin) {
		return false
	}

	return sie.TimeEnd.Equals(other.TimeEnd)
}

// String returns the string representation of the ServiceItemEvent
func (sie *ServiceItemEvent) String() string {
	return sie.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemEvent using the provided indentation level
func (sie *ServiceItemEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEvent{\n")
	b.WriteString(fmt.Sprintf("%sEventID: %s,\n", indentationValues, sie.EventID))
	b.WriteString(fmt.Sprintf("%sParamInt: %s,\n", indentationValues, sie.ParamInt))
	b.WriteString(fmt.Sprintf("%sParamString: %s,\n", indentationValues, sie.ParamString))
	b.WriteString(fmt.Sprintf("%sParamBinary: %s,\n", indentationValues, sie.ParamBinary))
	b.WriteString(fmt.Sprintf("%sPresentTicketType: %s,\n", indentationValues, sie.PresentTicketType))
	b.WriteString(fmt.Sprintf("%sPresentTicketNum: %s,\n", indentationValues, sie.PresentTicketNum))
	b.WriteString(fmt.Sprintf("%sTimeBegin: %s,\n", indentationValues, sie.TimeBegin.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTimeEnd: %s,\n", indentationValues, sie.TimeEnd.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEvent returns a new ServiceItemEvent
func NewServiceItemEvent() *ServiceItemEvent {
	sie := &ServiceItemEvent{
		EventID:           types.NewPrimitiveU64(0),
		ParamInt:          types.NewPrimitiveS32(0),
		ParamString:       types.NewString(""),
		ParamBinary:       types.NewQBuffer(nil),
		PresentTicketType: types.NewPrimitiveU32(0),
		PresentTicketNum:  types.NewPrimitiveU32(0),
		TimeBegin:         types.NewDateTime(0),
		TimeEnd:           types.NewDateTime(0),
	}

	return sie
}
