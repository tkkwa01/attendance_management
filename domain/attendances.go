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
	CheckInTime      time.Time `json:"check_in_time"`
	CheckOutTime     time.Time `json:"check_out_time"`
	AttendanceNumber uint      `json:"attendance_number" gorm:"unique"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
}

func NewAttendance(req *request.CreateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmploymentID:     req.EmploymentID,
		CheckInTime:      req.CheckInTime,
		CheckOutTime:     req.CheckOutTime,
		AttendanceNumber: req.AttendanceNumber,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
	}
	//err := validation.Validate(attendance)エラーが出たから下に変更
	validator := validation.Validate()
	err := validator.Struct(attendance)

	if err != nil {
		return nil, err
	}
	return attendance, nil
}
