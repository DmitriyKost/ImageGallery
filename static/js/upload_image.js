const uploadBtn = document.getElementById("btn1");
const imageInput = document.getElementById("imageInput");
uploadBtn.addEventListener("click", function () {
    // Trigger click on the hidden file input when the button is clicked
    imageInput.click();
});

imageInput.addEventListener("change", function () {
    // Handle file selection and upload
    uploadImage(this.files[0]);
});

function uploadImage(file) {
    const formData = new FormData();
    formData.append("image", file);

    fetch("http://127.0.0.1:8080/upload/", {
        method: "POST",
        body: formData,
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        alert("Image uploaded");
    })
    .catch(error => {
        console.error(error);
        alert("Error while uploading image");
    });
}
