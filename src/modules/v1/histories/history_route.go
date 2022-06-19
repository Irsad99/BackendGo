package histories

import (
	"BackendGo/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/date", middleware.Do(ctrl.SortByStartDate, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/{id}", middleware.Do(ctrl.GetHistoryByID, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/register", middleware.Do(ctrl.AddData, "user", middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Delete, "admin", middleware.CheckAuth)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, "user", middleware.CheckAuth)).Methods("PUT")
}
