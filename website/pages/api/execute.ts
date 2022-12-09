import type { NextApiRequest, NextApiResponse } from 'next'

export default (req: NextApiRequest, res: NextApiResponse) => {
    console.log(req.body)
    res.status(200).json({ name: 'John Doe' })

    var options = {
        host: 'localhost',
        path: '/execute/enqueue',
        port: '8080',
        method: 'POST'
    };
    
    let cb = function(response) {
        var str = ''
        response.on('data', function (chunk) {
            str += chunk;
        });
    
        response.on('end', function () {
            console.log(str);
        });
    }

    var http = require('http');
    var http_req = http.request(options, cb);
    
    http_req.write("{\"name\":\"huina\"}");
    http_req.end();
}