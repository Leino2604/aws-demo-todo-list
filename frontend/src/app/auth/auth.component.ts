import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit {
  authForm: FormGroup;
  isLogin = true;
  error = '';

  constructor(
    private fb: FormBuilder,
    private authService: AuthService
  ) {
    this.authForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.authForm.valid) {
      const { email, password } = this.authForm.value;

      if (this.isLogin) {
        this.authService.login(email, password).subscribe(
          (response: any) => {
            localStorage.setItem('token', response.token);
          },
          (error: any) => {
            this.error = error.error.message || 'Login failed';
          }
        );
      } else {
        this.authService.register(email, password).subscribe(
          (response: any) => {
            localStorage.setItem('token', response.token);
          },
          (error: any) => {
            this.error = error.error.message || 'Registration failed';
          }
        );
      }
    }
  }

  toggleAuthMode(): void {
    this.isLogin = !this.isLogin;
    this.error = '';
  }
}