import NavBar from "../components/navbar";
import Link from 'next/link'

export default function Tasks(props) {
    let task_table_entries = []

    let TagList = (props)=> {
        let tag_entries = []
        let tags = props.tags;
        for(let tag of tags) {
            let bg = "bg-secondary"
            if(tag=="ProblemCon++") bg = "bg-primary"
            tag_entries.push(<span className={`badge ${bg} m-1`}>{tag}</span>)
        }
        return (
            <>
                {tag_entries}
            </>
        )
    }
    if (props.tasks) {
        let tasks = props.tasks
        tasks.forEach((task) => {
            task_table_entries.push(
                <tr key={task["code"]}>
                    <th scope="row"><Link href={"/tasks/" + task["code"]}><a className="nav-link">{task["code"]}</a></Link></th>
                    <td><Link href={"/tasks/" + task["code"]}><a className="nav-link">{task["name"]}</a></Link></td>
                    <td><TagList tags={task["tags"]}/></td>
                    <td><span className="badge bg-danger">6.9</span></td>
                    <td>2</td>
                    <td>13</td>
                </tr>
            )
        })
    }
    let ErrorAlert = ({ msg }) => {
        if (msg) return (
            <div className="alert alert-danger text-center" role="alert">
                {msg}
            </div>)
        else return <></>
    }
    return (
        <div>
            <NavBar active_page={"tasks"} />
            <main className="container">
                <h1 className="my-4 text-center">uzdevumi</h1>
                <ErrorAlert msg={props.error}/>
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


export async function getServerSideProps() {
    try {
        const res = await fetch(`${process.env.API_URL}/tasks/list`)
        const tasks = await res.json()
        // Pass data to the page via props
        return { props: { tasks } }
    } catch (err) {
        return { props: { error: "failed to fetch tasks from the API :(" } }
    }
}
