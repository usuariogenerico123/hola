
let h3 = document.getElementById("jijo");

h3.innerText="aksjdlajsdk";

let button= document.getElementById("button");
button.addEventListener("click", ()=>{
    
    

    h3.innerText=Math.round(Math.random(1)*100) ;
})

