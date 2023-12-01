import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-doctor-home',
  templateUrl: './doctor-home.component.html',
  styleUrls: ['./doctor-home.component.css']
})
export class DoctorHomeComponent implements OnInit {
  
  doctorSlots: any[] = [];
  doctorIdToShowSlots: number | null = null;
  slotForm: FormGroup;

  constructor(private apiService: ApiService, private formBuilder: FormBuilder, private cdr: ChangeDetectorRef, private activatedRoute: ActivatedRoute) {
    this.slotForm = this.formBuilder.group({
      dateTime: ['', Validators.required]
    });
  }

  ngOnInit(): void {}

  // onSubmit function in your Angular component
onSubmit(): void {
  if (this.slotForm.valid) {
    const dateTimeValue = this.slotForm.get('dateTime')?.value;
    const formattedDateTime = this.formatDateTime(dateTimeValue);

    this.SetSchedule(formattedDateTime);
  } else {
    console.error('Form is invalid');
  }
}

formatDateTime(dateTime: string): string {
  const date = new Date(dateTime);
  return date.toISOString(); // Convert date to ISO string format
}


SetSchedule(slotDateTime: string): void {
  const doctorID = this.retrieveDoctorID();
  if (doctorID !== null) {
    // Pass slotDateTime in the format YYYY-MM-DD HH:mm:ss
    this.apiService.setDoctorSchedule(doctorID, slotDateTime).subscribe(
      (response: any) => {
        console.log('Schedule set successfully:', response);
      },
      (error: any) => {
        console.error('Error while setting schedule:', error);
      }
    );
  } else {
    console.error('Doctor ID is null');
  }
}

  retrieveDoctorID(): string | null {
    // Implement your logic to retrieve the doctor ID from the URL or any other source
    const urlParts = window.location.href.split('/');
    const DoctorID = urlParts[urlParts.length - 1]; // Retrieve doctor ID from the URL
    return DoctorID ? DoctorID.toString() : null;
  }

  showDoctorSlots(): void {
    const doctorId = this.retrieveDoctorID();
    if (doctorId !== null) {
      const doctorIdNumber = parseInt(doctorId, 10); // Parse the string to a number
      if (!isNaN(doctorIdNumber)) {
        this.apiService.showDoctorSlots(doctorIdNumber).subscribe(
          (response: any) => {
            console.log(response);
            this.doctorSlots = response.data;
            console.log('Doctor Slots:', this.doctorSlots);
          },
          (error: any) => {
            console.error(error);
          }
        );
      } else {
        console.error('Invalid Doctor ID:', doctorId);
      }
    } else {
      console.error('Doctor ID is null');
    }
  }
  
 
  
  
}
