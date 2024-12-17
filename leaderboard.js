document.addEventListener('DOMContentLoaded', function() {
    // Simulated leaderboard data (would typically come from a backend)
    const leaderboardData = [
        {
            nickname: "DreadMaster",
            points: 1250,
            totalRiddlesSolved: 75,
            accuracy: 92.5,
            category: "all"
        },
        {
            nickname: "RiddleShadow",
            points: 1100,
            totalRiddlesSolved: 68,
            accuracy: 88.3,
            category: "mathematics"
        },
        {
            nickname: "EnigmaPuzzler",
            points: 1050,
            totalRiddlesSolved: 62,
            accuracy: 90.1,
            category: "general"
        },
        {
            nickname: "LogicNinja",
            points: 980,
            totalRiddlesSolved: 55,
            accuracy: 95.2,
            category: "logic"
        },
        {
            nickname: "WordWizard",
            points: 920,
            totalRiddlesSolved: 50,
            accuracy: 87.6,
            category: "word"
        }
        // More entries can be added
    ];

    const leaderboardBody = document.getElementById('leaderboard-body');
    const timeframeSelect = document.getElementById('timeframe');
    const categorySelect = document.getElementById('category');
    const prevPageBtn = document.getElementById('prev-page');
    const nextPageBtn = document.getElementById('next-page');
    const currentPageSpan = document.getElementById('current-page');
    const backToMainBtn = document.getElementById('back-to-main');
    
    // New elements for user search and position
    const userSearchInput = document.getElementById('user-search');
    const findUserBtn = document.getElementById('find-user');
    const showMyPositionBtn = document.getElementById('show-my-position');
    const userPositionResult = document.getElementById('user-position-result');

    let currentPage = 1;
    const itemsPerPage = 10;

    function renderLeaderboard(data) {
        leaderboardBody.innerHTML = ''; // Clear previous entries

        data.forEach((entry, index) => {
            const row = document.createElement('tr');
            row.classList.add(`rank-${index + 1}`);
            
            row.innerHTML = `
                <td>${index + 1}</td>
                <td>${entry.nickname}</td>
                <td>${entry.points}</td>
                <td>${entry.totalRiddlesSolved}</td>
                <td>${entry.accuracy.toFixed(1)}%</td>
            `;

            leaderboardBody.appendChild(row);
        });
    }

    function filterLeaderboard() {
        const selectedTimeframe = timeframeSelect.value;
        const selectedCategory = categorySelect.value;

        let filteredData = leaderboardData.filter(entry => 
            (selectedCategory === 'all' || entry.category === selectedCategory)
        );

        // Sort by points in descending order
        filteredData.sort((a, b) => b.points - a.points);

        renderLeaderboard(filteredData);
    }

    function findUserPosition(nickname) {
        // Find user's position in the current filtered leaderboard
        const sortedData = leaderboardData
            .sort((a, b) => b.points - a.points)
            .filter(entry => entry.nickname.toLowerCase().includes(nickname.toLowerCase()));

        if (sortedData.length > 0) {
            const user = sortedData[0];
            const userRank = leaderboardData.indexOf(user) + 1;
            
            userPositionResult.innerHTML = `
                <div class="user-position-card">
                    <h3>User Position</h3>
                    <p><strong>Nickname:</strong> ${user.nickname}</p>
                    <p><strong>Rank:</strong> ${userRank}</p>
                    <p><strong>Points:</strong> ${user.points}</p>
                    <p><strong>Total Riddles Solved:</strong> ${user.totalRiddlesSolved}</p>
                    <p><strong>Accuracy:</strong> ${user.accuracy.toFixed(1)}%</p>
                </div>
            `;
        } else {
            userPositionResult.innerHTML = '<p>No user found with that nickname.</p>';
        }
    }

    function showCurrentUserPosition() {
        // In a real app, this would use the currently logged-in user
        const currentUser = "DreadMaster";
        findUserPosition(currentUser);
    }

    // Initial render
    renderLeaderboard(leaderboardData);

    // Filter event listeners
    timeframeSelect.addEventListener('change', filterLeaderboard);
    categorySelect.addEventListener('change', filterLeaderboard);

    // User search event listeners
    findUserBtn.addEventListener('click', function() {
        const nickname = userSearchInput.value.trim();
        if (nickname) {
            findUserPosition(nickname);
        }
    });

    showMyPositionBtn.addEventListener('click', showCurrentUserPosition);

    // Back to main menu
    backToMainBtn.addEventListener('click', function() {
        window.location.href = 'index2.html'; // Adjust path as needed
    });

    // Pagination (simplified for this example)
    prevPageBtn.addEventListener('click', function() {
        if (currentPage > 1) {
            currentPage--;
            currentPageSpan.textContent = `Page ${currentPage}`;
        }
    });

    nextPageBtn.addEventListener('click', function() {
        // In a real scenario, check against total number of entries
        currentPage++;
        currentPageSpan.textContent = `Page ${currentPage}`;
    });
});