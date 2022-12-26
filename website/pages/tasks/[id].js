import NavBar from "../../components/navbar";
import Link from 'next/link'
import { useRouter } from 'next/router'

export default function Task() {
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