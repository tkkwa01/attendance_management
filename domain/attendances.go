package domain

import (
	"attendance-management/resource/request"
	"time"
)
import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	EmployeeNumber   uint       `json:"employee_number" gorm:"foreignKey:EmployeeNumber"`
	EmploymentID     uint       `json:"employment_id" gorm:"foreignKey:EmploymentID"`
	CheckInTime      time.Time  `json:"check_in_time"`
	CheckOutTime     *time.Time `json:"check_out_time"`
	AttendanceNumber uint       `json:"attendance_number"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
}

func NewAttendance(dto *request.CreateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmployeeNumber:   dto.EmployeeNumber,
		EmploymentID:     dto.EmploymentID,
		CheckInTime:      dto.CheckInTime,
		CheckOutTime:     nil,
		AttendanceNumber: dto.AttendanceNumber,
		Latitude:         dto.Latitude,
		Longitude:        dto.Longitude,
	}
	return attendance, nil
}

func UpdateAttendance(dto *request.UpdateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmployeeNumber:   dto.EmployeeNumber,
		EmploymentID:     dto.EmploymentID,
		CheckInTime:      dto.CheckInTime,
		CheckOutTime:     dto.CheckOutTime,
		AttendanceNumber: dto.AttendanceNumber,
		Latitude:         dto.Latitude,
		Longitude:        dto.Longitude,
	}
	return attendance, nil
}
