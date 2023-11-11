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
    // Initialize the form group in the constructor
    this.appointmentForm = this.formBuilder.group({
      newSlotDateTime: ['', Validators.required],
      newDoctorId: [0, Validators.required],
      // Add other necessary form controls here
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
        // Reload appointments or update the UI as needed
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
        // Handle error or show an error message to the user
      }
    );
  }

  cancelAppointment(slotId: number): void {
    this.apiService.cancelAppointment(slotId).subscribe(
      (response: any) => {
        console.log(response);
        // Reload appointments or update the UI as needed
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
        // Handle error or show an error message to the user
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
        // Reload appointments or update the UI as needed
        this.loadAppointments();
      },
      (error: any) => {
        console.error(error);
        // Handle error or show an error message to the user
      }
    );
  }

  showAllDoctors(): void {
  this.apiService.showAllDoctors().subscribe(
    (response: any) => {
      console.log(response); // Log the API response
      this.doctors = response.data; // Update to response.data
      console.log('After API call:', this.doctors); // Confirm data in doctors array
    },
    (error: any) => {
      console.error(error);
      // Handle error or show an error message to the user
    }
  );
}

  
  
  
showSlotsForDoctor(): void {
  this.apiService.showDoctorSlots(this.doctorIdToShowSlots).subscribe(
    (response: any) => {
      console.log(response);
      this.doctorSlots = response.data; // Assuming the data is in a property named 'data'
      console.log('Doctor Slots:', this.doctorSlots); // Confirm data in doctorSlots array
    },
    (error: any) => {
      console.error(error);
      // Handle error or show an error message to the user
    }
  );
}


  createAppointment(): void {
    if (this.appointmentForm.valid) {
      const newSlotDateTime = this.appointmentForm.value.newSlotDateTime;
      const newDoctorId = this.appointmentForm.value.newDoctorId;

      // Perform the create appointment logic here
      // Use newSlotDateTime and newDoctorId

      // After creating the appointment, you may want to reload appointments or update the UI as needed
      // this.loadAppointments();
    } else {
      // Handle form validation errors if needed
    }
  }
}
