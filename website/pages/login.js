import {useState} from 'react';
import {Envelope, Lock} from 'react-bootstrap-icons';
import Navbar from "../components/Navbar";
import Image from 'next/image'
import Fractal_canopy from '../public/fractal-canopy.png'

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        // Handle form submission
    };

    return (<>
        <Navbar/>
        <div className="container">
            <div className={"d-flex justify-content-around mt-5"}>
                <div className="col-5 border p-4">
                    <h4 className="text-center mb-3">Pieslēdzies savam kontam</h4>
                    <form onSubmit={handleSubmit}>
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
                        <div className="d-flex justify-content-between">
                            <div>
                                <input className="form-check-input" type="checkbox" value="" id="remember"/>
                                <label className="form-check-label ms-2 mb-3" htmlFor="remember">
                                    Atcerēties mani
                                </label>
                            </div>
                            <div>
                                <a href="#" className="text-decoration-none text-danger">Aizmirsi paroli?</a>
                            </div>
                        </div>
                        <button type="submit" className="btn btn-primary w-100">pieslēgties</button>
                        <div className="py-3">
                            Neesi piereģistrējies? <a href="#"
                                                      className="text-decoration-none text-success">Reģistrēties</a>
                        </div>
                    </form>

                </div>
                <div className="col-4 d-flex bg-white border py-4">
                    <Image src={Fractal_canopy} objectFit={"contain"} alt="fractal canopy" height={"100%"}/>
                </div>
            </div>
        </div>
    </>);
};

export default LoginForm;
