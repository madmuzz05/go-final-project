package repositoryPhoto

import (
	"context"
	"net/http"

	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoPhoto "github.com/madmuzz05/go-final-project/service/photo/dto"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
)

func (r PhotoRepository) GetAll(ctx context.Context) (res []entityPhoto.Photo, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("SELECT * FROM public.photo ORDER BY id DESC").Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Photo Not Found.")
		return
	}

	return
}

func (r PhotoRepository) GetOne(ctx context.Context, id int) (res entityPhoto.Photo, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("SELECT * FROM public.photo where id = ? ORDER BY id DESC", id).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Photo Not Found.")
		return
	}

	return
}

func (r PhotoRepository) CreatePhoto(ctx context.Context, req dtoPhoto.PhotoRequest) (res entityPhoto.Photo, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("INSERT INTO public.photo (created_at, updated_at, title, caption, photo_url, user_id) VALUES (now(), now(), ?, ?, ?,?) RETURNING *",
		req.Title, req.Caption, req.PhotoUrl, req.UserId).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Photo Not Found.")
		return
	}

	return
}

func (r PhotoRepository) UpdatePhoto(ctx context.Context, req dtoPhoto.PhotoRequest, id int) (res entityPhoto.Photo, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("UPDATE public.photo SET updated_at = now(), title = ?, caption = ?, photo_url = ?, user_id = ? WHERE id = ? RETURNING *",
		req.Title, req.Caption, req.PhotoUrl, req.UserId, id).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Photo Not Found.")
		return
	}

	return
}

func (r PhotoRepository) DeletePhoto(ctx context.Context, id int) (err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Exec("DELETE FROM public.photo WHERE id = ?", id)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	return
}
