// Package types implements all the types used by the FriendsWiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NNAInfo is a type within the FriendsWiiU protocol
type NNAInfo struct {
	types.Structure
	*types.Data
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           *types.PrimitiveU8
	Unknown2           *types.PrimitiveU8
}

// WriteTo writes the NNAInfo to the given writable
func (nnai *NNAInfo) WriteTo(writable types.Writable) {
	nnai.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	nnai.PrincipalBasicInfo.WriteTo(contentWritable)
	nnai.Unknown1.WriteTo(contentWritable)
	nnai.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nnai.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NNAInfo from the given readable
func (nnai *NNAInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = nnai.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Data. %s", err.Error())
	}

	err = nnai.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo header. %s", err.Error())
	}

	err = nnai.PrincipalBasicInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.PrincipalBasicInfo. %s", err.Error())
	}

	err = nnai.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown1. %s", err.Error())
	}

	err = nnai.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NNAInfo
func (nnai *NNAInfo) Copy() types.RVType {
	copied := NewNNAInfo()

	copied.StructureVersion = nnai.StructureVersion
	copied.Data = nnai.Data.Copy().(*types.Data)
	copied.PrincipalBasicInfo = nnai.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.Unknown1 = nnai.Unknown1.Copy().(*types.PrimitiveU8)
	copied.Unknown2 = nnai.Unknown2.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given NNAInfo contains the same data as the current NNAInfo
func (nnai *NNAInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*NNAInfo); !ok {
		return false
	}

	other := o.(*NNAInfo)

	if nnai.StructureVersion != other.StructureVersion {
		return false
	}

	if !nnai.Data.Equals(other.Data) {
		return false
	}

	if !nnai.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if !nnai.Unknown1.Equals(other.Unknown1) {
		return false
	}

	return nnai.Unknown2.Equals(other.Unknown2)
}

// String returns the string representation of the NNAInfo
func (nnai *NNAInfo) String() string {
	return nnai.FormatToString(0)
}

// FormatToString pretty-prints the NNAInfo using the provided indentation level
func (nnai *NNAInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NNAInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, nnai.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrincipalBasicInfo: %s,\n", indentationValues, nnai.PrincipalBasicInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, nnai.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, nnai.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNNAInfo returns a new NNAInfo
func NewNNAInfo() *NNAInfo {
	nnai := &NNAInfo{
		Data:               types.NewData(),
		PrincipalBasicInfo: NewPrincipalBasicInfo(),
		Unknown1:           types.NewPrimitiveU8(0),
		Unknown2:           types.NewPrimitiveU8(0),
	}

	return nnai
}
