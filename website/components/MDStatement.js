import renderMathInElement from "katex/contrib/auto-render";
import "katex/dist/katex.min.css";
import {useEffect} from "react";

export default function MDStatement({mdStatement}) {
    let description = mdStatement["desc"];
    let input = mdStatement["input"];
    let output = mdStatement["output"];
    let examples = mdStatement["examples"];
    let scoring = mdStatement["scoring"];
    let notes = mdStatement["notes"];
    useEffect(() => {
        // RENDER KATEX MATH
        renderMathInElement(document.getElementById("task-statement"), {
            throwOnError: true, delimiters: [{left: '$$', right: '$$', display: true}, {
                left: '$', right: '$', display: false
            }, {left: '\\(', right: '\\)', display: false}, {left: '\\[', right: '\\]', display: true}],
        });
    }, []);
    return (<>
        <section className="my-4">
            <h4 className={"my-3"}>formulējums</h4>
            <div dangerouslySetInnerHTML={{__html: description}}></div>
        </section>
        <section className="my-4">
            <h4 className={"my-3"}>ievaddati</h4>
            <div dangerouslySetInnerHTML={{__html: input}}></div>
        </section>
        <section className="my-4">
            <h4 className={"my-3"}>izvaddati</h4>
            <div dangerouslySetInnerHTML={{__html: output}}></div>
        </section>
        <section className="my-4">
            <h4 className={""}>piemēri</h4>
            <div className={"d-flex"}>
                {examples.map((example, index) =>
                    <ExampleTable key={index} input={example["input"]} output={example["output"]}/>)}
            </div>
        </section>
        {scoring && <section className="my-4 md-statement-scoring">
            <h4 className={"my-3"}>vērtēšana</h4>
            <div dangerouslySetInnerHTML={{__html: scoring}}></div>
        </section>}
        {notes && <section className="my-4">
            <h4 className={"my-3"}>piezīmes</h4>
            <div dangerouslySetInnerHTML={{__html: notes}}></div>
        </section>}
    </>)
}

function ExampleTable({input, output}) {
    return (<table className={"table table-bordered col my-3"}>
        <thead>
        <tr>
            <th>ievaddati</th>
            <th>izvaddati</th>
        </tr>
        </thead>
        <tbody>
        <tr>
            <td><code>{input}</code></td>
            <td><code>{output}</code></td>
        </tr>
        </tbody>
    </table>)
}