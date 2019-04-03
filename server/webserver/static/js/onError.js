'use strict';

function errorHandler(message, source, lineNumber, colNumber, error) {

    const errorDiv = document.getElementById("error");

    errorDiv.classList.remove("hide");

    console.log("Error happened, message:", message);
    console.log("On source file: ", source);
    console.log("On line - col: ", lineNumber, "-", colNumber);
    console.log("Error:", error);

    try {
        const ua = navigator.userAgent;
        console.log("UA: ", ua);
    } catch (e) {
        console.log("Unable to get UA");
    }

}

window.onerror = errorHandler;
