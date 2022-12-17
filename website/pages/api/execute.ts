import type { NextApiRequest, NextApiResponse } from 'next'
import axios from 'axios'

export default async function execute(req: NextApiRequest, res: NextApiResponse) {
    let data = {
        user_code: req.body.code,
        lang_id: "cpp"
    }
    
    let scheduler_res = await axios.post('http://localhost:8080/submissions/enqueue', data)
    console.log(scheduler_res.data)

    res.status(200).json(scheduler_res.data)
}