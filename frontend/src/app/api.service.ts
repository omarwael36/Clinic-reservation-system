import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  doctorSignIn(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/DoctorSignIn`, data);
  }

  doctorSignUp(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/DoctorSignUp`, data);
  }

  setDoctorSchedule(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/DoctorSetSchedule`, data);
  }


  patientSignUp(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/PatientSignUp`, data);
  }

  patientSignIn(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/api/PatientSignIn`, { params: data });
  }

  showAllDoctors(): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowAllDoctors`);
  }

  showDoctorSlots(doctorID: number): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowDoctorSlots`, { params: { DoctorID: doctorID.toString() } });
  }

  reserveSlot(data: any): Observable<any> {
    return this.http.put(`${this.apiUrl}/api/PatientReserveSlot`, data);
  }

  updateAppointment(data: any): Observable<any> {
    return this.http.put(`${this.apiUrl}/api/PatientUpdateAppointment`, data);
  }

  cancelAppointment(slotID: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/api/PatientCancelAppointment`, { params: { slotId: slotID.toString() } });
  }

  showAllReservations(): Observable<any> {
    return this.http.get(`${this.apiUrl}/api/PatientShowAppointments`);
  }
}
