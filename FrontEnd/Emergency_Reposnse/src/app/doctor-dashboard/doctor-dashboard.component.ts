import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

interface Patient {
  name: string;
  age: number;
  blood: string;
  location: string;
}

@Component({
  selector: 'app-doctor-dashboard',
  imports: [CommonModule],
  templateUrl: './doctor-dashboard.component.html',
  styleUrl: './doctor-dashboard.component.css'
})
export class DoctorDashboardComponent {
  constructor(private route:Router){}
  showModal = false;
  showSuccessMessage = false;
  selectedPatient: Patient | null = null;
  
  
   ngOnInit(){
 console.log("hello")
this.isTokenAvailable()

 }

 isTokenAvailable(){
  let token= localStorage.getItem('token');
  if(token){

  }else{
    this.route.navigate(['/login'])
  }
 }
  patients: Patient[] = [
    { name: 'Laxman', age: 25, blood: 'A+', location: 'Pune' },
    { name: 'Shravani', age: 30, blood: 'B+', location: 'Pimpri' },
    { name: 'Mansi', age: 41, blood: 'O+', location: 'Hadapsar' },
    { name: 'Abhishek', age: 30, blood: 'AB-', location: 'Pune' }
  ];

  openAlert() {
    if (this.patients.length > 0) {
      this.selectedPatient = this.patients[0];
      this.showModal = true;
    }
  }

  closeModal() {
    this.showModal = false;
    this.selectedPatient = null;
  }

  acceptPatient() {
    this.showSuccessMessage = true;
    this.showModal = false;
    // Auto close success message after 3 seconds
    setTimeout(() => {
      this.showSuccessMessage = false;
    }, 3000);
  }
}
