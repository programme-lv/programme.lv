import Navbar from '../components/Navbar'
import FractalCanopy from '../public/fractal-canopy.svg'
import Image from 'next/image'
import LinkedList from '../public/linked list.drawio.svg'

export default function Home() {
    // generate a landing page

    return (
        <>
            <Navbar/>
            <main className="container">
                <div className="d-flex my-5">
                    <h1 className="my-4 col-8 mx-auto text-center align-self-center font-monospace">
                        <strong>programme.lv</strong> - jauns
                        nostūris<br/>
                        informātikas un
                        matemātikas <br/>cienītājiem un
                        iesācējiem!</h1>
                    <div className="col-4">
                        <Image src={FractalCanopy} alt="fractal canopy" height={"600px"} objectFit={"contain"}/>
                    </div>
                </div>
                <div className="d-flex my-4 px-5">
                    <div className="col-5">
                        <Image src={LinkedList} alt="linked list" width={"300px"} height={"350px"}/>
                    </div>
                    <div className="col-7">
                        <h2>Kāpēc programme.lv?</h2>
                        <div className="fs-5 mt-3">
                            <ul>
                                <li>Automātiska risinājumu testēšana ar reāllaika atgriezenisko saiti;</li>
                                <li>Modernu programmēšanas valodu atbalsts;</li>
                                <li>Integrēta programmēšanas vide ar zemu latentumu;</li>
                                <li>Iespēja iegūt daļēju punktu skaitu par risinājumu;</li>
                                <li>Latvijas informātikas olimpiādes uzdevumu arhīvs;</li>
                                <li>NP, kā arī interaktīvo un citu uzdevumu veidu atbalsts;</li>
                                <li>Iespēja veidot savus uzdevumus un dalīties ar tiem;</li>
                                <li>Iespēja skatīt citu cilvēku risinājums pēc uzd. atrisināšanas;</li>
                                <li>Uzdevumu filtrēšana pēc avota, nepieciešamajām zināšanām;</li>
                                <li>Augošs klāsts ar algoritmu, datu struktūru un matemātikas teoriju;</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </main>
            <footer>
            </footer>
        </>
    )
}