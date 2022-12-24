import Link from 'next/link'

export default function Home() {
    return (
        <div>
            <nav className="navbar bg-white fixed-top">
                <div className="navbar-expand-md container">
                    <Link href="#"><a className="navbar-brand">programme.lv</a></Link>
                    <button className="navbar-toggler" type="button" data-bs-toggle="offcanvas" data-bs-target="#navbar-offcanvas">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbar-main">
                        <ul className="navbar-nav">
                            <li className="nav-item">
                                <Link href="/tasks"><a className='nav-link'>uzdevumi</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/results"><a className='nav-link disabled'>rezultāti</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/submissions"><a className='nav-link disabled'>iesūtījumi</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/competitions"><a className='nav-link disabled'>sacensības</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/editor"><a className='nav-link disabled'>redaktors</a></Link>
                            </li>
                        </ul>
                    </div>
                </div>
                <div className="offcanvas offcanvas-end" id="navbar-offcanvas">
                    <div className="offcanvas-body">
                        <ul className="navbar-nav me-auto">
                            <li className="nav-item">
                                <Link href="/tasks"><a className='nav-link'>uzdevumi</a></Link>
                            </li>
                            <li className="nav-item mt-2">
                                <Link href="/results"><a className='nav-link disabled'>rezultāti</a></Link>
                            </li>
                            <li className="nav-item mt-2">
                                <Link href="/submissions"><a className='nav-link disabled'>iesūtījumi</a></Link>
                            </li>
                            <li className="nav-item mt-2">
                                <Link href="/competitions"><a className='nav-link disabled'>sacensības</a></Link>
                            </li>
                            <li className="nav-item mt-2">
                                <Link href="/editor"><a className='nav-link disabled'>redaktors</a></Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
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