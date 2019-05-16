package main

import (
    "fmt"
    "os"
    "errors"
    "regexp"
    "github.com/sharpvik/scoops/Package/Import"
)





// FUNCTIONS
func TypeOfArg(arg string) (string, error) {
    validFlag := regexp.MustCompile(`^-[ace]$`)
    validFilename := regexp.MustCompile(`^(.+\.scp[ab]?$)|(help)$`)
    
    if validFlag.MatchString(arg) {
        return "flag", nil
    } else if validFilename.MatchString(arg) {
        return "filename", nil
    } else {
        return "", errors.New("Invalid command line argument: '" + arg + "'")
    }
}


func ParseArgs(args []string) (byte, string, error) {
    argc := len(args)
    if argc > 2 {
        return 0, "", errors.New("Too many command line arguments.")
    }
    
    var (
        flag byte
        filename string
    )
    for _, arg := range args {
        argType, err := TypeOfArg(arg)
        if err != nil {
            return 0, "", err
        }
        
        if argType == "flag" {
            flag = arg[1]
        } else {
            filename = arg
        }
    }
    
    if filename == "" {
        return 0, "", errors.New("Filename not specified.")
    }
    
    return flag, filename, nil
}





func main() {
    // Processing command line arguments...
    flag, filename, err := ParseArgs(os.Args[1:])
    if err != nil {
        util.Err(err)
        fmt.Println(`
│ To use embedded helper:
└──── scoops help
        `)
        os.Exit(1)
    }
    fmt.Printf("Flag: %c\nFilename: %s\n", flag, filename)
    
    // Checking for command line keywords...
    switch filename {
    case "help":
        util.Help()
    }
    
    
}
