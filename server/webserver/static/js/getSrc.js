'use strict';
function AJAX(method, url) {
    return new Promise(function (resolve, reject){
        var req = new XMLHttpRequest();
        var handler = function(state) {
            if (req.status == 200) {
                resolve(req.response);
            } else {
                var errorMsg = req.response;
                try {
                    var errorJson = JSON.parse(errorMsg);
                    errorMsg = errorJson["Error"];
                } catch(e) {
                    errorMsg = e.message;
                }
                reject(new Error(errorMsg));
            }
        }
        req.open(method, url, true);
        req.onload = handler;
        req.onerror = () => reject(new Error("Network error."));
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
    var stateElement = document.querySelector("p#state");
    var imageElement = document.querySelector("img.preview");
    var formElement = document.querySelector("form");
    AJAX("GET", "/api/state/" + taskID).then(function(res){
        var json = JSON.parse(res);
        if ("Error" in json) {
            throw new Error(json["Error"]);
        }
        var state = json['StateCode'];
        if (state == -1) {  // Error state
            throw new Error(json['Description']);
        } else if (state == 0) {    // Successful state
            return AJAX("GET", '/api/url/' + json["TaskID"]);
        } else {    // Other state
            stateElement.textContent = "Request " + times.toString() + ", upload state: " + json['Description'] + ", please wait...";
            setTimeout(getPhotoState, 1000, taskID, times + 1);
            return new Promise(() => {});
        }
    }).then(function(res){
        var json = JSON.parse(res);
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
    })
}
