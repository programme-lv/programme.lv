import Sidebar from '../components/sidebar'
import CodeMirror from '@uiw/react-codemirror';
import { sublime } from '@uiw/codemirror-theme-sublime'
import { cpp } from '@codemirror/lang-cpp'

const hello_cpp = `#include <iostream>

using namespace std;

int main() {
  cout<<"Hello, world!"<<endl;
}
`

var written_code = hello_cpp;

function execute_code() {
  const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ code: written_code })
  };
  fetch('http://localhost:8080/execute/enqueue', requestOptions)
      .then(response => response.json())
}

export default function Editor() {
  return (
    <div className="d-flex">
        <div className="bg-dark">
          <Sidebar/>
        </div>
        <div className="flex-grow-1">
          <div className="h-75">
            <CodeMirror
              className="h-100"
              value={hello_cpp}
              theme={sublime}
              height='100%'
              extensions={[cpp()]}
              onChange={(value) => {
                written_code = value;
              }}
            />
          </div>
          <div className="h-25 bg-dark p-2 pb-3 d-flex">
            <div className='col-4 px-3 d-flex flex-column'>
              <label className="form-label text-white">Standard Input ( stdin )</label>
              <textarea className="form-control flex-grow-1" spellCheck="false"></textarea>
            </div>
            <div className="col-4 px-3 d-flex flex-column">
              <label className="form-label text-white">Expected Output ( stdout )</label>
              <textarea className="form-control flex-grow-1" spellCheck="false"></textarea>
            </div>
            <div className="col-4 px-3">
              <button onClick={execute_code} className="col btn btn-success" type="button">Run</button>
            </div>
          </div>
        </div>
    </div>
  )
}
