from sys import argv



def main(ifile):
    ifile = open(ifile, 'r').read().split()
    ofile = open('generated_scoop.scpa', 'w')

    ofile.write('MAKE_SCOOP 0\n')
    for word in ifile:
        ofile.write(f'SCOOP_PUSH {word}\n')

    ofile.close()



if __name__ == '__main__':
    main(argv[1])
