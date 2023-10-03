document.getElementById("input")
    .addEventListener("submit", async e => {
        e.preventDefault()

        const data = new FormData(e.target);
        const id = Object.fromEntries(data.entries());
        const output = document.getElementById("info__json");
        const response = await fetch(`/orders/${id.uid}`);

        if (response.ok) {
            const data = await response.json();
            output.textContent = JSON.stringify(data, undefined, 2);
        } else {
            output.innerHTML = `Invalid order uid`
        }
    })