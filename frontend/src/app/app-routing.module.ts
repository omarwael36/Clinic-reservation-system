import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { DoctorSignupComponent } from './sign-up/sign-up.component';
import { PatientHomeComponent } from './patient-home/patient-home.component';
import { DoctorHomeComponent } from './doctor-home/doctor-home.component';

const routes: Routes = [{path:"",component:LoginComponent}, {path:"signup", component:DoctorSignupComponent}, {path:"UserPatient", component:PatientHomeComponent}, {path:"UserDoctor", component:DoctorHomeComponent}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
