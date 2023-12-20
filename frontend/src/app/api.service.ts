import { HttpClient, HttpErrorResponse, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, throwError } from 'rxjs';
import { environment } from '../environments/environment';
@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl: string = ''; 

  constructor(private http: HttpClient) {
    this.fetchApiUrl();
  }

  private fetchApiUrl(): void {
    this.http.get<any>('your/environment/endpoint')
      .subscribe(
        (config: any) => {
          this.apiUrl = config.apiUrl; // Update the apiUrl
        },
        (error) => {
          console.error('Error fetching API_URL', error);
        }
      );
  }
  

  doctorSignIn(data: any): Observable<any> {
    console.log("hello worlddd");
    console.log(this.apiUrl);
    return this.http.post(`${this.apiUrl}/api/DoctorSignIn`, data);
  }

  doctorSignUp(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/DoctorSignUp`, data);
  }

  


  patientSignUp(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/PatientSignUp`, data);
  }

  patientSignIn(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/PatientSignIn`,  data);
  }

  showAllDoctors(): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowAllDoctors`);
  }

  showDoctorSlots(doctorID: number): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowDoctorSlots`, { params: { DoctorID: doctorID.toString() } });
  }

  reserveSlot(patientID: string, data: any): Observable<any> {
  
    return this.http.put(`${this.apiUrl}/api/PatientReserveSlot/${patientID}`, data);
}

  

  updateAppointment(patientID: string, data: any) {
    return this.http.put(`${this.apiUrl}/api/PatientUpdateAppointment/${patientID}`, data)
      .pipe(
        catchError(this.handleError)
      );
  }
  setDoctorSchedule(doctorID: string, slotDateTime: string): Observable<any> {
    const url = `${this.apiUrl}/api/DoctorSetSchedule/${doctorID}`;
  
    const requestBody = { slotDateTime };
  
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
  
    return this.http.post(url, requestBody, { headers });
  }
  

  cancelAppointment(patientID: string, slotID: number): Observable<any> {
    const url = `${this.apiUrl}/api/PatientCancelAppointment/${patientID}`;
    const params = new HttpParams().set('slotId', slotID.toString());
  
    return this.http.delete(url, { params });
  }
  
  

  showPatientAppointments(patientID: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowAppointments/${patientID}`);
  }

  private handleError(error: HttpErrorResponse) {
    let errorMessage = 'Unknown error occurred!';
    if (error.error instanceof ErrorEvent) {
      // Client-side error
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side error
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    console.error(errorMessage);
    return throwError(errorMessage);
  }
  
}
