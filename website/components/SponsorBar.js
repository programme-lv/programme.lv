import nms_logo from "../public/sponsors_logo/nms.jpg";
import pps_logo from "../public/sponsors_logo/pps.png";
import start_it_logo from "../public/sponsors_logo/start-it.png";
import Image from "next/image";

export default function SponsorBar() {
    return (
        <div className="container-fluid">
            <div className="row">
                <div className="col-sm align-self-center text-center">
                    <h4>Projektu atbalsta</h4>
                </div>
                <div className="col align-self-center">
                    <Image className="img" src={pps_logo} alt="Pirmā programmēšanas skola" objectFit="contain" height={"300px"}></Image>
                </div>
                <div className="col align-self-center">
                        <Image className="img" src={start_it_logo} alt="IT Izglītības fonds - start(it)" objectFit="contain" height={"100px"}></Image>
                </div>
                <div className="col align-self-center">
                        <Image className="img" src={nms_logo} alt="LU A. Liepas Neklātienes matemātikas skola" objectFit="contain" height={"100px"}></Image>
                </div>
            </div>
        </div>
    );
}