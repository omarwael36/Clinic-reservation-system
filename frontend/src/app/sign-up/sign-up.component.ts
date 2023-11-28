import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ApiService } from "../app.component";
import { Component } from "@angular/core";
import { Router } from "@angular/router";

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent {
  signUpForm: FormGroup;
  errorMessage: string = '';

  constructor(private apiService: ApiService, private fb: FormBuilder, private router: Router) {
    this.signUpForm = this.fb.group({
      userName: ['', Validators.required],
      userEmail: ['', [Validators.required, Validators.email]],
      userPassword: ['', Validators.required],
      userType: ['Doctor', Validators.required]
    });
  }

  signUp(): void {
    if (this.signUpForm.invalid) {
      return;
    }
  
    const postData = this.signUpForm.value;
    const userType = postData.userType;

    if (userType === 'Doctor') {
      this.apiService.doctorSignUp(postData).subscribe(
        (response: any) => {
          console.log(response);
          // Redirect to login page on success
          this.router.navigate(['/']);
        },
        (error: any) => {
          console.error(error);
          this.errorMessage = error.error.error; // Assuming the error message structure
        }
      );
    } else if (userType === 'Patient') {
      this.apiService.patientSignUp(postData).subscribe(
        (response: any) => {
          console.log(response);
          // Redirect to login page on success
          this.router.navigate(['/']);
        },
        (error: any) => {
          console.error(error);
          this.errorMessage = error.error.error; // Assuming the error message structure
        }
      );
    }
  }
}
