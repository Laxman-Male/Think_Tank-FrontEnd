import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-emergency-form',
  imports: [CommonModule, FormsModule],
  templateUrl: './emergency-form.component.html',
  styleUrl: './emergency-form.component.css'
})
export class EmergencyFormComponent {

  latitude: number | null = null;
  longitude: number | null = null;
  error: string | null = null;
locationRequested: boolean = false;
loding:boolean = false;
  constructor() {}

  requestLocation(): void {
    console.log('Requesting location...');
  this.locationRequested = true; // show message while requesting
 this.loding=true
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          this.latitude = position.coords.latitude;
          this.longitude = position.coords.longitude;
          this.loding=false;
          console.log('Latitude:', this.latitude, 'Longitude:', this.longitude);
          this.error = null;
        },
        (err) => {
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
    }
  }
}
