window.addEventListener("load", function load(event){

    var raw = document.getElementById("raw");
    var feedback = document.getElementById("feedback");
    
    var do_export = function(){

	feedback.innerText = ""

	try {
	    var f = JSON.parse(raw.innerText);
	} catch(err) {
	    feedback.innerText = "Failed to parse feature: " + err;
	    return;
	}

	var str_f = JSON.stringify(f);
	
	wof_validate(str_f).then(rsp => {
	    feedback.innerText = "Document is valid.";
	}).catch(err => {
	    feedback.innerText = "Failed to validate feature: " + err;
	});
    };
    
    var init = function(){

	var btn = document.getElementById("submit");

	if (! btn){
	    console.log("Unable to load submit button");
	    return;
	}

	btn.onclick = function(){
	    do_export();
	    return false;
	};

	btn.innerText = "Validate";	
	btn.removeAttribute("disabled");
    };

     // https://github.com/sfomuseum/js-sfomuseum-golang-wasm    
    sfomuseum.golang.wasm.fetch("/wasm/wof_validate.wasm").then(rsp => {
	init();	
    });
    
});
