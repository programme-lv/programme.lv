#include <iostream>
#include <vector>
#include <cstring>
#include <cmath>

using namespace std;
using ll = long long;
using u64 = uint64_t;
using u128 = __uint128_t;

bool prime[(ll)1e7+1];

void get_primes(vector<ll>& primes, ll upper_bound) {
    memset(prime,1,sizeof(prime));
    for(ll i=2;i*i<=upper_bound;i++) {
        if(!prime[i]) continue;
        for(ll j=i*i;j<=upper_bound;j+=i) {
            prime[j] = false;
        }
    }
    for(ll i=2;i<=upper_bound;i++)
        if(prime[i])
            primes.push_back(i);
}


bool is_prime(ll x) {
    if(x<=1) return false;
    for(ll i=2;i*i<=x;i++)
        if(x%i==0)
            return false;
    return true;
}

u128 binpower(u128 base, u128 e) {
    u128 result = 1;
    while (e) {
        if (e & 1)
            result = (u128)result * base;
        base = (u128)base * base;
        e >>= 1;
    }
    return result;
}

u64 binpower(u64 base, u64 e, u64 mod) {
    u64 result = 1;
    base %= mod;
    while (e) {
        if (e & 1)
            result = (u128)result * base % mod;
        base = (u128)base * base % mod;
        e >>= 1;
    }
    return result;
}

bool check_composite(u64 n, u64 a, u64 d, int s) {
    u64 x = binpower(a, d, n);
    if (x == 1 || x == n - 1)
        return false;
    for (int r = 1; r < s; r++) {
        x = (u128)x * x % n;
        if (x == n - 1)
            return false;
    }
    return true;
};

bool MillerRabin(u64 n) { // returns true if n is prime, else returns false.
    if (n < 2)
        return false;

    int r = 0;
    u64 d = n - 1;
    while ((d & 1) == 0) {
        d >>= 1;
        r++;
    }

    for (int a : {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}) {
        if (n == a)
            return true;
        if (check_composite(n, a, d, r))
            return false;
    }
    return true;
}

int main() {
    ll a, b;
    cin>>a>>b;
    
    ll res = 0;
    if(a<=2&&b>=2)
        res = 1; // add 2 to baobabs
    vector<ll> primes;
    primes.reserve(7e5); // there are 664579 primes in [1;1e7]
    get_primes(primes, 1e7);
    for(ll p=2;p<40;p+=2) {
        ll j = 0;
        while(binpower(primes[j],p)<a) j++;
        ll iters = 0;
        while(binpower(primes[j],p)<=b) {
            iters++;
            //cout<<primes[j]<<" "<<p<<endl;
            ll x = primes[j];
            ll summa = (binpower(primes[j],p+1)-1)/(primes[j]-1);
            if(MillerRabin(summa))
                res++;
            j++;
        }
        //if(iters==0) break;
    }
    cout<<res<<endl;
}