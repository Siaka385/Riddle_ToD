document.addEventListener('DOMContentLoaded', () => {
    const slides = document.querySelectorAll('.slide-container');
    const dots = document.querySelectorAll('.dot');
    const nextButtons = [
        document.getElementById('next-slide1'),
        document.getElementById('next-slide2'),
        document.getElementById('next-slide3')
    ];
    const skipButton = document.getElementById('skip-button');

    function switchSlide(currentSlide, nextSlide) {
        slides[currentSlide].classList.remove('active');
        slides[nextSlide].classList.add('active');

        dots.forEach(dot => dot.classList.remove('active'));
        dots[nextSlide].classList.add('active');
    }

    // Skip button functionality
    skipButton.addEventListener('click', () => {
        slides.forEach(slide => slide.classList.remove('active'));
        slides[3].classList.add('active');

        dots.forEach(dot => dot.classList.remove('active'));
        dots[3].classList.add('active');
    });

    // Rest of the JavaScript remains the same as in the previous version
    // ... (previous event listeners)

    nextButtons.forEach((button, index) => {
        button.addEventListener('click', () => {
            switchSlide(index, index + 1);
        });
    });

    dots.forEach(dot => {
        dot.addEventListener('click', function() {
            const slideIndex = parseInt(this.dataset.slide) - 1;
            slides.forEach(slide => slide.classList.remove('active'));
            slides[slideIndex].classList.add('active');

            dots.forEach(d => d.classList.remove('active'));
            this.classList.add('active');
        });
    });

    // Login/Signup/Guest Play Handlers (Placeholder)
    document.getElementById('login').addEventListener('click', () => {
        alert('Login functionality to be implemented');
    });

    document.getElementById('signup').addEventListener('click', () => {
        alert('Signup functionality to be implemented');
    });

    document.getElementById('guest-play').addEventListener('click', () => {
        alert('Guest play functionality to be implemented');
    });
});