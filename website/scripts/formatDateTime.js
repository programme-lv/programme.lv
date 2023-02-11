export function formatDateTime(dateTime) {
    let d = new Date(dateTime);
    let date = d.toISOString().split("T")[0];
    date = date.replace("-", ".");
    date = date.replace("-", ".");
    let time = d.toTimeString().split(" ")[0];
    return date + " " + time;
}