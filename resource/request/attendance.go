package request

import "time"

type CreateAttendance struct {
	EmploymentID     uint      `json:"employment_id"`
	Date             time.Time `json:"date"`
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number"`
}

type UpdateAttendance struct {
	EmploymentID     uint      `json:"employment_id"`
	Date             time.Time `json:"date"`
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number"`
}
