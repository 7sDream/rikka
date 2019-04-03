'use strict';

function AJAX(method, url) {
    // noinspection JSUnresolvedFunction
    return new Promise(function (resolve, reject){

        const req = new XMLHttpRequest();

        req.onload = function () {
            if (req.status === 200) {
                resolve(req.response);
            } else {
                let errorMsg = req.response;
                try {
                    // noinspection JSCheckFunctionSignatures
                    const errorJson = JSON.parse(errorMsg);
                    errorMsg = errorJson["Error"];
                } catch (e) {
                    errorMsg = e.message;
                }
                reject(new Error(errorMsg));
            }
        };

        req.onerror = function(){ reject(new Error("Network error.")); };

        req.open(method, url, true);
        req.send();
    });
}

function hide(elem) {
    elem.classList.add("hide");
}

function show(elem) {
    elem.classList.remove("hide");
}

function getPhotoState(taskID, times) {

    times = times || 0;

    const stateElement = document.querySelector("p#state");
    const imageElement = document.querySelector("img.preview");
    const formElement = document.querySelector("form");

    // noinspection JSUnresolvedFunction
    AJAX("GET", "/api/state/" + taskID).then(function (res) {

        const json = JSON.parse(res);

        if ("Error" in json) {
            throw new Error(json["Error"]);
        }

        const state = json['StateCode'];

        if (state === -1) {  // Error state
            throw new Error(json['Description']);
        } else if (state === 0) {    // Successful state
            return AJAX("GET", '/api/url/' + json["TaskID"]);
        } else {    // Other state
            stateElement.textContent = "Request " + times.toString() + ", upload state: " + json['Description'] + ", please wait...";
            setTimeout(getPhotoState, 1000, taskID, times + 1);
            // noinspection JSUnresolvedFunction
            return new Promise(function(){});
        }

    }).then(function(res){

        const json = JSON.parse(res);

        if ("Error" in json) {
            throw new Error(json["Error"]);
        }

        return json["URL"];

    }).then(function (url){

        imageElement.src = url;

        hide(stateElement);
        show(imageElement);
        show(formElement);

        addCopyEventListener(url);

    }).catch(function(err){

        stateElement.textContent = "Error: " + err.message + ", please close page";

        show(stateElement);
        hide(imageElement);
        hide(formElement);

    });
}
