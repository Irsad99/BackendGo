package histories

import (
	"BackendGo/src/helpers"
	"gorm.io/gorm"
)

var histories Histories
var response helpers.Response
var results Results

type history_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *history_repo {
	return &history_repo{grm}
}

func (r *history_repo) FindAll() (*helpers.Response, error) {

	result := r.db.Order("history_id desc").Find(&histories)

	if result.Error != nil {
		res := response.ResponseJSON(500, histories)
		return res, nil
	}

	res := response.ResponseJSON(200, histories)
	return res, nil
}

func (r *history_repo) FindHistoryByID(data *int) (*helpers.Response, error) {

	result := r.db.Raw(
	" select u.name as users, v.name as vehicle, h.start_date , h.end_date , h.prepayment , h.status " + 
	" from histories h , users u , vehicles v" + 
	" where u.user_id = ?" + 
	" and u.user_id = h.id_user and v.vehicle_id = h.id_vehicle", data).Scan(&results)

	if result.RowsAffected < 1 {
		res := response.ResponseJSON(404, results)
		return res, nil
	}

	if result.Error != nil {
		res := response.ResponseJSON(500, results)
		return res, nil
	}

	res := response.ResponseJSON(200, results)
	return res, nil
}

func (r *history_repo) SortByStart() (*helpers.Response, error) {

	result := r.db.Order("start_date").Find(&histories)

	if result.Error != nil {
		res := response.ResponseJSON(500, histories)
		return res, nil
	}

	res := response.ResponseJSON(200, histories)
	return res, nil
}

func (r *history_repo) Add(data *History) (*helpers.Response, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		res := response.ResponseJSON(400, histories)
		return res, nil
	}

	getData := r.db.First(&histories, &data.History_Id)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, histories)
		return res, nil
	}

	res := response.ResponseJSON(201, histories)
	return res, nil
}

func (r *history_repo) Delete(data *int) (*helpers.Response, error) {

	getData := r.db.First(&histories, data)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, histories)
		return res, nil
	}

	result := r.db.Delete(&History{}, data)

	if result.Error != nil {
		res := response.ResponseJSON(400, histories)
		return res, nil
	}

	res := response.ResponseJSON(204, histories)
	return res, nil
}

func (r *history_repo) Update(id *int, data *string) (*helpers.Response, error) {

	result := r.db.Model(&History{}).Where("history_id = ?", &id).Update("status", &data)

	if result.Error != nil {
		res := response.ResponseJSON(400, histories)
		return res, nil
	}

	getData := r.db.First(&histories, &id)
	if getData.RowsAffected < 1 {
		res := response.ResponseJSON(404, histories)
		return res, nil
	}

	res := response.ResponseJSON(201, histories)
	return res, nil
}