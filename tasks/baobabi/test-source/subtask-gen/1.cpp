#include "testlib.h"

using namespace std;
using ll = long long;

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);
    ll a, b;

    // subtask 1
    a = 1;
    b = rnd.next((ll)1, (ll)1e3);
    
    println(a,b);
}