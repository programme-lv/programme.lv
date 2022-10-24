// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default function handler(req, res) {
    let code = req.body.code;
    console.log(code);
    res.status(200).json({ name: 'Mike Oxlong' });
}
