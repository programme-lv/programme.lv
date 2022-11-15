#include "testlib.h"

using namespace std;
using ll = long long;

void write_test(int test) {
    startTest(test);
    ll a, b;
    if(test<=5) { // subtask 1
        a = 1;
        b = rnd.next((ll)1, (ll)1e3);
    }
    else if(test<=10) { // subtask 2
        a = rnd.next((ll)1, (ll)1e3);
        b = rnd.next((ll)a, (ll)1e5);
    }
    else if(test<=15) { // subtask 3
        a = rnd.next((ll)1,(ll)1e16-1005);
        b = rnd.next((ll)a, (ll)a+1000);
    }
    else if(test<=20) { // subtask 4
        a = rnd.next((ll)1,(ll)1e7);
        b = rnd.next((ll)a, (ll)1e9);
    }
    else if(test<=25) { // subtask 5
        a = rnd.next((ll)1,(ll)1e8);
        b = rnd.next((ll)a, (ll)1e11);
    }
    else if(test<=30) { // subtask 6
        a = rnd.next((ll)1,(ll)1e12);
        b = rnd.next((ll)a, (ll)1e14);
    }
    println(a,b);
}

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);

    for(int i=1;i<=30;i++)
        write_test(i);
}