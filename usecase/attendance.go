package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	request "attendance-management/resource/request"
	"github.com/pkg/errors"
	"time"
)

type (
	AttendanceInputPort interface {
		CheckIn(ctx context.Context, req *request.CreateAttendance, number uint) error
		CheckOut(ctx context.Context, number uint) error
		GetByID(ctx context.Context, number uint) error
		Update(ctx context.Context, req *request.UpdateAttendance) error
		Delete(ctx context.Context, number uint) error
	}
)

type AttendanceOutputPort interface {
	CheckIn(id uint) error
	CheckOut(res *domain.Attendance) error
	GetByID(res *domain.Attendance) error
	Update(res *domain.Attendance) error
	Delete() error
	Create(id uint) error
}

type AttendanceRepository interface {
	CheckIn(ctx context.Context, attendance *domain.Attendance) (uint, error)
	CheckOut(ctx context.Context, number uint) error
	Create(ctx context.Context, attendance *domain.Attendance) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Attendance, error)
	Update(ctx context.Context, attendance *domain.Attendance) error
	Delete(ctx context.Context, number uint) error
	NumberExist(ctx context.Context, number uint) (bool, error)
	GetByEmployeeNumberAndEmptyCheckout(ctx context.Context, number uint) (*domain.Attendance, error)
	GetByDate(ctx context.Context, date time.Time) ([]*domain.Attendance, error)
}

type attendance struct {
	outputPort     AttendanceOutputPort
	attendanceRepo AttendanceRepository
}

type AttendanceInputFactory func(outputPort AttendanceOutputPort) AttendanceInputPort

func NewAttendanceInputFactory(ar AttendanceRepository) AttendanceInputFactory {
	return func(o AttendanceOutputPort) AttendanceInputPort {
		return &attendance{
			outputPort:     o,
			attendanceRepo: ar,
		}
	}
}

func (a attendance) CheckIn(ctx context.Context, req *request.CreateAttendance, number uint) error {
	// 指定されたemployee_numberでcheckout_timeが空のレコードが存在するかチェック
	existingAttendance, err := a.attendanceRepo.GetByEmployeeNumberAndEmptyCheckout(ctx, number)
	if err != nil {
		return err
	}

	// CheckOutTimeが空のレコードが存在する場合、エラーを返す
	if existingAttendance != nil {
		return errors.New("既に出勤済みです。退勤を先に行ってください。")
	}
	// 新しい出勤データを作成
	newAttendance := &domain.Attendance{
		EmploymentID:     req.EmploymentID,
		AttendanceNumber: req.AttendanceNumber,
		Latitude:         req.Latitude,
		Longitude:        req.Longitude,
		CheckInTime:      time.Now(),
	}

	// データベースに新しい出勤データを保存
	_, err = a.attendanceRepo.CheckIn(ctx, newAttendance)
	if err != nil {
		return err
	}

	return a.outputPort.CheckIn(newAttendance.ID)
}

func (a attendance) CheckOut(ctx context.Context, number uint) error {
	attendance, err := a.attendanceRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	// 出勤データが存在しない場合の確認
	if attendance == nil {
		return errors.New("出勤データが存在しません。先に出勤してください。")
	}

	attendance.CheckOutTime = time.Now()

	err = a.attendanceRepo.CheckOut(ctx, number)
	if err != nil {
		return err
	}

	return a.outputPort.CheckOut(attendance)
}

func (a attendance) GetByID(ctx context.Context, number uint) error {
	attendance, err := a.attendanceRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	return a.outputPort.GetByID(attendance)
}

func (a attendance) Update(ctx context.Context, req *request.UpdateAttendance) error {
	attendance, err := a.attendanceRepo.GetByID(ctx, req.AttendanceNumber)
	if err != nil {
		return err
	}

	err = ctx.Transaction(func(ctx context.Context) error {

		oldID := req.AttendanceNumber
		err = a.attendanceRepo.Delete(ctx, oldID)
		if err != nil {
			return err
		}

		newAttendance, err := domain.UpdateAttendance(req)
		if err != nil {
			return err
		}

		_, err = a.attendanceRepo.Create(ctx, newAttendance)
		if err != nil {
			return err
		}

		err = a.attendanceRepo.Update(ctx, attendance)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return a.outputPort.Update(attendance)
}

func (a attendance) Delete(ctx context.Context, number uint) error {
	err := a.attendanceRepo.Delete(ctx, number)
	if err != nil {
		return err
	}

	return a.outputPort.Delete()
}
