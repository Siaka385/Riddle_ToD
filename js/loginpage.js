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

        // Simulated login check 
        const validUsers = [
            { username: 'admin', password: 'password123' },
            { username: 'player', password: 'gamepass456' }
        ];

        const userMatch = validUsers.find(
            user => user.username === username && user.password === password
        );

        if (userMatch) {
            console.log('Login successful', username);
            errorMessage.style.display = 'none';
        } else {
            errorMessage.style.display = 'block';
        }
    });

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