'use strict';

function isImageType(typeStr) {
    if (typeStr.startsWith("image") === false) {
        return false;
    }
    let accepted = ["jpeg", "bmp", "gif", "png"];
    return accepted.some((type) => typeStr.endsWith("/" + type))
}

function check(maxSizeByMb) {
    let passwordInput = document.querySelector("input#password")
    let fileInput = document.querySelector("input#uploadFile");
    let file = fileInput.files[0];
    if (passwordInput.value === "") {
        alert("Please input password");
        return false;
    }
    if (file === undefined) {
        alert("Plesae choose a image to upload");
        return false;
    }
    let fileType = file.type
    if (!isImageType(fileType)) {
        fileType = fileType || "unknown";
        alert("Can't upload a " + fileType + " type file");
        return false;
    }
    console.log("Accept a", fileType, "file")
    if (file.size > (maxSizeByMb * 1024 * 1024)) {
        let fileSizeByMb = Math.round(file.size / 1024 / 1024 * 100) / 100;
        alert("Max file size is " + maxSizeByMb + " Mb, input file is " + fileSizeByMb.toString() + " Mb");
        return false;
    }
    let imgElement = document.querySelector("img#uploading");
    imgElement.classList.add("show");
    return true;
}
