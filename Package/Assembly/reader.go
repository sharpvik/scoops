package assembly

import (
    "bufio"
    "io"
    "os"
)

func Read(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    
    rdr := bufio.NewReader(file)
    var assemblyСode []string
    for {
        line, err := rdr.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                assemblyСode = append(assemblyСode, line)
                break
            }
            return nil, err
        }
        assemblyСode = append(assemblyСode, line)
    }
    
    return assemblyСode, nil
}
