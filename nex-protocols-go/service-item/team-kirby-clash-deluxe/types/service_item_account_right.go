// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAccountRight is a type within the ServiceItem protocol
type ServiceItemAccountRight struct {
	types.Structure
	PID           *types.PID
	Limitation    *ServiceItemLimitation
	RightBinaries *types.List[*ServiceItemRightBinary]
}

// WriteTo writes the ServiceItemAccountRight to the given writable
func (siar *ServiceItemAccountRight) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siar.PID.WriteTo(contentWritable)
	siar.Limitation.WriteTo(contentWritable)
	siar.RightBinaries.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siar.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAccountRight from the given readable
func (siar *ServiceItemAccountRight) ExtractFrom(readable types.Readable) error {
	var err error

	err = siar.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight header. %s", err.Error())
	}

	err = siar.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.PID. %s", err.Error())
	}

	err = siar.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.Limitation. %s", err.Error())
	}

	err = siar.RightBinaries.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.RightBinaries. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAccountRight
func (siar *ServiceItemAccountRight) Copy() types.RVType {
	copied := NewServiceItemAccountRight()

	copied.StructureVersion = siar.StructureVersion
	copied.PID = siar.PID.Copy().(*types.PID)
	copied.Limitation = siar.Limitation.Copy().(*ServiceItemLimitation)
	copied.RightBinaries = siar.RightBinaries.Copy().(*types.List[*ServiceItemRightBinary])

	return copied
}

// Equals checks if the given ServiceItemAccountRight contains the same data as the current ServiceItemAccountRight
func (siar *ServiceItemAccountRight) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRight); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRight)

	if siar.StructureVersion != other.StructureVersion {
		return false
	}

	if !siar.PID.Equals(other.PID) {
		return false
	}

	if !siar.Limitation.Equals(other.Limitation) {
		return false
	}

	return siar.RightBinaries.Equals(other.RightBinaries)
}

// String returns the string representation of the ServiceItemAccountRight
func (siar *ServiceItemAccountRight) String() string {
	return siar.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAccountRight using the provided indentation level
func (siar *ServiceItemAccountRight) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRight{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, siar.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLimitation: %s,\n", indentationValues, siar.Limitation.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRightBinaries: %s,\n", indentationValues, siar.RightBinaries))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRight returns a new ServiceItemAccountRight
func NewServiceItemAccountRight() *ServiceItemAccountRight {
	siar := &ServiceItemAccountRight{
		PID:           types.NewPID(0),
		Limitation:    NewServiceItemLimitation(),
		RightBinaries: types.NewList[*ServiceItemRightBinary](),
	}

	siar.RightBinaries.Type = NewServiceItemRightBinary()

	return siar
}
