package services

import (
	"fmt"
	"splitexpense/internal/models"
	"strings"
)

type WhatsAppService struct {
	expenseService *ExpenseService
	userService    *UserService
}

func NewWhatsAppService(expenseService *ExpenseService, userService *UserService) *WhatsAppService {
	return &WhatsAppService{
		expenseService: expenseService,
		userService:    userService,
	}
}

func (s *WhatsAppService) ProcessMessage(msg models.WhatsAppMessage) error {
	if len(msg.Entry) > 0 && len(msg.Entry[0].Changes) > 0 && len(msg.Entry[0].Changes[0].Value.Messages) > 0 {
		message := msg.Entry[0].Changes[0].Value.Messages[0]
		command := strings.ToLower(strings.TrimSpace(message.Text.Body))

		// Placeholder for command processing logic
		fmt.Printf("Processing command: %s\n", command)

		// Example of creating an expense
		if strings.HasPrefix(command, "expense") {
			// This is a very basic example. A real implementation would parse the command
			// to get the expense details.
			_, err := s.expenseService.CreateExpense(models.CreateExpenseRequest{
				Amount:      100,
				Currency:    "USD",
				Description: "Lunch",
				Shares: []models.CreateExpenseShareRequest{
					{
						ShareAmount: 50,
					},
				},
			})
			if err != nil {
				return fmt.Errorf("failed to create expense: %w", err)
			}
		}

		if strings.HasPrefix(command, "link") {
			parts := strings.Split(command, " ")
			if len(parts) == 2 {
				// In a real application, you would get the user ID from the linked account.
				// For this example, we'll just use a placeholder.
				userID := "some-user-id"
				if err := s.LinkUser(message.From, userID); err != nil {
					return fmt.Errorf("failed to link user: %w", err)
				}
			}
		}

		if strings.HasPrefix(command, "consent") {
			// In a real application, you would get the user ID from the linked account.
			// For this example, we'll just use a placeholder.
			userID := "some-user-id"
			if err := s.GiveConsent(userID); err != nil {
				return fmt.Errorf("failed to give consent: %w", err)
			}
		}
	}
	return nil
}

func (s *WhatsAppService) LinkUser(whatsappID, userID string) error {
	// In a real application, you would insert the mapping into the whatsapp_users table.
	fmt.Printf("Linking whatsapp user %s to splitexpense user %s\n", whatsappID, userID)
	return nil
}

func (s *WhatsAppService) GiveConsent(userID string) error {
	// In a real application, you would insert or update the user_consent table.
	fmt.Printf("User %s has given consent\n", userID)
	return nil
}
