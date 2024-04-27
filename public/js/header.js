let menubtn = document.querySelector(".menu")
let nav=document.querySelector("nav")

function slide(){
  nav.classList.toggle("slideIn")
  
}

menubtn.addEventListener("click", slide)