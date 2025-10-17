// Main JavaScript file for CAPM Mock Exam System

// Utility functions
const VALID_ALERT_TYPES = new Set(['primary', 'secondary', 'success', 'danger', 'warning', 'info', 'light', 'dark']);

function getGlobalAlertContainer() {
    let container = document.getElementById('globalAlertContainer');

    if (!container) {
        container = document.createElement('div');
        container.id = 'globalAlertContainer';
        container.className = 'global-alert-wrapper';
        container.setAttribute('aria-live', 'polite');
        container.setAttribute('aria-atomic', 'true');
        document.body.appendChild(container);
    }

    return container;
}

function showAlert(message, type = 'info', options = {}) {
    const { autoHide = true, timeout = 6000 } = options;
    const resolvedType = VALID_ALERT_TYPES.has(type) ? type : 'info';
    const container = getGlobalAlertContainer();

    const alertDiv = document.createElement('div');
    alertDiv.className = `alert alert-${resolvedType} alert-dismissible fade show`;
    alertDiv.setAttribute('role', 'alert');
    alertDiv.innerHTML = `
        <div>${message}</div>
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    `;

    container.appendChild(alertDiv);

    if (autoHide) {
        const hideDelay = Number.isFinite(timeout) ? timeout : 6000;
        setTimeout(() => {
            if (typeof bootstrap !== 'undefined' && bootstrap.Alert) {
                try {
                    bootstrap.Alert.getOrCreateInstance(alertDiv).close();
                    return;
                } catch (error) {
                    console.warn('Failed to auto-dismiss alert via Bootstrap:', error);
                }
            }
            alertDiv.remove();
        }, hideDelay);
    }

    return alertDiv;
}

// Form validation helpers
function validateEmail(email) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}

function validateForm(formData) {
    const errors = [];

    if (!formData.name || formData.name.trim().length < 2) {
        errors.push('Name must be at least 2 characters long');
    }

    if (!formData.email || !validateEmail(formData.email)) {
        errors.push('Please enter a valid email address');
    }

    return errors;
}

// API helper functions
async function apiRequest(url, options = {}) {
    const defaultOptions = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const config = { ...defaultOptions, ...options };

    try {
        const response = await fetch(url, config);

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`HTTP ${response.status}: ${errorText}`);
        }

        const contentType = response.headers.get('content-type');
        if (contentType && contentType.includes('application/json')) {
            return await response.json();
        }

        return response;
    } catch (error) {
        console.error('API Request failed:', error);
        throw error;
    }
}

// Local storage helpers for saving exam progress
function saveExamProgress(attemptId, answers) {
    try {
        localStorage.setItem(`exam_${attemptId}`, JSON.stringify({
            answers,
            timestamp: Date.now()
        }));
    } catch (error) {
        console.warn('Failed to save exam progress:', error);
    }
}

function loadExamProgress(attemptId) {
    try {
        const saved = localStorage.getItem(`exam_${attemptId}`);
        if (saved) {
            const data = JSON.parse(saved);
            // Return progress if it's less than 24 hours old
            if (Date.now() - data.timestamp < 24 * 60 * 60 * 1000) {
                return data.answers;
            }
        }
    } catch (error) {
        console.warn('Failed to load exam progress:', error);
    }
    return {};
}

function clearExamProgress(attemptId) {
    try {
        localStorage.removeItem(`exam_${attemptId}`);
    } catch (error) {
        console.warn('Failed to clear exam progress:', error);
    }
}

// Exam timer functionality
class ExamTimer {
    constructor(duration = 180) { // 3 hours in minutes
        this.duration = duration * 60; // Convert to seconds
        this.remaining = this.duration;
        this.interval = null;
        this.callbacks = [];
    }

    start() {
        if (this.interval) return;

        this.interval = setInterval(() => {
            this.remaining--;
            this.notifyCallbacks();

            if (this.remaining <= 0) {
                this.stop();
                this.onTimeUp();
            }
        }, 1000);
    }

    stop() {
        if (this.interval) {
            clearInterval(this.interval);
            this.interval = null;
        }
    }

    addCallback(callback) {
        this.callbacks.push(callback);
    }

    notifyCallbacks() {
        this.callbacks.forEach(callback => callback(this.remaining));
    }

    onTimeUp() {
        // Override this method to handle time up
        showAlert('Time is up! The exam will be submitted automatically.', 'warning');
    }

    formatTime(seconds) {
        const hours = Math.floor(seconds / 3600);
        const minutes = Math.floor((seconds % 3600) / 60);
        const secs = seconds % 60;

        return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
    }
}

// Question shuffling utility
function shuffleArray(array) {
    const shuffled = [...array];
    for (let i = shuffled.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
    }
    return shuffled;
}

// Keyboard navigation for exam
function setupKeyboardNavigation() {
    document.addEventListener('keydown', (event) => {
        // Only handle keyboard navigation on exam page
        if (!window.location.pathname.includes('/exam/')) return;

        switch (event.key) {
            case 'ArrowLeft':
                if (event.ctrlKey && typeof previousQuestion === 'function') {
                    event.preventDefault();
                    previousQuestion();
                }
                break;
            case 'ArrowRight':
                if (event.ctrlKey && typeof nextQuestion === 'function') {
                    event.preventDefault();
                    nextQuestion();
                }
                break;
            case '1':
            case '2':
            case '3':
            case '4':
                if (event.ctrlKey) {
                    event.preventDefault();
                    const choiceIndex = parseInt(event.key) - 1;
                    selectChoiceByIndex(choiceIndex);
                }
                break;
        }
    });
}

function selectChoiceByIndex(index) {
    const container = document.querySelector('#choicesContainer');
    if (!container) {
        return;
    }

    const inputs = container.querySelectorAll('input.form-check-input');
    const input = inputs[index];
    if (!input) {
        return;
    }

    if (input.type === 'checkbox') {
        input.checked = !input.checked;
        input.dispatchEvent(new Event('change', { bubbles: true }));
    } else {
        input.click();
    }
}

// Analytics and tracking
class ExamAnalytics {
    constructor(attemptId) {
        this.attemptId = attemptId;
        this.events = [];
        this.startTime = Date.now();
    }

    trackEvent(event, data = {}) {
        this.events.push({
            event,
            data,
            timestamp: Date.now(),
            timeFromStart: Date.now() - this.startTime
        });
    }

    trackQuestionView(questionId, questionIndex) {
        this.trackEvent('question_view', { questionId, questionIndex });
    }

    trackAnswerChange(questionId, oldChoiceId, newChoiceId) {
        this.trackEvent('answer_change', { questionId, oldChoiceId, newChoiceId });
    }

    trackNavigation(from, to, method) {
        this.trackEvent('navigation', { from, to, method });
    }

    getAnalytics() {
        return {
            attemptId: this.attemptId,
            totalTime: Date.now() - this.startTime,
            events: this.events
        };
    }
}

// Initialize common functionality when DOM is loaded
document.addEventListener('DOMContentLoaded', function() {
    // Setup keyboard navigation
    setupKeyboardNavigation();

    // Add tooltips to navigation buttons
    const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    const tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl);
    });

    // Add confirmation for page reload/close during exam
    if (window.location.pathname.includes('/exam/')) {
        window.addEventListener('beforeunload', function (event) {
            event.preventDefault();
            event.returnValue = 'You have an exam in progress. Are you sure you want to leave?';
        });
    }
});

// Export utilities for use in other scripts
window.ExamUtils = {
    showAlert,
    validateEmail,
    validateForm,
    apiRequest,
    saveExamProgress,
    loadExamProgress,
    clearExamProgress,
    ExamTimer,
    shuffleArray,
    ExamAnalytics
};
