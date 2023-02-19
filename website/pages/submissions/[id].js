import Navbar from "../../components/Navbar";
import Editor from "@monaco-editor/react";
import Link from "next/link";
import {formatDateTime} from "../../scripts/formatDateTime.js";

export default function Submission({submission}) {
    return (
        <div className="vw-100 mw-100">
            <Navbar active_page={"submissions"}/>
            <div className={"container my-3"}>
                <table className="table table-hover table-bordered text-center">
                    <thead>
                    <tr>
                        <th scope="col">iesūtījums</th>
                        <th scope="col">iesūtījuma laiks</th>
                        <th scope="col">lietotājs</th>
                        <th scope="col">uzdevums</th>
                        <th scope="col">valoda</th>
                        <th scope="col">statuss</th>
                        <th scope="col">izpildes laiks</th>
                        <th scope="col">izmantotā atmiņa [MB]</th>
                        <th scope="col">tiesāšanas laiks [ms]</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <th scope="row"><a>{submission["subm_id"]}</a></th>
                        <td>{formatDateTime(submission["created_time"])}</td>
                        <td><Link href={"/users/" + submission["user_id"]}><a
                            className="nav-link">{submission["user_id"]}</a></Link></td>
                        <td><Link
                            href={"/tasks/" + submission["task_code"]}><a>{submission["task_code"]}</a></Link>
                        </td>
                        <td>{submission["lang_id"]}</td>
                        <td>IQS</td>
                        <td>?</td>
                        <td>?</td>
                        <td>?</td>
                    </tr>
                    </tbody>
                </table>
                <Editor
                    className={"border"}
                    height="40vh"
                    defaultLanguage="cpp"
                    defaultValue={submission["src_code"]}
                    options={{
                        readOnly: true,
                        minimap: {enabled: false}
                    }}
                />
                <table className="table table-hover table-bordered text-center mt-4">
                    <thead>
                    <tr>
                        <th scope="col">tests</th>
                        <th scope="col">statuss</th>
                        <th scope="col">apakšuzdevumi</th>
                        <th scope="col">tiesāšanas laiks [ms]</th>
                        <th scope="col">izmantotā atmiņa [MB]</th>
                        <th scope="col">ievaddati, izvaddati</th>
                    </tr>
                    </thead>
                    <tbody>
                    {submission["task_subm_evals"][0] && submission["task_subm_evals"][0]["task_subm_job_tests"] &&
                        // create a row for each test
                        submission["task_subm_evals"][0]["task_subm_job_tests"].map((test, index) => {
                            return (
                                <tr key={index}
                                    className={TestStatusRowClasses(test["status"])}>
                                    <th scope="row"><a>{index + 1}</a></th>
                                    <td title={TestStatusTooltip(test["status"])}
                                        className={TestStatusTextClasses(test["status"])}>{test["status"]}</td>
                                    <td><span className="badge bg-secondary m-1">1</span></td>
                                    <td>{test["time"]}</td>
                                    <td>{test["memory"]}</td>
                                    <td>
                                        <button className="btn btn-primary btn-sm">skatīt</button>
                                    </td>
                                </tr>
                            )
                        })
                    }
                    </tbody>
                </table>
            </div>
        </div>

    )
}

function TestStatusRowClasses(status) {
    switch (status) {
        case "OK":
            return "bg-success bg-opacity-25"
        default:
            return ""
    }
}

function TestStatusTooltip(status) {
    switch (status) {
        case "OK":
            return "testa izpilde ir veiksmīga (okay)"
        case "WA":
            return "izvaddati nav pareizi (wrong answer)"
        case "TLE":
            return "testa izpilde pārsniedza laika limitu (time limit exceeded)"
        default:
            return ""
    }
}

function TestStatusTextClasses(status) {
    switch (status) {
        case "OK":
            return "text-success"
        case "WA":
            return "text-danger"
        case "TLE":
            return "text-warning"
        default:
            return ""
    }
}

export async function getServerSideProps(context) {
    try {
        let url = `${process.env.API_URL}/submissions/view/${context.params.id}`
        const submission = await fetch(url).then(res => res.json())

        return {
            props: {
                submission: submission,
                error: null
            }
        }
    } catch (err) {
        console.log(err)
        return {
            props: {submission: {}, error: "failed to fetch task info from the API :("}
        }
    }
}