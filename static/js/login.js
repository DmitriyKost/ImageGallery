function openPopup() {
    document.getElementById("overlay").style.display = "flex";
}

function closePopup() {
    document.getElementById("overlay").style.display = "none";
}

document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        closePopup();
    }
});


function login() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    var credentials = {
        username: username,
        password: password
    };

    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Invalid credentials');
        }
        return response;
    })
    .then(data => {
        console.log('Login successful:', data);
        closePopup();
        window.location.href = "/admin";
    })
    .catch(error => {
        console.error('Login failed:', error.message);
        alert(error.message);
    });
}
