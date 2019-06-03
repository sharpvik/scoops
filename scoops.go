package main

import (
    "errors"
    "fmt"
    "github.com/sharpvik/scoops/Package/Util"
    "github.com/sharpvik/scoops/Package/Assembly"
    "github.com/sharpvik/scoops/Package/Bytes"
    "os"
    "regexp"
)



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


func UsageHint() {
    fmt.Println(`
│ To use embedded helper:
└──── scoops help
    `)
}



func main() {
    // Processing command line arguments...
    flag, filename, err := ParseArgs(os.Args[1:])
    if err != nil {
        util.Error(err)
        UsageHint()
        os.Exit(1)
    }
    util.Log( "Flag: " + string(flag) )
    util.Log("Filename: " + filename)

    fileExtention := GetFileExtention(filename)
    var (
        //sourceСode []string
        assemblyСode []string
        byteCode []*bytes.Instruction
    )
    
    // Use filename extention to determine execution process and catch errors...
    switch fileExtention {
    case "scp":
        util.Error(
            errors.New("This file format is not yet supported. Sorry."),
        )
        os.Exit(1)
        // The 4 lines above will be removed upon completion of compiler.
        // They will be replaced with the code snippet below.
        /*
        sourceCode, err = source.Read(filename)
        if err != nil {
            util.Error(err)
            os.Exit(1)
        }
        assemblyCode = source.Compile(sourceCode)
        if flag == 'a' {
            assembly.Write(assemblyCode)
            os.Exit(0)
        }
        */
        fallthrough

    case "scpa":
        if flag == 'a' {
            util.Error(
                errors.New("Invalid flag for *.scpa input file."),
            )
            os.Exit(1)
        }
        if assemblyСode == nil {
            assemblyСode, err = assembly.Read(filename)
            if err != nil {
                util.Error(err)
                os.Exit(1)
            }
        }
        byteCode, err = assembly.Assemble(assemblyСode)
        if err != nil {
            util.Error(err)
            os.Exit(1)
        }
        if flag == 'c' {
            bytes.Write(byteCode)
            os.Exit(0)
        }
        fallthrough

    case "scpb":
        if flag != 'e' && flag != 0 {
            util.Error(
                errors.New("Invalid flag for *.scpb input file."),
            )
            os.Exit(1)
        }
        if byteCode == nil {
            byteCode, err = bytes.Read(filename)
            if err != nil {
                util.Error(err)
                os.Exit(1)
            }
        }
        interpreter := bytes.NewInterpreter(byteCode)
        interpreter.Execute()
    
    case "":
        // If filename doesn't have extention, it is either a command line
        // keyword (like "help"), or an invalid input file.
        switch filename {
        case "help":
            util.Help()
            
        default:
            util.Error( errors.New("Input file does not have exntention.") )
            UsageHint()
            os.Exit(1)
        }
        
    default:
        // Any other filename extentions are to be considered invalid.
        util.Error( errors.New("Input file has an invalid exntention.") )
        UsageHint()
        os.Exit(1)
    }
}
