import { Component } from '@angular/core';

@Component({
  selector: 'app-login',
  imports: [],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  userType: 'User' | 'Hospital' = 'User'; // Default selection

  setUserType(type: 'User' | 'Hospital') {
    this.userType = type;
  }
}
