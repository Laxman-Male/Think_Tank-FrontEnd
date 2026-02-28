import { CommonModule } from '@angular/common';
import { Component, ElementRef, ViewChild } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { LocationService } from '../service/services/location.service';
import { Hospital } from '../model/getLocation';
import { Router } from '@angular/router';

@Component({
  selector: 'app-emergency-form',
  imports: [CommonModule, FormsModule],
  templateUrl: './emergency-form.component.html',
  styleUrls: ['./emergency-form.component.css']
})
export class EmergencyFormComponent {

  latitude: number | null = null;
  longitude: number | null = null;
  error: string | null = null;
  locationRequested: boolean = false;
  loading: boolean = false;
  nearestHospitals: Hospital[] = [];

   @ViewChild('pName') pName!: ElementRef;
  @ViewChild('pAge') pAge!: ElementRef;
  @ViewChild('pGender') pGender!: ElementRef;
  @ViewChild('pBlood') pBlood!: ElementRef;
  @ViewChild('pType') pType!: ElementRef;

  constructor(private hospitalService: LocationService, private route:Router) {}

  // Request user's location
  requestLocation(): void {
    console.log('Requesting location...');
    this.locationRequested = true;
    this.loading = true;

    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          this.latitude = position.coords.latitude;
          this.longitude = position.coords.longitude;
          this.loading = false;
          console.log('Latitude:', this.latitude, 'Longitude:', this.longitude);
          this.error = null;
        },
        (err) => {
          this.loading = false;
          switch(err.code){
            case err.PERMISSION_DENIED:
              this.error = "You denied the location request. Please allow location access.";
              break;
            case err.POSITION_UNAVAILABLE:
              this.error = "Location information is unavailable.";
              break;
            case err.TIMEOUT:
              this.error = "The request to get your location timed out.";
              break;
            default:
              this.error = "An unknown error occurred.";
          }
          console.error(err);
        },
        { enableHighAccuracy: true, timeout: 10000 }
      );
    } else {
      this.error = 'Geolocation is not supported by this browser.';
      console.error(this.error);
      this.loading = false;
    }
  }

  // Call backend to get nearest hospitals
  getNearHospital(): void {
    // Assign to local variables to narrow types
    const lat = this.latitude;
    const lng = this.longitude;

    if (lat == null || lng == null) {
      alert("Location not available yet.");
      return;
    }

    this.hospitalService.findNearestHospital(lat, lng)
      .subscribe({
        next: (res:any) => {
          this.nearestHospitals=res;
          console.log('Nearest hospital:', res);
        },
        error: (err) => {
          console.error('Error fetching hospitals:', err);
        }
      });
  }

 goToHospitalDetails(hospital: Hospital) {
    const patientData = {
      name: this.pName.nativeElement.value,
      age: this.pAge.nativeElement.value,
      gender: this.pGender.nativeElement.value,
      blood: this.pBlood.nativeElement.value,
      type: this.pType.nativeElement.value,
      hospital: hospital
    };

        localStorage.setItem('patientData', JSON.stringify(patientData));
        
  this.route.navigate(['/confirm-booking'], { state: { hospital } });
}
}