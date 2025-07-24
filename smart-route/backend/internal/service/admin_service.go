package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"smart-route/internal/model"
	"smart-route/pkg/convert"
	"smart-route/pkg/data/entity"
	"smart-route/pkg/data/repository"

	"gorm.io/gorm"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

func (s *AdminService) GetAdminByUsername(ctx context.Context, username string) (*model.Admin, error) {
	adminEntity, err := repository.GetAdminByUsername(ctx, s.db, username)
	if err != nil {
		return nil, err
	}
	var admin *model.Admin
	convert.CopyStructFields(admin, adminEntity)
	return admin, nil
}

func (s *AdminService) GetAdminByUsernameAndPassword(ctx context.Context, username, password string) (*model.Admin, error) {
	adminEntity, err := repository.GetAdminByUsernameAndPassword(ctx, s.db, username, password)
	if err != nil {
		return nil, err
	}
	var admin *model.Admin
	convert.CopyStructFields(admin, adminEntity)
	return admin, nil
}

func (s *AdminService) CreateAdmin(ctx context.Context, admin *model.Admin, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	var adminEntity *entity.Admin
	convert.CopyStructFields(adminEntity, admin)
	adminEntity.Password = string(hashed)
	return repository.CreateAdmin(ctx, s.db, adminEntity)
}

func (s *AdminService) CheckAdmin(ctx context.Context, username, password string) (*model.Admin, error) {
	adminEntity, err := repository.GetAdminByUsername(ctx, s.db, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid username or password"})
			return nil, errors.New("invalid username or password")
		} else {
			//c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid username or password"})
			return nil, errors.New("database error")
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(adminEntity.Password), []byte(password)) != nil {
		//c.JSON(http.StatusUnauthorized, api.ErrorResponse{Error: "Invalid username or password"})
		return nil, errors.New("invalid username or password")
	}
	// 这里可生成 JWT
	//token, err := h.jtwAuth.GenerateTokenAdmin(adminModel.ID.String())
	admin := model.Admin{}
	convert.CopyStructFields(&admin, adminEntity)
	return &admin, nil
}
