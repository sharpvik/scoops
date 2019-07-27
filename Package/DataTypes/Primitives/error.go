package primitives

import (
    "bufio"
    "errors"
    "fmt"
)



type Error struct {
    Sort  string
    Value error
}


func NewError(_type string, value error) *Error {
    return &Error{_type, value}
}


func (e *Error) Clone() *Error {
    return NewError(e.Sort, e.Value)
}


func (e *Error) Print(w *bufio.Writer) {
    w.WriteString(
        e.ToGoString(),
    )
}


func (e *Error) ToGoString() string {
    return fmt.Sprintf("%s: %s", e.Sort, e.Value)
}


func (e *Error) ToGoError() error {
    return errors.New( e.ToGoString() )
}


func (e *Error) Type() string {
    return "error"
}