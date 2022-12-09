import type { NextApiRequest, NextApiResponse } from 'next'

export default (req: NextApiRequest, res: NextApiResponse) => {
    console.log(req.body)
    res.status(200).json("{}")

    var options = {
        host: 'localhost',
        path: '/submissions/enqueue',
        port: '8080',
        method: 'POST'
    };
    
    var data = {
        task_name: "hello",
        user_code: req.body.code,
        lang_id: "cpp"
    }

    let cback = function(response) {
        var str = ''
        response.on('data', function (chunk) {
            str += chunk;
        });
    
        response.on('end', function () {
            console.log(str);
        });
    }

    var http = require('http');
    var http_req = http.request(options, cback);
    
    
    http_req.write(JSON.stringify(data));
    http_req.end();
}