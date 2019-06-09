package bytes

import (
    "errors"
    "fmt"
    //"github.com/sharpvik/scoops/Package/Util"
)


func (interpreter *Interpreter) Execute() {
    for interpreter.running && interpreter.err == nil {
        fmt.Println("It's all good.")
        interpreter.err = errors.New("Just to show off.")
    }
}
