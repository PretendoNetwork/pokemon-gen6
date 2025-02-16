// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MyProfile is a type within the Friends3DS protocol
type MyProfile struct {
	types.Structure
	*types.Data
	Region   *types.PrimitiveU8
	Country  *types.PrimitiveU8
	Area     *types.PrimitiveU8
	Language *types.PrimitiveU8
	Platform *types.PrimitiveU8
	Unknown1 *types.PrimitiveU64
	Unknown2 *types.String
	Unknown3 *types.String
}

// WriteTo writes the MyProfile to the given writable
func (mp *MyProfile) WriteTo(writable types.Writable) {
	mp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mp.Region.WriteTo(contentWritable)
	mp.Country.WriteTo(contentWritable)
	mp.Area.WriteTo(contentWritable)
	mp.Language.WriteTo(contentWritable)
	mp.Platform.WriteTo(contentWritable)
	mp.Unknown1.WriteTo(contentWritable)
	mp.Unknown2.WriteTo(contentWritable)
	mp.Unknown3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MyProfile from the given readable
func (mp *MyProfile) ExtractFrom(readable types.Readable) error {
	var err error

	err = mp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Data. %s", err.Error())
	}

	err = mp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile header. %s", err.Error())
	}

	err = mp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Region. %s", err.Error())
	}

	err = mp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Country. %s", err.Error())
	}

	err = mp.Area.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Area. %s", err.Error())
	}

	err = mp.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Language. %s", err.Error())
	}

	err = mp.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Platform. %s", err.Error())
	}

	err = mp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown1. %s", err.Error())
	}

	err = mp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown2. %s", err.Error())
	}

	err = mp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MyProfile
func (mp *MyProfile) Copy() types.RVType {
	copied := NewMyProfile()

	copied.StructureVersion = mp.StructureVersion
	copied.Data = mp.Data.Copy().(*types.Data)
	copied.Region = mp.Region.Copy().(*types.PrimitiveU8)
	copied.Country = mp.Country.Copy().(*types.PrimitiveU8)
	copied.Area = mp.Area.Copy().(*types.PrimitiveU8)
	copied.Language = mp.Language.Copy().(*types.PrimitiveU8)
	copied.Platform = mp.Platform.Copy().(*types.PrimitiveU8)
	copied.Unknown1 = mp.Unknown1.Copy().(*types.PrimitiveU64)
	copied.Unknown2 = mp.Unknown2.Copy().(*types.String)
	copied.Unknown3 = mp.Unknown3.Copy().(*types.String)

	return copied
}

// Equals checks if the given MyProfile contains the same data as the current MyProfile
func (mp *MyProfile) Equals(o types.RVType) bool {
	if _, ok := o.(*MyProfile); !ok {
		return false
	}

	other := o.(*MyProfile)

	if mp.StructureVersion != other.StructureVersion {
		return false
	}

	if !mp.Data.Equals(other.Data) {
		return false
	}

	if !mp.Region.Equals(other.Region) {
		return false
	}

	if !mp.Country.Equals(other.Country) {
		return false
	}

	if !mp.Area.Equals(other.Area) {
		return false
	}

	if !mp.Language.Equals(other.Language) {
		return false
	}

	if !mp.Platform.Equals(other.Platform) {
		return false
	}

	if !mp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !mp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return mp.Unknown3.Equals(other.Unknown3)
}

// String returns the string representation of the MyProfile
func (mp *MyProfile) String() string {
	return mp.FormatToString(0)
}

// FormatToString pretty-prints the MyProfile using the provided indentation level
func (mp *MyProfile) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MyProfile{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, mp.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, mp.Country))
	b.WriteString(fmt.Sprintf("%sArea: %s,\n", indentationValues, mp.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, mp.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, mp.Platform))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, mp.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, mp.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, mp.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	mp := &MyProfile{
		Data:     types.NewData(),
		Region:   types.NewPrimitiveU8(0),
		Country:  types.NewPrimitiveU8(0),
		Area:     types.NewPrimitiveU8(0),
		Language: types.NewPrimitiveU8(0),
		Platform: types.NewPrimitiveU8(0),
		Unknown1: types.NewPrimitiveU64(0),
		Unknown2: types.NewString(""),
		Unknown3: types.NewString(""),
	}

	return mp
}
