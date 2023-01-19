import NavBar from "../components/navbar";

export default function Submissions({submissions}) {
    console.log(submissions)
    return (
        <>
            <NavBar active_page={"submissions"}/>
        </>
    )
}

export async function getServerSideProps() {
    try {
        const res = await fetch(`${process.env.API_URL}/submissions/list`)
        const submissions = await res.json()
        // Pass data to the page via props
        return { props: { submissions } }
    } catch (err) {
        return { props: { error: "failed to fetch submissions from the API :(" } }
    }
}