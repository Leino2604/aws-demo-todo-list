import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from './auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css']
})
export class AuthComponent {
  authForm: FormGroup;
  isSignIn: boolean = true;

  constructor(private fb: FormBuilder, private authService: AuthService) {
    this.authForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  toggleMode() {
    this.isSignIn = !this.isSignIn;
    this.authForm.reset();
  }

  onSubmit() {
    if (this.authForm.valid) {
      if (this.isSignIn) {
        this.authService.signIn(this.authForm.value).subscribe(response => {
          // Handle successful sign-in
        }, error => {
          // Handle sign-in error
        });
      } else {
        this.authService.signUp(this.authForm.value).subscribe(response => {
          // Handle successful sign-up
        }, error => {
          // Handle sign-up error
        });
      }
    }
  }
}