import NavBar from "../../components/navbar";
import TagList from "../../components/taglist";


export default function Task(props) {
    let pdfURL = props.apiURL + "/tasks/statement/" + props.task["code"] + "/" + props.task["pdf_statements"][0].filename

    let mdStatement = props.task["md_statements"][0]

    console.log(props)
    return (
        <>
            <NavBar active_page={"tasks"}/>
            <main className="container">
                <div className={"row my-5"}>
                    <div className="col-9">
                        <h2>{props.task["name"]}</h2>
                        <hr></hr>
                        <section className="my-4">
                            <h3>formulējums</h3>
                            <p>{mdStatement["desc"]}</p>
                        </section>
                    </div>
                    <div className="col-3 card shadow-sm">
                        <div className="card-body">
                            <h5 className="card-title text-center">uzd. Informācija</h5>
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
                            <h6 className="card-subtitle mt-3 mb-2">birkas</h6>
                            <TagList tags={props.task["tags"]}/>
                            <h6 className="card-subtitle mt-3 mb-2">statistika</h6>
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
                            <h5 className="card-title text-center">iesūtīšana</h5>

                            <div className="my-3 text-center">
                                <button type="button" className="btn btn-sm btn-primary">atvērt sūtījuma redaktoru
                                </button>
                            </div>
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