'use strict';

function isImageType(typeStr) {
    if (typeStr === "") {
        return false;
    }
    if (typeStr.startsWith("image") === false) {
        return false;
    }
    if (typeStr.includes("vnd")) {
        return false;
    }
    let accepted = ["/jpeg", "/bmp", "/gif", "/png"];
    return accepted.some((type) => typeStr.includes(type))
}

function check(maxSizeByMb) {
    let fileInput = document.querySelector("input#uploadFile");
    let file = fileInput.files[0];
    let fileType = file.type
    if (!isImageType(fileType)) {
        fileType = fileType || "unknown";
        alert("Can't upload a " + fileType + " type file");
        return false;
    }
    console.log("Accept a", fileTypem, "file")
    if (file.size > (maxSizeByMb * 1024 * 1024)) {
        let fileSizeByMb = Math.round(file.size / 1024 / 1024 * 100) / 100;
        alert("Max file size is " + maxSizeByMb + " Mb, input file is " + fileSizeByMb.toString() + " Mb");
        return false;
    }
    let imgElement = document.querySelector("img#uploading");
    imgElement.classList.add("show");
    return true;
}
