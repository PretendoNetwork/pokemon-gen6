// Package types implements all the types used by the MessageDelivery protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// BinaryMessage is a type within the MessageDelivery protocol
type BinaryMessage struct {
	types.Structure
	*UserMessage
	BinaryBody *types.QBuffer
}

// WriteTo writes the BinaryMessage to the given writable
func (bm *BinaryMessage) WriteTo(writable types.Writable) {
	bm.UserMessage.WriteTo(writable)

	contentWritable := writable.CopyNew()

	bm.BinaryBody.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	bm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BinaryMessage from the given readable
func (bm *BinaryMessage) ExtractFrom(readable types.Readable) error {
	var err error

	err = bm.UserMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BinaryMessage.UserMessage. %s", err.Error())
	}

	err = bm.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BinaryMessage header. %s", err.Error())
	}

	err = bm.BinaryBody.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BinaryMessage.BinaryBody. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BinaryMessage
func (bm *BinaryMessage) Copy() types.RVType {
	copied := NewBinaryMessage()

	copied.StructureVersion = bm.StructureVersion
	copied.UserMessage = bm.UserMessage.Copy().(*UserMessage)
	copied.BinaryBody = bm.BinaryBody.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given BinaryMessage contains the same data as the current BinaryMessage
func (bm *BinaryMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*BinaryMessage); !ok {
		return false
	}

	other := o.(*BinaryMessage)

	if bm.StructureVersion != other.StructureVersion {
		return false
	}

	if !bm.UserMessage.Equals(other.UserMessage) {
		return false
	}

	return bm.BinaryBody.Equals(other.BinaryBody)
}

// String returns the string representation of the BinaryMessage
func (bm *BinaryMessage) String() string {
	return bm.FormatToString(0)
}

// FormatToString pretty-prints the BinaryMessage using the provided indentation level
func (bm *BinaryMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BinaryMessage{\n")
	b.WriteString(fmt.Sprintf("%sUserMessage (parent): %s,\n", indentationValues, bm.UserMessage.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sBinaryBody: %s,\n", indentationValues, bm.BinaryBody))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBinaryMessage returns a new BinaryMessage
func NewBinaryMessage() *BinaryMessage {
	bm := &BinaryMessage{
		UserMessage: NewUserMessage(),
		BinaryBody:  types.NewQBuffer(nil),
	}

	return bm
}
