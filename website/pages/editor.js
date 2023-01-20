import NavBar from "../components/navbar";
import Editor from "@monaco-editor/react";

export default function EditorPage() {
    return (
        <div className="d-flex flex-column vh-100">
            <NavBar active_page={"editor"}/>
            <div className={"container flex-grow-1 my-3"}>
                <Editor
                    height="100%"
                    defaultLanguage="cpp"
                    defaultValue={defaultCPPCode}
                />
            </div>
        </div>
    )
}

let defaultCPPCode = `#include <iostream>

using namespace std;
using ll = long long;

ll binpow(ll b, ll e) {
    if (e == 0)
        return 1;
    ll res = binpow(b, e / 2);
    if (e % 2) return res * res * b;
    else return res * res;
}

int main() {
    cout<<binpow(3,4)<<endl;
}
`