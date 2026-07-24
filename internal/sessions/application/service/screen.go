package service

import (
	"fmt"
	"strings"

	"github.com/cleidison-barradas/korvia.api/internal/sessions/application/dto"
)

type RenderFunc func(ctx dto.Context) (text string, options map[string]dto.Option, err error)

var screens = map[dto.State]RenderFunc{
	dto.STATE_ONBOARDING_NAME: func(ctx dto.Context) (string, map[string]dto.Option, error) {
		text := "Olá Bem vindo! \nantes de começar, qual o seu nome?"

		return text, nil, nil
	},

	dto.STATE_MENU: func(ctx dto.Context) (string, map[string]dto.Option, error) {
		text := fmt.Sprintf("Olá %s! 👋\nBem-vindo à %s. \nComo posso ajudar?\n\n1️⃣ Agendar horário\n2️⃣ Meus agendamentos", ctx.CustomerName, ctx.Establishment.Name)
		
		options := map[string]dto.Option{
			"1": {
				Event: dto.EVENT_SCHEDULE_APPOINTMENT,
				Payload: dto.Payload{
					Value: "1",
				},
			},
			"2": {
				Event: dto.EVENT_APPOINTMENT_LIST,
				Payload: dto.Payload{
					Value: "1",
				},
			},
		}

		return text, options, nil
	},

	dto.STATE_CHOOSE_SERVICE: func (ctx dto.Context) (string, map[string]dto.Option, error) {
		var text strings.Builder
		var options map[string]dto.Option

		text.WriteString("Selecione um serviço:\n\n")

		for i, service := range ctx.Services {
			fmt.Fprintf(&text, "%d. %s\n", i+1, service.Name)

			options = map[string]dto.Option{
				fmt.Sprintf("%d", i+1): {
					Event: dto.EVENT_SELECTED_SERVICE,
					Payload: dto.Payload{
						Value: service.Id.String(),
					},
				},
			}
		}

		return text.String(), options, nil
	},

	dto.STATE_HOW_TO_SCHEDULE: func (ctx dto.Context) (string, map[string]dto.Option, error) {
		text := "Como gostaria de agendar?\n\n1️⃣ Primeiro horário disponível\n 2️⃣ Escolher um horário\n0️⃣ Cancelar"

		options := map[string]dto.Option{
			"0": {
				Event: dto.EVENT_CANCEL,
			},
			"1": {
				Event: dto.EVENT_SCHEDULE_FIRST_TIME_AVAILABLE,
			},
			"2": {
				Event: dto.EVENT_CHOOSE_TIME,
			},
		}

		return text, options, nil
	},
	dto.STATE_FIRST_TIME_AVAILABLE: func(ctx dto.Context) (string, map[string]dto.Option, error) {
		text := "Encontrei o primeiro horário disponível:\n\n"

		options := map[string]dto.Option{
			"0": {
				Event: dto.EVENT_CANCEL,
			},
		}

		return text, options, nil
	},
	dto.STATE_CHOOSE_TIME: func(ctx dto.Context) (string, map[string]dto.Option, error) {
		text := "Escolha um horário:\n\n"

		options := map[string]dto.Option{
			"0": {
				Event: dto.EVENT_CANCEL,
				Payload: dto.Payload{
					Value: "1",
				},
			},
		}

		return text, options, nil
	},

	dto.STATE_APPOINTMENT_CONFIRMATION: func(ctx dto.Context) (string, map[string]dto.Option, error) {
		text := "Confira seu agendamento:\n\n"

		text += fmt.Sprintf("Serviço: %s\nData: %s\nHorário: %s", ctx.Day, ctx.Date, ctx.Time)

		text += "Confirma esse agendamento?\n\n1️⃣ Confirmar\n0️⃣ Cancelar"

		options := map[string]dto.Option{
			"0": {
				Event: dto.EVENT_CANCEL,
				Payload: dto.Payload{
					Value: "1",
				},
			},
			"1": {
				Event: dto.EVENT_CONFIRM_APPOINTMENT,
				Payload: dto.Payload{
					Value: "1",
				},
			},
		}

		return text, options, nil
	},
}

func Screen(state dto.State, ctx dto.Context) (string, map[string]dto.Option, error) {
	text, options, err := screens[state](ctx)

	if err != nil {
		return "", nil, err
	}

	return text, options, nil
}

