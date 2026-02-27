import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { LocationService } from '../service/services/location.service';

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

  constructor(private hospitalService: LocationService) {}

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
        next: (res) => {
          console.log('Nearest hospital:', res);
        },
        error: (err) => {
          console.error('Error fetching hospitals:', err);
        }
      });
  }
}