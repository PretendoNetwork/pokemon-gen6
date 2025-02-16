// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightConsumptionInfo is a type within the ServiceItem protocol
type ServiceItemRightConsumptionInfo struct {
	types.Structure
	*ServiceItemRightInfo
	AccountRights *types.List[*ServiceItemAccountRightConsumption]
}

// WriteTo writes the ServiceItemRightConsumptionInfo to the given writable
func (sirci *ServiceItemRightConsumptionInfo) WriteTo(writable types.Writable) {
	sirci.ServiceItemRightInfo.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sirci.AccountRights.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sirci.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightConsumptionInfo from the given readable
func (sirci *ServiceItemRightConsumptionInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = sirci.ServiceItemRightInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightConsumptionInfo.ServiceItemRightInfo. %s", err.Error())
	}

	err = sirci.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightConsumptionInfo header. %s", err.Error())
	}

	err = sirci.AccountRights.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightConsumptionInfo.AccountRights. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightConsumptionInfo
func (sirci *ServiceItemRightConsumptionInfo) Copy() types.RVType {
	copied := NewServiceItemRightConsumptionInfo()

	copied.StructureVersion = sirci.StructureVersion
	copied.ServiceItemRightInfo = sirci.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)
	copied.AccountRights = sirci.AccountRights.Copy().(*types.List[*ServiceItemAccountRightConsumption])

	return copied
}

// Equals checks if the given ServiceItemRightConsumptionInfo contains the same data as the current ServiceItemRightConsumptionInfo
func (sirci *ServiceItemRightConsumptionInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightConsumptionInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightConsumptionInfo)

	if sirci.StructureVersion != other.StructureVersion {
		return false
	}

	if !sirci.ServiceItemRightInfo.Equals(other.ServiceItemRightInfo) {
		return false
	}

	return sirci.AccountRights.Equals(other.AccountRights)
}

// String returns the string representation of the ServiceItemRightConsumptionInfo
func (sirci *ServiceItemRightConsumptionInfo) String() string {
	return sirci.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightConsumptionInfo using the provided indentation level
func (sirci *ServiceItemRightConsumptionInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightConsumptionInfo{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemRightInfo (parent): %s,\n", indentationValues, sirci.ServiceItemRightInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAccountRights: %s,\n", indentationValues, sirci.AccountRights))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightConsumptionInfo returns a new ServiceItemRightConsumptionInfo
func NewServiceItemRightConsumptionInfo() *ServiceItemRightConsumptionInfo {
	sirci := &ServiceItemRightConsumptionInfo{
		ServiceItemRightInfo: NewServiceItemRightInfo(),
		AccountRights:        types.NewList[*ServiceItemAccountRightConsumption](),
	}

	sirci.AccountRights.Type = NewServiceItemAccountRightConsumption()

	return sirci
}
