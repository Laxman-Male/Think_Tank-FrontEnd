import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environment/environment.component';
import { Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LocationService {

  constructor(private http:HttpClient) { }
 private apiUrl = environment.apiUrl;

    findNearestHospital(lat: number, lng: number): Observable<Object> {
    const token = localStorage.getItem('token');

    if (!token) {
      // Instead of returning undefined, return an Observable error
      return throwError(() => new Error('You must login first'));
    }

    const headers = new HttpHeaders({
      'Authorization': `Bearer ${token}`
    });

    return this.http.post(`${this.apiUrl}/nearest-hospital`, 
      { latitude: lat, longitude: lng }, 
      { headers }
    );
  }

    ConfirmBooking(patientData:any): Observable<Object> {
    const token = localStorage.getItem('token');

    if (!token) {
      // Instead of returning undefined, return an Observable error
      return throwError(() => new Error('You must login first'));
    }

    const headers = new HttpHeaders({
      'Authorization': `Bearer ${token}`
    });

    return this.http.post(`${this.apiUrl}/confirm-booking`, 
      {  patientData }, 
      { headers }
    );
  }
}
