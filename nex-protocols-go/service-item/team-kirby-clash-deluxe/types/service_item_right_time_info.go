// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightTimeInfo is a type within the ServiceItem protocol
type ServiceItemRightTimeInfo struct {
	types.Structure
	*ServiceItemRightInfo
	AccountRights *types.List[*ServiceItemAccountRightTime]
}

// WriteTo writes the ServiceItemRightTimeInfo to the given writable
func (sirti *ServiceItemRightTimeInfo) WriteTo(writable types.Writable) {
	sirti.ServiceItemRightInfo.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sirti.AccountRights.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sirti.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightTimeInfo from the given readable
func (sirti *ServiceItemRightTimeInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = sirti.ServiceItemRightInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightTimeInfo.ServiceItemRightInfo. %s", err.Error())
	}

	err = sirti.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightTimeInfo header. %s", err.Error())
	}

	err = sirti.AccountRights.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightTimeInfo.AccountRights. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightTimeInfo
func (sirti *ServiceItemRightTimeInfo) Copy() types.RVType {
	copied := NewServiceItemRightTimeInfo()

	copied.StructureVersion = sirti.StructureVersion
	copied.ServiceItemRightInfo = sirti.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)
	copied.AccountRights = sirti.AccountRights.Copy().(*types.List[*ServiceItemAccountRightTime])

	return copied
}

// Equals checks if the given ServiceItemRightTimeInfo contains the same data as the current ServiceItemRightTimeInfo
func (sirti *ServiceItemRightTimeInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightTimeInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightTimeInfo)

	if sirti.StructureVersion != other.StructureVersion {
		return false
	}

	if !sirti.ServiceItemRightInfo.Equals(other.ServiceItemRightInfo) {
		return false
	}

	return sirti.AccountRights.Equals(other.AccountRights)
}

// String returns the string representation of the ServiceItemRightTimeInfo
func (sirti *ServiceItemRightTimeInfo) String() string {
	return sirti.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightTimeInfo using the provided indentation level
func (sirti *ServiceItemRightTimeInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightTimeInfo{\n")
	b.WriteString(fmt.Sprintf("%sServiceItemRightInfo (parent): %s,\n", indentationValues, sirti.ServiceItemRightInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAccountRights: %s,\n", indentationValues, sirti.AccountRights))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightTimeInfo returns a new ServiceItemRightTimeInfo
func NewServiceItemRightTimeInfo() *ServiceItemRightTimeInfo {
	sirti := &ServiceItemRightTimeInfo{
		ServiceItemRightInfo: NewServiceItemRightInfo(),
		AccountRights:        types.NewList[*ServiceItemAccountRightTime](),
	}

	sirti.AccountRights.Type = NewServiceItemAccountRightTime()

	return sirti
}
