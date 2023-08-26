package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
	"time"
)

type AttendanceInputPort interface {
	Create(ctx context.Context, req *request.CreateAttendance) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdateAttendance) error
	Delete(ctx context.Context, number uint) error
}

type AttendanceOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Attendance) error
	Update(res *domain.Attendance) error
	Delete() error
}

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *domain.Attendance) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Attendance, error)
	Update(ctx context.Context, attendance *domain.Attendance) error
	NumberExist(ctx context.Context, number uint) error
	Delete(ctx context.Context, number uint) error
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

func (a attendance) Create(ctx context.Context, req *request.CreateAttendance) error {
	res, err := domain.NewAttendance(ctx, req)
	// req.AttendanceNumberをkeyにしてattendanceが存在するか確認
	err = a.attendanceRepo.NumberExist(ctx, req.AttendanceNumber)
	if err == nil {
		return ctx.Error().BadRequest("attendance already exist")
	}

	id, err := a.attendanceRepo.Create(ctx, res)
	if err != nil {
		return err
	}

	return a.outputPort.Create(id)
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

		newAttendance, err := domain.NewAttendance(ctx, (*request.CreateAttendance)(req))
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

func (a attendance) CheckOut(ctx context.Context, number uint) error {
	attendance, err := a.attendanceRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	attendance.CheckOutTime = time.Now()

	err = a.attendanceRepo.Update(ctx, attendance)
	if err != nil {
		return err
	}

	return a.outputPort.Update(attendance)
}
