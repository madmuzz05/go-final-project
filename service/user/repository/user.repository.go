package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/madmuzz05/go-final-project/pkg/helper/bcrypt"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	userDto "github.com/madmuzz05/go-final-project/service/user/dto"
	userEntity "github.com/madmuzz05/go-final-project/service/user/entity"
)

func (r UserRepository) GetDataUser(ctx context.Context, req userEntity.User) (res userEntity.User, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	var condition []string
	sql := `SELECT * FROM public.user `
	if req.Age != nil {
		condition = append(condition, fmt.Sprintf("age = %s", *req.Age))
	}
	if req.Email != nil {
		condition = append(condition, fmt.Sprintf("email like '%s'", *req.Email))
	}
	if req.Username != nil {
		condition = append(condition, fmt.Sprintf("username like '%s'", *req.Username))
	}
	if req.Id != 0 {
		condition = append(condition, fmt.Sprintf("id = %d", req.Id))
	}

	if len(condition) > 0 {
		sql += "where " + strings.Join(condition, " AND ")
	}

	model := db.Raw(sql).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "User Not Found.")
		return
	}
	return
}

func (r UserRepository) Register(ctx context.Context, req userDto.UserRequest) (res userEntity.User, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)

	hashPassword := bcrypt.HashPassword(*req.Password)

	sql := `INSERT INTO public.user (username, email, password, age, created_at, updated_at) values
			(?,?,?,?,now(),now())
		RETURNING *`

	user := db.Raw(sql, req.Username, req.Email, hashPassword, int(*req.Age)).Scan(&res)

	if user.Error != nil {
		err = sysresponse.GetErrorMessage(user.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	return
}
