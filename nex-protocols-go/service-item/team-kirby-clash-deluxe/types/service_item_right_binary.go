// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightBinary is a type within the ServiceItem protocol
type ServiceItemRightBinary struct {
	types.Structure
	UseType     *types.PrimitiveU8
	RightBinary *types.QBuffer
}

// WriteTo writes the ServiceItemRightBinary to the given writable
func (sirb *ServiceItemRightBinary) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sirb.UseType.WriteTo(contentWritable)
	sirb.RightBinary.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sirb.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightBinary from the given readable
func (sirb *ServiceItemRightBinary) ExtractFrom(readable types.Readable) error {
	var err error

	err = sirb.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary header. %s", err.Error())
	}

	err = sirb.UseType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.UseType. %s", err.Error())
	}

	err = sirb.RightBinary.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightBinary.RightBinary. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightBinary
func (sirb *ServiceItemRightBinary) Copy() types.RVType {
	copied := NewServiceItemRightBinary()

	copied.StructureVersion = sirb.StructureVersion
	copied.UseType = sirb.UseType.Copy().(*types.PrimitiveU8)
	copied.RightBinary = sirb.RightBinary.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given ServiceItemRightBinary contains the same data as the current ServiceItemRightBinary
func (sirb *ServiceItemRightBinary) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightBinary); !ok {
		return false
	}

	other := o.(*ServiceItemRightBinary)

	if sirb.StructureVersion != other.StructureVersion {
		return false
	}

	if !sirb.UseType.Equals(other.UseType) {
		return false
	}

	return sirb.RightBinary.Equals(other.RightBinary)
}

// String returns the string representation of the ServiceItemRightBinary
func (sirb *ServiceItemRightBinary) String() string {
	return sirb.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightBinary using the provided indentation level
func (sirb *ServiceItemRightBinary) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightBinary{\n")
	b.WriteString(fmt.Sprintf("%sUseType: %s,\n", indentationValues, sirb.UseType))
	b.WriteString(fmt.Sprintf("%sRightBinary: %s,\n", indentationValues, sirb.RightBinary))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightBinary returns a new ServiceItemRightBinary
func NewServiceItemRightBinary() *ServiceItemRightBinary {
	sirb := &ServiceItemRightBinary{
		UseType:     types.NewPrimitiveU8(0),
		RightBinary: types.NewQBuffer(nil),
	}

	return sirb
}
