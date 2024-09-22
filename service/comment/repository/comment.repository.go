package repositoryComment

import (
	"context"
	"net/http"

	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	entityComment "github.com/madmuzz05/go-final-project/service/comment/entity"
)

func (r CommentRepository) GetAll(ctx context.Context) (res []entityComment.Comment, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("SELECT * FROM public.comment ORDER BY id DESC").Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Not Found.")
		return
	}

	return
}

func (r CommentRepository) GetOne(ctx context.Context, id int) (res entityComment.Comment, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("SELECT * FROM public.comment where id = ? ORDER BY id DESC", id).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Not Found.")
		return
	}

	return
}

func (r CommentRepository) CreateComment(ctx context.Context, req entityComment.Comment) (res entityComment.Comment, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("INSERT INTO public.comment (created_at, updated_at, message, photo_id, user_id) VALUES (now(), now(),  ?, ?,?) RETURNING *",
		req.Message, req.PhotoId, req.UserId).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Not Found.")
		return
	}

	return
}

func (r CommentRepository) UpdateComment(ctx context.Context, req entityComment.Comment, id int) (res entityComment.Comment, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Raw("UPDATE public.comment SET updated_at = now(), message = ?, photo_id = ?, user_id = ? WHERE id = ? RETURNING *",
		req.Message, req.PhotoId, req.UserId, id).Scan(&res)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Not Found.")
		return
	}

	return
}

func (r CommentRepository) DeleteComment(ctx context.Context, id int) (err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	model := db.Exec("DELETE FROM public.comment WHERE id = ?", id)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	return
}
