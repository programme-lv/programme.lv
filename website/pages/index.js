import Link from 'next/link'
import NavBar from '../components/navbar'

export default function Home() {
    return (
        <div>
            <NavBar/>
            <main style={{ backgroundImage: "url(space1.jpg)", height: "100vh", backgroundSize: "contain" }}>
                <div className="container text-white py-5">
                    <h2 className="my-5">jauns nostūris informātikas un matemātikas cienītājiem</h2>
                </div>
            </main>
            <footer>
            </footer>
        </div>
    )
}