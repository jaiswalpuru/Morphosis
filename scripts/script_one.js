let headers = new Headers();
headers.append("Content-Type","application/json")

fetch("http://localhost:8000/images", {
    method : 'GET',
    headers: headers,
    body : null,
})
.then(response => response.json())
.then(data => {
    document.getElementById('img_one').src = data
})
.catch((error) => {
    console.error('Error', error)
});