# Python3 required

def opcodes_pars_template_fill(file):
    file.write(
        "package assembly\n\n" +
        "import \"github.com/sharpvik/scoops/Package/Shared\"\n\n" +
        "// Based on scoops/Package/Shared/opcodes.go/const(opcodes)\n" +
        "var OpcodeMap = map[string]uint8{\n"
    )
    
def operands_template_fill(file):
    file.write(
        "package shared\n" +
        "\n" +
        "var NumberOfOperands = [256]uint32{\n"
    )

def update():
    opcodes_main = open("../Package/Shared/opcodes.go", 'r')
    opcodes_pars = open("../Package/Assembly/opcodes.go", 'w')
    operands     = open("../Package/Shared/operands.go", 'w')
    opcodes_pars_template_fill(opcodes_pars)
    operands_template_fill(operands)
    
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
            operands.write('}\n')
            
        if STATE == USEIT:
            line_split = line.split()
            opcode = line_split[0]
            number_of_operands = line_split[-2]
            opcodes_pars.write(f'    "{opcode}": shared.{opcode},\n')
            operands_paste = f'    /* {opcode} */'
            operands.write( 
                operands_paste + 
                ( 32 - len(operands_paste) ) * ' ' +
                f'{number_of_operands},\n'
            )
    
    opcodes_main.close()
    opcodes_pars.close()

if __name__ == "__main__":
    update()
