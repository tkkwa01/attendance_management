package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
	"time"
)

type Attendance struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	EmploymentID uint       `json:"employment_id" gorm:"foreignKey;references:ID"`
	CheckInTime  time.Time  `json:"check_in_time" gorm:"type:time;not null"`
	CheckOutTime *time.Time `json:"check_out_time"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
}

func NewAttendance(ctx context.Context, dto *request.CreateAttendance) (*Attendance, error) {
	attendance := Attendance{
		EmploymentID: dto.EmploymentID,
		CheckInTime:  dto.CheckInTime,
		CheckOutTime: nil,
		Latitude:     dto.Latitude,
		Longitude:    dto.Longitude,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return &attendance, nil
}

func UpdateAttendance(dto *request.UpdateAttendance) (*Attendance, error) {
	attendance := &Attendance{
		EmploymentID: dto.EmploymentID,
		CheckInTime:  dto.CheckInTime,
		CheckOutTime: dto.CheckOutTime,
		Latitude:     dto.Latitude,
		Longitude:    dto.Longitude,
	}
	return attendance, nil
}
