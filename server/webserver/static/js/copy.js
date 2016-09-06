'use strict';
function addCopyEventListener(url){
    let divs = document.querySelectorAll("div.copyAsText");
    for (let div of divs) {
        let input = div.querySelector("input");
        let btn = div.querySelector("label");
        if (url !== "") {
            let template = input.getAttribute("data-template");
            input.value = template.replace("${url}", url);
        }
        btn.addEventListener("click", function(){
            if (btn.disabled) {
                return;
            }
            let res = false;
            try {
                input.disabled = false;
                let section = window.getSelection();
                section.removeAllRanges();
                input.focus()
                input.setSelectionRange(0, input.value.length);
                res = document.execCommand("copy");
                console.log("res =", res)
                input.disabled = true
            } catch(e) {
                res = false;
            }
            if (res) {
                let origin = btn.textContent;
                btn.textContent = "Copied!";
                btn.disabled = true;
                setTimeout(function(){
                    btn.textContent = origin;
                    btn.disabled = false;
                }, 2000);
            } else {
                window.prompt("Copy to clipboard: Ctrl+C, Enter", input.value);
            }
        });
    }
}
