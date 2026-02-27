export interface Hospital {
  id: number;
  hospital_name: string;
  latitude: number;
  longitude: number;
  available_ambulances: number;
  distance: number; // distance in km
}