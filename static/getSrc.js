function AJAX(method, url) {
    return new Promise(function (resolve, reject){
        let req = new XMLHttpRequest();
        let handler = (state) => {
            if (req.status == 200) {
                resolve(req.response);
            } else {
                reject(new Error(req.responseText));
            }
        }
        req.open(method, url, true);
        req.onload = handler;
        req.onerror = () => reject(new Error("Network error."));
        req.send();
    });
}

function hide(elem) {
    elem.classList.add("hide")
}

function show(elem) {
    elem.classList.remove("hide")
}

function getPhotoState(taskID) {
    let stateElement = document.querySelector("p#state");
    let imageElement = document.querySelector("img.preview")
    let formElement = document.querySelector("form");

    AJAX("GET", "/api/state/" + taskID).then(function(res){
        let json = JSON.parse(res);
        let state = json['StateCode'];

        if (state == -1) {  // Error state
            throw new Error(json['Description']);
        } else if (state == 0) {    // Successful state
            return json;
        } else {    // Other state
            stateElement.textContent = "State:" + json['Description'] + ", please wait...";
            setTimeout(getPhotoState, 1000, taskID);
            return new Promise(() => {})
        }
    }).then(function(json){
        return json["TaskID"];
    }).then(function (taskID){
        return AJAX("GET", '/api/url/' + taskID);
    }).then(function(res){
        let json = JSON.parse(res);
        console.log(json);
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
