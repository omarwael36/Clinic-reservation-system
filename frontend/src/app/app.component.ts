import { Component, OnInit } from '@angular/core';
import { ApiService } from './api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'frontend';

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    // Example: Fetch data from Golang backend
    this.apiService.doctorSignIn({ /* your data here */ }).subscribe((data: any) => {
      console.log(data);
    });
  
    // Example: Send data to Golang backend
    const postData = { /* your data here */ };
    this.apiService.doctorSignUp(postData).subscribe((response: any) => {
      console.log(response);
    });
  }
}
export { ApiService };

