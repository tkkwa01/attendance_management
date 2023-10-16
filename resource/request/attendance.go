package request

import "time"

type CreateAttendance struct {
	EmployeeNumber   uint      `json:"employee_number"`
	EmploymentID     uint      `json:"employment_id"`
	CheckInTime      time.Time `json:"check_in_time"`
	AttendanceNumber uint      `json:"attendance_number"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
}

type UpdateAttendance struct {
	EmployeeNumber   uint
	EmploymentID     uint       `json:"employment_id"`
	CheckInTime      time.Time  `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	AttendanceNumber uint       `json:"attendance_number"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
}

type CheckOutAttendance struct {
	AttendanceNumber uint      `json:"attendance_number"`
	CheckOutTime     time.Time `json:"check_out_time"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
}
