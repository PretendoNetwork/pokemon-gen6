// Package types implements all the types used by the MessageDelivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MessageRecipient is a type within the MessageDelivery protocol
type MessageRecipient struct {
	types.Structure
	UIRecipientType *types.PrimitiveU32
	PrincipalID     *types.PID
	GatheringID     *types.PrimitiveU32
}

// WriteTo writes the MessageRecipient to the given writable
func (mr *MessageRecipient) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mr.UIRecipientType.WriteTo(contentWritable)
	mr.PrincipalID.WriteTo(contentWritable)
	mr.GatheringID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MessageRecipient from the given readable
func (mr *MessageRecipient) ExtractFrom(readable types.Readable) error {
	var err error

	err = mr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient header. %s", err.Error())
	}

	err = mr.UIRecipientType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.UIRecipientType. %s", err.Error())
	}

	err = mr.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.PrincipalID. %s", err.Error())
	}

	err = mr.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.GatheringID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (mr *MessageRecipient) Copy() types.RVType {
	copied := NewMessageRecipient()

	copied.StructureVersion = mr.StructureVersion
	copied.UIRecipientType = mr.UIRecipientType.Copy().(*types.PrimitiveU32)
	copied.PrincipalID = mr.PrincipalID.Copy().(*types.PID)
	copied.GatheringID = mr.GatheringID.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given MessageRecipient contains the same data as the current MessageRecipient
func (mr *MessageRecipient) Equals(o types.RVType) bool {
	if _, ok := o.(*MessageRecipient); !ok {
		return false
	}

	other := o.(*MessageRecipient)

	if mr.StructureVersion != other.StructureVersion {
		return false
	}

	if !mr.UIRecipientType.Equals(other.UIRecipientType) {
		return false
	}

	if !mr.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	return mr.GatheringID.Equals(other.GatheringID)
}

// String returns the string representation of the MessageRecipient
func (mr *MessageRecipient) String() string {
	return mr.FormatToString(0)
}

// FormatToString pretty-prints the MessageRecipient using the provided indentation level
func (mr *MessageRecipient) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MessageRecipient{\n")
	b.WriteString(fmt.Sprintf("%sUIRecipientType: %s,\n", indentationValues, mr.UIRecipientType))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, mr.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %s,\n", indentationValues, mr.GatheringID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	mr := &MessageRecipient{
		UIRecipientType: types.NewPrimitiveU32(0),
		PrincipalID:     types.NewPID(0),
		GatheringID:     types.NewPrimitiveU32(0),
	}

	return mr
}
