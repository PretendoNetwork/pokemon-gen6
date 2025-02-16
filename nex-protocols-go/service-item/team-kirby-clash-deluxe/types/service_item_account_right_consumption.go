// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAccountRightConsumption is a type within the ServiceItem protocol
type ServiceItemAccountRightConsumption struct {
	types.Structure
	*ServiceItemAccountRight
	UsedCount    *types.PrimitiveU32
	ExpiredCount *types.PrimitiveU32
	ExpiryCounts *types.List[*types.PrimitiveU32]
}

// WriteTo writes the ServiceItemAccountRightConsumption to the given writable
func (siarc *ServiceItemAccountRightConsumption) WriteTo(writable types.Writable) {
	siarc.ServiceItemAccountRight.WriteTo(writable)

	contentWritable := writable.CopyNew()

	siarc.UsedCount.WriteTo(contentWritable)
	siarc.ExpiredCount.WriteTo(contentWritable)
	siarc.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siarc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAccountRightConsumption from the given readable
func (siarc *ServiceItemAccountRightConsumption) ExtractFrom(readable types.Readable) error {
	var err error

	err = siarc.ServiceItemAccountRight.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ServiceItemAccountRight. %s", err.Error())
	}

	err = siarc.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption header. %s", err.Error())
	}

	err = siarc.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.UsedCount. %s", err.Error())
	}

	err = siarc.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiredCount. %s", err.Error())
	}

	err = siarc.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiryCounts. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAccountRightConsumption
func (siarc *ServiceItemAccountRightConsumption) Copy() types.RVType {
	copied := NewServiceItemAccountRightConsumption()

	copied.StructureVersion = siarc.StructureVersion
	copied.ServiceItemAccountRight = siarc.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)
	copied.UsedCount = siarc.UsedCount.Copy().(*types.PrimitiveU32)
	copied.ExpiredCount = siarc.ExpiredCount.Copy().(*types.PrimitiveU32)
	copied.ExpiryCounts = siarc.ExpiryCounts.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given ServiceItemAccountRightConsumption contains the same data as the current ServiceItemAccountRightConsumption
func (siarc *ServiceItemAccountRightConsumption) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRightConsumption); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRightConsumption)

	if siarc.StructureVersion != other.StructureVersion {
		return false
	}

	if !siarc.ServiceItemAccountRight.Equals(other.ServiceItemAccountRight) {
		return false
	}

	if !siarc.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !siarc.ExpiredCount.Equals(other.ExpiredCount) {
		return false
	}

	return siarc.ExpiryCounts.Equals(other.ExpiryCounts)
}

// String returns the string representation of the ServiceItemAccountRightConsumption
func (siarc *ServiceItemAccountRightConsumption) String() string {
	return siarc.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAccountRightConsumption using the provided indentation level
func (siarc *ServiceItemAccountRightConsumption) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRightConsumption{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemAccountRight (parent): %s,\n", indentationValues, siarc.ServiceItemAccountRight.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUsedCount: %s,\n", indentationValues, siarc.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %s,\n", indentationValues, siarc.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %s,\n", indentationValues, siarc.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRightConsumption returns a new ServiceItemAccountRightConsumption
func NewServiceItemAccountRightConsumption() *ServiceItemAccountRightConsumption {
	siarc := &ServiceItemAccountRightConsumption{
		ServiceItemAccountRight: NewServiceItemAccountRight(),
		UsedCount:               types.NewPrimitiveU32(0),
		ExpiredCount:            types.NewPrimitiveU32(0),
		ExpiryCounts:            types.NewList[*types.PrimitiveU32](),
	}

	siarc.ExpiryCounts.Type = types.NewPrimitiveU32(0)

	return siarc
}
