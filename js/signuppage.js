 // Avatar Selection Logic
 document.addEventListener('DOMContentLoaded', () => {
    const avatarOptions = document.querySelectorAll('.avatar-option');
    const hiddenAvatarInput = document.getElementById('selected-avatar');
    const errorMessage = document.getElementById('error-message');
    const signupForm = document.getElementById('signup-form');

    hiddenAvatarInput.value="1";
    
    // Avatar Selection
    avatarOptions.forEach(avatar => {
        avatar.addEventListener('click', () => {
            // Remove selected class from all avatars
            avatarOptions.forEach(a => a.classList.remove('selected'));
            
            // Add selected class to clicked avatar
            avatar.classList.add('selected');
            
            // Set the hidden input value to the selected avatar's data-avatar
            hiddenAvatarInput.value = avatar.dataset.avatar;
        });
    });

    // Form Submission Handler
    signupForm.addEventListener('submit', (e) => {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const confirmPassword = document.getElementById('confirm-password').value;
        const selectedAvatar = hiddenAvatarInput.value;

        // Basic client-side validation
        if (password !== confirmPassword) {
           errorMessage.textContent="passwords do not match, make";
           errorMessage.style.display='block';
            return;
        }

        // Simulated username/email check 
        // In a real application, this would be a server-side check
        const existingUsers = ['admin', 'test@example.com'];
        if (existingUsers.includes(username) || existingUsers.includes(email)) {
            errorMessage.style.display = 'block';
            return;
        }

        var User={
            Username:username,
            Email:email,
            Password:password,
            ConfirmPass:confirmPassword,
            SelectedAvatar:selectedAvatar
        }

        fetch("/register",{
              method:"POST",
              headers: { 'Content-Type': 'application/json' },
              body:JSON.stringify(User)

        }).then(response => response.json()).then(data => handleServerResponse(data.response)).catch(error => console.log(error))
        // If all checks pass, you would normally send data to server
       

        // Reset error message
        errorMessage.style.display = 'none';
    });


//check msg 

function handleServerResponse(message){

  if(message == "Email already exists"){
    errorMessage.textContent="Email Already Exist"
    errorMessage.style.display = 'block';
    scrollToTop()
    return
  }
  if(message == "Username already exists"){
      errorMessage.textContent="Username Already Exist"
      errorMessage.style.display = 'block';
      scrollToTop()
    return
  }

  if(message == "Email is not valid"){
    errorMessage.textContent="Email is not valid"
    errorMessage.style.display = 'block';
    scrollToTop()
  return
}

  if (message == "Passwords do not match"){
      errorMessage.textContent="Passwords do not match"
      errorMessage.style.display = 'block';
      scrollToTop()
      return
  }

  if (message == "ok"){
      window.location.href="/loginpage"
  }

// setTimeout(()=>{
//     window.location.href="/dashboard"
// },1000)
}

function scrollToTop() {
    window.scrollTo({
        top: 0,        // Scroll to the top
        behavior: 'smooth' // Smooth scrolling
    });
}




});