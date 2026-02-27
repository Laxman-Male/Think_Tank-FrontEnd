import { Component } from '@angular/core';
import { AuthService } from '../../service/services/auth.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  imports: [CommonModule,FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  constructor(private authService: AuthService, private router: Router) {}
  

  userType: 'User' | 'Hospital' = 'User'; // Default selection

  loginData = {
    mobile:'',
    email: '',
    password: '',
    role: 'User',
    isNewUser: false
  };

  setUserType(type: 'User' | 'Hospital') {
    this.userType = type;
  }

 onSubmit() {
  if(this.loginData.isNewUser==true){
    console.log(this.loginData)
    // return
   this.authService.register(this.loginData).subscribe({
      next: (response) => {
        // Save the JWT sent by your Go backend
        localStorage.setItem('token', response.token);
        
        // Redirect based on the role to solve "real-time coordination"
        if (this.userType === 'Hospital') {
          this.router.navigate(['/hospital-dashboard']);
        } else {
          this.router.navigate(['/sos-dashboard']);
        }
      },
      error: (err) => {
        console.error('Login Error:', err);
        alert('Invalid credentials. Check your Go server!');
      }
    });
  } else{
   this.authService.login(this.loginData).subscribe({
      next: (response) => {
        // Save the JWT sent by your Go backend

        localStorage.setItem('token', response.token);
        
        // Redirect based on the role to solve "real-time coordination"
        if (this.userType === 'Hospital') {
          this.router.navigate(['/hospital-dashboard']);
        } else {
          this.router.navigate(['']);
        }
      },
      error: (err) => {
        console.error('Login Error:', err);
        alert('Invalid credentials. Check your Go server!');
      }
    });
  }
  console.log(this.loginData.isNewUser)
 
  }

}
