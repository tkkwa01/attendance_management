package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
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
}

func NewAttendance(ctx context.Context, req *request.CreateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmploymentID:     req.EmploymentID,
		Date:             req.Date,
		CheckInTime:      req.CheckInTime,
		CheckOutTime:     req.CheckOutTime,
		AttendanceNumber: req.AttendanceNumber,
	}
	err := validate.Validate(attendance)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
