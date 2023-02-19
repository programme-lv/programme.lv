import Navbar from "../components/Navbar";
import Error from "../components/Error";
import Link from "next/link";
import {formatDateTime} from "../scripts/formatDateTime.js";

export default function Submissions({submissions, error}) {
    if (error) {
        return (
            <div className="vw-100 mw-100">
                <Navbar active_page={"submissions"}/>
                <main className="container">
                    <h1 className="my-4 text-center">iesūtījumi</h1>
                    <Error msg={error}/>
                </main>
            </div>
        )
    }
    return (
        <div className="vw-100 mw-100">
            <Navbar active_page={"submissions"}/>
            <main className="container">
                <h1 className="my-4 text-center">iesūtījumi</h1>
                <Error msg={error}/>
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
                    </tr>
                    </thead>
                    <tbody>
                    {submissions && submissions.map((submission, index) => {
                        return (
                            <tr key={index}>
                                <th scope="row"><Link href={"/submissions/" + submission["subm_id"]}><a
                                >{submission["subm_id"]}</a></Link></th>
                                <td>{formatDateTime(submission["created_time"])}</td>
                                <td><Link href={"/users/" + submission["user_id"]}><a
                                    className="nav-link">{submission["user_id"]}</a></Link></td>
                                <td><Link
                                    href={"/tasks/" + submission["task_code"]}><a>{submission["task_code"]}</a></Link>
                                </td>
                                <td>{submission["lang_id"]}</td>
                                <td>{submission["status"]}</td>
                                <td>?</td>
                                <td>?</td>
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
        const res = await fetch(`${process.env.API_URL}/submissions/list`)
        const submissions = await res.json()
        // Pass data to the page via props
        return {props: {submissions}}
    } catch (err) {
        return {props: {error: err.toString()}}
    }
}