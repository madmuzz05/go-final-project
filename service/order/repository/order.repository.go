package repository

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	"github.com/madmuzz05/go-final-project/service/order/dto/request"
	"github.com/madmuzz05/go-final-project/service/order/dto/response"
	"github.com/madmuzz05/go-final-project/service/order/entity"
)

func (r OrderRepository) GetOrders(ctx context.Context) (res []response.OrdersDtoResponse, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	var order []entity.Order
	model := db.Raw("SELECT * FROM assignment2.order ORDER BY order_id DESC").Scan(&order)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Tidak Ditemukan.")
		return
	}

	for _, x := range order {
		tempRes := response.OrdersDtoResponse{}
		listItem := []response.ItemsDtoResponse{}

		copier.Copy(&tempRes, &x)

		modelItem := db.Raw("SELECT * FROM assignment2.item Where order_id = ? ORDER BY item_id ASC", x.OrderId).Scan(&listItem)

		if modelItem.Error != nil {
			err = sysresponse.GetErrorMessage(modelItem.Error, http.StatusInternalServerError, "Internal Server Error.")
			return
		}
		tempRes.Items = listItem

		res = append(res, tempRes)

	}

	return
}

func (r OrderRepository) GetOrder(ctx context.Context, id string) (res response.OrdersDtoResponse, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)
	var order entity.Order
	model := db.Raw("SELECT * FROM assignment2.order where order_id = ? ORDER BY order_id DESC", id).Scan(&order)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if model.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Tidak Ditemukan.")
		return
	}

	tempRes := response.OrdersDtoResponse{}
	listItem := []response.ItemsDtoResponse{}
	copier.Copy(&tempRes, &order)

	modelItem := db.Raw("SELECT * FROM assignment2.item Where order_id = ? ORDER BY item_id ASC", id).Scan(&listItem)

	if modelItem.Error != nil {
		err = sysresponse.GetErrorMessage(modelItem.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	tempRes.Items = listItem

	res = tempRes

	return
}

func (r OrderRepository) StoreOrder(ctx context.Context, req request.OrderDtoRequest, id string) (res response.OrdersDtoResponse, err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)

	var sql string
	var tempId string
	if id == "" {
		sql = `INSERT INTO assignment2.order (customer_name, ordered_at, created_at, updated_at)
		VALUES ('` + req.CustomerName + `','` + req.OrderedAt + `',NOW(),NOW())
		RETURNING order_id`
	} else {
		model := db.Exec("SELECT * FROM assignment2.order where order_id = ? ORDER BY order_id DESC", id)

		if model.Error != nil {
			err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
			return
		} else if model.RowsAffected == 0 {
			err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Tidak Ditemukan.")
			return
		}

		sql = `UPDATE assignment2.order set customer_name = '` + req.CustomerName + `', updated_at = NOW() 
		WHERE order_id = ` + id + `
		RETURNING order_id
		`
	}
	order := db.Raw(sql).Scan(&tempId)

	if order.Error != nil {
		err = sysresponse.GetErrorMessage(order.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}
	if id != "" {
		delItem := db.Exec(`DELETE FROM assignment2.item where order_id = ?`, &tempId)
		if delItem.Error != nil {
			err = sysresponse.GetErrorMessage(delItem.Error, http.StatusInternalServerError, "Internal Server Error.")
			return
		}
	}

	listItem := []request.ItemDtoRequest{}
	copier.Copy(&listItem, &req.Items)
	for _, x := range listItem {
		resItem := db.Exec(`INSERT INTO assignment2.item (item_code, description, quantity, created_at, updated_at, order_id)
			values(?,?,?,NOW(),NOW(),?)`, x.ItemCode, x.Description, x.Quantity, &tempId)
		if resItem.Error != nil {
			err = sysresponse.GetErrorMessage(resItem.Error, http.StatusInternalServerError, "Internal Server Error.")
			return
		}
	}

	res, _ = r.GetOrder(ctx, id)

	return
}

func (r OrderRepository) DeleteOrder(ctx context.Context, id string) (err sysresponse.IError) {
	db := r.gormDb.GetDB().WithContext(ctx)

	findData := db.Exec("SELECT * FROM assignment2.order where order_id = ? ORDER BY order_id DESC", id)

	if findData.Error != nil {
		err = sysresponse.GetErrorMessage(findData.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	} else if findData.RowsAffected == 0 {
		err = sysresponse.GetErrorMessage(nil, http.StatusNotFound, "Data Tidak Ditemukan.")
		return
	}

	modelItem := db.Exec("DELETE FROM assignment2.item Where order_id = ?", id)

	if modelItem.Error != nil {
		err = sysresponse.GetErrorMessage(modelItem.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	model := db.Exec("DELETE FROM assignment2.order where order_id = ?", id)

	if model.Error != nil {
		err = sysresponse.GetErrorMessage(model.Error, http.StatusInternalServerError, "Internal Server Error.")
		return
	}

	return
}
