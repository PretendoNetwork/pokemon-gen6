// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemStartChallengeParam is a type within the ServiceItem protocol
type ServiceItemStartChallengeParam struct {
	types.Structure
	ChallengeScheduleID *types.PrimitiveU32
	TicketType          *types.PrimitiveU32
	NumTicket           *types.PrimitiveU32
}

// WriteTo writes the ServiceItemStartChallengeParam to the given writable
func (siscp *ServiceItemStartChallengeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siscp.ChallengeScheduleID.WriteTo(contentWritable)
	siscp.TicketType.WriteTo(contentWritable)
	siscp.NumTicket.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siscp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemStartChallengeParam from the given readable
func (siscp *ServiceItemStartChallengeParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = siscp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam header. %s", err.Error())
	}

	err = siscp.ChallengeScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.ChallengeScheduleID. %s", err.Error())
	}

	err = siscp.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.TicketType. %s", err.Error())
	}

	err = siscp.NumTicket.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.NumTicket. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemStartChallengeParam
func (siscp *ServiceItemStartChallengeParam) Copy() types.RVType {
	copied := NewServiceItemStartChallengeParam()

	copied.StructureVersion = siscp.StructureVersion
	copied.ChallengeScheduleID = siscp.ChallengeScheduleID.Copy().(*types.PrimitiveU32)
	copied.TicketType = siscp.TicketType.Copy().(*types.PrimitiveU32)
	copied.NumTicket = siscp.NumTicket.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the given ServiceItemStartChallengeParam contains the same data as the current ServiceItemStartChallengeParam
func (siscp *ServiceItemStartChallengeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemStartChallengeParam); !ok {
		return false
	}

	other := o.(*ServiceItemStartChallengeParam)

	if siscp.StructureVersion != other.StructureVersion {
		return false
	}

	if !siscp.ChallengeScheduleID.Equals(other.ChallengeScheduleID) {
		return false
	}

	if !siscp.TicketType.Equals(other.TicketType) {
		return false
	}

	return siscp.NumTicket.Equals(other.NumTicket)
}

// String returns the string representation of the ServiceItemStartChallengeParam
func (siscp *ServiceItemStartChallengeParam) String() string {
	return siscp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemStartChallengeParam using the provided indentation level
func (siscp *ServiceItemStartChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemStartChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %s,\n", indentationValues, siscp.ChallengeScheduleID))
	b.WriteString(fmt.Sprintf("%sTicketType: %s,\n", indentationValues, siscp.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %s,\n", indentationValues, siscp.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemStartChallengeParam returns a new ServiceItemStartChallengeParam
func NewServiceItemStartChallengeParam() *ServiceItemStartChallengeParam {
	siscp := &ServiceItemStartChallengeParam{
		ChallengeScheduleID: types.NewPrimitiveU32(0),
		TicketType:          types.NewPrimitiveU32(0),
		NumTicket:           types.NewPrimitiveU32(0),
	}

	return siscp
}
