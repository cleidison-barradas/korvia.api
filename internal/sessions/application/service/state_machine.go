package service

import (
	"github.com/cleidison-barradas/korvia.api/internal/sessions/application/dto"
	"github.com/cleidison-barradas/korvia.api/internal/sessions/domain/model"
)

type TargetFunc func(ctx dto.Context, payload dto.Payload) dto.State
type Action func(ctx dto.Context, payload dto.Payload) dto.Context

type Rule struct {
	Invalid bool
	Action Action
	Target dto.State
	TargetFunc TargetFunc
}

type StateDef struct {
	On map[dto.Event]Rule
}

type TransitionResult struct {
	State dto.State
	Context dto.Context
	History []dto.State
	Invalid bool
}

var transitions = map[dto.State]StateDef{
	dto.STATE_MENU: {
		On: map[dto.Event]Rule{
			dto.EVENT_SCHEDULE_APPOINTMENT: {
				Target: dto.STATE_CHOOSE_SERVICE,
			},
			dto.EVENT_APPOINTMENT_LIST: {
				Target: dto.STATE_MENU,
			},
		},
	},

	dto.STATE_CHOOSE_SERVICE: {
		On: map[dto.Event]Rule{
			dto.EVENT_SELECTED_SERVICE: {
				Target: dto.STATE_HOW_TO_SCHEDULE,
				Action: func(ctx dto.Context, payload dto.Payload) dto.Context {
					ctx.ServiceID = payload.Value

					return ctx
				},
			},
			dto.EVENT_CANCEL: {
				Target: dto.STATE_MENU,
			},
		},
	},

	dto.STATE_HOW_TO_SCHEDULE: {
		On: map[dto.Event]Rule{
			dto.EVENT_SCHEDULE_FIRST_TIME_AVAILABLE: {
				Target: dto.STATE_FIRST_TIME_AVAILABLE,
			},
			dto.EVENT_CHOOSE_TIME: {
				Target: dto.STATE_CHOOSE_TIME,
			},
			dto.EVENT_CANCEL: {
				Target: dto.STATE_MENU,
			},
		},
	},

	dto.STATE_FIRST_TIME_AVAILABLE: {
		On: map[dto.Event]Rule{
			dto.EVENT_SCHEDULE_FIRST_TIME_AVAILABLE: {
				Target: dto.STATE_APPOINTMENT_CONFIRMATION,
			},
			dto.EVENT_CANCEL: {
				Target: dto.STATE_MENU,
			},
		},
	},

	dto.STATE_CHOOSE_TIME: {
		On: map[dto.Event]Rule{
			dto.EVENT_CANCEL: {
				Target: dto.STATE_MENU,
				Action: func (ctx dto.Context, payload dto.Payload) dto.Context {
					ctx.ServiceID = payload.Value

					return ctx
				},
			},
		},
	},

	dto.STATE_APPOINTMENT_CONFIRMATION: {
		On: map[dto.Event]Rule{
			dto.EVENT_CONFIRM_APPOINTMENT: {
				Target: dto.STATE_MENU,
			},
			dto.EVENT_CANCEL: {
				Target: dto.STATE_MENU,
			},
		},
	},
}

func Transition(session model.Session, event dto.Event, payload dto.Payload) TransitionResult {
	stateDef, ok := transitions[session.State]

	if event == dto.EVENT_CANCEL {
		return TransitionResult{
			State: dto.STATE_MENU,
			Context: dto.Context{},
			History: []dto.State{},
		}
	}

	if event == dto.EVENT_BACK {
		if len(session.History) == 0 {
			return TransitionResult{
				State: session.State,
				Context: session.Context,
				History: session.History,
			}
		}

		session.History = session.History[:len(session.History) - 1]

		return TransitionResult{
			State: session.History[len(session.History) - 1],
			Context: session.Context,
			History: session.History,
		}
	}

	if !ok {
		return TransitionResult{
			State: session.State,
			Context: session.Context,
			History: session.History,
			Invalid: true,
		}
	}

	rule, ok := stateDef.On[event]

	if !ok || rule.Invalid {
		return TransitionResult{
			State: session.State,
			Context: session.Context,
			History: session.History,
			Invalid: true,
		}
	}

	target := rule.Target

	if rule.TargetFunc != nil {
		target = rule.TargetFunc(session.Context, payload)
	}

	ctx := session.Context

	if rule.Action != nil {
		ctx = rule.Action(ctx, payload)
	}

	updateHistory := append(append([]dto.State{}, session.History...), session.State)

	return TransitionResult{
		Context: ctx,
		State: target,
		History: updateHistory,
	}
}