export default function MDStatement({mdStatement}) {
    let description = mdStatement["desc"];
    let input = mdStatement["input"];
    let output = mdStatement["output"];
    let examples = mdStatement["examples"];
    let scoring = mdStatement["scoring"];
    let notes = mdStatement["notes"];
    return (<>
        <section className="my-4">
            <h5 className={"my-3"}>formulējums</h5>
            <div dangerouslySetInnerHTML={{__html: description}}></div>
        </section>
        <section className="my-4">
            <h5>ievaddati</h5>
            <div dangerouslySetInnerHTML={{__html: input}}></div>
        </section>
        <section className="my-4">
            <h5>izvaddati</h5>
            <div dangerouslySetInnerHTML={{__html: output}}></div>
        </section>
        <section className="my-4">
            <h5>piemēri</h5>
            <div className={"row"}>
                {examples.map((example, index) =>
                    <ExampleTable key={index} input={example["input"]} output={example["output"]}/>)}
            </div>
        </section>
        {scoring && <section className="my-4 md-statement-scoring">
            <h5>vērtēšana</h5>
            <div dangerouslySetInnerHTML={{__html: scoring}}></div>
        </section>}
        {notes && <section className="my-4">
            <h5>piezīmes</h5>
            <div dangerouslySetInnerHTML={{__html: notes}}></div>
        </section>}
    </>)
}

function ExampleTable({input, output}) {
    return (<table className={"table table-bordered col m-3"}>
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