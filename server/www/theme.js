// Immediately apply theme before page renders
(function () {
    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
        return null;
    }

    // Check for existing preference
    let prefersDark = getCookie('prefers-dark-mode');

    // If no cookie exists, detect system preference and set cookie
    if (prefersDark === null) {
        prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        document.cookie = `prefers-dark-mode=${prefersDark ? 'true' : 'false'}; path=/; max-age=86400`; // 24 hours
    } else {
        prefersDark = prefersDark === 'true';
    }

    // Apply theme class immediately to prevent flash
    if (prefersDark) {
        document.documentElement.classList.add('dark-theme');
    }

    // Make page visible again
    document.addEventListener('DOMContentLoaded', function () {
        document.documentElement.style.visibility = 'visible';
    });

    // Listen for system preference changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
        const newPrefersDark = e.matches;
        document.cookie = `prefers-dark-mode=${newPrefersDark ? 'true' : 'false'}; path=/; max-age=86400`;
        if (newPrefersDark) {
            document.documentElement.classList.add('dark-theme');
        } else {
            document.documentElement.classList.remove('dark-theme');
        }
    });
})();