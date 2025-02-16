// Package types implements all the types used by the NintendoNotifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoNotificationEvent is a type within the NintendoNotifications protocol
type NintendoNotificationEvent struct {
	types.Structure
	Type       *types.PrimitiveU32
	SenderPID  *types.PID
	DataHolder *types.AnyDataHolder
}

// WriteTo writes the NintendoNotificationEvent to the given writable
func (nne *NintendoNotificationEvent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nne.Type.WriteTo(contentWritable)
	nne.SenderPID.WriteTo(contentWritable)
	nne.DataHolder.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nne.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoNotificationEvent from the given readable
func (nne *NintendoNotificationEvent) ExtractFrom(readable types.Readable) error {
	var err error

	err = nne.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEvent header. %s", err.Error())
	}

	err = nne.Type.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEvent.Type. %s", err.Error())
	}

	err = nne.SenderPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEvent.SenderPID. %s", err.Error())
	}

	err = nne.DataHolder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEvent.DataHolder. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoNotificationEvent
func (nne *NintendoNotificationEvent) Copy() types.RVType {
	copied := NewNintendoNotificationEvent()

	copied.StructureVersion = nne.StructureVersion
	copied.Type = nne.Type.Copy().(*types.PrimitiveU32)
	copied.SenderPID = nne.SenderPID.Copy().(*types.PID)
	copied.DataHolder = nne.DataHolder.Copy().(*types.AnyDataHolder)

	return copied
}

// Equals checks if the given NintendoNotificationEvent contains the same data as the current NintendoNotificationEvent
func (nne *NintendoNotificationEvent) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoNotificationEvent); !ok {
		return false
	}

	other := o.(*NintendoNotificationEvent)

	if nne.StructureVersion != other.StructureVersion {
		return false
	}

	if !nne.Type.Equals(other.Type) {
		return false
	}

	if !nne.SenderPID.Equals(other.SenderPID) {
		return false
	}

	return nne.DataHolder.Equals(other.DataHolder)
}

// String returns the string representation of the NintendoNotificationEvent
func (nne *NintendoNotificationEvent) String() string {
	return nne.FormatToString(0)
}

// FormatToString pretty-prints the NintendoNotificationEvent using the provided indentation level
func (nne *NintendoNotificationEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEvent{\n")
	b.WriteString(fmt.Sprintf("%sType: %s,\n", indentationValues, nne.Type))
	b.WriteString(fmt.Sprintf("%sSenderPID: %s,\n", indentationValues, nne.SenderPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sDataHolder: %s,\n", indentationValues, nne.DataHolder.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEvent returns a new NintendoNotificationEvent
func NewNintendoNotificationEvent() *NintendoNotificationEvent {
	nne := &NintendoNotificationEvent{
		Type:       types.NewPrimitiveU32(0),
		SenderPID:  types.NewPID(0),
		DataHolder: types.NewAnyDataHolder(),
	}

	return nne
}
