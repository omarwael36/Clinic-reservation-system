import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm = this.formBuilder.group({
    username: '',
    password: '',
  });

  constructor(
    private formBuilder: FormBuilder,
  ) {}

  onSubmit() {
    console.log('Login clicked', this.loginForm.value);
  }
}