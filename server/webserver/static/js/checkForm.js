'use strict';

function isImageType(typeStr) {
    if (typeStr.startsWith("image") === false) {
        return false;
    }
    var accepted = ["jpeg", "bmp", "gif", "png"];
    return accepted.some(function(type){return typeStr.endsWith("/" + type)})
}

function check(maxSizeByMb) {
    var passwordInput = document.querySelector("input#password")
    var fileInput = document.querySelector("input#uploadFile");
    var file = fileInput.files[0];
    if (passwordInput.value === "") {
        alert("Please input password");
        return false;
    }
    if (file === undefined) {
        alert("Plesae choose a image to upload");
        return false;
    }
    var fivarype = file.type
    if (!isImageType(fivarype)) {
        fivarype = fivarype || "unknown";
        alert("Can't upload a " + fivarype + " type file");
        return false;
    }
    console.log("Accept a", fivarype, "file")
    if (file.size > (maxSizeByMb * 1024 * 1024)) {
        var fileSizeByMb = Math.round(file.size / 1024 / 1024 * 100) / 100;
        alert("Max file size is " + maxSizeByMb + " Mb, input file is " + fileSizeByMb.toString() + " Mb");
        return false;
    }
    var imgElement = document.querySelector("img#uploading");
    imgElement.classList.add("show");
    return true;
}
