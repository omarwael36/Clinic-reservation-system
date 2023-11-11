import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ApiService } from "../app.component";
import { Component } from "@angular/core";

@Component({
  selector: 'app-sign-up', // This selector is important
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})


export class SignUpComponent {
  signUpForm: FormGroup;

  constructor(private apiService: ApiService, private fb: FormBuilder) {
    this.signUpForm = this.fb.group({
      doctorName: ['', Validators.required],
      doctorEmail: ['', [Validators.required, Validators.email]],
      doctorPassword: ['', Validators.required],
    });
  }

  signUp(): void {
    if (this.signUpForm.invalid) {
      // Handle form validation errors
      return;
    }
  
    const postData = this.signUpForm.value;
  
    this.apiService.doctorSignUp(postData).subscribe(
      (response: any) => {
        console.log(response);
        // Handle success or show a message to the user
      },
      (error: any) => {
        console.error(error);
        // Handle error or show an error message to the user
      }
    );
  }
}
