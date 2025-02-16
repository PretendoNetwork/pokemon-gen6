// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayedGame is a type within the Friends3DS protocol
type PlayedGame struct {
	types.Structure
	GameKey *GameKey
	Unknown *types.DateTime
}

// WriteTo writes the PlayedGame to the given writable
func (pg *PlayedGame) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pg.GameKey.WriteTo(contentWritable)
	pg.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pg.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayedGame from the given readable
func (pg *PlayedGame) ExtractFrom(readable types.Readable) error {
	var err error

	err = pg.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayedGame header. %s", err.Error())
	}

	err = pg.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayedGame.GameKey. %s", err.Error())
	}

	err = pg.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayedGame.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayedGame
func (pg *PlayedGame) Copy() types.RVType {
	copied := NewPlayedGame()

	copied.StructureVersion = pg.StructureVersion
	copied.GameKey = pg.GameKey.Copy().(*GameKey)
	copied.Unknown = pg.Unknown.Copy().(*types.DateTime)

	return copied
}

// Equals checks if the given PlayedGame contains the same data as the current PlayedGame
func (pg *PlayedGame) Equals(o types.RVType) bool {
	if _, ok := o.(*PlayedGame); !ok {
		return false
	}

	other := o.(*PlayedGame)

	if pg.StructureVersion != other.StructureVersion {
		return false
	}

	if !pg.GameKey.Equals(other.GameKey) {
		return false
	}

	return pg.Unknown.Equals(other.Unknown)
}

// String returns the string representation of the PlayedGame
func (pg *PlayedGame) String() string {
	return pg.FormatToString(0)
}

// FormatToString pretty-prints the PlayedGame using the provided indentation level
func (pg *PlayedGame) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayedGame{\n")
	b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, pg.GameKey.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, pg.Unknown.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPlayedGame returns a new PlayedGame
func NewPlayedGame() *PlayedGame {
	pg := &PlayedGame{
		GameKey: NewGameKey(),
		Unknown: types.NewDateTime(0),
	}

	return pg
}
