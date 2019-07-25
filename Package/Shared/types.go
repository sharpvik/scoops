package shared

/*
 * It is important to note that the 'Object' interface type is satisfied by all
 * of the pointer types in the DataTypes folder.
 *
 * In other words, *primitives.Rune can be retreived from an Object type using
 * type assertion as *primitives.Rune perfectly satisfies the Object interface:
 *     var a Object = interpreter.scope.data.Pop()
 *     var b *primitives.Rune = a.(*primitives.Rune)
 *
 * However, primitives.Rune cannot. It doesn't satisfy the Object interface.
 * Here's why... We declare the Print and Type functions as follows:
 *     func (r *Rune) Print() { ... }
 * so type *Rune satisfies the Object interface but type Rune doesn't.
 * This goes for every type in the DataTypes folder. Similar logic applies to 
 * the Collection interface.
 */
type Object interface {
    Clone() Object
    Print()
    Type() string
}


type Collection interface {
    Size() uint64
}