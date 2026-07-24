package usecases

import (
	"context"
	"fmt"
	"time"

	customerRepo "github.com/cleidison-barradas/korvia.api/internal/customers/domain/repository"
	establishmentRepo "github.com/cleidison-barradas/korvia.api/internal/establishments/domain/repository"
	"github.com/cleidison-barradas/korvia.api/internal/messages/application/dto"
	messageService "github.com/cleidison-barradas/korvia.api/internal/messages/domain/service"
	serviceRepo "github.com/cleidison-barradas/korvia.api/internal/services/domain/repository"
	sessionDto "github.com/cleidison-barradas/korvia.api/internal/sessions/application/dto"
	sessionService "github.com/cleidison-barradas/korvia.api/internal/sessions/application/service"
	"github.com/cleidison-barradas/korvia.api/internal/sessions/domain/model"
	sessionModel "github.com/cleidison-barradas/korvia.api/internal/sessions/domain/model"
	sessionRepo "github.com/cleidison-barradas/korvia.api/internal/sessions/domain/repository"
)

type IncomingMessageUseCase struct {
	Ctx context.Context
	SessionRepo sessionRepo.SessionRepository
	CustomerRepo customerRepo.CustomerRepository
	EstablishmentRepo establishmentRepo.EstablishmentRepository
	WhatsappSender messageService.WhatsappSender
	ServiceRepo serviceRepo.ServiceRepository
}

func NewIncomingMessageUseCase(p *IncomingMessageUseCase) *IncomingMessageUseCase {
	return &IncomingMessageUseCase{
		Ctx: p.Ctx,
		ServiceRepo: p.ServiceRepo,
		SessionRepo: p.SessionRepo,
		CustomerRepo: p.CustomerRepo,
		WhatsappSender: p.WhatsappSender,
		EstablishmentRepo: p.EstablishmentRepo,
	}
}

func (uc *IncomingMessageUseCase) Execute(req dto.IncomingMessage) error {

	establishment, err := uc.EstablishmentRepo.GetByWabaID(req.WabaID)

	if err != nil {
		return err
	}

	if establishment == nil {
		return nil
	}

	services, err := uc.ServiceRepo.GetAll(establishment.Id)

	if err != nil {
		return err
	}

	session, err := uc.SessionRepo.GetSession(uc.Ctx, req.PhoneNumber)

	if err != nil {
		return err
	}

	if session == nil {
		session = sessionModel.NewSession(&model.NewSessionParams{
			PhoneNumber: req.PhoneNumber,
			Context: sessionDto.Context{
				CustomerName: req.Name,
				Services: services,
				Establishment: establishment,
			},
		})

		text, options, err := sessionService.Screen(session.State, session.Context)

		if err != nil {
			fmt.Println("Error render: ", err)
			return err
		}

		session.LastText = text
		session.LastOptions = options
	
		if err := uc.SessionRepo.SetSession(uc.Ctx, session); err != nil {
			return err
		}

		if err := uc.WhatsappSender.SendText(req.PhoneNumber, text); err != nil {
			return err
		}

		return nil
	}

	options, ok := session.LastOptions[req.Message]

	if !ok {
		fmt.Println("invalid entry")
		return nil
	}

	result := sessionService.Transition(*session, options.Event, options.Payload)

	if result.Invalid {
		text, _, err := sessionService.Screen(session.State,session.Context)

		if err != nil {
			fmt.Println("Error render: ", err)
			return err
		}

		if err := uc.WhatsappSender.SendText(req.PhoneNumber, text); err != nil {
			return err
		}
	}

	updatedAt := time.Now()
	session.State = result.State
	session.Context = result.Context
	session.History = result.History
	session.UpdatedAt = &updatedAt

	text, op, err := sessionService.Screen(session.State,session.Context)

	session.LastText = text
	session.LastOptions = op

	if err := uc.SessionRepo.SetSession(uc.Ctx, session); err != nil {
		return err
	}

	if err := uc.WhatsappSender.SendText(req.PhoneNumber, text); err != nil {
		return err
	}

	return nil
}