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
    this.apiService.doctorSignIn({ }).subscribe((data: any) => {
      console.log(data);
    });
  
    const postData = {  };
    this.apiService.doctorSignUp(postData).subscribe((response: any) => {
      console.log(response);
    });
  }
}
export { ApiService };

