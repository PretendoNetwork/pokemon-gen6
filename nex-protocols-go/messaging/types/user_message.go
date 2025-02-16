// Package types implements all the types used by the MessageDelivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserMessage is a type within the MessageDelivery protocol
type UserMessage struct {
	types.Structure
	*types.Data
	UIID             *types.PrimitiveU32
	UIIDRecipient    *types.PrimitiveU32
	UIRecipientType  *types.PrimitiveU32
	UIParentID       *types.PrimitiveU32
	PIDSender        *types.PID
	Receptiontime    *types.DateTime
	UILifeTime       *types.PrimitiveU32
	UIFlags          *types.PrimitiveU32
	StrSubject       *types.String
	StrSender        *types.String
}

// WriteTo writes the UserMessage to the given writable
func (um *UserMessage) WriteTo(writable types.Writable) {
	um.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	um.UIID.WriteTo(contentWritable)
	um.UIIDRecipient.WriteTo(contentWritable)
	um.UIRecipientType.WriteTo(contentWritable)
	um.UIParentID.WriteTo(contentWritable)
	um.PIDSender.WriteTo(contentWritable)
	um.Receptiontime.WriteTo(contentWritable)
	um.UILifeTime.WriteTo(contentWritable)
	um.UIFlags.WriteTo(contentWritable)
	um.StrSubject.WriteTo(contentWritable)
	um.StrSender.WriteTo(contentWritable)
	//um.MessageRecipient.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	um.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserMessage from the given readable
func (um *UserMessage) ExtractFrom(readable types.Readable) error {
	var err error

	err = um.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.Data. %s", err.Error())
	}

	err = um.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage header. %s", err.Error())
	}

	err = um.UIID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIID. %s", err.Error())
	}

	err = um.UIIDRecipient.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIID. %s", err.Error())
	}

	err = um.UIRecipientType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIID. %s", err.Error())
	}

	err = um.UIParentID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIParentID. %s", err.Error())
	}

	err = um.PIDSender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.PIDSender. %s", err.Error())
	}

	err = um.Receptiontime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.Receptiontime. %s", err.Error())
	}

	err = um.UILifeTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UILifeTime. %s", err.Error())
	}

	err = um.UIFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.UIFlags. %s", err.Error())
	}

	err = um.StrSubject.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSubject. %s", err.Error())
	}

	err = um.StrSender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.StrSender. %s", err.Error())
	}

	/*err = um.MessageRecipient.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.MessageRecipient. %s", err.Error())
	}*/

	return nil
}

// Copy returns a new copied instance of UserMessage
func (um *UserMessage) Copy() types.RVType {
	copied := NewUserMessage()

	copied.StructureVersion = um.StructureVersion
	copied.Data = um.Data.Copy().(*types.Data)
	copied.UIID = um.UIID.Copy().(*types.PrimitiveU32)
	copied.UIIDRecipient = um.UIID.Copy().(*types.PrimitiveU32)
	copied.UIRecipientType = um.UIID.Copy().(*types.PrimitiveU32)
	copied.UIParentID = um.UIParentID.Copy().(*types.PrimitiveU32)
	copied.PIDSender = um.PIDSender.Copy().(*types.PID)
	copied.Receptiontime = um.Receptiontime.Copy().(*types.DateTime)
	copied.UILifeTime = um.UILifeTime.Copy().(*types.PrimitiveU32)
	copied.UIFlags = um.UIFlags.Copy().(*types.PrimitiveU32)
	copied.StrSubject = um.StrSubject.Copy().(*types.String)
	copied.StrSender = um.StrSender.Copy().(*types.String)

	return copied
}

// Equals checks if the given UserMessage contains the same data as the current UserMessage
func (um *UserMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*UserMessage); !ok {
		return false
	}

	other := o.(*UserMessage)

	if um.StructureVersion != other.StructureVersion {
		return false
	}

	if !um.Data.Equals(other.Data) {
		return false
	}

	if !um.UIID.Equals(other.UIID) {
		return false
	}

	if !um.UIIDRecipient.Equals(other.UIIDRecipient) {
		return false
	}

	if !um.UIRecipientType.Equals(other.UIRecipientType) {
		return false
	}

	if !um.UIParentID.Equals(other.UIParentID) {
		return false
	}

	if !um.PIDSender.Equals(other.PIDSender) {
		return false
	}

	if !um.Receptiontime.Equals(other.Receptiontime) {
		return false
	}

	if !um.UILifeTime.Equals(other.UILifeTime) {
		return false
	}

	if !um.UIFlags.Equals(other.UIFlags) {
		return false
	}

	if !um.StrSubject.Equals(other.StrSubject) {
		return false
	}

	return um.StrSender.Equals(other.StrSender)
}

// String returns the string representation of the UserMessage
func (um *UserMessage) String() string {
	return um.FormatToString(0)
}

// FormatToString pretty-prints the UserMessage using the provided indentation level
func (um *UserMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserMessage{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, um.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUIID: %s,\n", indentationValues, um.UIID))
	b.WriteString(fmt.Sprintf("%sUIParentID: %s,\n", indentationValues, um.UIParentID))
	b.WriteString(fmt.Sprintf("%sUIIDRecipient: %s,\n", indentationValues, um.UIIDRecipient))
	b.WriteString(fmt.Sprintf("%sUIRecipientType: %s,\n", indentationValues, um.UIRecipientType))
	b.WriteString(fmt.Sprintf("%sPIDSender: %s,\n", indentationValues, um.PIDSender.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sReceptiontime: %s,\n", indentationValues, um.Receptiontime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUILifeTime: %s,\n", indentationValues, um.UILifeTime))
	b.WriteString(fmt.Sprintf("%sUIFlags: %s,\n", indentationValues, um.UIFlags))
	b.WriteString(fmt.Sprintf("%sStrSubject: %s,\n", indentationValues, um.StrSubject))
	b.WriteString(fmt.Sprintf("%sStrSender: %s,\n", indentationValues, um.StrSender))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUserMessage returns a new UserMessage
func NewUserMessage() *UserMessage {
	um := &UserMessage{
		Data:             types.NewData(),
		UIID:             types.NewPrimitiveU32(0),
		UIIDRecipient:    types.NewPrimitiveU32(0),
		UIRecipientType:  types.NewPrimitiveU32(0),
		UIParentID:       types.NewPrimitiveU32(0),
		PIDSender:        types.NewPID(0),
		Receptiontime:    types.NewDateTime(0),
		UILifeTime:       types.NewPrimitiveU32(0),
		UIFlags:          types.NewPrimitiveU32(0),
		StrSubject:       types.NewString(""),
		StrSender:        types.NewString(""),
	}

	return um
}
