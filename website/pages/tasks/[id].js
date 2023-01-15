import NavBar from "../../components/navbar";
import { useRouter } from 'next/router'

export default function Task(props) {
    console.log(props)
    const router = useRouter()
    const { id } = router.query

    console.log(props.task["pdf_statements"][0].filename)

    let pdfURL = props.apiURL+"/tasks/statement/" + props.task["code"]+"/"+props.task["pdf_statements"][0].filename

    return (
        <>
            <NavBar active_page={"tasks"} />
            <main className="container" style={{height: "80vh"}}>
                <h1 className="my-4 text-center">{id}</h1>
                <embed src={pdfURL} type="application/pdf" width="100%" height="100%"/>
            </main>
        </>
    )
}

export async function getServerSideProps(context) {
    try {
        const reqRes = await fetch(`${process.env.API_URL}/tasks/view/${context.params.id}`)
        const task = await reqRes.json()
        return { props: { task: task,
                apiURL: process.env.API_URL} }
    } catch (err) {
        console.log(err)
        return { props: { error: "failed to fetch task info from the API :(" } }
    }
}