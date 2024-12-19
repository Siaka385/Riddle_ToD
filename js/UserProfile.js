document.addEventListener('DOMContentLoaded', () => {
    const avatarOptions = document.querySelectorAll('.avatar-option');
    const hiddenAvatarInput = document.getElementById('selected-avatar');
    const usernameInput = document.getElementById('username-input');
    const emailInput = document.getElementById('email-input');
    const currentPasswordInput = document.getElementById('current-password');
    const newPasswordInput = document.getElementById('new-password');
    const confirmPasswordInput = document.getElementById('confirm-password');
    const profileEditContainer = document.querySelector('.profile-edit-container');

    // Create success message div
    function createSuccessMessage() {
        const successDiv = document.createElement('div');
        successDiv.className = 'success-overlay';
        successDiv.innerHTML = `
            <div class="success-message">
                <h2>Profile Updated</h2>
                <p>Your profile changes have been saved successfully!</p>
                <div class="profile-actions">
                    <button class="btn btn-save-profile dismiss-success">Continue</button>
                </div>
            </div>
        `;
        
        // Style the success message overlay
        successDiv.style.cssText = `
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0, 0, 0, 0.8);
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 1000;
        `;

        // Style the success message content
        const successMessage = successDiv.querySelector('.success-message');
        successMessage.style.cssText = `
            background: rgba(0, 0, 0, 0.7);
            border: 2px solid #8B0000;
            border-radius: 15px;
            padding: 30px;
            max-width: 500px;
            width: 90%;
            text-align: center;
            box-shadow: 0 0 30px rgba(139, 0, 0, 0.5);
        `;

        // Style the heading
        const heading = successMessage.querySelector('h2');
        heading.style.cssText = `
            font-family: 'Nosifer', cursive;
            color: #8B0000;
            text-shadow: 0 0 15px rgba(139, 0, 0, 0.7);
            margin-bottom: 20px;
        `;

        // Style the paragraph
        const paragraph = successMessage.querySelector('p');
        paragraph.style.cssText = `
            color: white;
            margin-bottom: 20px;
        `;

        // Add dismiss functionality
        successDiv.querySelector('.dismiss-success').addEventListener('click', () => {
            successDiv.remove();
        });

        // Add to the document
        document.body.appendChild(successDiv);
    }

    // Avatar Selection Logic
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

    // Username Edit
    usernameInput.addEventListener('focus', (e) => {
        e.target.select();
    });

    // Save Profile Functionality
    const saveProfileBtn = document.querySelector('.btn-save-profile');
    const cancelBtn = document.querySelector('.btn-cancel');

    saveProfileBtn.addEventListener('click', () => {
        // Validate inputs
        const newUsername = usernameInput.value.trim();
        const newEmail = emailInput.value.trim();
        const currentPassword = currentPasswordInput.value;
        const newPassword = newPasswordInput.value;
        const confirmPassword = confirmPasswordInput.value;
        const selectedAvatar = hiddenAvatarInput.value;

        // Basic validation
        if (!newUsername) {
            alert('Username cannot be empty');
            return;
        }

        if (!newEmail || !newEmail.includes('@')) {
            alert('Please enter a valid email address');
            return;
        }

        // Password change validation (optional)
        if (newPassword || confirmPassword) {
            if (newPassword !== confirmPassword) {
                alert('New passwords do not match');
                return;
            }
            
            if (newPassword.length < 8) {
                alert('New password must be at least 8 characters long');
                return;
            }
        }

        // In a real application, you would send this data to a backend
        const profileData = {
            username: newUsername,
            email: newEmail,
            avatar: `https://api.dicebear.com/7.x/adventurer/svg?seed=${selectedAvatar}&backgroundType=gradientLinear`,
            currentPassword: currentPassword,
            newPassword: newPassword || null
        };

        console.log('Saving Profile:', profileData);
        
        // Create and show success message
        createSuccessMessage();

        // Clear password fields after submission
        currentPasswordInput.value = '';
        newPasswordInput.value = '';
        confirmPasswordInput.value = '';
    });

    cancelBtn.addEventListener('click', () => {
        // Reset to original values or navigate back
        window.location.href = 'main-page.html'; // Adjust the redirect as needed
    });
});