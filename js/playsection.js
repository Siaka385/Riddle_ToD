document.addEventListener('DOMContentLoaded', () => {
    let riddleNum = {
        count: 0, // Dynamic value for the count
    };
    
    function loadRiddles() {
        fetch("http://localhost:8089/loadriddle", {
            method: "POST", // Sending data to the server
            headers: {
                "Content-Type": "application/json", // Specify the content type as JSON
            },
            body: JSON.stringify(riddleNum), // Send `riddleNum` as the request payload
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.json(); // Parse the response JSON
            })
            .then(data => {
                reddit(data); // Log the loaded data
            })
            .catch(err => console.error("Error loading riddles:", err)); // Log any errors
    }
    
    // Call the function to load riddles
   loadRiddles();

 

    const riddledisplay=document.getElementById("riddleQustion")
    const Choices=document.querySelectorAll(".choices")
    const hints=document.getElementById("riddleHint")
    
   function reddit(m){
             riddledisplay.innerHTML=m.Riddle
             for (var i = 0;i<4;i++){
                   Choices[i].innerHTML=m.Choices[i];
             }
             hints.innerHTML=m.Hint;
   }          

    
    const hintLink = document.querySelector('.hint-link');
    const hintElement = document.getElementById('hint');
    const answerButtons = document.querySelectorAll('.answer-buttons .btn');
    const feedbackElement = document.getElementById('feedback');


    
    // Custom answer elements
 //   const customAnswerLink = document.querySelector('.custom-answer-link');
    // const customAnswerContainer = document.getElementById('custom-answer-container');
    // const customAnswerInput = document.getElementById('custom-answer');
    // const customAnswerSubmit = document.querySelector('.custom-answer-submit');

    // Add mobile-friendly feedback styles
    const feedbackStyles = `
        <style>
            .feedback-container {
                width: 100%;
                padding: 15px;
                box-sizing: border-box;
                margin-top: 15px;
            }
            
            .feedback-message {
                margin-bottom: 15px;
                font-size: 16px;
                font-weight: bold;
            }
            
            .feedback-explanation {
                color: white;
                font-size: 14px;
                text-align: left;
                line-height: 1.6;
                padding: 15px;
                background: rgba(0, 0, 0, 0.3);
                border-radius: 10px;
                margin-top: 10px;
            }
            
            .feedback-explanation ul {
                margin: 0;
                padding-left: 20px;
                list-style-type: disc;
            }
            
            .feedback-explanation li {
                margin-bottom: 8px;
            }
            
            @media (min-width: 768px) {
                .feedback-explanation {
                    font-size: 16px;
                    padding: 20px;
                }
            }
        </style>
    `;
    
    // Add styles to document
    document.head.insertAdjacentHTML('beforeend', feedbackStyles);


    // Hint toggle
    hintLink.addEventListener('click', () => {
        hintElement.style.display = 
            hintElement.style.display === 'none' ? 'block' : 'none';
    });

    // Custom answer link toggle
    // customAnswerLink.addEventListener('click', () => {
    //     customAnswerContainer.style.display = 
    //         customAnswerContainer.style.display === 'none' ? 'flex' : 'none';
    // });

    // Function to show feedback with explanation
    const showFeedback = (isCorrect, answer) => {
        if (isCorrect) {
            const explanationPoints = riddleExplanations[answer];
            feedbackElement.innerHTML = `
                <div class="feedback-container">
                    <div class="feedback-message" style="color: green;">
                        Correct! Well done!
                    </div>
                    <div class="feedback-explanation">
                        <strong>Why is this correct?</strong>
                        <ul>
                            ${explanationPoints.map(point => `<li>${point}</li>`).join('')}
                        </ul>
                    </div>
                </div>
            `;
        } else {
            feedbackElement.innerHTML = `
                <div class="feedback-container">
                    <div class="feedback-message" style="color: red;">
                        Incorrect. Try again!
                    </div>
                </div>
            `;
        }
    };

    // Answer selection
    answerButtons.forEach(button => {
        button.addEventListener('click', () => {
            console.log("hello")
            const selectedAnswer = button.textContent.trim();
            showFeedback(selectedAnswer === 'An Echo', selectedAnswer);
            // Scroll to feedback smoothly
            feedbackElement.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
        });
    });


    // Custom answer submission
    // customAnswerSubmit.addEventListener('click', () => {
    //     const customAnswer = customAnswerInput.value.trim().toLowerCase();
    //     const isCorrect = customAnswer === 'echo';
    //     showFeedback(isCorrect, 'An Echo');
    //     // Scroll to feedback smoothly
    //     feedbackElement.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
        
    //     // Clear input
    //     customAnswerInput.value = '';
    // });
});