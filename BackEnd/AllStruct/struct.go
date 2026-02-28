package allstruct

type LocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Hospital struct {
	ID                  int     `json:"id"`
	HospitalName        string  `json:"hospital_name"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	AvailableAmbulances int     `json:"available_ambulances"`
	Distance            float64 `json:"distance"`
}

type PatientData struct {
	Name     string   `json:"name"`
	Age      string   `json:"age"` // note: Angular sends age as string
	Gender   string   `json:"gender"`
	Blood    string   `json:"blood"`
	Type     string   `json:"type"`
	Hospital Hospital `json:"hospital"`
}

// Wrapper for Angular payload { patientData: { ... } }
type PatientRequest struct {
	PatientData PatientData `json:"patientData"`
}
