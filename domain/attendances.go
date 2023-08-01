package domain

import "time"
import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	ID           uint      `json:"id"`
	EmploymentID uint      `json:"employment_id"`
	Date         time.Time `json:"date"`
	CheckInTime  time.Time `json:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time"`
}
