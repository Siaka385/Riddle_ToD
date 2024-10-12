window.onload = () => {
    //button one player onclick
    var oneplayerbtn = document.getElementById("oneplayer");
    oneplayerbtn.onclick = () => {
      console.log("jjj")
      window.location.href = "one_player/oneplayer.html";
    };
  
    //button two player click
    var twoplayerbtn=document.getElementById("twoplayerbtn");
    twoplayerbtn.addEventListener("click",()=>{
      window.location.assign("two_player/boostraptwo player.html");
    })
   
    //cryptic
    var crypticplayerbtn=document.getElementById("couplebtn");
    crypticplayerbtn.addEventListener("click",()=>{
      window.location.assign("cryptic_horrors/boostrap tod riddlesmenu.html");
    })

  }
  