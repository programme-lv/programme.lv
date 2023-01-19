import NavBar from "../components/navbar";
import Link from 'next/link'
import TagList from "../components/taglist";
import ErrorAlert from "../components/error_alert";

export default function Tasks({tasks, error}) {
    return (
        <div>
            <NavBar active_page={"tasks"}/>
            <main className="container">
                <h1 className="my-4 text-center">uzdevumi</h1>
                <ErrorAlert msg={error}/>
                <table className="table table-hover table-bordered" style={{tableLayout: "fixed"}}>
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
                    {tasks.map((task, index) => {
                        return (
                            <tr key={index}>
                                <th scope="row"><Link href={"/tasks/" + task["code"]}><a
                                    className="nav-link">{task["code"]}</a></Link></th>
                                <td><Link href={"/tasks/" + task["code"]}><a
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
        // Pass data to the page via props
        return {props: {tasks: tasks, error: ""}}
    } catch (err) {
        return {props: {tasks: [], error: "failed to fetch tasks from the API :("}}
    }
}
