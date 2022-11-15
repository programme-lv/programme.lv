#include <iostream>

using namespace std;
using ll = long long;

// true if square with center x, y is inside r
bool inside(ll x, ll y, ll r) { 
    ll vertices[4][2] = {{x+1,y+1},{x+1,y-1},{x-1,y-1},{x-1,y+1}};
    for(ll i=0;i<4;i++) {
        ll d = vertices[i][0]*vertices[i][0]+vertices[i][1]*vertices[i][1];
        if(d>r*r)
            return false;
    }
    return true;
}

// true if square with cetner x,y is outside r
bool outside(ll x, ll y, ll r) { 
    ll vertices[4][2] = {{x+1,y+1},{x+1,y-1},{x-1,y-1},{x-1,y+1}};
    for(ll i=0;i<4;i++) {
        ll d = vertices[i][0]*vertices[i][0]+vertices[i][1]*vertices[i][1];
        if(d<r*r)
            return false;
    }
    return true;
}

int main() {
    ll a, b;
    cin>>a>>b;
    ll y = 1, r = 1;
    while(inside(r+2,y,2*b)) r+=2;
    ll l = r+2;
    ll res = 0;
    for(y = 1; y <= 2*b; y+=2) {
        while(r>=1&&!inside(r,y,2*b)) r-=2;
        while(l-2>=1&&outside(l-2,y,2*a)) l-=2;
        res += (r-l)/2+1;
    }
    cout<<res*4<<endl;
}