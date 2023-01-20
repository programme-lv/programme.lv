import NavBar from "../../components/navbar";
import TagList from "../../components/taglist";
import "katex/dist/katex.min.css";
import renderMathInElement from "katex/dist/contrib/auto-render.mjs";
import {useEffect, useRef} from "react";
import {parseStatement} from "../../scripts/renderMD";
import Editor from "@monaco-editor/react";
import {useRouter} from "next/router";

export default function Task({task, apiURL}) {
    const router = useRouter();

    //let pdfURL = props.apiURL + "/tasks/statement/" + task["code"] + "/" + task["pdf_statements"][0].filename
    let mdStatement = task["md_statements"][0]

    useEffect(() => {
        // RENDER KATEX MATH
        renderMathInElement(document.getElementById("task-statement"), {
            throwOnError: true, delimiters: [{left: '$$', right: '$$', display: true}, {
                left: '$', right: '$', display: false
            }, {left: '\\(', right: '\\)', display: false}, {left: '\\[', right: '\\]', display: true}],
        });
    }, []);

    const submissionEditorRef = useRef(null);

    function handleSubmissionEditorDidMount(editor, monaco) {
        submissionEditorRef.current = editor;
    }

    return (<div className="vw-100">
        <NavBar active_page={"tasks"}/>
        <main className="container">
            <div className={"row my-5"}>
                <div className="col-9 pe-4" id="task-statement">
                    <h2>{task["name"]}</h2>
                    <hr></hr>
                    <section className="my-4">
                        <h5 className={"my-3"}>formulējums</h5>
                        <div dangerouslySetInnerHTML={{__html: mdStatement["desc"]}}></div>
                    </section>
                    <section className="my-4">
                        <h5>ievaddati</h5>
                        <div dangerouslySetInnerHTML={{__html: mdStatement["input"]}}></div>
                    </section>
                    <section className="my-4">
                        <h5>izvaddati</h5>
                        <div dangerouslySetInnerHTML={{__html: mdStatement["output"]}}></div>
                    </section>
                    <section className="my-4">
                        <h5>piemēri</h5>
                        <div className={"row"}>
                            {mdStatement["examples"].map((example, index) => {
                                return (<table className={"table table-bordered col m-3"} key={index}>
                                    <thead>
                                    <tr>
                                        <th>ievaddati</th>
                                        <th>izvaddati</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    <tr>
                                        <td>
                                            <code>{example["input"]}</code>
                                        </td>
                                        <td>
                                            <code>{example["output"]}</code>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>)
                            })}
                        </div>
                    </section>
                    <section className="my-4 md-statement-scoring">
                        <h5>vērtēšana</h5>
                        <div dangerouslySetInnerHTML={{__html: mdStatement["scoring"]}}></div>
                    </section>
                    {mdStatement["notes"] ? <section className="my-4">
                        <h5>piezīmes</h5>
                        <div dangerouslySetInnerHTML={{__html: mdStatement["notes"]}}></div>
                    </section> : <></>}

                </div>
                <div className="col-3 card shadow-sm h-100">
                    <div className="card-body">
                        <h5 className="card-title text-center">uzd. Informācija</h5>
                        <p className="card-text"></p>
                        <table className={"table table-hover"}>
                            <tbody>
                            <tr>
                                <th scope="col">kods:</th>
                                <td className={"text-start ps-2"}>{task["code"]}</td>
                            </tr>
                            <tr>
                                <th scope="col">laika limits:</th>
                                <td className={"text-start ps-2"}>{task["time_lim"]} sek.</td>
                            </tr>
                            <tr>
                                <th scope="col">atmiņa:</th>
                                <td className={"text-start ps-2"}>{task["mem_lim"]} MB</td>
                            </tr>
                            <tr>
                                <th scope="col">versija:</th>
                                <td className={"text-start ps-2"}>{task["version"]}</td>
                            </tr>
                            <tr>
                                <th scope="col">autors:</th>
                                <td className={"text-start ps-2"}>{task["author"]}</td>
                            </tr>
                            </tbody>
                        </table>
                        <h6 className="card-subtitle mt-3 mb-2">birkas</h6>
                        <TagList tags={task["tags"]}/>
                        <h6 className="card-subtitle mt-3 mb-2">statistika</h6>
                        <table className={"table table-hover"}>
                            <tbody>
                            <tr>
                                <th scope="col">iesūtījumi:</th>
                                <td className={"text-start ps-2"}>?</td>
                            </tr>
                            <tr>
                                <th scope="col">atrisinājumi:</th>
                                <td className={"text-start ps-2"}>?</td>
                            </tr>
                            <tr>
                                <th scope="col">grūtība:</th>
                                <td className={"text-start ps-2"}>?</td>
                            </tr>
                            </tbody>
                        </table>
                        <h5 className="card-title text-center">iesūtīšana</h5>

                        <div className="my-3 text-center">
                            <button type="button" className="btn btn-sm btn-outline-primary" data-bs-toggle="modal"
                                    data-bs-target="#submission-modal" id="submission-modal-toggle">atvērt sūtījuma
                                redaktoru
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <div className="modal modal-lg fade" id="submission-modal" tabIndex="-1">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title">{task["name"] + " - risinājuma iesūtīšana"}</h5>
                            <button type="button" className="btn-close" data-bs-dismiss="modal"
                                    aria-label="Close"></button>
                        </div>
                        <div className="modal-body">
                            <div className="row px-4">
                                <label className="col-4">programmēšanas valoda:</label>
                                <select className="col form-select form-select-sm mb-3" id="subm-lang-select"
                                        defaultValue="c++17">
                                    <option value="c++17">C++17 (GNU G++)</option>
                                    <option value="python3">Python 3.10.9</option>
                                    <option value="java19">Java 19.0.1 (OpenJDK)</option>
                                </select>
                            </div>
                            <Editor
                                height="50vh"
                                defaultLanguage="cpp"
                                defaultValue="hello"
                                onMount={handleSubmissionEditorDidMount}
                                options={{
                                    minimap: {enabled: false}
                                }}
                            />
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn btn-outline-secondary"
                                    data-bs-dismiss="modal">Aizvērt
                            </button>
                            <button type="button" className="btn btn-success" onClick={async () => {
                                const langCode = document.getElementById("subm-lang-select").value;
                                const submSrcCode = submissionEditorRef.current.getValue();
                                const dataSending = {
                                    "task_code": task["code"],
                                    "lang_code": langCode,
                                    "src_code": submSrcCode
                                }
                                const apiEndpoint = apiURL + "/submissions/enqueue";
                                try {
                                    const response = await fetch(apiEndpoint, {
                                        method: "POST",
                                        headers: {
                                            "Content-Type": "application/json",
                                        },
                                        body: JSON.stringify(dataSending)
                                    });
                                    if (response.ok) {
                                        const data = await response.json();
                                        console.log(data);

                                        // remove modal background
                                        document.getElementsByClassName("modal-backdrop")[0].remove();
                                        router.push("/submissions").then(() => {
                                        });
                                    }
                                } catch (e) {
                                    console.log(e);
                                }
                            }}>Iesūtīt
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </div>)
}

export async function getServerSideProps(context) {
    try {
        const reqRes = await fetch(`${process.env.API_URL}/tasks/view/${context.params.id}`)
        const task = await reqRes.json()

        for (let statement in task["md_statements"]) {
            task["md_statements"][statement] = await parseStatement(task["md_statements"][statement])
            console.log(task["md_statements"][statement])
        }
        return {
            props: {
                task: task, apiURL: process.env.API_URL
            }
        }
    } catch (err) {
        console.log(err)
        return {props: {error: "failed to fetch task info from the API :("}}
    }
}