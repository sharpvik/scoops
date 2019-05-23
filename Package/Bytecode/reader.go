package bytecode

import "io/ioutil"



// FUNCTIONS
func Read(filename string) ([]uint8, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, err
}

