import NavBar from "../../components/navbar";
import Editor from "@monaco-editor/react";
import Link from "next/link";
import {formatDateTime} from "../../scripts/format_datetime.js";

export default function Submission({submission}) {
    return (
        <div className="vw-100">
            <NavBar active_page={"submissions"}/>
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
                        <th scope="col">izmantotā atmiņa</th>
                        <th scope="col">tiesāšanas laiks</th>
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
            </div>
        </div>

    )
}


export async function getServerSideProps(context) {
    try {
        const reqRes = await fetch(`${process.env.API_URL}/submissions/view/${context.params.id}`)
        const submission = await reqRes.json()

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