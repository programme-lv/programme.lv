import NavBar from "../components/navbar";
import Link from 'next/link'

export default function Tasks({tasks}) {
    console.log(tasks)

    let task_table_entries = []
    tasks.forEach((task)=>{
        task_table_entries.push(
            <tr key={task["task_code"]}>
                <th scope="row"><Link href={"/tasks/"+task["task_code"]}><a className="nav-link">{task["task_code"]}</a></Link></th>
                <td><Link href={"/tasks/"+task["task_code"]}><a className="nav-link">{task["task_name"]}</a></Link></td>
                <td><span className="badge bg-primary">ProblemCon++</span></td>
                <td><span className="badge bg-danger">6.9</span></td>
                <td>2</td>
                <td>13</td>
            </tr>
        )
    })
    return (
        <div>
            <NavBar active_page={"tasks"} />
            <main className="container">
                <h1 className="my-4 text-center">uzdevumi</h1>
                <table className="table table-hover" style={{ tableLayout: "fixed" }}>
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
                    {task_table_entries}
                    </tbody>
                </table>
            </main>
        </div>
    )
}


// This gets called on every request
export async function getServerSideProps() {
    // Fetch data from external API
    const res = await fetch(`http://localhost:8080/tasks/list`)
    const tasks = await res.json()
    // Pass data to the page via props
    return { props: { tasks } }
}
