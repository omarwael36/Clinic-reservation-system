CREATE DATABASE clinic_reservation_system;
use clinic_reservation_system;
CREATE TABLE doctor (
    DoctorID INT AUTO_INCREMENT PRIMARY KEY,
    DoctorName VARCHAR(255) NOT NULL,
    DoctorEmail VARCHAR(255) NOT NULL UNIQUE,
    DoctorPassword VARCHAR(255) NOT NULL
);

CREATE TABLE patient (
    PatientID INT AUTO_INCREMENT PRIMARY KEY,
    PatientName VARCHAR(255) NOT NULL,
    PatientEmail VARCHAR(255) NOT NULL UNIQUE,
    PatientPassword VARCHAR(255) NOT NULL
);

CREATE TABLE slot (
    SlotID INT AUTO_INCREMENT PRIMARY KEY,
    SlotDateTime DATETIME NOT NULL,
    DoctorID INT,
    PatientID INT,
    FOREIGN KEY (DoctorID) REFERENCES doctor(DoctorID),
    FOREIGN KEY (PatientID) REFERENCES patient(PatientID)
);
