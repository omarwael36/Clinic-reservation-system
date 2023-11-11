import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  signInForm!: FormGroup;

  constructor(private apiService: ApiService, private formBuilder: FormBuilder) {}

  ngOnInit(): void {
    this.signInForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  signIn(): void {
    if (this.signInForm.invalid) {
      // Handle form validation errors
      return;
    }

    const getData = this.signInForm.value;

    this.apiService.doctorSignIn(getData).subscribe(
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
