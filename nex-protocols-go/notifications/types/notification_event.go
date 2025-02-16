// Package types implements all the types used by the Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NotificationEvent is a type within the Notifications protocol
type NotificationEvent struct {
	types.Structure
	PIDSource *types.PID
	Type      *types.PrimitiveU32
	Param1    *types.PrimitiveU32
	Param2    *types.PrimitiveU32
	StrParam  *types.String
	Param3    *types.PrimitiveU32
}

// WriteTo writes the NotificationEvent to the given writable
func (ne *NotificationEvent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ne.PIDSource.WriteTo(contentWritable)
	ne.Type.WriteTo(contentWritable)
	ne.Param1.WriteTo(contentWritable)
	ne.Param2.WriteTo(contentWritable)
	ne.StrParam.WriteTo(contentWritable)
	//ne.Param3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ne.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NotificationEvent from the given readable
func (ne *NotificationEvent) ExtractFrom(readable types.Readable) error {
	var err error

	err = ne.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent header. %s", err.Error())
	}

	err = ne.PIDSource.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.PIDSource. %s", err.Error())
	}

	err = ne.Type.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.Type. %s", err.Error())
	}

	err = ne.Param1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.Param1. %s", err.Error())
	}

	err = ne.Param2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.Param2. %s", err.Error())
	}

	err = ne.StrParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.StrParam. %s", err.Error())
	}

	err = ne.Param3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.Param3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NotificationEvent
func (ne *NotificationEvent) Copy() types.RVType {
	copied := NewNotificationEvent()

	copied.StructureVersion = ne.StructureVersion
	copied.PIDSource = ne.PIDSource.Copy().(*types.PID)
	copied.Type = ne.Type.Copy().(*types.PrimitiveU32)
	copied.Param1 = ne.Param1.Copy().(*types.PrimitiveU32)
	copied.Param2 = ne.Param2.Copy().(*types.PrimitiveU32)
	copied.StrParam = ne.StrParam.Copy().(*types.String)
	copied.Param3 = ne.Param3.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given NotificationEvent contains the same data as the current NotificationEvent
func (ne *NotificationEvent) Equals(o types.RVType) bool {
	if _, ok := o.(*NotificationEvent); !ok {
		return false
	}

	other := o.(*NotificationEvent)

	if ne.StructureVersion != other.StructureVersion {
		return false
	}

	if !ne.PIDSource.Equals(other.PIDSource) {
		return false
	}

	if !ne.Type.Equals(other.Type) {
		return false
	}

	if !ne.Param1.Equals(other.Param1) {
		return false
	}

	if !ne.Param2.Equals(other.Param2) {
		return false
	}

	if !ne.StrParam.Equals(other.StrParam) {
		return false
	}

	return ne.Param3.Equals(other.Param3)
}

// String returns the string representation of the NotificationEvent
func (ne *NotificationEvent) String() string {
	return ne.FormatToString(0)
}

// FormatToString pretty-prints the NotificationEvent using the provided indentation level
func (ne *NotificationEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NotificationEvent{\n")
	b.WriteString(fmt.Sprintf("%sPIDSource: %s,\n", indentationValues, ne.PIDSource.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sType: %s,\n", indentationValues, ne.Type))
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, ne.Param1))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, ne.Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %s,\n", indentationValues, ne.StrParam))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, ne.Param3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	ne := &NotificationEvent{
		PIDSource: types.NewPID(0),
		Type:      types.NewPrimitiveU32(0),
		Param1:    types.NewPrimitiveU32(0),
		Param2:    types.NewPrimitiveU32(0),
		StrParam:  types.NewString(""),
		Param3:    types.NewPrimitiveU32(0),
	}

	return ne
}
