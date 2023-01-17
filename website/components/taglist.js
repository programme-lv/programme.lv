export default function TagList(props) {
    let tag_entries = []
    let tags = props.tags;
    for (let tag of tags) {
        let bg = "bg-secondary"
        if (tag === "ProblemCon++") bg = "bg-primary"
        tag_entries.push(<span className={`badge ${bg} m-1`} key={tag}>{tag}</span>)
    }
    return (
        <>
            {tag_entries}
        </>
    )
}