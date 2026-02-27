import { Component } from '@angular/core';

@Component({
  selector: 'app-feature',
  imports: [],
  templateUrl: './feature.component.html',
  styleUrl: './feature.component.css'
})
export class FeatureComponent {
 

// Fake Hospital Data
// const hospitals=[
// {name:"CityCare Hospital",beds:5,distance:4},
// {name:"LifeLine Medical",beds:2,distance:2},
// {name:"Apollo Health Center",beds:7,distance:6}
// ];

// const bloodDatabase=["A+","B+","O+","AB+","O-"];

//   register(){
//     profileData.innerHTML=
//     `<b>Name:</b> ${name.value}<br>
//      <b>Age:</b> ${age.value}<br>
//      <b>Phone:</b> ${phone.value}<br>
//      <b>Blood:</b> ${blood.value}<br>
//      <b>Address:</b> ${address.value}`;

//     loginSection.classList.add("hidden");
//     profileSection.classList.remove("hidden");
// }

//   goDashboard(){
//     profileSection.classList.add("hidden");
//     dashboard.classList.remove("hidden");
// }

// function openEmergency(){
//     dashboard.classList.add("hidden");
//     emergencySection.classList.remove("hidden");
// }

// function submitEmergency(){

//     emergencySection.classList.add("emergency-mode");

//     const patientName=eName.value||"Unknown";
//     const patientAge=eAge.value;
//     const patientBlood=eBlood.value;
//     const patientLocation=eAddress.value;

//     // Doctor Pre-Alert
//     doctorAlert.innerHTML=
//     `<div style="background:#e3f2fd;padding:10px;border-radius:10px;margin-top:10px;">
//     🩺 <b>Doctor Alert:</b><br>
//     Age: ${patientAge} <br>
//     Blood: ${patientBlood} <br>
//     ETA: 10 mins <br>
//     Prepare Trauma Unit & Emergency Kit.
//     </div>`;

//     // Ambulance Simulation
//     tracking.classList.remove("hidden");
//     let progress=0;
//     const interval=setInterval(()=>{
//         progress+=10;
//         progressBar.style.width=progress+"%";
//         etaText.innerText="ETA: "+(10-progress/10)+" mins";
//         if(progress>=100){
//             clearInterval(interval);
//             etaText.innerText="Ambulance Arrived!";
//         }
//     },1000);

//     // Smart Hospital Prediction
//     let bestHospital=hospitals.sort((a,b)=>b.beds-a.beds)[0];
//     bedDashboard.innerHTML=
//     `Recommended Hospital: <b>${bestHospital.name}</b> (${bestHospital.beds} beds available)`;

//     // Blood Match
//     if(bloodDatabase.includes(patientBlood)){
//         bloodResult.innerHTML="✅ Blood Available";
//     }else{
//         bloodResult.innerHTML="⚠ Blood Not Available – Alerting Nearby Donors";
//     }

//     // Emergency Log
//     emergencyTable.innerHTML+=
//     `<tr><td>${patientName}</td><td>${patientAge}</td><td>${patientBlood}</td><td>${patientLocation}</td></tr>`;
// }

 
}
