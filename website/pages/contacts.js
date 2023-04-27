import Navbar from "../components/Navbar";

export default function EditorPage() {
    const team = [
        {
            name: "K. Petručeņa",
            position: "CTO",
            image: '/team_profiles/kp_profile.png'
        },
        {
            name: "V. Lohmanova",
            position: "CEO",
            image: '/team_profiles/vl_profile.png'
        },
        {
            name: "M. Timoņina",
            position: "CDO",
            image: '/team_profiles/mt_profile.png'
        },
        {
            name: "Aksels",
            position: "CFO",
            image: '/team_profiles/a_profile.png'
        }
    ]
    return (
        <div className="d-flex flex-column vh-100">
            <Navbar active_page={"contacts"}/>
            <div className={"container my-3"}>
                <div className="d-flex flex-row justify-content-between align-items-center">
                    <div className="d-flex flex-column h-auto">
                        <h1 className="fw-bold text-uppercase text-center" style={{writingMode: 'vertical-lr', transform: 'rotate(180deg)', }}>KOMANDA</h1>
                    </div>
                    {team.map((member) => {
                       return (
                           <div className="d-flex flex-column">
                                <div className="card border-0" style={{width: '10rem'}} >
                                    <img src={member.image} className="card-img-top rounded-circle" alt="..."/>
                                    <div className="card-body text-center">
                                        <h5 className="card-title">{member.name}</h5>
                                        <p className="card-text">{member.position}</p>
                                    </div>
                                </div>
                           </div>
                     )
                    })}
                </div>
            </div>
            <div className="container-fluid">
                <h3 className="text-center text-decoration-underline">Sazināties ar mums</h3>
                <div className="d-flex flex-row justify-content-center">
                    <div className="col border-end border-dark border-2 text-end" >
                        <div className="container-fluid">
                            <h7 className="fw-bold"> E-pasts </h7>
                            <br></br>
                            <br></br>
                            <h7 className="fw-bold"> Kompānija </h7>
                            <br></br>
                            <br></br>
                            <br></br>
                            <h7 className="fw-bold"> Adrese </h7>
                            <br></br>
                            <br></br>
                            <h7 className="fw-bold"> Rekvizīti </h7>
                        </div>
                    </div>
                    <div className="col text-start">
                        <div className="container-fluid">
                            <h7>programme.lv@gmail.com</h7>
                            {/*<h7><a className="text-decoration-none" href={"mailto:programme.lv@gmail.com"}>programme.lv@gmail.com</a></h7>*/}
                            <br></br>
                            <br></br>
                            <h7>SIA "Programme.lv"</h7>
                            <br></br>
                            <h7>Reģ. Nr. 40000000000</h7>
                            <br></br>
                            <br></br>
                            <h7>Brīvības iela 103, Rīga, LV-1001</h7>
                            <br></br>
                            <br></br>
                            <h7>N/K LV00HABA0000000000000</h7>
                            <br></br>
                            <h7>AS Swedbank</h7>
                            <br></br>
                            <h7>SWIFT kods HABALV22</h7>
                            <br></br>
                        </div>
                    </div>
                </div>


            </div>
        </div>
    )
}