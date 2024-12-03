document.addEventListener('DOMContentLoaded', () => {
    const hintLink = document.querySelector('.hint-link');
    const hintElement = document.getElementById('hint');
    const answerButtons = document.querySelectorAll('.answer-buttons .btn');
    const feedbackElement = document.getElementById('feedback');
    
    // Custom answer elements
    const customAnswerLink = document.querySelector('.custom-answer-link');
    const customAnswerContainer = document.getElementById('custom-answer-container');
    const customAnswerInput = document.getElementById('custom-answer');
    const customAnswerSubmit = document.querySelector('.custom-answer-submit');

    // Hint toggle
    hintLink.addEventListener('click', () => {
        hintElement.style.display = 
            hintElement.style.display === 'none' ? 'block' : 'none';
    });

    // Custom answer link toggle
    customAnswerLink.addEventListener('click', () => {
        customAnswerContainer.style.display = 
            customAnswerContainer.style.display === 'none' ? 'flex' : 'none';
    });

    // Answer selection
    answerButtons.forEach(button => {
        button.addEventListener('click', () => {
            const selectedAnswer = button.textContent.trim();
            
            // Check if the answer is correct
            if (selectedAnswer === 'An Echo') {
                feedbackElement.textContent = 'Correct! Well done!';
                feedbackElement.style.color = 'green';
            } else {
                feedbackElement.textContent = 'Incorrect. Try again!';
                feedbackElement.style.color = 'red';
            }
        });
    });

    // Custom answer submission
    customAnswerSubmit.addEventListener('click', () => {
        const customAnswer = customAnswerInput.value.trim();
        
        if (customAnswer.toLowerCase() === 'echo') {
            feedbackElement.textContent = 'Correct! Well done!';
            feedbackElement.style.color = 'green';
        } else {
            feedbackElement.textContent = 'Incorrect. Try again!';
            feedbackElement.style.color = 'red';
        }

        // Clear input
        customAnswerInput.value = '';
    });
});