const THEME_STORAGE_KEY = "theme";
const DEFAULT_THEME = "blank";

export function applyTheme(theme: string) {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem(THEME_STORAGE_KEY, theme);
}

export function getStoredTheme(): string | null {
    return localStorage.getItem(THEME_STORAGE_KEY);
}

export function initTheme(defaultTheme: string = DEFAULT_THEME) {
    const stored = getStoredTheme();
    const theme = stored && stored.trim().length > 0 ? stored : defaultTheme;

    if (theme) {
        document.documentElement.setAttribute("data-theme", theme);
    }
}