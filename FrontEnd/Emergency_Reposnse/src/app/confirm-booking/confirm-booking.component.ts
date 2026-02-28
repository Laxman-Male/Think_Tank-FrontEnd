import { Component } from '@angular/core';
import { Hospital } from '../model/getLocation';
import { Router } from '@angular/router';
import { LocationService } from '../service/services/location.service';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-confirm-booking',
  imports: [FormsModule, CommonModule],
  templateUrl: './confirm-booking.component.html',
  styleUrl: './confirm-booking.component.css'
})
export class ConfirmBookingComponent {

  
  hospital: Hospital | undefined;
  completeBooking:boolean=false;


  constructor(private router: Router, private hospitalService: LocationService) {}

   ngOnInit(): void {
    const state = history.state;
    if (state.hospital) {
      this.hospital = state.hospital;
    }
    this.isTokenAvailable()
  }

  

 isTokenAvailable(){
  let token= localStorage.getItem('token');
  if(token){

  }else{
    this.router.navigate(['/login'])
  }
 }
submitData(){
  let data = localStorage.getItem('patientData')
  console.log(data)
  if(data){
  let patientData= JSON.parse(data)
  console.log(patientData)
  
  
    this.hospitalService.ConfirmBooking(patientData)
      .subscribe({
        next: (res:any) => {
          // this.nearestHospitals=res;
          console.log('Nearest hospital:', res);
          this.completeBooking=true
        },
        error: (err) => {
          console.error('Error fetching hospitals:', err);
        }
      });

}
}


}
