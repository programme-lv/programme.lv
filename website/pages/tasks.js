import Navbar from "../components/Navbar";
import Link from 'next/link'
import TagList from "../components/TagList";
import Error from "../components/Error";

export default function Tasks({tasks, error}) {
    if (error) {
        return (
            <div className="vw-100 mw-100">
                <Navbar active_page={"tasks"}/>
                <main className="container">
                    <h1 className="my-4 text-center">uzdevumi</h1>
                    <Error msg={error}/>
                </main>
            </div>
        )
    }
    return (
        <div className="vw-100 mw-100">
            <Navbar active_page={"tasks"}/>
            <main className="container">
                <h1 className="my-4 text-center">uzdevumi</h1>
                <table className="table table-hover table-bordered text-center" style={{tableLayout: "fixed"}}>
                    <thead>
                    <tr>
                        <th scope="col">kods</th>
                        <th scope="col">nosaukums</th>
                        <th scope="col">birkas</th>
                        <th scope="col">grūtība</th>
                        <th scope="col">atrisinājumi</th>
                        <th scope="col">iesūtījumi</th>
                    </tr>
                    </thead>
                    <tbody>
                    {tasks && tasks.map((task, index) => {
                        return (
                            <tr key={index}>
                                <th scope="row"><Link href={"/tasks/" + task["task_id"]}><a
                                    className="nav-link">{task["task_id"]}</a></Link></th>
                                <td><Link href={"/tasks/" + task["task_id"]}><a
                                    className="nav-link">{task["name"]}</a></Link></td>
                                <td><TagList tags={task["tags"]}/></td>
                                <td><span className="badge bg-danger">6.9</span></td>
                                <td>2</td>
                                <td>13</td>
                            </tr>
                        )
                    })}
                    </tbody>
                </table>
            </main>
        </div>
    )
}


export async function getServerSideProps() {
    try {
        const res = await fetch(`${process.env.API_URL}/tasks/list`)
        const tasks = await res.json()
        console.log(tasks)
        return {props: {tasks: tasks, error: null}}
    } catch (err) {
        return {props: {tasks: null, error: "failed to fetch tasks from the API: " + err}}
    }
}
