// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightInfos is a type within the ServiceItem protocol
type ServiceItemRightInfos struct {
	types.Structure
	SupportID                       *types.String
	ConsumptionRightInfos           *types.List[*ServiceItemRightConsumptionInfo]
	AdditionalTimeRightInfos        *types.List[*ServiceItemRightTimeInfo]
	PermanentRightInfos             *types.List[*ServiceItemRightTimeInfo]
	AlreadyPurchasedInitialOnlyItem *types.PrimitiveBool
}

// WriteTo writes the ServiceItemRightInfos to the given writable
func (siri *ServiceItemRightInfos) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siri.SupportID.WriteTo(contentWritable)
	siri.ConsumptionRightInfos.WriteTo(contentWritable)
	siri.AdditionalTimeRightInfos.WriteTo(contentWritable)
	siri.PermanentRightInfos.WriteTo(contentWritable)
	siri.AlreadyPurchasedInitialOnlyItem.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightInfos from the given readable
func (siri *ServiceItemRightInfos) ExtractFrom(readable types.Readable) error {
	var err error

	err = siri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos header. %s", err.Error())
	}

	err = siri.SupportID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.SupportID. %s", err.Error())
	}

	err = siri.ConsumptionRightInfos.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.ConsumptionRightInfos. %s", err.Error())
	}

	err = siri.AdditionalTimeRightInfos.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AdditionalTimeRightInfos. %s", err.Error())
	}

	err = siri.PermanentRightInfos.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.PermanentRightInfos. %s", err.Error())
	}

	err = siri.AlreadyPurchasedInitialOnlyItem.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AlreadyPurchasedInitialOnlyItem. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (siri *ServiceItemRightInfos) Copy() types.RVType {
	copied := NewServiceItemRightInfos()

	copied.StructureVersion = siri.StructureVersion
	copied.SupportID = siri.SupportID.Copy().(*types.String)
	copied.ConsumptionRightInfos = siri.ConsumptionRightInfos.Copy().(*types.List[*ServiceItemRightConsumptionInfo])
	copied.AdditionalTimeRightInfos = siri.AdditionalTimeRightInfos.Copy().(*types.List[*ServiceItemRightTimeInfo])
	copied.PermanentRightInfos = siri.PermanentRightInfos.Copy().(*types.List[*ServiceItemRightTimeInfo])
	copied.AlreadyPurchasedInitialOnlyItem = siri.AlreadyPurchasedInitialOnlyItem.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the given ServiceItemRightInfos contains the same data as the current ServiceItemRightInfos
func (siri *ServiceItemRightInfos) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfos); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfos)

	if siri.StructureVersion != other.StructureVersion {
		return false
	}

	if !siri.SupportID.Equals(other.SupportID) {
		return false
	}

	if !siri.ConsumptionRightInfos.Equals(other.ConsumptionRightInfos) {
		return false
	}

	if !siri.AdditionalTimeRightInfos.Equals(other.AdditionalTimeRightInfos) {
		return false
	}

	if !siri.PermanentRightInfos.Equals(other.PermanentRightInfos) {
		return false
	}

	return siri.AlreadyPurchasedInitialOnlyItem.Equals(other.AlreadyPurchasedInitialOnlyItem)
}

// String returns the string representation of the ServiceItemRightInfos
func (siri *ServiceItemRightInfos) String() string {
	return siri.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightInfos using the provided indentation level
func (siri *ServiceItemRightInfos) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfos{\n")
	b.WriteString(fmt.Sprintf("%sSupportID: %s,\n", indentationValues, siri.SupportID))
	b.WriteString(fmt.Sprintf("%sConsumptionRightInfos: %s,\n", indentationValues, siri.ConsumptionRightInfos))
	b.WriteString(fmt.Sprintf("%sAdditionalTimeRightInfos: %s,\n", indentationValues, siri.AdditionalTimeRightInfos))
	b.WriteString(fmt.Sprintf("%sPermanentRightInfos: %s,\n", indentationValues, siri.PermanentRightInfos))
	b.WriteString(fmt.Sprintf("%sAlreadyPurchasedInitialOnlyItem: %s,\n", indentationValues, siri.AlreadyPurchasedInitialOnlyItem))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfos returns a new ServiceItemRightInfos
func NewServiceItemRightInfos() *ServiceItemRightInfos {
	siri := &ServiceItemRightInfos{
		SupportID:                       types.NewString(""),
		ConsumptionRightInfos:           types.NewList[*ServiceItemRightConsumptionInfo](),
		AdditionalTimeRightInfos:        types.NewList[*ServiceItemRightTimeInfo](),
		PermanentRightInfos:             types.NewList[*ServiceItemRightTimeInfo](),
		AlreadyPurchasedInitialOnlyItem: types.NewPrimitiveBool(false),
	}

	siri.ConsumptionRightInfos.Type = NewServiceItemRightConsumptionInfo()
	siri.AdditionalTimeRightInfos.Type = NewServiceItemRightTimeInfo()
	siri.PermanentRightInfos.Type = NewServiceItemRightTimeInfo()

	return siri
}
