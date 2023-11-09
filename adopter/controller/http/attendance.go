package http

import (
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type attendance struct {
	inputFactory  usecase.AttendanceInputFactory
	outputFactory func(c *gin.Context) usecase.AttendanceOutputPort
}

func NewAttendance(r *router.Router, inputFactory usecase.AttendanceInputFactory, outputFactory func(c *gin.Context) usecase.AttendanceOutputPort) {
	handler := attendance{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("attendances", nil, func(r *router.Router) {
		r.Post("check-in/:employment_id", handler.CheckIn)
		r.Put("check-out/:id", handler.CheckOut)
		r.Get("", handler.GetAttendance)
		r.Get("all", handler.GetAll)
		r.Put(":id", handler.Update)
		r.Delete(":id", handler.Delete)
	})
}

func (a attendance) CheckIn(ctx context.Context, c *gin.Context) error {
	var req request.CreateAttendance

	// IDを文字列として取得
	employmentIDStr := c.Param("employment_id")
	if employmentIDStr == "" {
		return errors.New("employment_id parameter is missing")
	}

	// 文字列をuintに変換
	employmentId, err := strconv.ParseUint(employmentIDStr, 10, 64)
	if err != nil {
		return errors.New("invalid employment_id parameter")
	}
	req.EmploymentID = uint(employmentId)

	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return err
	}

	ID := req.ID

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckIn(ctx, &req, ID)
}

func (a attendance) CheckOut(ctx context.Context, c *gin.Context) error {
	var req request.CheckOutAttendance

	// IDを文字列として取得
	IDStr := c.Param("id")
	if IDStr == "" {
		return errors.New("id parameter is missing")
	}
	// 文字列をuintに変換
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		return errors.New("invalid id parameter")
	}
	req.ID = uint(ID)
	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return err
	}

	number := req.ID

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckOut(ctx, number)
}

func (a attendance) GetAttendance(ctx context.Context, c *gin.Context) error {
	IDStr := c.Query("id")
	ID, err := stringToUint(IDStr)
	if err != nil {
		return err
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.GetByID(ctx, ID)
}

func (a attendance) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdateAttendance

	// IDを文字列として取得
	IDStr := c.Param("id")
	if IDStr == "" {
		return errors.New("id parameter is missing")
	}
	// 文字列をuintに変換
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		return errors.New("invalid id parameter")
	}
	req.ID = uint(ID)
	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return err
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (a attendance) Delete(ctx context.Context, c *gin.Context) error {
	// IDをパスパラメータから取得
	IDStr := c.Param("id")
	if IDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is missing"})
		return errors.New("id parameter is missing")
	}

	// 文字列をuintに変換
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return err
	}

	// outputとinputポートを初期化
	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	// inputPortのDeleteメソッドを使用して従業員を削除
	return inputPort.Delete(ctx, uint(ID))
}

func (a attendance) GetAll(ctx context.Context, c *gin.Context) error {
	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.GetAll(ctx)
}
