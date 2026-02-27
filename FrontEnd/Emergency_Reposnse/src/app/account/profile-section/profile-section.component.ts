import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileService } from '../../service/services/profile.service';
import { Profile } from '../../model/profile';

@Component({
  selector: 'app-profile-section',
  imports: [],
  templateUrl: './profile-section.component.html',
  styleUrl: './profile-section.component.css'
})
export class ProfileSectionComponent {
 constructor (private route:Router, private profileService:ProfileService){}
UserProfile?: Profile
 ngOnInit(){
 console.log("hello")
this.getProfileDetails()

 }
goToHome(){
  this.route.navigate([''])
}

getProfileDetails(){
  this.profileService.GetProfileDetails().subscribe({
      next: (response:any) => {
        this.UserProfile=response;
       console.log("response",response)
      },
      error: (err:any) => {
        console.error('Error:', err);
        alert('no profile details');
      }
    });
}
 

}
