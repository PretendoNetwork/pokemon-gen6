// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemLimitation is a type within the ServiceItem protocol
type ServiceItemLimitation struct {
	types.Structure
	LimitationType  *types.PrimitiveU32
	LimitationValue *types.PrimitiveU32
}

// WriteTo writes the ServiceItemLimitation to the given writable
func (sil *ServiceItemLimitation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sil.LimitationType.WriteTo(contentWritable)
	sil.LimitationValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sil.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemLimitation from the given readable
func (sil *ServiceItemLimitation) ExtractFrom(readable types.Readable) error {
	var err error

	err = sil.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation header. %s", err.Error())
	}

	err = sil.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationType. %s", err.Error())
	}

	err = sil.LimitationValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLimitation.LimitationValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemLimitation
func (sil *ServiceItemLimitation) Copy() types.RVType {
	copied := NewServiceItemLimitation()

	copied.StructureVersion = sil.StructureVersion
	copied.LimitationType = sil.LimitationType.Copy().(*types.PrimitiveU32)
	copied.LimitationValue = sil.LimitationValue.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given ServiceItemLimitation contains the same data as the current ServiceItemLimitation
func (sil *ServiceItemLimitation) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemLimitation); !ok {
		return false
	}

	other := o.(*ServiceItemLimitation)

	if sil.StructureVersion != other.StructureVersion {
		return false
	}

	if !sil.LimitationType.Equals(other.LimitationType) {
		return false
	}

	return sil.LimitationValue.Equals(other.LimitationValue)
}

// String returns the string representation of the ServiceItemLimitation
func (sil *ServiceItemLimitation) String() string {
	return sil.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemLimitation using the provided indentation level
func (sil *ServiceItemLimitation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemLimitation{\n")
	b.WriteString(fmt.Sprintf("%sLimitationType: %s,\n", indentationValues, sil.LimitationType))
	b.WriteString(fmt.Sprintf("%sLimitationValue: %s,\n", indentationValues, sil.LimitationValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLimitation returns a new ServiceItemLimitation
func NewServiceItemLimitation() *ServiceItemLimitation {
	sil := &ServiceItemLimitation{
		LimitationType:  types.NewPrimitiveU32(0),
		LimitationValue: types.NewPrimitiveU32(0),
	}

	return sil
}
