import NavBar from '../components/navbar'

async function createTask(e) {
    e.preventDefault()
    const form = e.currentTarget;
    const url = form.action;

    const formData = new FormData(form);
    const plainFormData = Object.fromEntries(formData.entries());
    const formDataJsonString = JSON.stringify(plainFormData);

    const fetchOptions = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: formDataJsonString,
    };

    const response = await fetch(url, fetchOptions);

    if (!response.ok) {
        const errorMessage = await response.text();
        console.log(errorMessage)
        return
    }

    const responseData = response.json();
    console.log({responseData})
}

export default function Admin() {
    return (
        <div>
            <NavBar active_page={"admin"}/>
            <main className="container">
                <h1 className="my-4 text-center">administrācija</h1>
                <form action="http://localhost:8080/tasks/create" onSubmit={createTask}>
                    <div className="row">
                        <div className="mb-3 col">
                            <label htmlFor="task-code" className="form-label">uzdevuma kods:</label>
                            <input type="text" className="form-control" id="task-code" name="task_code"/>
                        </div>
                        <div className="mb-3 col">
                            <label htmlFor="task-name" className="form-label">uzdevuma nosaukums:</label>
                            <input type="text" className="form-control" id="task-name" name="task_name"/>
                        </div>
                    </div>
                    <button type="submit" className="btn btn-primary">pievienot uzdevumu</button>
                </form>
            </main>
        </div>
    )
}