#include "testlib.h"

using namespace std;
using ll = long long;

void write_test(int test) {
    startTest(test);
    ll a, b;
    if(test<=10) { // subtask 1
        a = rnd.next((ll)1, (ll)1e2-1);
        b = rnd.next((ll)a+1, (ll)1e2);
    }
    else if(test<=15) { // subtask 2
        a = rnd.next((ll)1e2, (ll)1e3-1);
        b = rnd.next((ll)a+1, (ll)1e3);
    }
    else if(test<=20) { // subtask 3
        a = rnd.next((ll)1e4, (ll)1e6-1);
        b = rnd.next((ll)a+1, (ll)1e6);
    }
    println(a,b);
}

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);

    for(int i=1;i<=20;i++)
        write_test(i);
}