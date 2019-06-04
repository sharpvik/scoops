# Python3 required

import subprocess as sp

DIRS = [
    "github.com/sharpvik/scoops/",
    "github.com/sharpvik/scoops/Package/Assembly/",
    "github.com/sharpvik/scoops/Package/Bytes/",
    "github.com/sharpvik/scoops/Package/DataTypes/Queue/",
]
OUTF = "../TESTS.txt"

def test():
    file = open(OUTF, 'wb')
    for dir in DIRS:
        output = sp.check_output(['go', 'test', dir])
        file.write(output)
    file.close()
    
if __name__ == "__main__":
    test()
    