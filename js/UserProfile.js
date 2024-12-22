document.addEventListener('DOMContentLoaded', () => {
    // Previous avatar gallery code remains the same
    const avatarGallery = document.getElementById('avatar-gallery');
    const currentAvatarImg = document.getElementById('current-avatar-img');
    const saveButton = document.getElementById('save-profile');
    const messageOverlay = document.getElementById('message-overlay');
    const showPasswordForm = document.getElementById('show-password-form');
    const passwordSection = document.getElementById('password-section');
    const savePassword = document.getElementById('save-password');
    const backbutton=document.getElementById("backbtn")
    let avatar;

    //back button
    backbutton.onclick=()=>{
        window.location.href="/"
    }
    
    // Generate avatar options
    for (let i = 1; i <= 6; i++) {
        const avatarOption = document.createElement('img');
        avatarOption.src = `https://api.dicebear.com/7.x/adventurer/svg?seed=${i}`;
        avatarOption.classList.add('avatar-option');
        avatarOption.dataset.avatarId = i;
        
        avatarOption.addEventListener('click', () => {
            document.querySelectorAll('.avatar-option').forEach(opt => {
                opt.classList.remove('selected');
            });
            avatarOption.classList.add('selected');
            currentAvatarImg.src = avatarOption.src;
            avatar=avatarOption.dataset.avatarId
        });
        
        avatarGallery.appendChild(avatarOption);
    }

    // Toggle password change section
    showPasswordForm.addEventListener('click', () => {
        passwordSection.classList.toggle('active');
        showPasswordForm.textContent = passwordSection.classList.contains('active') 
            ? 'Hide Password Form' 
            : 'Change Password';
    });

    // Handle profile save
    saveButton.addEventListener('click', () => {
        const username = document.getElementById('username').value;
        const email = document.getElementById('email').value;

        if (!username || !email) {
            showMessage('Error', 'Username and email are required.');
            return;
        }

        var UserUpdateDetails={
                NewUsername:username,
                NewEmail:email,
                NewAvatar:avatar

        }

        fetch("/updateprofile",{
            method:"POST",
            headers:{ 'Content-Type': 'application/json' },
            body:JSON.stringify(UserUpdateDetails)

        }).then(response => response.json()).then(data => Profile(data.message)).catch(err=> console.log(err))

    });

    function Profile(message){
        console.log(message)
         if(message == "Cannot change to this email, it already exists." ){
            showMessage("Failed to change email",message)
              return
         }
         if(message == "Cannot change to this username, it already exists."){
            showMessage("Failed to change Username",message)
                return
         }
         showMessage('Success', 'Profile updated successfully!');

    }

    // Handle password save
    savePassword.addEventListener('click', () => {
        const currentPassword = document.getElementById('current-password').value;
        const newPassword = document.getElementById('new-password').value;
        const confirmPassword = document.getElementById('confirm-password').value;

        if (!currentPassword || !newPassword || !confirmPassword) {
            showMessage('Error', 'All password fields are required.');
            return;
        }

        if (newPassword !== confirmPassword) {
            showMessage('Error', 'New passwords do not match.');
            return;
        }

        var UserUpdatePassword={
           OldPassword:currentPassword,
           NewPassword:newPassword,
           ConfirmPassword:confirmPassword

    }
    console.log(UserUpdatePassword)
       var isPasswordUpdate=false
    fetch("/updatepassword",{
        method:"POST",
        headers:{ 'Content-Type': 'application/json' },
        body:JSON.stringify(UserUpdatePassword)
   
    }).then(response => response.json()).then((data) =>{isPasswordUpdate=password(data.message)}).catch(err=> console.log(err))
          
        if (isPasswordUpdate){
        // Clear password fields
        document.getElementById('current-password').value = '';
        document.getElementById('new-password').value = '';
        document.getElementById('confirm-password').value = '';
        
        // Hide password section
        passwordSection.classList.remove('active');
        showPasswordForm.textContent = 'Change Password';
        }
    });
});

function password(message){
    if (message == "Your old password is incorrect, password update failed"){
          showMessage("failed to update password",message)
          return false
    }
    showMessage("success",message)
    return true
}



function showMessage(title, text) {
    const messageOverlay = document.getElementById('message-overlay');
    const messageTitle = document.getElementById('message-title');
    const messageText = document.getElementById('message-text');
    
    messageTitle.textContent = title;
    messageText.textContent = text;
    messageOverlay.style.display = 'flex';
}

function closeMessage() {
    document.getElementById('message-overlay').style.display = 'none';
}
