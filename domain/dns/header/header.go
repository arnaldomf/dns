package header

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arnaldomf/dns/utils/bytesutil"
)

const (
	// HeaderLength is the number of bytes of a DNS Header
	HeaderLength = 12
)

var (
	ErrHeaderLength = errors.New("unexpected header length")
)

// Header represets the DNS message header and all its fields
// represented as uint16 for easier bit manipulation as each
// field of the DNS header as defined in rfc1035 is 2 bytes
// long.
type Header struct {
	// ID is a 2 bytes field that represents the request ID
	// as set by the DNS client.
	ID    uint16
	Flags uint16
	// QDCount represents the number of entries in the question
	// section
	QDCount uint16
	// ANCount represents the number of resource records in the
	// answer section
	ANCount uint16
	// NSCount specifies the number of name server resource
	// records in the authority records section
	NSCount uint16
	// ARCount specifies the number of resource records in the
	// additional records section
	ARCount uint16
}

// New receives the whole package as a byte slice and returns
// a Header pointer or an error
func New(pkg []byte) (*Header, error) {
	pkgLen := len(pkg)
	if pkgLen < HeaderLength {
		return nil, fmt.Errorf("%w: %d", ErrHeaderLength, pkgLen)
	}
	return &Header{
		ID:      bytesutil.ToUInt16(pkg[0:2]),
		Flags:   bytesutil.ToUInt16(pkg[2:4]),
		QDCount: bytesutil.ToUInt16(pkg[4:6]),
		ANCount: bytesutil.ToUInt16(pkg[6:8]),
		NSCount: bytesutil.ToUInt16(pkg[8:10]),
		ARCount: bytesutil.ToUInt16(pkg[10:12]),
	}, nil
}

// QR returns the first bit of Flags which specifies if a message is
// a query (0) or a response (1)
func (h *Header) QR() byte {
	return byte(h.Flags >> 15)
}

// SetQR sets the QR bit to 1
func (h *Header) SetQR() {
	h.Flags = h.Flags | 0x8000
}

// ClearQR sets the QR bit to 0
func (h *Header) ClearQR() {
	h.Flags = h.Flags & 0x7fff
}

// OPCode is a 4 bit field from Flags that specifies the kind of Query
func (h *Header) OPCode() byte {
	return byte(h.Flags>>11) & 0x0f
}

// AA Authoritative Answer - this bit is valid in responses,
// and specifies that the responding name server is an
// authority for the domain name in question section.
// Note that the contents of the answer section may have
// multiple owner names because of aliases. The AA bit
// corresponds to the name which matches the query name, or
// the first owner name in the answer section.
func (h *Header) AA() byte {
	return byte(h.Flags>>10) & 0x01
}

// SetAA sets the AA bit to 1
func (h *Header) SetAA() {
	h.Flags = h.Flags | 0x0400
}

// ClearAA clears the AA bit
func (h *Header) ClearAA() {
	h.Flags = h.Flags & 0xfbff
}

// TC bit specifies that the message was truncated due to
// length greater than permitted on the transmission channel.
func (h *Header) TC() byte {
	return byte(h.Flags>>9) & 0x01
}

// SetTC sets the TC bit to 1
func (h *Header) SetTC() {
	h.Flags = h.Flags | 0x0200
}

// ClearTC clears the TC bit
func (h *Header) ClearTC() {
	h.Flags = h.Flags & 0xfdff
}

// RD bit specifies if recursion is desired (1) or not (0)
func (h *Header) RD() byte {
	return byte(h.Flags>>8) & 0x01
}

// SetRD sets RD bit to 1
func (h *Header) SetRD() {
	h.Flags = h.Flags | 0x0100
}

// ClearRD clears the RD bit
func (h *Header) ClearRD() {
	h.Flags = h.Flags & 0xfe11
}

// RA recusion is available (1) or not (0)
func (h *Header) RA() byte {
	return byte(h.Flags>>7) & 0x01
}

// SetRA sets RA bit to 1
func (h *Header) SetRA() {
	h.Flags = h.Flags | 0x0080
}

// ClarRA RA bit
func (h *Header) ClearRA() {
	h.Flags = h.Flags & 0xff7f
}

// DA bit specifies if data is authenticated
func (h *Header) DA() byte {
	return byte(h.Flags>>5) & 0x01
}

// SetDA bit to 1
func (h *Header) SetDA() {
	h.Flags = h.Flags | 0x0020
}

// ClearDA bit
func (h *Header) ClearDA() {
	h.Flags = h.Flags & 0xffdf
}

// CD check disabled bit
func (h *Header) CD() byte {
	return byte(h.Flags>>4) & 0x01
}

// SetCD sets CD bit to 1
func (h *Header) SetCD() {
	h.Flags = h.Flags | 0x0010
}

// ClearCD clears CD bit
func (h *Header) ClearCD() {
	h.Flags = h.Flags & 0xffef
}

// RCODE four bit field part of the response
func (h *Header) RCODE() byte {
	return byte(h.Flags) & 0x0f
}

// SetRCODE to rcode
func (h *Header) SetRCODE(rcode byte) {
	h.Flags = h.Flags | uint16(rcode&0x0f)
}

func (h *Header) String() string {
	var sb strings.Builder

	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf("ID=%d, ", h.ID))
	sb.WriteString(fmt.Sprintf("QR=%q, ", QR(h.QR())))
	sb.WriteString(fmt.Sprintf("OPCode=%q, ", OPCode(h.OPCode())))
	sb.WriteString(fmt.Sprintf("AA=%d, ", h.AA()))
	sb.WriteString(fmt.Sprintf("TC=%d, ", h.TC()))
	sb.WriteString(fmt.Sprintf("RD=%d, ", h.RD()))
	sb.WriteString(fmt.Sprintf("RA=%d, ", h.RA()))
	sb.WriteString(fmt.Sprintf("DA=%d, ", h.DA()))
	sb.WriteString(fmt.Sprintf("CD=%d, ", h.CD()))
	sb.WriteString(fmt.Sprintf("RCODE=%q, ", RCode(h.RCODE())))
	sb.WriteString(fmt.Sprintf("QDCount=%d, ", h.QDCount))
	sb.WriteString(fmt.Sprintf("ANCount=%d, ", h.ANCount))
	sb.WriteString(fmt.Sprintf("NSCount=%d, ", h.NSCount))
	sb.WriteString(fmt.Sprintf("ARCount=%d", h.ARCount))
	sb.WriteRune('}')

	return sb.String()
}
