#include "testlib.h"

using namespace std;
using ll = long long;

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);
    ll a, b;

    // subtask 5
    a = rnd.next((ll)1,(ll)1e8);
    b = rnd.next((ll)a, (ll)1e11);

    println(a,b);
}