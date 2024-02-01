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
const replaceBtn = document.getElementById("replacebtn");
const replaceInp = document.getElementById("replaceImg");
replaceBtn.addEventListener("click", function () {
    // Trigger click on the hidden file input when the button is clicked
    replaceInp.click();
});

replaceInp.addEventListener("change", function () {
    // Handle file selection and upload
    replaceImage(this.files[0]);
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
        window.location.reload();
        alert("Image uploaded");
    })
    .catch(error => {
        console.error(error);
        alert("Error while uploading image");
    });
}

function replaceImage(file) {
    const formData = new FormData();
    formData.append("image", file);

    fetch("/replace_idx", {
        method: "POST",
        body: formData,
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        window.location.reload();
        alert("Image uploaded");
    })
    .catch(error => {
        console.error(error);
        alert("Error uploading image");
    });
}

function deleteImage(imgPath) {
    const userConfirmed = window.confirm("Are you sure you want to delete this image?");

    if (!userConfirmed) {
        return;
    }
    const path = imgPath
    fetch("/delete_image", {
        method: "DELETE",
        body: path,
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        alert("Image deleted");
    })
    .catch(error => {
        window.location.reload();
    })
}
