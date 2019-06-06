# Scoops Assembly Specs

*Scoops Assembly* is a language based on simple and clear instructions that are
meant to be executed by the *Scoops Interpreter*. In this file, you can find all
the specs that you should know if you wish to

1. Contribute to this project
2. Create your own *Scoops* compliant compiler or interpreter
3. Branch this repo and modify some core components

## Instruction format

Assembly instruction consists of two parts:

1. Opcode
2. Operand(s)

Regular expression that matches a valid *Scoops Assembly* instruction:

```
^[A-Z_]+( b[01]+| '.'| -?\d+\.\d+| x[\dA-F]+| -?\d+| ".+")*\s*$
```

## Opcodes

*Scoops Interpreter (SI)* can support up to 256 different opcodes as it is a
[bytecode] based [virtual machine]. Each opcode is a simple command that *SI*
can understand and execute. An example of a trivial opcode would be the
`THE_END` opcode that tells *SI* that program has come to an end and it is time
to stop execution.

Some opcodes may require one or more operand(s) -- additional parameter(s)
needed to make *SI* more flexible. For example, you may want to `MAKE_INT 42`.
This opcode tells *SI* to push 42 into the *chamber* ([queue]) of the currently
active execution environment as 8 bytes and then use those bytes to create an
instance of *Integer* and push it onto the *data stack* ([stack]).

You can find the list of all the opcodes used by the *Scoops Interpreter* 
[in this file]. Regular expression that matches a valid *SI* opcode is as
follows: `[A-Z_]+`

[bytecode]: https://en.wikipedia.org/wiki/Bytecode
[virtual machine]: https://en.wikipedia.org/wiki/Virtual_machine
[queue]: ../DataTypes/Queue/README.md
[stack]: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
[in this file]: ../Shared/opcodes.go

## Operands

There are 4 types of operands that *Scoops Assembler (SA)* accepts as 
syntactically valid:

| №  | Operand Type | Definitions               | Regular Expression | Example |
|---:|:-------------|:--------------------------|:-------------------|:--------|
| 1. | Boolean      | 64-bit long boolean       | `b[01]+`           | b101010 |
| 2. | Rune         | UTF-8 compliant character | `'.'`              | '🍨'    |
| 3. | Float        | 64 bit long float         | `-?\d+\.\d+`       | 3.14159 |
| 4. | Hexadecimal  | 64 bit long hexadecimal   | `x[\dA-F]+`        | x2A     |
| 4. | Integer      | 64 bit long integer       | `-?\d+`            | -42     |
| 5. | String       | UTF-8 compliant string    | `".+"`             | "I❤️U"  |

It is important to note that booleans, floating point numbers, hexadecimals and 
integers **must** be so that *SI* will be able to store them in 64 bits. For 
example, all integers must fall in range of **±9 223 372 036 854 775 807**.

All opcodes require **specific** number of operands to be served with them. You 
can find these numbers [in here].

[in here]: ../Shared/operands.go