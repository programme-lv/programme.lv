import CodeMirror from '@uiw/react-codemirror';
import { sublime } from '@uiw/codemirror-theme-sublime'
import { javascript } from '@codemirror/lang-javascript'

export default function Markdown(){
    return (
      <CodeMirror
        value="console.log('hello world!');"
        height="100px"
        theme={sublime}
        extensions={[javascript({ jsx: true })]}
        onChange={(value, viewUpdate) => {
          console.log('value:', value);
        }}
      />
    );
}