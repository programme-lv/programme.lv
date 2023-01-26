import TagList from "./taglist";

export default function TaskInfoCard({task}) {
    return (
        <div className="col-3 card shadow-sm h-100">
            <div className="card-body">
                <h5 className="card-title text-center">uzd. informācija</h5>
                <p className="card-text"></p>
                <table className={"table table-hover"}>
                    <tbody>
                    <tr>
                        <th scope="col">kods:</th>
                        <td className={"text-start ps-2"}>{task["task_id"]}</td>
                    </tr>
                    <tr>
                        <th scope="col">laika limits:</th>
                        <td className={"text-start ps-2"}>{task["time_lim"]} sek.</td>
                    </tr>
                    <tr>
                        <th scope="col">atmiņa:</th>
                        <td className={"text-start ps-2"}>{task["mem_lim"]} MB</td>
                    </tr>
                    <tr>
                        <th scope="col">versija:</th>
                        <td className={"text-start ps-2"}>{task["version"]}</td>
                    </tr>
                    <tr>
                        <th scope="col">autors:</th>
                        <td className={"text-start ps-2"}>{task["author"]}</td>
                    </tr>
                    </tbody>
                </table>
                <h6 className="card-subtitle mt-3 mb-2">birkas</h6>
                <TagList tags={task["tags"]}/>
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
                    <button type="button" className="btn btn-sm btn-outline-primary" data-bs-toggle="modal"
                            data-bs-target="#submission-modal" id="submission-modal-toggle">atvērt sūtījuma
                        redaktoru
                    </button>
                </div>
            </div>
        </div>
    )
}