// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAmount is a type within the ServiceItem protocol
type ServiceItemAmount struct {
	types.Structure
	FormattedAmount *types.String
	Currency        *types.String
	RawValue        *types.String
}

// WriteTo writes the ServiceItemAmount to the given writable
func (sia *ServiceItemAmount) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sia.FormattedAmount.WriteTo(contentWritable)
	sia.Currency.WriteTo(contentWritable)
	sia.RawValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sia.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAmount from the given readable
func (sia *ServiceItemAmount) ExtractFrom(readable types.Readable) error {
	var err error

	err = sia.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount header. %s", err.Error())
	}

	err = sia.FormattedAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.FormattedAmount. %s", err.Error())
	}

	err = sia.Currency.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.Currency. %s", err.Error())
	}

	err = sia.RawValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.RawValue. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAmount
func (sia *ServiceItemAmount) Copy() types.RVType {
	copied := NewServiceItemAmount()

	copied.StructureVersion = sia.StructureVersion
	copied.FormattedAmount = sia.FormattedAmount.Copy().(*types.String)
	copied.Currency = sia.Currency.Copy().(*types.String)
	copied.RawValue = sia.RawValue.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemAmount contains the same data as the current ServiceItemAmount
func (sia *ServiceItemAmount) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAmount); !ok {
		return false
	}

	other := o.(*ServiceItemAmount)

	if sia.StructureVersion != other.StructureVersion {
		return false
	}

	if !sia.FormattedAmount.Equals(other.FormattedAmount) {
		return false
	}

	if !sia.Currency.Equals(other.Currency) {
		return false
	}

	return sia.RawValue.Equals(other.RawValue)
}

// String returns the string representation of the ServiceItemAmount
func (sia *ServiceItemAmount) String() string {
	return sia.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAmount using the provided indentation level
func (sia *ServiceItemAmount) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAmount{\n")
	b.WriteString(fmt.Sprintf("%sFormattedAmount: %s,\n", indentationValues, sia.FormattedAmount))
	b.WriteString(fmt.Sprintf("%sCurrency: %s,\n", indentationValues, sia.Currency))
	b.WriteString(fmt.Sprintf("%sRawValue: %s,\n", indentationValues, sia.RawValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAmount returns a new ServiceItemAmount
func NewServiceItemAmount() *ServiceItemAmount {
	sia := &ServiceItemAmount{
		FormattedAmount: types.NewString(""),
		Currency:        types.NewString(""),
		RawValue:        types.NewString(""),
	}

	return sia
}
