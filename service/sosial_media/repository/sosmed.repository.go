package repositorySosmed

import (
	"context"
	"net/http"

	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	dtoSosmed "github.com/madmuzz05/go-final-project/service/sosial_media/dto"
	entitySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/entity"
)

func (r SosmedRepository) GetAll(ctx context.Context) (res []entitySosmed.SosialMedia, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	sosmed := db.Raw("SELECT * FROM public.sosialmedia ORDER BY id DESC").Scan(&res)

	if sosmed.Error != nil {
		err = sysresponse.GetErrorMessage(sosmed.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if sosmed.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Social Media Not Found.")
		return
	}

	return
}

func (r SosmedRepository) GetOne(ctx context.Context, id int) (res entitySosmed.SosialMedia, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	sosmed := db.Raw("SELECT * FROM public.sosialmedia where id = ? ORDER BY id DESC", id).Scan(&res)

	if sosmed.Error != nil {
		err = sysresponse.GetErrorMessage(sosmed.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if sosmed.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Social Media Not Found.")
		return
	}

	return
}

func (r SosmedRepository) CreateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest) (res entitySosmed.SosialMedia, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	sosmed := db.Raw("INSERT INTO public.sosialmedia (created_at, updated_at, name, sosial_media_url, user_id) VALUES (now(), now(), ?, ?, ?) RETURNING *",
		req.Name, req.SosialMediaUrl, req.UserId).Scan(&res)

	if sosmed.Error != nil {
		err = sysresponse.GetErrorMessage(sosmed.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if sosmed.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Social Media Not Found.")
		return
	}

	return
}

func (r SosmedRepository) UpdateSocialMedia(ctx context.Context, req dtoSosmed.SosmedRequest, id int) (res entitySosmed.SosialMedia, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	sosmed := db.Raw("UPDATE public.sosialmedia SET updated_at = now(), name = ?, sosial_media_url = ?, user_id = ? WHERE id = ? RETURNING *",
		req.Name, req.SosialMediaUrl, req.UserId, id).Scan(&res)

	if sosmed.Error != nil {
		err = sysresponse.GetErrorMessage(sosmed.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if sosmed.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Social Media Not Found.")
		return
	}

	return
}

func (r SosmedRepository) DeleteSocialMedia(ctx context.Context, id int) (err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	sosmed := db.Exec("DELETE FROM public.sosialmedia WHERE id = ?", id)

	if sosmed.Error != nil {
		err = sysresponse.GetErrorMessage(sosmed.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	return
}
