package question

import (
	"errors"
	"fmt"
	"strings"

	resourcerecord "github.com/arnaldomf/dns/domain/dns/resource_record"
	"github.com/arnaldomf/dns/utils/bytesutil"
)

const (
	nullChar byte = 0
)

var (
	ErrQuestionIncomplete = errors.New("question section is incomplete")
)

type Question struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

// New creates a Question from pkg, starting reading it from index = offset
func New(pkg []byte, offset int) (*Question, error) {
	q := &Question{}
	qtypeOffset := q.SetQName(pkg, offset)
	if qtypeOffset == -1 {
		return nil, ErrQuestionIncomplete
	}
	if len(pkg[qtypeOffset:]) < 4 {
		return nil, ErrQuestionIncomplete
	}
	q.QType = bytesutil.ToUInt16(pkg[qtypeOffset : qtypeOffset+2])
	q.QClass = bytesutil.ToUInt16(pkg[qtypeOffset+2 : qtypeOffset+4])
	return q, nil
}

func (q *Question) SetQName(pkg []byte, offset int) int {
	for i := offset; i < len(pkg); i++ {
		if pkg[i] == nullChar {
			return i + 1
		}
		q.QName = append(q.QName, pkg[i])
	}

	return -1
}

func (q *Question) DecodedName() string {
	var sb strings.Builder

	var nextPart byte

	for i, c := range q.QName {
		char := c
		if i == int(nextPart) {
			nextPart = c + byte(i+1)
			if i == 0 {
				continue
			} else {
				char = '.'
			}
		}
		sb.WriteByte(char)
	}
	return sb.String()
}

func (q *Question) String() string {
	var sb strings.Builder

	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf("QName=%s, ", q.DecodedName()))
	sb.WriteString(fmt.Sprintf("QType=%s, ", resourcerecord.Type(q.QType)))
	sb.WriteString(fmt.Sprintf("QClass=%s", resourcerecord.Class(q.QClass)))
	sb.WriteRune('}')

	return sb.String()
}
