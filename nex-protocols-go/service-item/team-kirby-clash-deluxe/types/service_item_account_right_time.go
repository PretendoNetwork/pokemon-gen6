// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAccountRightTime is a type within the ServiceItem protocol
type ServiceItemAccountRightTime struct {
	types.Structure
	*ServiceItemAccountRight
}

// WriteTo writes the ServiceItemAccountRightTime to the given writable
func (siart *ServiceItemAccountRightTime) WriteTo(writable types.Writable) {
	siart.ServiceItemAccountRight.WriteTo(writable)

	contentWritable := writable.CopyNew()

	content := contentWritable.Bytes()

	siart.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAccountRightTime from the given readable
func (siart *ServiceItemAccountRightTime) ExtractFrom(readable types.Readable) error {
	var err error

	err = siart.ServiceItemAccountRight.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightTime.ServiceItemAccountRight. %s", err.Error())
	}

	err = siart.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightTime header. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAccountRightTime
func (siart *ServiceItemAccountRightTime) Copy() types.RVType {
	copied := NewServiceItemAccountRightTime()

	copied.StructureVersion = siart.StructureVersion
	copied.ServiceItemAccountRight = siart.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)

	return copied
}

// Equals checks if the given ServiceItemAccountRightTime contains the same data as the current ServiceItemAccountRightTime
func (siart *ServiceItemAccountRightTime) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRightTime); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRightTime)

	if siart.StructureVersion != other.StructureVersion {
		return false
	}

	if !siart.ServiceItemAccountRight.Equals(other.ServiceItemAccountRight) {
		return false
	}

	return true
}

// String returns the string representation of the ServiceItemAccountRightTime
func (siart *ServiceItemAccountRightTime) String() string {
	return siart.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAccountRightTime using the provided indentation level
func (siart *ServiceItemAccountRightTime) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRightTime{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemAccountRight (parent): %s,\n", indentationValues, siart.ServiceItemAccountRight.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRightTime returns a new ServiceItemAccountRightTime
func NewServiceItemAccountRightTime() *ServiceItemAccountRightTime {
	siart := &ServiceItemAccountRightTime{
		ServiceItemAccountRight: NewServiceItemAccountRight(),
	}

	return siart
}
