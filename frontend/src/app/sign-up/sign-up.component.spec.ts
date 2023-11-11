import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DoctorSignupComponent } from './sign-up.component';

describe('DoctorSignupComponent', () => {
  let component: DoctorSignupComponent;
  let fixture: ComponentFixture<DoctorSignupComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DoctorSignupComponent]
    });
    fixture = TestBed.createComponent(DoctorSignupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
