import {remark} from "remark";
import remarkGfm from 'remark-gfm'
import remarkRehype from 'remark-rehype'
import rehypeStringify from 'rehype-stringify'

export async function renderMD(md) {
    const processedContent = await remark()
        .use(remarkGfm)
        .use(remarkRehype)
        .use(rehypeStringify).process(md);
    return processedContent.toString();
}

export async function parseStatement(statement) {
    statement["desc"] = await renderMD(statement["desc"])
    statement["input"] = await renderMD(statement["input"])
    statement["output"] = await renderMD(statement["output"])
    statement["scoring"] = await renderMD(statement["scoring"])
    return statement
}