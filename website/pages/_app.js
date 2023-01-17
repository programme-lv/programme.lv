import 'bootstrap/dist/css/bootstrap.css'
import '../styles/globals.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
import Head from "next/head";

import {useEffect} from 'react'


function MyApp({Component, pageProps}) {
    useEffect(() => {
        require("bootstrap/dist/js/bootstrap");
    }, [])

    return (
        <>
            <Head>
                <script id="MathJax-script" async
                        src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-chtml.js">
                </script>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <title>programme.lv</title>
            </Head>
            <Component {...pageProps} />
        </>
    );
}

export default MyApp
