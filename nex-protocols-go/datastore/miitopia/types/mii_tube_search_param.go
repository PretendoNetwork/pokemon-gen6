// Package types implements all the types used by the DataStoreMiitopia protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MiiTubeSearchParam is a type within the DataStoreMiitopia protocol
type MiiTubeSearchParam struct {
	types.Structure
	Name         *types.String
	Page         *types.PrimitiveU32
	Category     *types.PrimitiveU8
	Gender       *types.PrimitiveU8
	Country      *types.PrimitiveU8
	SearchType   *types.PrimitiveU8
	ResultOption *types.PrimitiveU8
}

// WriteTo writes the MiiTubeSearchParam to the given variable
func (mtsp *MiiTubeSearchParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mtsp.Name.WriteTo(contentWritable)
	mtsp.Page.WriteTo(contentWritable)
	mtsp.Category.WriteTo(contentWritable)
	mtsp.Gender.WriteTo(contentWritable)
	mtsp.Country.WriteTo(contentWritable)
	mtsp.SearchType.WriteTo(contentWritable)
	mtsp.ResultOption.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mtsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiTubeSearchParam from the given readable
func (mtsp *MiiTubeSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = mtsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam header. %s", err.Error())
	}

	err = mtsp.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.Name. %s", err.Error())
	}

	err = mtsp.Page.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.Page. %s", err.Error())
	}

	err = mtsp.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.Category. %s", err.Error())
	}

	err = mtsp.Gender.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.Gender. %s", err.Error())
	}

	err = mtsp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.Country. %s", err.Error())
	}

	err = mtsp.SearchType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.SearchType. %s", err.Error())
	}

	err = mtsp.ResultOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiTubeSearchParam.ResultOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiTubeSearchParam
func (mtsp *MiiTubeSearchParam) Copy() types.RVType {
	copied := NewMiiTubeSearchParam()

	copied.Name = mtsp.Name
	copied.Page = mtsp.Page
	copied.Category = mtsp.Category
	copied.Gender = mtsp.Gender
	copied.Country = mtsp.Country
	copied.SearchType = mtsp.SearchType
	copied.ResultOption = mtsp.ResultOption

	return copied
}

// Equals checks if the given MiiTubeSearchParam contains the same data as the current MiiTubeSearchParam
func (mtsp *MiiTubeSearchParam) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiTubeSearchParam); !ok {
		return false
	}

	other := o.(*MiiTubeSearchParam)

	if !mtsp.Name.Equals(other.Name) {
		return false
	}

	if !mtsp.Page.Equals(other.Page) {
		return false
	}

	if !mtsp.Category.Equals(other.Category) {
		return false
	}

	if !mtsp.Gender.Equals(other.Gender) {
		return false
	}

	if !mtsp.Country.Equals(other.Country) {
		return false
	}

	if !mtsp.SearchType.Equals(other.SearchType) {
		return false
	}

	return mtsp.ResultOption.Equals(other.ResultOption)
}

// String returns the string representation of the MiiTubeSearchParam
func (mtsp *MiiTubeSearchParam) String() string {
	return mtsp.FormatToString(0)
}

// FormatToString pretty-prints the MiiTubeSearchParam using the provided indentation level
func (mtsp *MiiTubeSearchParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiTubeSearchParam{\n")
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, mtsp.Name))
	b.WriteString(fmt.Sprintf("%sPage: %s,\n", indentationValues, mtsp.Page))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, mtsp.Category))
	b.WriteString(fmt.Sprintf("%sGender: %s,\n", indentationValues, mtsp.Gender))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, mtsp.Country))
	b.WriteString(fmt.Sprintf("%sSearchType: %s,\n", indentationValues, mtsp.SearchType))
	b.WriteString(fmt.Sprintf("%sResultOption: %s,\n", indentationValues, mtsp.ResultOption))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiTubeSearchParam returns a new MiiTubeSearchParam
func NewMiiTubeSearchParam() *MiiTubeSearchParam {
	mtsp := &MiiTubeSearchParam{
		Name:         types.NewString(""),
		Page:         types.NewPrimitiveU32(0),
		Category:     types.NewPrimitiveU8(0),
		Gender:       types.NewPrimitiveU8(0),
		Country:      types.NewPrimitiveU8(0),
		SearchType:   types.NewPrimitiveU8(0),
		ResultOption: types.NewPrimitiveU8(0),
	}

	return mtsp
}
