document.addEventListener("DOMContentLoaded",()=>{
let classicBtn=document.getElementById("classic")
classicBtn.onclick=()=>{
    window.location.href="/playsection?gamemode=classic";
}

let timeBtn=document.getElementById("time")
timeBtn.onclick=()=>{
    window.location.href="/playsection?gamemode=timeattack"
}

let survivalBtn=document.getElementById("survival")
survivalBtn.onclick=()=>{
    window.location.href="/playsection?gamemode=survival"
}

let difficultyBtn=document.getElementById("difficulty")
difficultyBtn.onclick=()=>{
    window.location.href="/DifficultySetting"
}

})