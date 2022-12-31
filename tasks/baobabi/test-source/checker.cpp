#include "testlib.h"

int main(int argc, char * argv[]) {
    setName("compares two signed integers");
    registerTestlibCmd(argc, argv);
    std::string ja = ans.readWord();
    std::string pa = ouf.readWord();
    if (ja != pa)
        quitf(_wa, "expected %s, found %s", ja.c_str(), pa.c_str());
    quitf(_ok, "answer is %s", ja.c_str());
}