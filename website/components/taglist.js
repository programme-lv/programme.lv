export default function TagList({tags}) {
    console.log(tags)
    if (!tags) return null
    let tag_entries = []
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