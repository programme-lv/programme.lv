import Link from 'next/link'
import NavBar from '../components/navbar'

export default function Home() {
    // generate a landing page

    return (
        <div>
            <NavBar />
            <main className="container">
                <h2 className="my-4 text-center">jauns nostūris informātikas un matemātikas cienītājiem</h2>
            </main>
            <footer>
            </footer>
        </div>
    )
}