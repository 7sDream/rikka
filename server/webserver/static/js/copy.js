'use strict';

function addCopyEventListener(url){
    const divs = document.querySelectorAll("div.copyAsText");

    for (const index in divs) {

        if(!divs.hasOwnProperty(index)) {
            continue;
        }

        const div = divs[index];
        const input = div.querySelector("input");
        const btn = div.querySelector("label");

        if (url !== "") {
            const template = input.getAttribute("data-template");
            input.value = template.replace("${url}", url);
        }

        void function(btn, input) {
            btn.addEventListener("click", function(){
                if (btn.disabled) {
                    return;
                }

                let res = false;
                try {
                    const section = window.getSelection();
                    section.removeAllRanges();

                    input.disabled = false;

                    input.focus();
                    input.setSelectionRange(0, input.value.length);

                    res = document.execCommand("copy");

                    input.disabled = true
                } catch(e) {
                    res = false;
                }
                if (res) {
                    const origin = btn.textContent;

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
        }(btn, input);
    }
}
