let host = "http://127.0.0.1:8000"

async function pass() {
    let url = host + "/pass"
    fetch(url)
        .then(response => response.json())
        .then(data => {
            document.getElementById("step0").value = data["PrivateKey"]
            document.getElementById("step1").value = data["PublicKey"]
        })
}