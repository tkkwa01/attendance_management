package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type attendance struct{}

func (a attendance) GetByDate(ctx context.Context, date time.Time) ([]*domain.Attendance, error) {
	//TODO implement me
	panic("implement me")
}

func (a attendance) CheckIn(ctx context.Context, attendance *domain.Attendance) error {
	//TODO implement me
	panic("implement me")
}

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

func (a attendance) CheckOut(ctx context.Context, number uint) error {
	db := ctx.DB()
	// 出席記録を検索
	var targetAttendance domain.Attendance
	res := db.Where("attendance_number = ?", number).First(&targetAttendance)
	if res.Error != nil {
		return dbError(res.Error)
	}

	if res.RowsAffected == 0 {
		return dbError(fmt.Errorf("No attendance record found"))
	}

	// チェックアウト時間を追加
	targetAttendance.CheckOutTime = time.Now()

	// 更新を保存
	res = db.Save(&targetAttendance)
	if res.Error != nil {
		return dbError(res.Error)
	}
	if res.RowsAffected == 0 {
		return dbError(fmt.Errorf("Failed to update attendance record"))
	}

	return nil
}

func (a attendance) GetByEmployeeNumberAndEmptyCheckout(ctx context.Context, number uint) (*domain.Attendance, error) {
	db := ctx.DB()

	var attendance domain.Attendance
	err := db.Where("employee_number = ? AND checkout_time IS NULL", number).First(&attendance).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}
