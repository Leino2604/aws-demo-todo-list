import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from './services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'To-Do List App';

  constructor(private authService: AuthService, private router: Router) {
    this.checkAuthentication();
  }

  checkAuthentication() {
    if (!this.authService.isAuthenticated()) {
      this.router.navigate(['/login']);
    }
    setInterval(() => {
      if (!this.authService.isAuthenticated()) {
        this.router.navigate(['/login']);
      }
    }, 1800000); // 30 minutes
  }
}