import 'bootstrap/dist/css/bootstrap.min.css'
import '../styles/globals.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
import Head from "next/head";
import SponsorBar from "../components/SponsorBar";

import {useEffect} from 'react'

function MyApp({Component, pageProps}) {
    useEffect(() => {
        require("bootstrap/dist/js/bootstrap");
    }, [])

    useEffect(()=>{
	    console.log("hello");
    },[Component])

    return (
        <>
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <meta name="description" content="Moderna programmēšanas izglītības platforma."/>
                <title>programme.lv</title>

            </Head>
            <Component {...pageProps} />
            <SponsorBar/>
        </>
    );
}

export default MyApp
