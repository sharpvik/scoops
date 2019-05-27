import subprocess as sp

DIRS = [
    "github.com/sharpvik/scoops/",
    "github.com/sharpvik/scoops/Package/Assembly/",
    "github.com/sharpvik/scoops/Package/Bytes/",
]
OUTFILE = "TESTS.txt"

def test():
    file = open(OUTFILE, 'w')
    for dir in DIRS:
        output = sp.check_output(['go', 'test', dir]).decode('utf-8')
        file.write(output)
    
if __name__ == "__main__":
    test()
    