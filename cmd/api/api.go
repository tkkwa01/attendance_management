package api

import (
	httpController "attendance-management/adopter/controller/http"
	"attendance-management/adopter/gateway/mail"
	mysqlRepository "attendance-management/adopter/gateway/mysql"
	"attendance-management/adopter/presenter"
	"attendance-management/config"
	"attendance-management/driver"
	"attendance-management/packages/http/middleware"
	"attendance-management/packages/http/router"
	"attendance-management/packages/log"
	"attendance-management/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Execute() {
	logger := log.Logger()
	defer logger.Sync()

	engine := gin.New()

	engine.GET("health", func(c *gin.Context) { c.Status(http.StatusOK) })

	// cors
	engine.Use(middleware.Cors(nil))

	r := router.New(engine, driver.GetRDB)

	// dependencies injection
	// ----- gateway -----
	mailAdapter := mail.New()

	//mysql
	employeeRepository := mysqlRepository.NewEmployee()
	companyRepository := mysqlRepository.NewCompany()
	employmentRepository := mysqlRepository.NewEmployment()
	positionRepository := mysqlRepository.NewPosition()
	salaryTypeRepository := mysqlRepository.NewSalaryType()
	attendanceRepository := mysqlRepository.NewAttendance()

	//usecase
	employeeInputFactory := usecase.NewEmployeeInputFactory(employeeRepository, mailAdapter)
	employeeOutputFactory := presenter.NewEmployeeOutputFactory()
	companyInputFactory := usecase.NewCompanyInputFactory(companyRepository)
	companyOutputFactory := presenter.NewCompanyOutputFactory()
	employmentInputFactory := usecase.NewEmploymentInputFactory(employmentRepository)
	employmentOutputFactory := presenter.NewEmploymentOutputFactory()
	positionInputFactory := usecase.NewPositionInputFactory(positionRepository)
	positionOutputFactory := presenter.NewPositionOutputFactory()
	salaryTypeInputFactory := usecase.NewSalaryTypeInputFactory(salaryTypeRepository)
	salaryTypeOutputFactory := presenter.NewSalaryTypeOutputFactory()
	attendanceInputFactory := usecase.NewAttendanceInputFactory(attendanceRepository)
	attendanceOutputFactory := presenter.NewAttendanceOutputFactory()

	//controller
	httpController.NewEmployee(r, employeeInputFactory, employeeOutputFactory)
	httpController.NewCompany(r, companyInputFactory, companyOutputFactory)
	httpController.NewEmployment(r, employmentInputFactory, employmentOutputFactory)
	httpController.NewPosition(r, positionInputFactory, positionOutputFactory)
	httpController.NewSalaryType(r, salaryTypeInputFactory, salaryTypeOutputFactory)
	httpController.NewAttendance(r, attendanceInputFactory, attendanceOutputFactory)
	//serve
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.Port),
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
			// logger.Error(fmt.Sprintf("Server forced to shutdown: %+v", err))
		}
	}()

	logger.Info("Succeeded in listen and serve.")

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("Server forced to shutdown: %+v", err))
	}

	logger.Info("Server existing")
}
