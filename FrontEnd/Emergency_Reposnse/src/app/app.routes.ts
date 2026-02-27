import { Routes } from '@angular/router';
import { HomepageComponent } from './homepage/homepage.component';
import { LoginComponent } from './account/login/login.component';
import { Component } from 'lucide-angular';
import { ProfileSectionComponent } from './account/profile-section/profile-section.component';
import { FeatureComponent } from './feature/feature.component';

export const routes: Routes = [
    {
        path:'',
        component:HomepageComponent
    },
    {
        path: 'login',
        component:LoginComponent
    },
    {
        path:'profile',
        component:ProfileSectionComponent
    },
    {
        path:'feature',
        component:FeatureComponent
    }
];
