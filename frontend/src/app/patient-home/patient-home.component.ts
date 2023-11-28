import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-patient-home',
  templateUrl: './patient-home.component.html',
  styleUrls: ['./patient-home.component.css']
})
export class PatientHomeComponent implements OnInit {
  appointments: any[] = [];
  doctors: any[] = [];
  doctorSlots: any[] = [];
  newSlotDateTime: string = '';
  newDoctorId: number = 0;
  doctorIdToShowSlots: number = 0;
  appointmentForm: FormGroup;

  constructor(private apiService: ApiService, private formBuilder: FormBuilder) {
    this.appointmentForm = this.formBuilder.group({
      newSlotDateTime: ['', Validators.required],
      newDoctorId: [0, Validators.required],
    });
  }

  ngOnInit(): void {
    this.loadAppointments();
    this.loadDoctors();
  }

  loadAppointments(): void {
    this.apiService.showAllReservations().subscribe(
      (response: any) => {
        this.appointments = response.Data;
      },
      (error: any) => {
        console.error(error);
        
      }
    );
  }

  loadDoctors(): void {
    this.apiService.showAllDoctors().subscribe(
      (response: any) => {
        this.doctors = response.Data;
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  showDoctorSlots(doctorId: number): void {
    this.apiService.showDoctorSlots(doctorId).subscribe(
      (response: any) => {
        this.doctorSlots = response.Data;
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  reserveSlot(slotId: number): void {
    this.apiService.reserveSlot({ SlotID: slotId }).subscribe(
      (response: any) => {
        console.log(response);
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  cancelAppointment(slotId: number): void {
    this.apiService.cancelAppointment(slotId).subscribe(
      (response: any) => {
        console.log(response);
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  updateAppointment(appointmentId: number, newSlotId: number, newDoctorId: number, newPatientId: number): void {
    this.apiService.updateAppointment({
      AppointmentID: appointmentId,
      NewSlotID: newSlotId,
      NewDoctorID: newDoctorId,
      NewPatientID: newPatientId
    }).subscribe(
      (response: any) => {
        console.log(response);
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  showAllDoctors(): void {
  this.apiService.showAllDoctors().subscribe(
    (response: any) => {
      console.log(response);
      this.doctors = response.data;
      console.log('After API call:', this.doctors);
    },
    (error: any) => {
      console.error(error);
    }
  );
}

  
  
  
showSlotsForDoctor(): void {
  this.apiService.showDoctorSlots(this.doctorIdToShowSlots).subscribe(
    (response: any) => {
      console.log(response);
      this.doctorSlots = response.data;
      console.log('Doctor Slots:', this.doctorSlots);
    },
    (error: any) => {
      console.error(error);
    }
  );
}


  createAppointment(): void {
    if (this.appointmentForm.valid) {
      const newSlotDateTime = this.appointmentForm.value.newSlotDateTime;
      const newDoctorId = this.appointmentForm.value.newDoctorId;

      
    } else {
    }
  }
}