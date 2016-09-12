'use strict';
function addCopyEventListener(url){
    var divs = document.querySelectorAll("div.copyAsText");
    for (var index in divs) {
        if(!divs.hasOwnProperty(index)) {
            continue;
        }
        var div = divs[index];
        var input = div.querySelector("input");
        var btn = div.querySelector("label");
        if (url !== "") {
            var template = input.getAttribute("data-template");
            input.value = template.replace("${url}", url);
        }
        btn.addEventListener("click", function(){
            if (btn.disabled) {
                return;
            }
            var res = false;
            try {
                input.disabled = false;
                var section = window.getSelection();
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
                var origin = btn.textContent;
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
