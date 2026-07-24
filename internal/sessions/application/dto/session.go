package dto

import (
	establishmentModel "github.com/cleidison-barradas/korvia.api/internal/establishments/domain/model"
	serviceModel "github.com/cleidison-barradas/korvia.api/internal/services/domain/model"
)

type State string

type Event string

const (
	STATE_INITIAL State = "initial"
	STATE_ONBOARDING_NAME State = "onboarding_name"
	STATE_MENU State = "menu"
	STATE_CHOOSE_SERVICE State = "choose_service"
	STATE_HOW_TO_SCHEDULE State = "how_to_schedule"
	STATE_FIRST_TIME_AVAILABLE State = "first_time_available"
	STATE_CHOOSE_TIME State = "choose_time"
	STATE_SCHEDULED State = "scheduled"
	STATE_APPOINTMENT_CONFIRMATION State = "appointment_confirmation"
)

const (
	EVENT_BACK Event = "back"
	EVENT_CANCEL Event = "cancel"
	EVENT_CHOOSE_TIME Event = "choose_time"
	EVENT_SCHEDULE_SERVICE Event = "schedule_service"
	EVENT_APPOINTMENT_LIST Event = "appointment_list"
	EVENT_SCHEDULE_APPOINTMENT Event = "schedule_appointment"
	EVENT_CONFIRM_APPOINTMENT Event = "confirm_appointment"
	EVENT_SELECTED_SERVICE Event = "selected_service"
	EVENT_SCHEDULE_FIRST_TIME_AVAILABLE = "schedule_first_time_available"
)

type Context struct {
	Day string
	Time string
	Date string
	ServiceID string
	OutputText string
	PhoneNumber string
	CustomerName string
	Professional string
	Services []serviceModel.Services
	Establishment *establishmentModel.Establishment 
}

type Payload struct {
	Value string `json:"value"`
}

type Option struct {
	Event Event `json:"event"`
	Payload Payload `json:"payload"`
}

