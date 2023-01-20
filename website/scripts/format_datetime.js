export function formatDateTime(dateTime) {
    console.log("dateTime: ", dateTime);
    let d = new Date(dateTime);
    console.log(d);
    let date = d.toISOString().split("T")[0];
    date = date.replace("-", ".");
    date = date.replace("-", ".");
    let time = d.toTimeString().split(" ")[0];
    return date + " " + time;
}