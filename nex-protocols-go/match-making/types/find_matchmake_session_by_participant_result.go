// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FindMatchmakeSessionByParticipantResult is a type within the Matchmaking protocol
type FindMatchmakeSessionByParticipantResult struct {
	types.Structure
	PrincipalID *types.PID
	Session     *MatchmakeSession
}

// WriteTo writes the FindMatchmakeSessionByParticipantResult to the given writable
func (fmsbpr *FindMatchmakeSessionByParticipantResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fmsbpr.PrincipalID.WriteTo(contentWritable)
	fmsbpr.Session.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fmsbpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FindMatchmakeSessionByParticipantResult from the given readable
func (fmsbpr *FindMatchmakeSessionByParticipantResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = fmsbpr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult header. %s", err.Error())
	}

	err = fmsbpr.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.PrincipalID. %s", err.Error())
	}

	err = fmsbpr.Session.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantResult.Session. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantResult
func (fmsbpr *FindMatchmakeSessionByParticipantResult) Copy() types.RVType {
	copied := NewFindMatchmakeSessionByParticipantResult()

	copied.StructureVersion = fmsbpr.StructureVersion
	copied.PrincipalID = fmsbpr.PrincipalID.Copy().(*types.PID)
	copied.Session = fmsbpr.Session.Copy().(*MatchmakeSession)

	return copied
}

// Equals checks if the given FindMatchmakeSessionByParticipantResult contains the same data as the current FindMatchmakeSessionByParticipantResult
func (fmsbpr *FindMatchmakeSessionByParticipantResult) Equals(o types.RVType) bool {
	if _, ok := o.(*FindMatchmakeSessionByParticipantResult); !ok {
		return false
	}

	other := o.(*FindMatchmakeSessionByParticipantResult)

	if fmsbpr.StructureVersion != other.StructureVersion {
		return false
	}

	if !fmsbpr.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	return fmsbpr.Session.Equals(other.Session)
}

// String returns the string representation of the FindMatchmakeSessionByParticipantResult
func (fmsbpr *FindMatchmakeSessionByParticipantResult) String() string {
	return fmsbpr.FormatToString(0)
}

// FormatToString pretty-prints the FindMatchmakeSessionByParticipantResult using the provided indentation level
func (fmsbpr *FindMatchmakeSessionByParticipantResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FindMatchmakeSessionByParticipantResult{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, fmsbpr.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sSession: %s,\n", indentationValues, fmsbpr.Session.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFindMatchmakeSessionByParticipantResult returns a new FindMatchmakeSessionByParticipantResult
func NewFindMatchmakeSessionByParticipantResult() *FindMatchmakeSessionByParticipantResult {
	fmsbpr := &FindMatchmakeSessionByParticipantResult{
		PrincipalID: types.NewPID(0),
		Session:     NewMatchmakeSession(),
	}

	return fmsbpr
}
