import NavBar from "../../components/navbar";
import { useRouter } from 'next/router'

export default function Task(props) {
    console.log(props)
    const router = useRouter()
    const { id } = router.query


    return (
        <div>
            <NavBar active_page={"tasks"} />
            <main className="container">
                <h1 className="my-4 text-center">{id}</h1>
            </main>
        </div>
    )
}

export async function getServerSideProps(context) {
    try {
        const reqRes = await fetch(`${process.env.API_URL}/tasks/view/${context.params.id}`)
        const task = await reqRes.json()
        return { props: { task } }
    } catch (err) {
        console.log(err)
        return { props: { error: "failed to fetch task info from the API :(" } }
    }
}