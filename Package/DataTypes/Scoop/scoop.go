package scoop

import (
	"bufio"
	"fmt"
	"github.com/sharpvik/scoops/Package/Bytes"
)



type Scoop struct {
	name string
	size uint64
	code []*bytes.Instruction
}


func New(name string, code []*bytes.Instruction) *Scoop {
	return &Scoop{name, uint64( len(code) ), code}
}


func (s *Scoop) Clone() *Scoop {
	return New(s.name, s.code)
}


func (s *Scoop) Print(w *bufio.Writer) {
	w.WriteString( 
		fmt.Sprintf("[ Scoop %s at %v ]", s.name, s),
	)
}


func (s *Scoop) Type() string {
	return "scoop"
}