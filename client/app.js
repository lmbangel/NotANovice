class QuizApp {
    constructor() {
        this.API_BASE_URL = 'http://localhost:8000/v1';
        this.currentQuestion = null;
        this.selectedAnswer = null;

        // Page navigation
        this.landingPage = document.getElementById('landingPage');
        this.quizPage = document.getElementById('quizPage');

        // Navigation elements
        this.startQuizBtn = document.getElementById('startQuizBtn');
        this.backBtn = document.getElementById('backBtn');

        // Quiz DOM elements
        this.loadingState = document.getElementById('loadingState');
        this.quizContainer = document.getElementById('quizContainer');
        this.errorState = document.getElementById('errorState');
        this.questionText = document.getElementById('questionText');
        this.answerOptions = document.getElementById('answerOptions');
        this.submitBtn = document.getElementById('submitBtn');
        this.resultsSection = document.getElementById('resultsSection');
        this.retryBtn = document.getElementById('retryBtn');

        this.init();
    }

    init() {
        this.setupEventListeners();
        this.showLandingPage();
    }

    setupEventListeners() {
        // Navigation event listeners
        this.startQuizBtn.addEventListener('click', () => this.showQuizPage());
        this.backBtn.addEventListener('click', () => this.showLandingPage());

        // Quiz event listeners
        this.submitBtn.addEventListener('click', () => this.handleSubmit());
        this.retryBtn.addEventListener('click', () => this.loadQuestion());
    }

    // Navigation methods
    showLandingPage() {
        this.landingPage.classList.remove('hidden');
        this.quizPage.classList.add('hidden');
    }

    showQuizPage() {
        this.landingPage.classList.add('hidden');
        this.quizPage.classList.remove('hidden');
        this.loadQuestion();
    }

    async loadQuestion() {
        this.showLoading();

        try {
            const response = await fetch(`${this.API_BASE_URL}/questions`);

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const questions = await response.json();

            if (!questions || questions.length === 0) {
                throw new Error('No questions available');
            }

            // For testing, we'll show only the first question
            this.currentQuestion = questions[0];
            this.displayQuestion();

        } catch (error) {
            console.error('Error loading question:', error);
            this.showError();
        }
    }

    displayQuestion() {
        this.hideAllStates();
        this.quizContainer.classList.remove('hidden');

        // Display the question text
        this.questionText.textContent = this.currentQuestion.question;

        // Clear previous answers
        this.answerOptions.innerHTML = '';
        this.selectedAnswer = null;
        this.submitBtn.disabled = true;

        // Create answer options
        this.createAnswerOptions();
    }

    createAnswerOptions() {
        // Extract real answer options from the API response
        const answers = [
            { id: 'A', text: this.currentQuestion.a_answer, isCorrect: this.currentQuestion.correct_answer === 'A' },
            { id: 'B', text: this.currentQuestion.b_answer, isCorrect: this.currentQuestion.correct_answer === 'B' },
            { id: 'C', text: this.currentQuestion.c_answer, isCorrect: this.currentQuestion.correct_answer === 'C' },
            { id: 'D', text: this.currentQuestion.d_answer, isCorrect: this.currentQuestion.correct_answer === 'D' }
        ];

        answers.forEach((answer, index) => {
            const optionDiv = document.createElement('div');
            optionDiv.className = 'relative';

            optionDiv.innerHTML = `
                <input
                    type="radio"
                    id="answer_${answer.id}"
                    name="quiz_answer"
                    value="${answer.id}"
                    class="sr-only peer"
                    data-correct="${answer.isCorrect}"
                >
                <label
                    for="answer_${answer.id}"
                    class="flex items-center p-4 bg-white border-2 border-notion-200 rounded-xl cursor-pointer hover:border-coral-300 peer-checked:border-coral-500 peer-checked:bg-coral-50 transition-all duration-200"
                >
                    <div class="flex items-center justify-center w-6 h-6 border-2 border-notion-300 rounded-full mr-4 peer-checked:border-coral-500 peer-checked:bg-coral-500">
                        <div class="w-2 h-2 bg-white rounded-full opacity-0 peer-checked:opacity-100"></div>
                    </div>
                    <span class="text-notion-900 font-medium">${answer.text}</span>
                </label>
            `;

            this.answerOptions.appendChild(optionDiv);

            // Add event listener to the radio button
            const radioButton = optionDiv.querySelector('input[type="radio"]');
            radioButton.addEventListener('change', () => this.handleAnswerSelection(answer.id));
        });
    }

    handleAnswerSelection(answerId) {
        this.selectedAnswer = answerId;
        this.submitBtn.disabled = false;

        // Update visual feedback for custom radio buttons
        const labels = this.answerOptions.querySelectorAll('label');
        labels.forEach(label => {
            const radio = label.previousElementSibling || label.querySelector('input[type="radio"]');
            const circle = label.querySelector('div');
            const dot = circle.querySelector('div');

            if (radio && radio.checked) {
                circle.classList.add('border-coral-500', 'bg-coral-500');
                circle.classList.remove('border-notion-300');
                dot.classList.remove('opacity-0');
                label.classList.add('border-coral-500', 'bg-coral-50');
                label.classList.remove('border-notion-200');
            } else {
                circle.classList.remove('border-coral-500', 'bg-coral-500');
                circle.classList.add('border-notion-300');
                dot.classList.add('opacity-0');
                label.classList.remove('border-coral-500', 'bg-coral-50');
                label.classList.add('border-notion-200');
            }
        });
    }

    handleSubmit() {
        if (!this.selectedAnswer) return;

        // Get the selected answer element
        const selectedInput = document.querySelector(`input[value="${this.selectedAnswer}"]`);
        const isCorrect = selectedInput.dataset.correct === 'true';

        this.showResults(isCorrect);
    }

    showResults(isCorrect) {
        this.resultsSection.classList.remove('hidden');
        const resultDiv = this.resultsSection.querySelector('#correctAnswer');
        const icon = resultDiv.querySelector('svg');
        const iconContainer = icon.parentElement;
        const title = resultDiv.querySelector('h3');
        const message = resultDiv.querySelector('p');

        if (isCorrect) {
            iconContainer.classList.add('bg-teal-100');
            icon.classList.add('text-teal-600');
            title.textContent = 'Great job!';
            title.classList.add('text-teal-900');
            message.textContent = 'You got it right!';
        } else {
            iconContainer.classList.add('bg-red-100');
            icon.classList.add('text-red-600');
            icon.innerHTML = '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>';
            title.textContent = 'Not quite right';
            title.classList.add('text-red-900');
            message.textContent = 'Keep studying and try again tomorrow!';
        }

        // Disable submit button after submission
        this.submitBtn.disabled = true;
        this.submitBtn.textContent = 'Submitted';
        this.submitBtn.classList.remove('bg-coral-500', 'hover:bg-coral-600');
        this.submitBtn.classList.add('bg-notion-400');
    }

    showLoading() {
        this.hideAllStates();
        this.loadingState.classList.remove('hidden');
    }

    showError() {
        this.hideAllStates();
        this.errorState.classList.remove('hidden');
    }

    hideAllStates() {
        this.loadingState.classList.add('hidden');
        this.quizContainer.classList.add('hidden');
        this.errorState.classList.add('hidden');
        this.resultsSection.classList.add('hidden');
    }
}

// Initialize the app when the DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new QuizApp();
});