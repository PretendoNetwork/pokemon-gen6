// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAttribute is a type within the ServiceItem protocol
type ServiceItemAttribute struct {
	types.Structure
	Name  *types.String
	Value *types.String
}

// WriteTo writes the ServiceItemAttribute to the given writable
func (sia *ServiceItemAttribute) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sia.Name.WriteTo(contentWritable)
	sia.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sia.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAttribute from the given readable
func (sia *ServiceItemAttribute) ExtractFrom(readable types.Readable) error {
	var err error

	err = sia.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute header. %s", err.Error())
	}

	err = sia.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Name. %s", err.Error())
	}

	err = sia.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAttribute.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAttribute
func (sia *ServiceItemAttribute) Copy() types.RVType {
	copied := NewServiceItemAttribute()

	copied.StructureVersion = sia.StructureVersion
	copied.Name = sia.Name.Copy().(*types.String)
	copied.Value = sia.Value.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemAttribute contains the same data as the current ServiceItemAttribute
func (sia *ServiceItemAttribute) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAttribute); !ok {
		return false
	}

	other := o.(*ServiceItemAttribute)

	if sia.StructureVersion != other.StructureVersion {
		return false
	}

	if !sia.Name.Equals(other.Name) {
		return false
	}

	return sia.Value.Equals(other.Value)
}

// String returns the string representation of the ServiceItemAttribute
func (sia *ServiceItemAttribute) String() string {
	return sia.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAttribute using the provided indentation level
func (sia *ServiceItemAttribute) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAttribute{\n")
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, sia.Name))
	b.WriteString(fmt.Sprintf("%sValue: %s,\n", indentationValues, sia.Value))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAttribute returns a new ServiceItemAttribute
func NewServiceItemAttribute() *ServiceItemAttribute {
	sia := &ServiceItemAttribute{
		Name:  types.NewString(""),
		Value: types.NewString(""),
	}

	return sia
}
