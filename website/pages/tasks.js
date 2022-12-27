import NavBar from "../components/navbar";
import Link from 'next/link'

export default function Tasks() {
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
                        <tr>
                            <th scope="row"><Link href="/tasks/baobabi"><a className="nav-link">baobabi</a></Link></th>
                            <td><Link href="/tasks/baobabi"><a className="nav-link">baobabi</a></Link></td>
                            <td><span className="badge bg-primary">ProblemCon++</span></td>
                            <td><span className="badge bg-danger">6.9</span></td>
                            <td>2</td>
                            <td>13</td>
                        </tr>
                    </tbody>
                </table>
            </main>
        </div>
    )
}