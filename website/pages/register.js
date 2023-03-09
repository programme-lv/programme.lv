import {useState} from 'react';
import {Envelope, Lock, Person} from 'react-bootstrap-icons';
import Link from "next/link";
import LogoWithText from '../public/logo-with-text.png'
import Image from "next/image";

const RegisterForm = ({apiUrl}) => {
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (password !== password2) {
            alert("Paroles nesakrīt");
            return;
        }
        // Handle form submission
        const data = {"first_name": firstName, "last_name": lastName, username, email, password};
        const response = await fetch(apiUrl + "/users/register", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(data)
        });

        if (response.ok) {
            const data = await response.json();
            console.log(data);
            alert("Reģistrācija veiksmīga!");
        } else {
            alert("Kļūda: " + response.status + " " + response.statusText);
        }
    };

    return (<>
            <div className="container vh-100 w-100 d-flex align-items-center pb-5">
                <div className={"d-flex flex-column w-100 align-items-center mt-5 pb-5"}>
                    <Link href="/">
                        <a className="my-3 d-none d-lg-block">
                            <Image src={LogoWithText} objectFit={"contain"} alt="logo with text" height={"80%"}/>
                        </a>
                    </Link>
                    <Link href="/">
                        <a className="my-3 d-md-none d-block">
                            <Image src={LogoWithText} objectFit={"contain"} alt="logo with text" height={"200px"}/>
                        </a>
                    </Link>
                    <Link href="/">
                        <a className="my-3 d-lg-none d-none d-md-block">
                            <Image src={LogoWithText} objectFit={"contain"} alt="logo with text" height={"130px"}/>
                        </a>
                    </Link>
                    <div className="col-lg-5 col-sm-10 col-11  border p-4">
                        <h4 className="text-center mb-3 bold">Reģistrācija :)</h4>
                        <form onSubmit={handleSubmit}>

                            <div className="mb-3">
                                <label htmlFor="username" className="form-label">Lietotājvārds</label>
                                <div className="input-group">
                                    <input type="text" className="form-control" id="username" value={username}
                                           onChange={(e) => setUsername(e.target.value)} required/>
                                    <span className="input-group-text bg-white"><Person/></span>
                                </div>
                            </div>
                            <div className="mb-3 d-flex mt-4">
                                {/*first name, last name*/}
                                <div className={"col-6 pe-3"}>
                                    <label htmlFor="first_name" className="form-label">Jūsu vārds</label>
                                    <div className="input-group">
                                        <input type="text" className="form-control" id="first_name" value={firstName}
                                               onChange={(e) => setFirstName(e.target.value)} required/>
                                    </div>
                                </div>
                                <div className={"col-6 ps-3"}>
                                    <label htmlFor="last_name" className="form-label">Uzvārds</label>
                                    <div className="input-group">
                                        <input type="text" className="form-control" id="last_name" value={lastName}
                                               onChange={(e) => setLastName(e.target.value)} required/>
                                    </div>
                                </div>
                            </div>
                            <div className="mb-3">
                                <label htmlFor="email" className="form-label">Epasta adrese</label>
                                <div className="input-group">
                                    <input type="email" className="form-control" id="email" value={email}
                                           onChange={(e) => setEmail(e.target.value)} required/>
                                    <span className="input-group-text bg-white"><Envelope/></span>
                                </div>
                            </div>
                            <div className="mb-3">
                                <label htmlFor="password" className="form-label">Parole</label>
                                <div className="input-group">
                                    <input type="password" className="form-control" id="password" value={password}
                                           onChange={(e) => setPassword(e.target.value)} required/>
                                    <span className="input-group-text bg-white"><Lock/></span>
                                </div>
                            </div>
                            <div className="mb-3">
                                <label htmlFor="password2" className="form-label">Apstipriniet paroli</label>
                                <div className="input-group">
                                    <input type="password" className="form-control" id="password2" value={password2}
                                           onChange={(e) => setPassword2(e.target.value)} required/>
                                    <span className="input-group-text bg-white"><Lock/></span>
                                </div>
                            </div>
                            <button type="submit" className="btn btn-success w-100 my-2">reģistrēties</button>
                            <div className="py-3">
                                Jau esi piereģistrējies? <Link href="/login"><a
                                className="text-decoration-none text-primary">Pieslēgties</a></Link>
                            </div>
                        </form>

                    </div>
                    {/*<div className="col-4 d-flex bg-white border py-4">*/}
                    {/*    <Image src={Fractal_canopy} objectFit={"contain"} alt="fractal canopy" height={"100%"}/>*/}
                    {/*</div>*/}
                </div>
            </div>
        </>
    );
};

export default RegisterForm;

export async function getServerSideProps() {
    return {
        props: {
            apiUrl: process.env.API_URL
        }
    }
}