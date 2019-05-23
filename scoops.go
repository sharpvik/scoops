package main

import (
    "fmt"
    "os"
    "errors"
    "regexp"
    "github.com/sharpvik/scoops/Package/Util"
    "github.com/sharpvik/scoops/Package/Bytecode"
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


func GetFileExtention(filename string) string {
    for i, char := range filename {
        if char == '.' {
            return filename[i + 1:]
        }
    }
    return ""
}





func main() {
    // Processing command line arguments...
    flag, filename, err := ParseArgs(os.Args[1:])
    if err != nil {
        util.Error(err)
        fmt.Println(`
│ To use embedded helper:
└──── scoops help
        `)
        os.Exit(1)
    }
    util.Log( "Flag: " + string(flag) )
    util.Log("Filename: " + filename)

    // Checking for command line keywords...
    switch filename {
    case "help":
        util.Help()

    default:
        fileExtention := GetFileExtention(filename)
        switch fileExtention {
        case "scp":
            fmt.Println("This file format is not yet supported. Sorry.")

        case "scpa":
            fmt.Println("This file format is not yet supported. Sorry.")

        case "scpb":
            if flag != 0 && flag != 'e' {
                util.Warning(
                    "Input file is of *.scpb format. Bytecode files " +
                    "can only be executed. Omitting the flag...",
                )
            }
            data, err := bytecode.Read(filename)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            fmt.Println(data)
        }
    }
}

