import Link from 'next/link'

export default function Navbar({active_page, admin}) {
    admin = true; // TODO: this is a temporary hack to make the admin page show up
    return (
        <nav className="navbar">
            <div className="navbar-expand-md container border p-1 px-3">
                <div className="d-flex">
                    <Link href="/"><a className="navbar-brand">programme.lv</a></Link>
                    <button className="navbar-toggler" type="button" data-bs-toggle="offcanvas"
                            data-bs-target="#navbar-offcanvas">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbar-main">
                        <ul className="navbar-nav">
                            <li className="nav-item">
                                <Link href="/tasks"><a
                                    className={'nav-link ' + (active_page === 'tasks' ? 'active' : '')}>uzdevumi</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/submissions"><a
                                    className={'nav-link ' + (active_page === 'submissions' ? 'active' : '')}>iesūtījumi</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/results"><a
                                    className={'nav-link disabled ' + (active_page === 'results' ? 'active' : '')}>rezultāti</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/competitions"><a
                                    className={'nav-link disabled ' + (active_page === 'competitions' ? 'active' : '')}>sacensības</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/editor"><a
                                    className={'nav-link disabled ' + (active_page === 'editor' ? 'active' : '')}>redaktors</a></Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/theory"><a
                                    className={'nav-link disabled ' + (active_page === 'theory' ? 'active' : '')}>teorija</a></Link>
                            </li>
                            {admin ? <li className="nav-item">
                                <Link href="/admin"><a
                                    className={'nav-link ' + (active_page === 'admin' ? 'active' : '')}>administrācija</a></Link>
                            </li> : <></>}

                        </ul>
                    </div>
                </div>
                <div className="navbar-nav">
                    <li className="nav-item">
                        <Link href="#"><a className='nav-link active'>pierakstīties <i
                            className="bi bi-box-arrow-in-right"></i></a></Link>
                    </li>

                </div>

            </div>
            <div className="offcanvas offcanvas-end" id="navbar-offcanvas">
                <div className="offcanvas-body">
                    <ul className="navbar-nav me-auto">
                        <li className="nav-item">
                            <Link href="/tasks"><a className='nav-link'>uzdevumi</a></Link>
                        </li>
                        <li className="nav-item mt-2">
                            <Link href="/submissions"><a className='nav-link'>iesūtījumi</a></Link>
                        </li>
                        <li className="nav-item mt-2">
                            <Link href="/results"><a className='nav-link disabled'>rezultāti</a></Link>
                        </li>
                        <li className="nav-item mt-2">
                            <Link href="/competitions"><a className='nav-link disabled'>sacensības</a></Link>
                        </li>
                        <li className="nav-item mt-2">
                            <Link href="/editor"><a className='nav-link disabled'>redaktors</a></Link>
                        </li>
                        <li className="nav-item mt-2">
                            <Link href="/theory"><a className='nav-link disabled'>teorija</a></Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>

    )
}