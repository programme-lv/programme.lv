import {remark} from "remark";
import html from "remark-html"

export async function renderMD(md) {
    const processedContent = await remark().use(html).process(md);
    return processedContent.toString();
}

export async function parseStatement(statement) {
    statement["desc"]=await renderMD(statement["desc"])
    statement["input"]=await renderMD(statement["input"])
    statement["output"]=await renderMD(statement["output"])
    statement["scoring"]=await renderMD(statement["scoring"])
    return statement
}