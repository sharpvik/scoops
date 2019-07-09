package slice

import (
    "fmt"
    "github.com/sharpvik/scoops/Package/Shared"
)



type Slice struct {
    size        uint64
    value       []shared.Object
}


func New() *Slice {
    return &Slice{0, nil}
}


func (s *Slice) Print() {
    fmt.Print("[ ")
    for _, o := range s.value {
        o.Print()
    }
    fmt.Print(" ]")
}


func (s *Slice) Type() string {
    return "slice"
}


func (s *Slice) Size() uint64 {
    return s.size
}


func (s *Slice) Append(item shared.Object) {
    s.value = append(s.value, item)
}


func (s *Slice) Pop(index uint64) {
    tmp := s.value[:index]
    len := uint64( len(s.value) )
    if index < len - 1 {
        for i := index + 1; i < len; i++ {
            tmp = append(tmp, s.value[i])
        }
    }
    s.value = tmp
}
