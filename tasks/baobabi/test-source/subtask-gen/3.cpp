#include "testlib.h"

using namespace std;
using ll = long long;

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);
    ll a, b;

    // subtask 3
    a = rnd.next((ll)1,(ll)1e16-1005);
    b = rnd.next((ll)a, (ll)a+1000);

    println(a,b);
}