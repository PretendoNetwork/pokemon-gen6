// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingUploadScoreParam is a type within the Ranking protocol
type CompetitionRankingUploadScoreParam struct {
	types.Structure
	Unknown  *types.PrimitiveU32
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU32
	Unknown5 *types.PrimitiveU8
	Unknown6 *types.PrimitiveU32
	Unknown7 *types.PrimitiveBool
	Metadata *types.QBuffer
}

// WriteTo writes the CompetitionRankingUploadScoreParam to the given writable
func (crusp *CompetitionRankingUploadScoreParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crusp.Unknown.WriteTo(contentWritable)
	crusp.Unknown2.WriteTo(contentWritable)
	crusp.Unknown3.WriteTo(contentWritable)
	crusp.Unknown4.WriteTo(contentWritable)
	crusp.Unknown5.WriteTo(contentWritable)
	crusp.Unknown6.WriteTo(contentWritable)
	crusp.Unknown7.WriteTo(contentWritable)
	crusp.Metadata.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crusp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingUploadScoreParam from the given readable
func (crusp *CompetitionRankingUploadScoreParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = crusp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam header. %s", err.Error())
	}

	err = crusp.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown. %s", err.Error())
	}

	err = crusp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown2. %s", err.Error())
	}

	err = crusp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown3. %s", err.Error())
	}

	err = crusp.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown4. %s", err.Error())
	}

	err = crusp.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown5. %s", err.Error())
	}

	err = crusp.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown6. %s", err.Error())
	}

	err = crusp.Unknown7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown7. %s", err.Error())
	}

	err = crusp.Metadata.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Metadata. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingUploadScoreParam
func (crusp *CompetitionRankingUploadScoreParam) Copy() types.RVType {
	copied := NewCompetitionRankingUploadScoreParam()

	copied.StructureVersion = crusp.StructureVersion
	copied.Unknown = crusp.Unknown.Copy().(*types.PrimitiveU32)
	copied.Unknown2 = crusp.Unknown2.Copy().(*types.PrimitiveU32)
	copied.Unknown3 = crusp.Unknown3.Copy().(*types.PrimitiveU32)
	copied.Unknown4 = crusp.Unknown4.Copy().(*types.PrimitiveU32)
	copied.Unknown5 = crusp.Unknown5.Copy().(*types.PrimitiveU8)
	copied.Unknown6 = crusp.Unknown6.Copy().(*types.PrimitiveU32)
	copied.Unknown7 = crusp.Unknown7.Copy().(*types.PrimitiveBool)
	copied.Metadata = crusp.Metadata.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the given CompetitionRankingUploadScoreParam contains the same data as the current CompetitionRankingUploadScoreParam
func (crusp *CompetitionRankingUploadScoreParam) Equals(o types.RVType) bool {
	if _, ok := o.(*CompetitionRankingUploadScoreParam); !ok {
		return false
	}

	other := o.(*CompetitionRankingUploadScoreParam)

	if crusp.StructureVersion != other.StructureVersion {
		return false
	}

	if !crusp.Unknown.Equals(other.Unknown) {
		return false
	}

	if !crusp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !crusp.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !crusp.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !crusp.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !crusp.Unknown6.Equals(other.Unknown6) {
		return false
	}

	if !crusp.Unknown7.Equals(other.Unknown7) {
		return false
	}

	return crusp.Metadata.Equals(other.Metadata)
}

// String returns the string representation of the CompetitionRankingUploadScoreParam
func (crusp *CompetitionRankingUploadScoreParam) String() string {
	return crusp.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingUploadScoreParam using the provided indentation level
func (crusp *CompetitionRankingUploadScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingUploadScoreParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, crusp.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, crusp.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, crusp.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %s,\n", indentationValues, crusp.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, crusp.Unknown5))
	b.WriteString(fmt.Sprintf("%sUnknown6: %s,\n", indentationValues, crusp.Unknown6))
	b.WriteString(fmt.Sprintf("%sUnknown7: %s,\n", indentationValues, crusp.Unknown7))
	b.WriteString(fmt.Sprintf("%sMetadata: %s,\n", indentationValues, crusp.Metadata))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingUploadScoreParam returns a new CompetitionRankingUploadScoreParam
func NewCompetitionRankingUploadScoreParam() *CompetitionRankingUploadScoreParam {
	crusp := &CompetitionRankingUploadScoreParam{
		Unknown:  types.NewPrimitiveU32(0),
		Unknown2: types.NewPrimitiveU32(0),
		Unknown3: types.NewPrimitiveU32(0),
		Unknown4: types.NewPrimitiveU32(0),
		Unknown5: types.NewPrimitiveU8(0),
		Unknown6: types.NewPrimitiveU32(0),
		Unknown7: types.NewPrimitiveBool(false),
		Metadata: types.NewQBuffer(nil),
	}

	return crusp
}
