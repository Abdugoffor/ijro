package form_service

import (
	"context"
	form_dto "ijro-nazorat/modul/form/dto"
	form_model "ijro-nazorat/modul/form/model"

	"git.sriss.uz/shared/shared_service/pg"
	"git.sriss.uz/shared/shared_service/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppService interface {
	All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.AppPage, error)
	Create(ctx echo.Context, req form_dto.ApplicationCreate) (int, error)
}

type appService struct {
	db *gorm.DB
}

func NewAppService(db *gorm.DB) AppService {
	return &appService{
		db: db,
	}
}

func (service *appService) All(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*form_dto.AppPage, error) {
	return pg.PageWithScan[form_model.App, form_dto.AppInfo](service.db, paginate, filter)
}

func (service *appService) Create(ctx echo.Context, req form_dto.ApplicationCreate) (int, error) {

	var appID int

	err := service.db.Transaction(func(tx *gorm.DB) error {

		app := form_model.App{
			AppCategoryID: req.AppCategoryID,
		}

		if err := tx.Create(&app).Error; err != nil {
			return err
		}

		appID = app.ID // ✅ shu yerda olib qo‘yamiz

		answers := make([]form_model.AppInfo, 0, len(req.Answers))
		for _, ans := range req.Answers {
			answers = append(answers, form_model.AppInfo{
				AppID:  app.ID,
				PageID: req.PageID,
				FormID: ans.FormID,
				Answer: ans.Answer,
			})
		}

		if len(answers) > 0 {
			if err := tx.Create(&answers).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return appID, nil
}
