import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';

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
  newSlotId: number | undefined;
slot: any;

  
  constructor(private apiService: ApiService, private formBuilder: FormBuilder, private cdr: ChangeDetectorRef, private activatedRoute: ActivatedRoute) {
    this.appointmentForm = this.formBuilder.group({
      newSlotDateTime: ['', Validators.required],
      newDoctorId: [0, Validators.required],
    });
  }

  ngOnInit(): void {
    
  }

  // Method to retrieve patient ID - Replace this with your logic
  retrievePatientID(): string | null {
    // Implement your logic to retrieve the patient ID from the URL or any other source
    const urlParts = window.location.href.split('/');
    const patientID = urlParts[urlParts.length - 1]; // Retrieve patient ID from the URL
    return patientID ? patientID.toString() : null;
  }

// Function to initiate loading appointments
showReservations(): void {
  const patientIDValue = this.retrievePatientID();
  if (patientIDValue) {
    this.loadAppointments(patientIDValue); // Provide the patientIDValue when calling
  } else {
    console.error('PatientID not found');
    // Handle the case when the patientIDValue is not available
  }
}

// Function to load appointments based on patient ID
loadAppointments(patientID: string): void {
  console.log("Loading appointments for PatientID:", patientID);
  
  this.apiService.showPatientAppointments(patientID).subscribe(
    (response: any) => {
      console.log(response);
      
      this.appointments = response.data; // Update this line to match the structure of your response
      // Trigger change detection manually after assigning data
      this.cdr.detectChanges();
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

  reserveSlot(patientID: string, slotId: number): void {
    this.apiService.reserveSlot({ PatientID: patientID, SlotID: slotId }).subscribe(
      (response: any) => {
        console.log(response);
        this.loadAppointments(patientID);
      },
      (error: any) => {
        console.error(error);
      }
    );
  }
  
  cancelAppointment(slotId: number): void {
    const patientID = this.retrievePatientID();
  
    if (patientID) {
      this.apiService.cancelAppointment(patientID, slotId).subscribe(
        (response: any) => {
          console.log(response);
          this.loadAppointments(patientID); // Load appointments if patientID is available
        },
        (error: any) => {
          console.error(error);
        }
      );
    } else {
      console.error('PatientID not found');
    }
  }
  
  
  updateAppointment(slotId: number) {
    const selectedSlot = this.doctorSlots.find(slot => slot.slotId === slotId);
    console.log("Selected Slot ID:", slotId);
    console.log("Doctor Slots:", this.doctorSlots);
    
    if (selectedSlot) {
      const { slotId: newSlotId, doctorId } = selectedSlot;
      const patientID = this.retrievePatientID();
  
      if (patientID) {
        const data = {
          PatientID: patientID,
          AppointmentID: slotId, // Use slotId from the function parameter
          NewSlotID: newSlotId,
          NewDoctorID: doctorId
        };
  
        this.apiService.updateAppointment(patientID, data).subscribe(
          (response: any) => {
            console.log(response);
            this.loadAppointments(patientID);
          },
          (error: any) => {
            console.error(error);
          }
        );
      } else {
        console.error("Invalid patientID");
      }
    } else {
      console.error("Selected slot not found.");
    }
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