package domain

import (
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
	"time"
)
import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	ID               uint      `json:"id"`
	EmploymentID     uint      `json:"employment_id"`
	Date             time.Time `json:"date"`
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number" gorm:"unique"`
	latitude         float64   `json:"latitude"`
	longitude        float64   `json:"longitude"`
}

func NewAttendance(req *request.CreateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmploymentID:     req.EmploymentID,
		Date:             req.Date,
		CheckInTime:      req.CheckInTime,
		CheckOutTime:     req.CheckOutTime,
		AttendanceNumber: req.AttendanceNumber,
		latitude:         req.Latitude,
		longitude:        req.Longitude,
	}
	err := validation.Validate(attendance)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
