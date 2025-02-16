// Package types implements all the types used by the MessageDelivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TextMessage is a type within the MessageDelivery protocol
type TextMessage struct {
	types.Structure
	*UserMessage
	StrTextBody *types.String
}

// WriteTo writes the TextMessage to the given writable
func (tm *TextMessage) WriteTo(writable types.Writable) {
	tm.UserMessage.WriteTo(writable)

	contentWritable := writable.CopyNew()

	tm.StrTextBody.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	tm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the TextMessage from the given readable
func (tm *TextMessage) ExtractFrom(readable types.Readable) error {
	var err error

	err = tm.UserMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TextMessage.UserMessage. %s", err.Error())
	}

	err = tm.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TextMessage header. %s", err.Error())
	}

	err = tm.StrTextBody.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract TextMessage.StrTextBody. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of TextMessage
func (tm *TextMessage) Copy() types.RVType {
	copied := NewTextMessage()

	copied.StructureVersion = tm.StructureVersion
	copied.UserMessage = tm.UserMessage.Copy().(*UserMessage)
	copied.StrTextBody = tm.StrTextBody.Copy().(*types.String)

	return copied
}

// Equals checks if the given TextMessage contains the same data as the current TextMessage
func (tm *TextMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*TextMessage); !ok {
		return false
	}

	other := o.(*TextMessage)

	if tm.StructureVersion != other.StructureVersion {
		return false
	}

	if !tm.UserMessage.Equals(other.UserMessage) {
		return false
	}

	return tm.StrTextBody.Equals(other.StrTextBody)
}

// String returns the string representation of the TextMessage
func (tm *TextMessage) String() string {
	return tm.FormatToString(0)
}

// FormatToString pretty-prints the TextMessage using the provided indentation level
func (tm *TextMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("TextMessage{\n")
	b.WriteString(fmt.Sprintf("%sUserMessage (parent): %s,\n", indentationValues, tm.UserMessage.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrTextBody: %s,\n", indentationValues, tm.StrTextBody))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewTextMessage returns a new TextMessage
func NewTextMessage() *TextMessage {
	tm := &TextMessage{
		UserMessage: NewUserMessage(),
		StrTextBody: types.NewString(""),
	}

	return tm
}
