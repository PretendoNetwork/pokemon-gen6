// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GatheringURLs is a type within the Matchmaking protocol
type GatheringURLs struct {
	types.Structure
	GID            *types.PrimitiveU32
	LstStationURLs *types.List[*types.StationURL]
}

// WriteTo writes the GatheringURLs to the given writable
func (gurl *GatheringURLs) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gurl.GID.WriteTo(contentWritable)
	gurl.LstStationURLs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gurl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GatheringURLs from the given readable
func (gurl *GatheringURLs) ExtractFrom(readable types.Readable) error {
	var err error

	err = gurl.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringURLs header. %s", err.Error())
	}

	err = gurl.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringURLs.GID. %s", err.Error())
	}

	err = gurl.LstStationURLs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GatheringURLs.LstStationURLs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GatheringURLs
func (gurl *GatheringURLs) Copy() types.RVType {
	copied := NewGatheringURLs()

	copied.StructureVersion = gurl.StructureVersion
	copied.GID = gurl.GID.Copy().(*types.PrimitiveU32)
	copied.LstStationURLs = gurl.LstStationURLs.Copy().(*types.List[*types.StationURL])

	return copied
}

// Equals checks if the given GatheringURLs contains the same data as the current GatheringURLs
func (gurl *GatheringURLs) Equals(o types.RVType) bool {
	if _, ok := o.(*GatheringURLs); !ok {
		return false
	}

	other := o.(*GatheringURLs)

	if gurl.StructureVersion != other.StructureVersion {
		return false
	}

	if !gurl.GID.Equals(other.GID) {
		return false
	}

	return gurl.LstStationURLs.Equals(other.LstStationURLs)
}

// String returns the string representation of the GatheringURLs
func (gurl *GatheringURLs) String() string {
	return gurl.FormatToString(0)
}

// FormatToString pretty-prints the GatheringURLs using the provided indentation level
func (gurl *GatheringURLs) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GatheringURLs{\n")
	b.WriteString(fmt.Sprintf("%sGID: %s,\n", indentationValues, gurl.GID))
	b.WriteString(fmt.Sprintf("%sLstStationURLs: %s,\n", indentationValues, gurl.LstStationURLs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGatheringURLs returns a new GatheringURLs
func NewGatheringURLs() *GatheringURLs {
	gurl := &GatheringURLs{
		GID:            types.NewPrimitiveU32(0),
		LstStationURLs: types.NewList[*types.StationURL](),
	}

	gurl.LstStationURLs.Type = types.NewStationURL("")

	return gurl
}
