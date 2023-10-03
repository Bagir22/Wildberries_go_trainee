package handlers

import (
	"net/http"
	"dev11/middleware"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/create_event", middleware.Logging(http.HandlerFunc(Create)))
	router.HandleFunc("/update_event", middleware.Logging(http.HandlerFunc(Update)))
	router.HandleFunc("/delete_event", middleware.Logging(http.HandlerFunc(Delete)))
	router.HandleFunc("/events_for_day", middleware.Logging(http.HandlerFunc(EventsForDay)))
	router.HandleFunc("/events_for_week", middleware.Logging(http.HandlerFunc(EventsForWeek)))
	router.HandleFunc("/events_for_month", middleware.Logging(http.HandlerFunc(EventsForMonth)))

	return router
}
