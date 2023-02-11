import NavBar from '../components/navbar'
import Link from "next/link";
import {useState} from 'react'
import TagList from "../components/taglist";
import {formatDateTime} from "../scripts/formatDateTime";
import Error from "../components/error";

async function deleteTask(taskId, apiURL) {
    const options = {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: JSON.stringify({task_id: taskId}),
    };

    const response = await fetch(`${apiURL}/tasks/delete/` + taskId, options);

    if (!response.ok) {
        const errorMessage = await response.text();
        console.error(errorMessage)
        throw errorMessage
    }

    const respText = response.text();
    console.log("create task response: ", respText)
    return respText
}

async function createTaskSubmit(form_event) {
    form_event.preventDefault()
    const form = form_event.currentTarget;
    const url = form.action;

    const formData = new FormData(form);
    const response = await fetch(url, {method: "POST", body: formData});

    if (!response.ok) {
        const errorMessage = await response.text();
        console.error(errorMessage)
        throw errorMessage
    }

    const respText = response.text();
    console.log("create task response: ", respText)
    return respText
}

export default function Admin(props) {

    const [tasks, setTasks] = useState(props.tasks)
    const [error, setError] = useState(null)

    let refreshTable = async () => {
        const res = await fetch(`${props.apiURL}/tasks/list`)
        const tasks = await res.json()
        setTasks(tasks)
    }

    let deleteTaskAndRefreshTable = async (taskId) => {
        try {
            await deleteTask(taskId, props.apiURL);
            await refreshTable();
        } catch (e) {
            setError(e)
        }
    }

    let displayTaskDeleteModal = async (taskId, taskName) => {
        document.getElementById("delete-task-modal-header").innerHTML = taskName
        document.getElementById("delete-task-modal-confirm").onclick = async () => {
            await deleteTaskAndRefreshTable(taskId);
            document.getElementById("delete-task-modal-close").click()
        }
        document.getElementById("delete-task-modal-show").click()
    }

    let createTaskSubmitAndRefresh = async (form_event) => {
        try {
            await createTaskSubmit(form_event)
            await refreshTable()
        } catch (e) {
            setError(e)
        }
    }

    return (
        <div>
            <NavBar active_page={"admin"}/>
            <main className="container">
                <h1 className="my-4 text-center">administrācija</h1>

                <Error msg={error}/>
                <form action={`${props.apiURL}/tasks/import`} onSubmit={createTaskSubmitAndRefresh}>
                    <div className="row">
                        <div className="mb-3 col">
                            <input className="form-control" type="file" name="task-file" accept={".zip"}/>
                        </div>
                        <div className={"col"}>
                            <button type="submit" className="btn btn-success">pievienot uzdevumu</button>
                        </div>
                    </div>
                </form>

                <table className="table table-hover table-bordered text-center" style={{tableLayout: "fixed"}}>
                    <thead>
                    <tr>
                        <th scope="col">kods</th>
                        <th scope="col">nosaukums</th>
                        <th scope="col">atjaunots</th>
                        <th scope="col">birkas</th>
                        <th scope="col">grūtība</th>
                        <th scope="col">atrisinājumi</th>
                        <th scope="col">iesūtījumi</th>
                        <th scope={"col"}>darbības</th>
                    </tr>
                    </thead>
                    <tbody>
                    {tasks.map((task) => (
                        <tr key={task["task_id"]}>
                            <th scope="row">
                                <Link href={"/tasks/" + task["task_id"]}>
                                    <a className="nav-link">{task["task_id"]}</a>
                                </Link>
                            </th>
                            <td>
                                <Link href={"/tasks/" + task["task_id"]}>
                                    <a className="nav-link">{task["name"]}</a>
                                </Link>
                            </td>
                            <td>{formatDateTime(task["updated_time"])}</td>
                            <td><TagList tags={task["tags"]}/></td>
                            <td><span className="badge bg-danger">6.9</span></td>
                            <td>2</td>
                            <td>13</td>
                            <td>
                                <button type="button" className="btn btn-sm btn-primary me-1 my-1 disabled">Rediģēt
                                </button>
                                <button type="button" className="btn btn-sm btn-danger ms-1 my-1"
                                        onClick={() => displayTaskDeleteModal(task["task_id"], task["name"])}>Izdzēst
                                </button>
                            </td>
                        </tr>
                    ))}
                    </tbody>
                </table>

            </main>

            <div className="modal fade" id="delete-task-modal" tabIndex="-1">
                <div className="modal-dialog">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h5 className="modal-title" id="delete-task-modal-header"></h5>
                            <button id="delete-task-modal-close" type="button" className="btn-close"
                                    data-bs-dismiss="modal"></button>
                        </div>
                        <div className="modal-body">
                            Vai esat pārliecināti, ka vēlaties dzēst šo uzdevumu?
                        </div>
                        <div className="modal-footer">
                            <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">aizvērt
                            </button>
                            <button type="button" className="btn btn-danger" id="delete-task-modal-confirm">dzēst
                                uzdevumu
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            <button type="button" className="btn btn-primary d-none" id="delete-task-modal-show" data-bs-toggle="modal"
                    data-bs-target="#delete-task-modal">
            </button>

        </div>
    )
}


// This gets called on every request
export async function getServerSideProps() {
    let result = {
        props: {
            apiURL: process.env.API_URL
        }
    }

    try {
        const res = await fetch(`${process.env.API_URL}/tasks/list`)
        result.props.tasks = await res.json()
    } catch (err) {
        result.props.error = "failed to fetch tasks from the API :("
    }

    return result
}
