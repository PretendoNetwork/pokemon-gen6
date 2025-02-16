// Package types implements all the types used by the AccountManagement protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-wiiu/types"
)

// NintendoCreateAccountData is a type within the AccountManagement protocol
type NintendoCreateAccountData struct {
	types.Structure
	NNAInfo  *friends_wiiu_types.NNAInfo
	Token    *types.String
	Birthday *types.DateTime
	Unknown  *types.PrimitiveU64
}

// WriteTo writes the NintendoCreateAccountData to the given writable
func (ncad *NintendoCreateAccountData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ncad.NNAInfo.WriteTo(contentWritable)
	ncad.Token.WriteTo(contentWritable)
	ncad.Birthday.WriteTo(contentWritable)
	ncad.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ncad.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoCreateAccountData from the given readable
func (ncad *NintendoCreateAccountData) ExtractFrom(readable types.Readable) error {
	var err error

	err = ncad.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData header. %s", err.Error())
	}

	err = ncad.NNAInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.NNAInfo. %s", err.Error())
	}

	err = ncad.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Token. %s", err.Error())
	}

	err = ncad.Birthday.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Birthday. %s", err.Error())
	}

	err = ncad.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoCreateAccountData
func (ncad *NintendoCreateAccountData) Copy() types.RVType {
	copied := NewNintendoCreateAccountData()

	copied.StructureVersion = ncad.StructureVersion
	copied.NNAInfo = ncad.NNAInfo.Copy().(*friends_wiiu_types.NNAInfo)
	copied.Token = ncad.Token.Copy().(*types.String)
	copied.Birthday = ncad.Birthday.Copy().(*types.DateTime)
	copied.Unknown = ncad.Unknown.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the given NintendoCreateAccountData contains the same data as the current NintendoCreateAccountData
func (ncad *NintendoCreateAccountData) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoCreateAccountData); !ok {
		return false
	}

	other := o.(*NintendoCreateAccountData)

	if ncad.StructureVersion != other.StructureVersion {
		return false
	}

	if !ncad.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if !ncad.Token.Equals(other.Token) {
		return false
	}

	if !ncad.Birthday.Equals(other.Birthday) {
		return false
	}

	return ncad.Unknown.Equals(other.Unknown)
}

// String returns the string representation of the NintendoCreateAccountData
func (ncad *NintendoCreateAccountData) String() string {
	return ncad.FormatToString(0)
}

// FormatToString pretty-prints the NintendoCreateAccountData using the provided indentation level
func (ncad *NintendoCreateAccountData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoCreateAccountData{\n")
	b.WriteString(fmt.Sprintf("%sNNAInfo: %s,\n", indentationValues, ncad.NNAInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, ncad.Token))
	b.WriteString(fmt.Sprintf("%sBirthday: %s,\n", indentationValues, ncad.Birthday.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, ncad.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoCreateAccountData returns a new NintendoCreateAccountData
func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	ncad := &NintendoCreateAccountData{
		NNAInfo:  friends_wiiu_types.NewNNAInfo(),
		Token:    types.NewString(""),
		Birthday: types.NewDateTime(0),
		Unknown:  types.NewPrimitiveU64(0),
	}

	return ncad
}
