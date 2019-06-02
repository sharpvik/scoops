# Python3 required

def opcodes_pars_template_fill(file):
    file.write(
        "package assembly\n\n" +
        "import \"github.com/sharpvik/scoops/Package/Shared\"\n\n" +
        "// Based on ../Shared/opcodes.go/const(opcodes)\n" +
        "var OpcodeMap = map[string]uint8{\n"
    )

def update():
    opcodes_main = open("../Package/Shared/opcodes.go", 'r')
    opcodes_pars = open("../Package/Assembly/opcodes.go", 'w')
    opcodes_pars_template_fill(opcodes_pars)
    
    AWAIT = -1
    USEIT = 0
    DONE  = 1
    STATE = AWAIT
    
    while STATE != DONE:
        line = opcodes_main.readline()
        
        if '(' in line:
            STATE = USEIT
            continue
        if ')' in line:
            STATE = DONE
            opcodes_pars.write('}\n')
            
        if STATE == USEIT:
            opcode = line.split()[0]
            opcodes_pars.write(f'\t"{opcode}": shared.{opcode},\n')
    
    opcodes_main.close()
    opcodes_pars.close()

if __name__ == "__main__":
    update()
