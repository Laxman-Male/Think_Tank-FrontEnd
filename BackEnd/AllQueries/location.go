package allqueries

const (
	GetNearestHospital = `
SELECT 
    id,
    hospital_name,
    latitude,
    longitude,
    available_ambulances,
    ROUND(
        6371 * ACOS(
            COS(RADIANS(?)) * COS(RADIANS(latitude)) *
            COS(RADIANS(longitude) - RADIANS(?)) +
            SIN(RADIANS(?)) * SIN(RADIANS(latitude))
        ), 2
    ) AS distance
FROM hospitals
WHERE available_ambulances > 0
ORDER BY distance ASC
LIMIT 5;
`

	ConfirmBooking = `
                INSERT INTO bookings
        (user_id,name, age, gender, blood_group, incident_type, hospital_details)
        VALUES (?, ?, ?, ?, ?, ?, ?)
        `
)
