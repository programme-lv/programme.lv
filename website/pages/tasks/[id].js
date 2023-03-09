import {parseStatement} from "../../scripts/renderMD";
import Navbar from "../../components/Navbar";
import MDStatement from "../../components/MDStatement";
import SubmitModal from "../../components/SubmitModal";
import TaskInfoCard from "../../components/TaskInfoCard";

export default function Task({languages, task, apiURL}) {
    let mdStatement = task["md_statements"][0] ?? null;

    return (<div className="vw-100 mw-100">
        <Navbar active_page={"tasks"}/>
        <main className="container p-0 my-4">
            <div className={"my-3 d-flex"}>
                <div className="col-9 p-0 pe-4" id="task-statement">
                    <div className="bg-white">
                        <h2>{task["name"]}</h2>
                        <hr></hr>
                        <MDStatement mdStatement={mdStatement}/>
                    </div>
                </div>
                <TaskInfoCard task={task}/>
            </div>
        </main>
        <SubmitModal languages={languages} task={task} apiURL={apiURL}/>
    </div>)
}

export async function getServerSideProps(context) {
    try {

        const languages = await fetch(`${process.env.API_URL}/languages/list`).then(res => res.json())

        const task = await fetch(`${process.env.API_URL}/tasks/view/${context.params.id}`).then(res => res.json())

        for (let statement in task["md_statements"]) {
            task["md_statements"][statement] = await parseStatement(task["md_statements"][statement])
        }

        return {
            props: {
                languages: languages,
                task: task,
                apiURL: process.env.API_URL
            }
        }
    } catch (err) {
        return {props: {error: "failed to fetch task info from the API :("}}
    }
}
