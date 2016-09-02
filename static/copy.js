function addCopyEventListener(){
    let divs = document.querySelectorAll("div.copyAsText");
    for (let div of divs) {
        let input = div.querySelector("input");
        let btn = div.querySelector("label");
        btn.addEventListener("click", function(){
            if (btn.disabled) {
                return
            }
            input.disable = false;
            let range = document.createRange();
            let section = window.getSelection();
            section.removeAllRanges();
            range.selectNode(input);
            section.addRange(range);
            let res = document.execCommand("copy");
            section.removeAllRanges();
            input.disable = true;
            if (res) {
                let origin = btn.textContent
                while (btn.firstChild) {
                    btn.removeChild(btn.firstChild);
                }
                btn.insertAdjacentText("afterbegin", "Copyed!");
                btn.disabled = true;
                setTimeout(function(){
                    while (btn.firstChild) {
                        btn.removeChild(btn.firstChild);
                    }
                    btn.insertAdjacentText("afterbegin", origin);
                    btn.disabled = false;
                }, 2000);
            } else {
                window.prompt("Copy to clipboard: Ctrl+C, Enter", input.value);
            }
        });
    }
}
