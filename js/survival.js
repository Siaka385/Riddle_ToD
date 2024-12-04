window.onload = () => {
    const timer = document.querySelector('.timer');
    const livesCount = document.querySelector('.lives-count');
    const answerButtons = document.querySelectorAll('.btn');
    const gameOverSection = document.querySelector('.game-over-section');
    const feedbackElement = document.getElementById('feedback');
    const hintLink = document.querySelector('.hint-link');
    const customAnswerLink = document.querySelector('.custom-answer-link');

    let timeLeft = 30;
    let lives = 3;
    let timerInterval;

    // Current riddle details
    const currentRiddle = {
        question: "What travels around the world but stays in one corner?",
        correctAnswer: "A Stamp",
        hint: "Think about something small that helps send messages globally.",
        possibleAnswers: ["A Stamp", "A Letter", "A Postcard", "An Envelope"]
    };

    function startTimer() {
        timerInterval = setInterval(() => {
            timeLeft--;
            timer.textContent = timeLeft;

            if (timeLeft <= 0) {
                endGame();
            }
        }, 1000);
    }

    function endGame() {
        clearInterval(timerInterval);
        document.querySelector('.challenge-section').style.display = 'none';
        document.querySelector('.riddle').style.display = 'none';
        document.querySelector('.answer-buttons').style.display = 'none';
        gameOverSection.style.display = 'block';
    }

    function handleAnswer(isCorrect) {
        if (!isCorrect) {
            lives--;
            livesCount.textContent = lives;
            feedbackElement.textContent = "Incorrect! You lost a life.";
            feedbackElement.style.color = "red";

            if (lives <= 0) {
                endGame();
            }
        } else {
            feedbackElement.textContent = "Correct! Well done!";
            feedbackElement.style.color = "green";
        }
    }

    // Hint functionality
    hintLink.addEventListener('click', (e) => {
        e.preventDefault();
        document.getElementById("hint").style.display="block";
       document.getElementById("hint").textContent = `Hint: ${currentRiddle.hint}`;

    });

    // Custom answer functionality
    customAnswerLink.addEventListener('click', (e) => {
        e.preventDefault();
      let answerContainer=document.getElementById("custom-answer-container")
      answerContainer.classList.remove(".displaycustom")

      answerContainer.classList.add("custom-answer-container")

        
        if (customAnswer) {
            const isCorrect = customAnswer.trim().toLowerCase() === currentRiddle.correctAnswer.toLowerCase();
            
            if (isCorrect) {
                feedbackElement.textContent = "Correct custom answer!";
                feedbackElement.style.color = "green";
            } else {
                feedbackElement.textContent = "Incorrect custom answer. You lost a life.";
                feedbackElement.style.color = "red";
                lives--;
                livesCount.textContent = lives;

                if (lives <= 0) {
                    endGame();
                }
            }
        }
    });

    // Answer button functionality
    answerButtons.forEach(button => {
        button.addEventListener('click', () => {
            const isCorrect = button.textContent === currentRiddle.correctAnswer;
            handleAnswer(isCorrect);
        });
    });

    startTimer();
};