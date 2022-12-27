import NavBar from "../../components/navbar";
import Link from 'next/link'
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
    return {
        props: {data:1234}, // will be passed to the page component as props
    }
}