package bytes

import (
    "errors"
    "fmt"
    "github.com/sharpvik/scoops/Package/Util"
)



func (interpreter *Interpreter) Evaluate() {
    //
}


func (interpreter *Interpreter) Execute() {
    for interpreter.running && interpreter.err == nil {
        interpreter.Evaluate()
    }
    if interpreter.err != nil {
        util.Error(interpreter.err)
        util.Log("Interpreter exited with non-zero return value.")
    }
}
