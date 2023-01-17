import NavBar from "../../components/navbar";
import TagList from "../../components/taglist";


export default function Task(props) {
    let pdfURL = props.apiURL + "/tasks/statement/" + props.task["code"] + "/" + props.task["pdf_statements"][0].filename

    return (
        <>
            <NavBar active_page={"tasks"}/>
            <main className="container">
                <div className={"row my-5"}>
                    <div className="col-9">asdf</div>
                    <div className="col-3 card">
                        <div className="card-body">
                            <h5 className="card-title">Uzd. Informācija</h5>
                            <p className="card-text"></p>
                            <table className={"table table-hover"}>
                                <tbody>
                                <tr>
                                    <th scope="col">kods:</th>
                                    <td className={"text-start ps-2"}>{props.task["code"]}</td>
                                </tr>
                                <tr>
                                    <th scope="col">laika limits:</th>
                                    <td className={"text-start ps-2"}>{props.task["time_lim"]} sek.</td>
                                </tr>
                                <tr>
                                    <th scope="col">atmiņa:</th>
                                    <td className={"text-start ps-2"}>{props.task["mem_lim"]} MB</td>
                                </tr>
                                <tr>
                                    <th scope="col">versija:</th>
                                    <td className={"text-start ps-2"}>{props.task["version"]}</td>
                                </tr>
                                <tr>
                                    <th scope="col">autors:</th>
                                    <td className={"text-start ps-2"}>{props.task["author"]}</td>
                                </tr>
                                </tbody>
                            </table>
                            <h6 className="card-subtitle mb-2">Birkas</h6>
                            <TagList tags={props.task["tags"]}/>
                            <h6 className="card-subtitle my-2">Statistika</h6>
                            <table className={"table table-hover"}>
                                <tbody>
                                <tr>
                                    <th scope="col">iesūtījumi:</th>
                                    <td className={"text-start ps-2"}>?</td>
                                </tr>
                                <tr>
                                    <th scope="col">atrisinājumi:</th>
                                    <td className={"text-start ps-2"}>?</td>
                                </tr>
                                <tr>
                                    <th scope="col">grūtība:</th>
                                    <td className={"text-start ps-2"}>?</td>
                                </tr>
                                </tbody>
                            </table>

                        </div>
                    </div>
                </div>
            </main>
        </>
    )
}

/*
<main className="container" style={{height: "80vh"}}>
                <h1 className="my-4 text-center">{id}</h1>

                <embed src={pdfURL} type="application/pdf" width="100%" height="100%"/>
            </main>
 */

export async function getServerSideProps(context) {
    try {
        const reqRes = await fetch(`${process.env.API_URL}/tasks/view/${context.params.id}`)
        const task = await reqRes.json()
        return {
            props: {
                task: task,
                apiURL: process.env.API_URL
            }
        }
    } catch (err) {
        console.log(err)
        return {props: {error: "failed to fetch task info from the API :("}}
    }
}