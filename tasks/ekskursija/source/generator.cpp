#include "testlib.h"
#include <bits/stdc++.h>
using namespace std;
using ll = long long;
using ii = pair<ll,ll>;

void write_test(int test) {
    startTest(test);

    int n = rnd.next(2, 5);
    int m = rnd.next(1,((n-1)*n)/2); // atempts at creating an edge
    
    set<ii> edges;
    for(int i=0;i<m;i++) {
        int a = rnd.next(1,n-1);
        int b = rnd.next(a+1,n);
        if(edges.count({a,b})) continue;
        edges.insert({a,b});
    }

    vector<int> id = rnd.perm(n);

    println(n, edges.size());
    for(ii edge : edges){
        int a = id[edge.first-1]+1;
        int b = id[edge.second-1]+1;
        int c = rnd.next(1,(int)n*n);
        println(a,b,c);
    }
}

int main(int argc, char* argv[]) {
    registerGen(argc, argv, 1);

    for(int i=1;i<=3;i++) {
        write_test(i);
    }
}
    // vector<int> perm = rnd.perm(n);
    // shuffle(edges.begin(), edges.end());