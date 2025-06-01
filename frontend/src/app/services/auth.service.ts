export class AuthService {
    private password: string;

    constructor() {
        this.password = this.getStoredPassword();
    }

    private getStoredPassword(): string {
        return process.env.APP_PASSWORD || '';
    }

    public validatePassword(inputPassword: string): boolean {
        return inputPassword === this.password;
    }

    public storePassword(password: string): void {
        this.password = password;
        // Logic to store password securely can be added here
    }

    public isAuthenticated(): boolean {
        return this.password !== '';
    }

    public redirectToLogin(): void {
        // Logic to redirect to login page
    }

    public autoRedirect(): void {
        setTimeout(() => {
            this.redirectToLogin();
        }, 1800000); // 30 minutes
    }
}