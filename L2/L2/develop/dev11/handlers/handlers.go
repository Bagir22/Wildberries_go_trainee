package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Event struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}

type Calendar struct {
	sync.RWMutex
	events map[string][]Event
}

var (
	c = NewCalendar()
)

func NewCalendar() *Calendar {
	c := Calendar{
		events: make(map[string][]Event),
	}
	return &c
}

//Обрабатываем запрос на создание нового события
func Create(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = createEvent(event)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(c)
}

func createEvent(e *Event) error {
	c.RLock()
	//Проверяем есть ли текущее событие
	for _, v := range c.events[e.UserId] {
		if v.Id == e.Id {
			return errors.New("Current event exsit")
		}
	}
	c.RUnlock()

	c.RLock()
	//Создаем новое событие
	c.events[e.UserId] = append(c.events[e.UserId], *e)
	c.RUnlock()

	return nil
}

//Обрабатываем запрос на обновление события
func Update(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	//Проверям есть ли событие
	if len(c.events[event.UserId]) == 0 {
		err = errors.New("Current event is not exsit")
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	c.RLock()
	//Обновление даты и заголовка события
	for i, v := range c.events[event.UserId] {
		if v.Id == event.Id {
			c.events[event.UserId][i].Title = event.Title
			c.events[event.UserId][i].Date = event.Date
		}
	}
	c.RUnlock()

	fmt.Println(c)
}

//Обрабатываем запрос на удаление события
func Delete(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	c.RLock()
	for i, v := range c.events[event.UserId] {
		//Ищем событие и вырезаем его
		if v.Id == event.Id {
			c.events[event.UserId] = append(c.events[event.UserId][:i], c.events[event.UserId][i+1:]...)
		}
	}
	c.RUnlock()

	fmt.Println(c)
}

//Обрабатываем запрос на события за день
func EventsForDay(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventFromRequest(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	result := getDayList(*event)

	writeOkResponse(w, result)
}

//Обрабатываем запрос на события за неделю
func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventFromRequest(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	result, err := getWeekList(*event)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
	}

	writeOkResponse(w, result)
}

//Обрабатываем запрос на события за месяц
func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventFromRequest(r)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	result, err := getMonthList(*event)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	writeOkResponse(w, result)
}


func getDayList(event Event) []Event {
	var result []Event

	c.RLock()

	for _, v := range c.events[event.UserId] {
		//Добавляем все события если дата совпадает полностью
		if event.Date == v.Date {
			result = append(result, v)
		}
	}
	c.RUnlock()

	return result
}

func getWeekList(event Event) ([]Event, error) {
	//Получаем год и неделю
	var result []Event
	timeS, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return nil, errors.New("Can't parse week")
	}
	searchedYear, searchedWeek := timeS.ISOWeek()

	c.RLock()

	for _, v := range c.events[event.UserId] {
		//Получаем год и неделю
		curr, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return nil, errors.New("Can't parse week")
		}
		year, week := curr.ISOWeek()
		//Если совпадает год и неделя, то добавляем
		if searchedYear == year && searchedWeek == week {
			result = append(result, v)
		}
	}
	c.RUnlock()

	return result, nil
}

func getMonthList(event Event) ([]Event, error) {
	var result []Event
	//Получаем год и месяц
	timeS, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return nil, errors.New("Can't parse month")
	}
	searchedYear, searchedMonth := timeS.Year(), timeS.Month()

	c.RLock()

	for _, v := range c.events[event.UserId] {
		//Получаем год и месяц
		curr, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return nil, errors.New("Can't parse week")
		}
		year, month := curr.Year(), curr.Month()
		//Если совпадает год и месяц, то добавляем
		if searchedYear == year && searchedMonth == month {
			result = append(result, v)
		}
	}

	c.RUnlock()

	return result, nil
}

//Парсим json из запроса в event
func parseEvent(r *http.Request) (*Event, error) {
	var event Event
	decodeEvent := json.NewDecoder(r.Body)
	
	if err := decodeEvent.Decode(&event); err != nil {
		return nil, err
	}

	err := validateEvent(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func writeErrorResponse(w http.ResponseWriter, code int, err error) {
	response := struct {
		Err string `json:"error"`
	}{Err: err.Error()}

	json, _ := json.MarshalIndent(&response, "", "  ")
	w.WriteHeader(code)
	_, _ = w.Write(json)	
}

func writeOkResponse(w http.ResponseWriter, result []Event) {
	response := struct {
		Result []Event `json:"result"`
	}{result}

	json, _ := json.MarshalIndent(&response, "", "  ")
	_, err := w.Write(json)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
	}
}

//Получаем id пользователя и дату
func parseEventFromRequest(r *http.Request) (*Event, error) {
	e := Event{
		UserId:    r.URL.Query().Get("user_id"),
		Date:  r.URL.Query().Get("date"),
	}

	fmt.Println(e)
	return &e, nil
}

//Проверяем все поля на корректность
func validateEvent(event *Event) error {
	if event.Id == "" {
	    return errors.New("Id is required")
	}
	if event.UserId == "" {
	    return errors.New("UserId is required")
	}
	if event.Title == "" {
	    return errors.New("Title is required")
	}
	_, err := time.Parse("2006-01-02", event.Date)
	if err != nil  {
	    return errors.New("Date is required")
	}

	return nil
}