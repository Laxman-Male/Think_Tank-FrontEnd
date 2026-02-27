import { Injectable } from '@angular/core';
import { environment } from '../../environment/environment.component';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

   private apiUrl = environment.apiUrl;
 
   constructor(private http: HttpClient) { }

  
     GetProfileDetails() {
    const token = localStorage.getItem('token');

    const headers = new HttpHeaders({
      'Authorization': `Bearer ${token}`
    });

    return this.http.get(`${this.apiUrl}/profile`, { headers });
  }
}
