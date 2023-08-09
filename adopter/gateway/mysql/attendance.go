package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type attendance struct{}

func NewAttendance() usecase.AttendanceRepository {
	return &attendance{}
}

func (a attendance) Create(ctx context.Context, attendance *domain.Attendance) (uint, error) {
	db := ctx.DB()

	if err := db.Create(attendance).Error; err != nil {
		return 0, dbError(err)
	}
	return attendance.ID, nil
}

func (a attendance) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var attendance domain.Attendance
	err := db.Where("number = ?", number).First(&attendance).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}

func (a attendance) GetByID(ctx context.Context, number uint) (*domain.Attendance, error) {
	db := ctx.DB()

	var attendance domain.Attendance
	err := db.Where("attendance_number = ?", number).First(&attendance).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &attendance, nil
}

func (a attendance) Update(ctx context.Context, attendance *domain.Attendance) error {
	db := ctx.DB()

	if err := db.Model(&attendance).Updates(attendance).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (a attendance) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var attendance domain.Attendance
	res := db.Where("attendance_number = ?", number).Delete(&attendance)
	if res.Error != nil {
		return dbError(res.Error)
	}
	if res.RowsAffected == 0 {
		return dbError(res.Error)
	}
	return nil
}
