// Import necessary modules from Angular
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';


// Define the component
@Component({
  selector: 'app-doctor-signup',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class DoctorSignupComponent implements OnInit {
  // Define the form group
  signupForm: FormGroup = this.formBuilder.group({
    // Define your form controls and validators here
    // For example:
    DoctorName: ['', Validators.required],
    DoctorEmail: ['', [Validators.required, Validators.email]],
    DoctorPassword: ['', [Validators.required, Validators.minLength(6)]],
  });

  // Inject FormBuilder and HttpClient in the constructor
  constructor(private formBuilder: FormBuilder, private http: HttpClient) { }

  ngOnInit() {
    // Initialize the form with validators
    this.signupForm = this.formBuilder.group({
      DoctorName: ['', Validators.required],
      DoctorEmail: ['', [Validators.required, Validators.email]],
      DoctorPassword: ['', Validators.required]
    });
  }

  // Function to handle form submission
  onSubmit() {
    // Check if the form is valid
    if (this.signupForm.invalid) {
      return;
    }

    // Get form values
    const formData = this.signupForm.value;

    // Make HTTP POST request to the backend API
    this.http.post<any>('http://localhost:8080/DoctorSignUp', formData)
      .subscribe(response => {
        // Handle the response from the server
        console.log(response);
        // You can add further logic based on the response if needed
      });
  }
}
