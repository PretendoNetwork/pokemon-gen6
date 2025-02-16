// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetServiceItemRightParam is a type within the ServiceItem protocol
type ServiceItemGetServiceItemRightParam struct {
	types.Structure
	ReferenceID *types.String
	TitleID     *types.String
}

// WriteTo writes the ServiceItemGetServiceItemRightParam to the given writable
func (sigsirp *ServiceItemGetServiceItemRightParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigsirp.ReferenceID.WriteTo(contentWritable)
	sigsirp.TitleID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigsirp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetServiceItemRightParam from the given readable
func (sigsirp *ServiceItemGetServiceItemRightParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigsirp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam header. %s", err.Error())
	}

	err = sigsirp.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ReferenceID. %s", err.Error())
	}

	err = sigsirp.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.TitleID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) Copy() types.RVType {
	copied := NewServiceItemGetServiceItemRightParam()

	copied.StructureVersion = sigsirp.StructureVersion
	copied.ReferenceID = sigsirp.ReferenceID.Copy().(*types.String)
	copied.TitleID = sigsirp.TitleID.Copy().(*types.String)

	return copied
}

// Equals checks if the given ServiceItemGetServiceItemRightParam contains the same data as the current ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetServiceItemRightParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetServiceItemRightParam)

	if sigsirp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigsirp.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	return sigsirp.TitleID.Equals(other.TitleID)
}

// String returns the string representation of the ServiceItemGetServiceItemRightParam
func (sigsirp *ServiceItemGetServiceItemRightParam) String() string {
	return sigsirp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetServiceItemRightParam using the provided indentation level
func (sigsirp *ServiceItemGetServiceItemRightParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightParam{\n")
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, sigsirp.ReferenceID))
	b.WriteString(fmt.Sprintf("%sTitleID: %s,\n", indentationValues, sigsirp.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightParam returns a new ServiceItemGetServiceItemRightParam
func NewServiceItemGetServiceItemRightParam() *ServiceItemGetServiceItemRightParam {
	sigsirp := &ServiceItemGetServiceItemRightParam{
		ReferenceID: types.NewString(""),
		TitleID:     types.NewString(""),
	}

	return sigsirp
}
