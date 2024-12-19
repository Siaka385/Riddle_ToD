document.addEventListener('DOMContentLoaded', () => {
    const avatarOptions = document.querySelectorAll('.avatar-option');
    const hiddenAvatarInput = document.getElementById('selected-avatar');
    const usernameInput = document.getElementById('username-input');

    // Create a reusable overlay message function
    function showOverlayMessage(title, message) {
        const overlayDiv = document.createElement('div');
        overlayDiv.className = 'message-overlay';
        overlayDiv.innerHTML = `
            <div class="success-message">
                <h2>${title}</h2>
                <p>${message}</p>
                <div class="profile-actions">
                    <button class="btn btn-save-profile dismiss-message">Continue</button>
                </div>
            </div>
        `;
        
        // Style the overlay
        overlayDiv.style.cssText = `
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

        // Style the message box
        const messageBox = overlayDiv.querySelector('.success-message');
        messageBox.style.cssText = `
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
        const heading = messageBox.querySelector('h2');
        heading.style.cssText = `
            font-family: 'Nosifer', cursive;
            color: #8B0000;
            text-shadow: 0 0 15px rgba(139, 0, 0, 0.7);
            margin-bottom: 20px;
        `;

        // Style the paragraph
        const paragraph = messageBox.querySelector('p');
        paragraph.style.cssText = `
            color: white;
            margin-bottom: 20px;
        `;

        // Add dismiss functionality
        overlayDiv.querySelector('.dismiss-message').addEventListener('click', () => {
            overlayDiv.remove();
        });

        // Add to document
        document.body.appendChild(overlayDiv);
    }

    // Avatar Selection Logic
    avatarOptions.forEach(avatar => {
        avatar.addEventListener('click', () => {
            avatarOptions.forEach(a => a.classList.remove('selected'));
            avatar.classList.add('selected');
            hiddenAvatarInput.value = avatar.dataset.avatar;
        });
    });

    // Email Update Logic
    const emailUpdateBtn = document.querySelector('.btn-save-email');
    emailUpdateBtn.addEventListener('click', () => {
        const currentEmail = document.getElementById('current-email').value;
        const newEmail = document.getElementById('new-email').value;
        const confirmEmail = document.getElementById('confirm-email').value;

        // Validate emails
        if (!newEmail || !confirmEmail) {
            showOverlayMessage('Error', 'Please fill in all email fields');
            return;
        }

        if (!newEmail.includes('@')) {
            showOverlayMessage('Error', 'Please enter a valid email address');
            return;
        }

        if (newEmail !== confirmEmail) {
            showOverlayMessage('Error', 'New emails do not match');
            return;
        }

        if (newEmail === currentEmail) {
            showOverlayMessage('Error', 'New email must be different from current email');
            return;
        }

        // If validation passes, show success message and clear fields
        showOverlayMessage('Email Updated', 'Your email has been successfully updated!');
        document.getElementById('new-email').value = '';
        document.getElementById('confirm-email').value = '';
    });

    // Password Update Logic
    const passwordUpdateBtn = document.querySelector('.btn-save-password');
    passwordUpdateBtn.addEventListener('click', () => {
        const currentPassword = document.getElementById('current-password').value;
        const newPassword = document.getElementById('new-password').value;
        const confirmPassword = document.getElementById('confirm-password').value;

        // Validate passwords
        if (!currentPassword || !newPassword || !confirmPassword) {
            showOverlayMessage('Error', 'Please fill in all password fields');
            return;
        }

        if (newPassword.length < 8) {
            showOverlayMessage('Error', 'New password must be at least 8 characters long');
            return;
        }

        if (newPassword !== confirmPassword) {
            showOverlayMessage('Error', 'New passwords do not match');
            return;
        }

        if (newPassword === currentPassword) {
            showOverlayMessage('Error', 'New password must be different from current password');
            return;
        }

        // If validation passes, show success message and clear fields
        showOverlayMessage('Password Updated', 'Your password has been successfully updated!');
        document.getElementById('current-password').value = '';
        document.getElementById('new-password').value = '';
        document.getElementById('confirm-password').value = '';
    });

    // Main Profile Save Logic
    const saveProfileBtn = document.querySelector('.btn-save-profile');
    const cancelBtn = document.querySelector('.btn-cancel');

    saveProfileBtn.addEventListener('click', () => {
        const newUsername = usernameInput.value.trim();
        const selectedAvatar = hiddenAvatarInput.value;

        // Validate username
        if (!newUsername) {
            showOverlayMessage('Error', 'Username cannot be empty');
            return;
        }

        // Save profile data
        const profileData = {
            username: newUsername,
            avatar: `https://api.dicebear.com/7.x/adventurer/svg?seed=${selectedAvatar}&backgroundType=gradientLinear`
        };

        console.log('Saving Profile:', profileData);
        showOverlayMessage('Profile Updated', 'Your profile has been successfully updated!');
    });

    cancelBtn.addEventListener('click', () => {
        window.location.href = 'main-page.html';
    });

    // Username Edit Focus
    usernameInput.addEventListener('focus', (e) => {
        e.target.select();
    });
});