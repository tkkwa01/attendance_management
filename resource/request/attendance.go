package request

import "time"

type CreateAttendance struct {
	EmploymentID     uint      `json:"employment_id"`
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
}

type UpdateAttendance struct {
	EmploymentID     uint      `json:"employment_id"`
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
}
