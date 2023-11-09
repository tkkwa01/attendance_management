package request

import "time"

type CreateAttendance struct {
	ID           uint      `json:"id"`
	EmploymentID uint      `json:"employment_id"`
	CheckInTime  time.Time `json:"check_in_time"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
}

type UpdateAttendance struct {
	ID           uint       `json:"id"`
	EmploymentID uint       `json:"employment_id"`
	CheckInTime  time.Time  `json:"check_in_time"`
	CheckOutTime *time.Time `json:"check_out_time"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
}

type CheckOutAttendance struct {
	ID           uint      `json:"id"`
	CheckOutTime time.Time `json:"check_out_time"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
}
