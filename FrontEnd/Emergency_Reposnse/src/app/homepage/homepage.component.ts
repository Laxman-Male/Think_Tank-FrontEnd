import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-homepage',
  imports: [],
  templateUrl: './homepage.component.html',
  styleUrl: './homepage.component.css'
})
export class HomepageComponent {
  //  scrollToSection(id){
  //       document.getElementById(id).scrollIntoView({behavior:"smooth"});
  //   }
  constructor(private route:Router){}
  go_to_profile(){
    this.route.navigate(['/profile'])

  }
  go_to_feature(){
    this.route.navigate(['/feature'])
  }
}
