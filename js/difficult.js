document.addEventListener('DOMContentLoaded', () => {
    const difficultyButtons = document.querySelectorAll('.difficulty-btn');
    const difficultyInfo = document.getElementById('difficulty-info');
       const diff=document.getElementById("playprefferedDifficulty").value;
       
    const difficultyMap = {
        easy: "Easy",
        medium: "Medium",
        hard: "Hard"
    };

    const difficultyDescriptions = {
        easy: "A gentle descent into the shadows. Riddles whisper softly, challenges are kind.",
        medium: "The darkness grows thick. Riddles bite with sharp teeth, testing your resolve.",
        hard: "Pure nightmare fuel. Only the bravest souls dare to tread this path of absolute terror."
    };
     
    if (diff == "Medium"){

            // Remove active state from all buttons
            difficultyButtons.forEach(btn => btn.classList.remove('active'));

        difficultyButtons[1].classList.add("active")
    }
    if(diff == "Hard"){

            // Remove active state from all buttons
            difficultyButtons.forEach(btn => btn.classList.remove('active'));
        difficultyButtons[2].classList.add("active")
    }



    difficultyButtons.forEach(button => {
        button.addEventListener('click', (e) => {
        
            e.preventDefault(); // Prevent default behavior

            const difficulty = button.dataset.difficulty;
            difficultyInfo.textContent = difficultyDescriptions[difficulty];

            const difficultyPayload = {
                SelectedDifficulty: difficultyMap[difficulty]
            };

            // Send the selected difficulty to the server
            fetch("/setdifficult", {
                headers: {
                    "Content-Type": "application/json"
                },
                method: "POST",
                body: JSON.stringify(difficultyPayload)
            })
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));

            // Remove active state from all buttons
            difficultyButtons.forEach(btn => btn.classList.remove('active'));

            // Add active state to clicked button
            button.classList.add('active');

           // location.reload(false);
        });
    });
});
