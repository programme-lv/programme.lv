import Navbar from "../components/Navbar";

export default function EditorPage() {
    return (
        <div className="d-flex flex-column vh-100">
            <Navbar active_page={"news"}/>
            <div className={"container flex-grow-1 my-3"}>
                <h1 className="my-4 text-center">jaunumi</h1>

            </div>
        </div>
    )
}