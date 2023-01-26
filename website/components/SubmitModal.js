import Editor from "@monaco-editor/react";
import {useRouter} from "next/router";
import {useRef} from "react";

export default function SubmitModal({task, languages, apiURL}) {
    console.log(languages)

    const submissionEditorRef = useRef(null);

    function handleSubmissionEditorDidMount(editor) {
        submissionEditorRef.current = editor;
    }

    const router = useRouter(); // used for automatic navigation to the submission page
    return (<div className="modal modal-lg fade" id="submission-modal" tabIndex="-1">
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
                                    defaultValue="C++17">
                                {
                                    languages.map((lang, index) => {
                                        return <option key={index} value={lang["lang_id"]}>{lang["name"]}</option>
                                    })
                                }
                            </select>
                        </div>
                        <Editor
                            height="50vh"
                            defaultLanguage="cpp"
                            defaultValue={defaultCode}
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
                                "task_code": task["task_id"],
                                "lang_id": langCode,
                                "src_code": submSrcCode
                            }
                            const apiEndpoint = apiURL + "/submissions/enqueue";
                            try {

                                const response = await fetch(apiEndpoint, {
                                    method: "POST",
                                    headers: {"Content-Type": "application/json"},
                                    body: JSON.stringify(dataSending)
                                });

                                if (response.ok) {
                                    const data = await response.json();
                                    console.log(data);

                                    // remove modal background
                                    // reset style on document.body
                                    document.body.removeAttribute("style");
                                    document.getElementsByClassName("modal-backdrop")[0].remove();
                                    router.push("/submissions").then(() => {
                                    });
                                } else {
                                    alert("Kļūda: " + response.status + " " + response.statusText);
                                }
                            } catch (e) {
                                alert("Kļūda: " + e);
                                console.log(e);
                            }
                        }}>Iesūtīt
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}

const defaultCode = `#include <iostream>

using namespace std;

int main() {
    cout<<"hello, world!";
}`;

