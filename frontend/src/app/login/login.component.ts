import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  signInForm!: FormGroup;
  userType: string = 'Doctor';
  errorMessage: string = '';

  constructor(private apiService: ApiService, private formBuilder: FormBuilder, private router: Router) {}

  ngOnInit(): void {
    this.signInForm = this.formBuilder.group({
      userEmail: ['', [Validators.required, Validators.email]],
      userPassword: ['', Validators.required],
    });
  }

  signIn(): void {
    if (this.signInForm.invalid) {
      return;
    }

    const postData = this.signInForm.value;

    if (this.userType === 'Doctor') {
      this.apiService.doctorSignIn(postData).subscribe(
        (response: any) => {
          console.log('Doctor Login:', response);
          this.router.navigate(['/UserDoctor']);
        },
        (error: any) => {
          console.error('Doctor Login Error:', error);
          this.errorMessage = error.error.error;
        }
      );
    } else if (this.userType === 'Patient') {
      this.apiService.patientSignIn(postData).subscribe(
        (response: any) => {
          console.log('Patient Login:', response);
          this.router.navigate(['/UserPatient']);
        },
        (error: any) => {
          console.error('Patient Login Error:', error);
          this.errorMessage = error.error.error;
        }
      );
    }
  }

  setUserType(userType: string): void {
    this.userType = userType;
  }
}
