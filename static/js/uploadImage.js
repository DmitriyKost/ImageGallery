const uploadBtn = document.getElementById("uploadbtn");
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

    fetch("/upload", {
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
