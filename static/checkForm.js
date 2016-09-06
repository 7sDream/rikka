'use strict';

function check(maxSizeByMb) {
    let fileInput = document.querySelector("input#uploadFile");
    let file = fileInput.files[0]
    if (file.type.startsWith("image") === false) {
        let fileType = file.type || "unknown"
        alert("Can't upload a " + file.type + " type file")
        return false;
    }
    if (file.size > (maxSizeByMb * 1024 * 1024)) {
        let fileSizeByMb = Math.round(file.size / 1024 / 1024 * 100) / 100;
        alert("Max file size is " + maxSizeByMb + " Mb, input file is " + fileSizeByMb.toString() + " Mb");
        return false;
    }
    let imgElement = document.querySelector("img#uploading");
    imgElement.classList.add("show");
    return true;
}
