window.onload = () => {
    //button one player onclick
    var oneplayerbtn = document.getElementById("oneplayer");
    oneplayerbtn.onclick = () => {
      ChooseGameMode()
    };

    //back to home button
    var backbtn=document.getElementById("back")
    backbtn.onclick=()=>{
       backtohome()
    }

    //button to logout
    var logout=document.getElementById("logout")
    logout.onclick=()=>{
       window.location.href="/logout"
            }

      //button to profile editing
      var profilebtn=document.getElementById("profile-settings")      
       profilebtn.onclick=()=>{
        window.location.href="/profileEdit"
       }



  
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

    //help
    var help=document.getElementById("helpbtn")
    help.addEventListener("click",()=>{
      window.location.href="/help"
    })

  }

//proceed to gamemode
  function ChooseGameMode(){
    var btn=document.getElementById("home")
    btn.classList.add("d-none")
    document.getElementById("selectriddlecategory").classList.remove("d-none")
  }

  //proceed to home
  function backtohome(){
    let back=document.getElementById("home")
    back.classList.remove("d-none")
    document.getElementById("selectriddlecategory").classList.add("d-none")
  }

 