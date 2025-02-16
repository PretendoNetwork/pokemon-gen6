// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemLawMessage is a type within the ServiceItem protocol
type ServiceItemLawMessage struct {
	types.Structure
	IsMessageRequired *types.PrimitiveBool
	LawMessage        *types.String
}

// WriteTo writes the ServiceItemLawMessage to the given writable
func (silm *ServiceItemLawMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	silm.IsMessageRequired.WriteTo(contentWritable)
	silm.LawMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	silm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemLawMessage from the given readable
func (silm *ServiceItemLawMessage) ExtractFrom(readable types.Readable) error {
	var err error

	err = silm.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage header. %s", err.Error())
	}

	err = silm.IsMessageRequired.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.IsMessageRequired. %s", err.Error())
	}

	err = silm.LawMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.LawMessage. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemLawMessage
func (silm *ServiceItemLawMessage) Copy() types.RVType {
	copied := NewServiceItemLawMessage()

	copied.StructureVersion = silm.StructureVersion
	copied.IsMessageRequired = silm.IsMessageRequired.Copy().(*types.PrimitiveBool)
	copied.LawMessage = silm.LawMessage.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemLawMessage contains the same data as the current ServiceItemLawMessage
func (silm *ServiceItemLawMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemLawMessage); !ok {
		return false
	}

	other := o.(*ServiceItemLawMessage)

	if silm.StructureVersion != other.StructureVersion {
		return false
	}

	if !silm.IsMessageRequired.Equals(other.IsMessageRequired) {
		return false
	}

	return silm.LawMessage.Equals(other.LawMessage)
}

// String returns the string representation of the ServiceItemLawMessage
func (silm *ServiceItemLawMessage) String() string {
	return silm.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemLawMessage using the provided indentation level
func (silm *ServiceItemLawMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemLawMessage{\n")
	b.WriteString(fmt.Sprintf("%sIsMessageRequired: %s,\n", indentationValues, silm.IsMessageRequired))
	b.WriteString(fmt.Sprintf("%sLawMessage: %s,\n", indentationValues, silm.LawMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLawMessage returns a new ServiceItemLawMessage
func NewServiceItemLawMessage() *ServiceItemLawMessage {
	silm := &ServiceItemLawMessage{
		IsMessageRequired: types.NewPrimitiveBool(false),
		LawMessage:        types.NewString(""),
	}

	return silm
}
