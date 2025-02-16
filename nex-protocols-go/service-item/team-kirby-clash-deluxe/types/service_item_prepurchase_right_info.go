// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPrepurchaseRightInfo is a type within the ServiceItem protocol
type ServiceItemPrepurchaseRightInfo struct {
	types.Structure
	LimitationType *types.PrimitiveU32
	AcquiredCount  *types.PrimitiveU32
	UsedCount      *types.PrimitiveU32
	ExpiryDate     *types.PrimitiveU32
	ExpiredCount   *types.PrimitiveU32
	ExpiryCounts   *types.List[*types.PrimitiveU32]
}

// WriteTo writes the ServiceItemPrepurchaseRightInfo to the given writable
func (sipri *ServiceItemPrepurchaseRightInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipri.LimitationType.WriteTo(contentWritable)
	sipri.AcquiredCount.WriteTo(contentWritable)
	sipri.UsedCount.WriteTo(contentWritable)
	sipri.ExpiryDate.WriteTo(contentWritable)
	sipri.ExpiredCount.WriteTo(contentWritable)
	sipri.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPrepurchaseRightInfo from the given readable
func (sipri *ServiceItemPrepurchaseRightInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = sipri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo header. %s", err.Error())
	}

	err = sipri.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.LimitationType. %s", err.Error())
	}

	err = sipri.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.AcquiredCount. %s", err.Error())
	}

	err = sipri.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.UsedCount. %s", err.Error())
	}

	err = sipri.ExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryDate. %s", err.Error())
	}

	err = sipri.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiredCount. %s", err.Error())
	}

	err = sipri.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryCounts. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPrepurchaseRightInfo
func (sipri *ServiceItemPrepurchaseRightInfo) Copy() types.RVType {
	copied := NewServiceItemPrepurchaseRightInfo()

	copied.StructureVersion = sipri.StructureVersion
	copied.LimitationType = sipri.LimitationType.Copy().(*types.PrimitiveU32)
	copied.AcquiredCount = sipri.AcquiredCount.Copy().(*types.PrimitiveU32)
	copied.UsedCount = sipri.UsedCount.Copy().(*types.PrimitiveU32)
	copied.ExpiryDate = sipri.ExpiryDate.Copy().(*types.PrimitiveU32)
	copied.ExpiredCount = sipri.ExpiredCount.Copy().(*types.PrimitiveU32)
	copied.ExpiryCounts = sipri.ExpiryCounts.Copy().(*types.List[*types.PrimitiveU32])

	return copied
}

// Equals checks if the given ServiceItemPrepurchaseRightInfo contains the same data as the current ServiceItemPrepurchaseRightInfo
func (sipri *ServiceItemPrepurchaseRightInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPrepurchaseRightInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPrepurchaseRightInfo)

	if sipri.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipri.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !sipri.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	if !sipri.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !sipri.ExpiryDate.Equals(other.ExpiryDate) {
		return false
	}

	if !sipri.ExpiredCount.Equals(other.ExpiredCount) {
		return false
	}

	return sipri.ExpiryCounts.Equals(other.ExpiryCounts)
}

// String returns the string representation of the ServiceItemPrepurchaseRightInfo
func (sipri *ServiceItemPrepurchaseRightInfo) String() string {
	return sipri.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPrepurchaseRightInfo using the provided indentation level
func (sipri *ServiceItemPrepurchaseRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPrepurchaseRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sLimitationType: %s,\n", indentationValues, sipri.LimitationType))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %s,\n", indentationValues, sipri.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %s,\n", indentationValues, sipri.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiryDate: %s,\n", indentationValues, sipri.ExpiryDate))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %s,\n", indentationValues, sipri.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %s,\n", indentationValues, sipri.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPrepurchaseRightInfo returns a new ServiceItemPrepurchaseRightInfo
func NewServiceItemPrepurchaseRightInfo() *ServiceItemPrepurchaseRightInfo {
	sipri := &ServiceItemPrepurchaseRightInfo{
		LimitationType: types.NewPrimitiveU32(0),
		AcquiredCount:  types.NewPrimitiveU32(0),
		UsedCount:      types.NewPrimitiveU32(0),
		ExpiryDate:     types.NewPrimitiveU32(0),
		ExpiredCount:   types.NewPrimitiveU32(0),
		ExpiryCounts:   types.NewList[*types.PrimitiveU32](),
	}

	sipri.ExpiryCounts.Type = types.NewPrimitiveU32(0)

	return sipri
}
