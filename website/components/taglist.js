export default function TagList({tags}) {
    console.log(tags)
    if (!tags) return null
    return (
        <>
            {tags.map((tag) => {
                let bg = "bg-secondary"
                if (tag["name"] === "ProblemCon++") {
                    bg = "bg-primary"
                }
                return <span className={`badge ${bg} m-1`} key={tag["name"]}>{tag["name"]}</span>
            })}
        </>
    )
}