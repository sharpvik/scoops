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


func (s *Slice) Clone() shared.Object {
    news := New()
    for _, o := range s.value {
        news.Append( o.Clone() )
    }
    return news
}


func (s *Slice) Print() {
    if s.size == 0 {
        fmt.Println("[ ]")
        return
    }
    fmt.Print("[")
    for _, o := range s.value {
        fmt.Print(" ")
        o.Print()
        fmt.Print(" ")
    }
    fmt.Print("]\n")
}


func (s *Slice) Type() string {
    return "slice"
}


func (s *Slice) Size() uint64 {
    return s.size
}


func (s *Slice) Append(item shared.Object) {
    s.value = append(s.value, item)
    s.size++
}


func (s *Slice) GetItemByIndex(index uint64) shared.Object {
    return s.value[index]
}


func (s *Slice) Pop(index uint64) shared.Object {
    obj := s.value[index]
    tmp := s.value[:index]
    len := uint64( len(s.value) )
    if index < len - 1 {
        for i := index + 1; i < len; i++ {
            tmp = append(tmp, s.value[i])
        }
    }
    s.value = tmp
    s.size--
    return obj
}
