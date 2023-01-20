import NavBar from "../../components/navbar";
import Editor from "@monaco-editor/react";
import {parseStatement} from "../../scripts/renderMD";

export default function Submission({submission, error}) {
    return (
        <div className="d-flex flex-column vh-100">
            <NavBar active_page={"submissions"}/>
            <div className={"container flex-grow-1 my-3"}>
                <Editor
                    height="100%"
                    defaultLanguage="cpp"
                    defaultValue={submission["subm_src_code"]}
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