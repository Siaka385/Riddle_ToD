document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('login-form');
    const errorMessage = document.getElementById('error-message');
    const btnGuest = document.getElementById('btn-guest');
    const btnSignup = document.getElementById('btn-signup');
    const forgotPasswordLink = document.getElementById('forgot-password');
    const sendRecoveryBtn = document.getElementById('send-recovery-btn');
    const recoveryEmailInput = document.getElementById('recovery-email');

    // Bootstrap Modal
    const forgotPasswordModal = new bootstrap.Modal(document.getElementById('forgotPasswordModal'));

    loginForm.addEventListener('submit', (e) => {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

      
        var user={
            Username:username,
            Password:password
        }


        fetch("/login",{
            headers: { 'Content-Type': 'application/json' },
            method:"POST",
            body:JSON.stringify(user)
        }).then(response => response.json()).then(data => handleServerResponse(data.response)).catch(error => console.log(error))

    });

function handleServerResponse(message){

    if(message == "Username does not exist"){
        errorMessage.textContent="Username does not exist"
        errorMessage.style.display = 'block';
        scrollToTop()
        return
      }
      if(message == "Incorrect password"){
          errorMessage.textContent="Incorrect password"
          errorMessage.style.display = 'block';
          scrollToTop()
        return
      }
     
      console.log(message)
    

}

function scrollToTop() {
    window.scrollTo({
        top: 0,        // Scroll to the top
        behavior: 'smooth' // Smooth scrolling
    });
}

    // Forgot Password Link
    forgotPasswordLink.addEventListener('click', (e) => {
        e.preventDefault();
        forgotPasswordModal.show();
    });

    // Send Recovery Link
    sendRecoveryBtn.addEventListener('click', () => {
        const email = recoveryEmailInput.value;
        
        // Simulated recovery process
        if (email) {
            console.log('Recovery link sent to:', email);
            alert('A password recovery link has been sent to ' + email);
            forgotPasswordModal.hide();
        } else {
            alert('Please enter a valid email address');
        }
    });

    // Signup and Guest buttons
    btnGuest.addEventListener('click', () => {
        console.log('Playing as Guest');
    });

    btnSignup.addEventListener('click', () => {
        window.location.href = 'signup2.html';
    });
});