document.addEventListener('DOMContentLoaded', () => {
    // Game Configuration
    const TOTAL_TIME = 30; // Total game time in seconds
    const TOTAL_CHALLENGES = 10; // Total number of challenges in time attack mode
    const CORRECT_POINTS = 10; // Points for each correct answer

    // DOM Elements
    const timerDisplay = document.querySelector('.timer');
    const scoreDisplay = document.querySelector('.current-score');
    const correctCountDisplay = document.querySelector('.correct-count');
    const totalChallengesDisplay = document.querySelector('.total-challenges');
    const riddleText = document.querySelector('.riddle-text');
    const answerButtons = document.querySelectorAll('.answer-btn');
    const feedbackElement = document.getElementById('feedback');
    const hintLink = document.querySelector('.hint-link');
    const customAnswerLink = document.querySelector('.custom-answer-link');

    // Game State
    let timeRemaining = TOTAL_TIME;
    let score = 0;
    let correctAnswers = 0;
    let gameActive = false;
    let timerInterval;

    // Riddle Bank for Time Attack
    const riddleBank = [
        {
            question: "What travels around the world but stays in one corner?",
            answers: ["A Stamp"],
            hint: "It's something you might find on an envelope"
        },
        {
            question: "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I?",
            answers: ["An Echo"],
            hint: "I repeat what you say without a mouth"
        },
        {
            question: "What has keys, but no locks; space, but no room; you can enter, but not go in?",
            answers: ["A Keyboard"],
            hint: "You use it to type"
        },
        {
            question: "What comes once in a minute, twice in a moment, but never in a thousand years?",
            answers: ["The Letter M"],
            hint: "Look carefully at the word 'moment'"
        },
        {
            question: "The more you take, the more you leave behind. What am I?",
            answers: ["Footsteps"],
            hint: "Think about walking"
        }
    ];

    // Initialize Game
    function initializeGame() {
        // Reset game state
        timeRemaining = TOTAL_TIME;
        score = 0;
        correctAnswers = 0;
        gameActive = true;

        // Update displays
        timerDisplay.textContent = timeRemaining;
        scoreDisplay.textContent = score;
        correctCountDisplay.textContent = correctAnswers;
        totalChallengesDisplay.textContent = TOTAL_CHALLENGES;
        feedbackElement.textContent = '';

        // Start timer
        startTimer();

        // Load first challenge
        loadNextChallenge();
    }

    // Timer Functionality
    function startTimer() {
        timerInterval = setInterval(() => {
            timeRemaining--;
            timerDisplay.textContent = timeRemaining;

            // Check if time is up
            if (timeRemaining <= 0) {
                endGame();
            }
        }, 1000);
    }

    // Load Next Challenge
    function loadNextChallenge() {
        if (correctAnswers >= TOTAL_CHALLENGES) {
            endGame();
            return;
        }

        // Randomly select a riddle
        const riddle = riddleBank[Math.floor(Math.random() * riddleBank.length)];
        
        // Update riddle text
        riddleText.textContent = riddle.question;

        // Shuffle answer buttons
        const answers = [...riddle.answers];
        const wrongAnswers = riddleBank
            .filter(r => r !== riddle)
            .map(r => r.answers[0])
            .sort(() => 0.5 - Math.random())
            .slice(0, 3);

        const allAnswers = [...answers, ...wrongAnswers].sort(() => 0.5 - Math.random());

        // Update answer buttons
        answerButtons.forEach((btn, index) => {
            btn.textContent = allAnswers[index];
            btn.classList.remove('correct', 'incorrect');
        });
    }

    // Answer Selection
    function handleAnswer(selectedButton) {
        if (!gameActive) return;

        const selectedAnswer = selectedButton.textContent.trim();
        const currentRiddle = riddleBank.find(r => 
            r.question === riddleText.textContent
        );

        if (currentRiddle.answers.includes(selectedAnswer)) {
            // Correct answer
            score += CORRECT_POINTS;
            correctAnswers++;
            
            feedbackElement.textContent = 'Correct! +' + CORRECT_POINTS + ' points';
            feedbackElement.style.color = 'green';
            
            scoreDisplay.textContent = score;
            correctCountDisplay.textContent = correctAnswers;

            // Load next challenge
            setTimeout(loadNextChallenge, 500);
        } else {
            // Incorrect answer
            feedbackElement.textContent = 'Incorrect. Try again!';
            feedbackElement.style.color = 'red';
        }
    }

    // End Game
    function endGame() {
        clearInterval(timerInterval);
        gameActive = false;

        // Display final score
        riddleText.textContent = 'Time\'s Up!';
        feedbackElement.innerHTML = `
            Final Score: ${score}<br>
            Correct Answers: ${correctAnswers}/${TOTAL_CHALLENGES}
        `;
        feedbackElement.style.color = 'white';

        // Disable answer buttons
        answerButtons.forEach(btn => {
            btn.disabled = true;
        });
    }

    // Event Listeners
    answerButtons.forEach(button => {
        button.addEventListener('click', () => handleAnswer(button));
    });

    // Hint Link (to be implemented)
    hintLink.addEventListener('click', () => {
        const currentRiddle = riddleBank.find(r => 
            r.question === riddleText.textContent
        );
        
        // Toggle hint visibility or show hint
        feedbackElement.textContent = currentRiddle.hint;
        feedbackElement.style.color = 'orange';
    });

    // Custom Answer Link (placeholder)
    customAnswerLink.addEventListener('click', () => {
        alert('Custom answer not available in Time Attack mode!');
    });

    // Start the game when the page loads
    initializeGame();
});