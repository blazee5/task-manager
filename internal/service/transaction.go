package service

import (
	"context"
	"github.com/blazee5/finance-tracker/internal/domain"
	"github.com/blazee5/finance-tracker/internal/models"
)

func (s *Service) GetTransactions(ctx context.Context, id string) ([]models.Transaction, error) {
	return s.Storage.TransactionDAO.GetTransactions(ctx, id)
}

func (s *Service) CreateTransaction(ctx context.Context, userId string, transaction domain.Transaction) (string, error) {
	id, err := s.Storage.TransactionDAO.Create(context.Background(), transaction)

	if err != nil {
		return "", err
	}

	err = s.Storage.UserDAO.AddBalance(ctx, transaction.UserID, transaction.Amount)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	return s.Storage.TransactionDAO.GetTransaction(ctx, id)
}

func (s *Service) UpdateTransaction(ctx context.Context, id string, transaction domain.Transaction) (int, error) {
	return s.Storage.TransactionDAO.Update(ctx, id, transaction)
}

func (s *Service) DeleteTransaction(ctx context.Context, id string) error {
	return s.Storage.TransactionDAO.Delete(ctx, id)
}

func (s *Service) AnalyzeTransactions(ctx context.Context, id string) ([]models.Analyze, error) {
	return s.Storage.TransactionDAO.GetAnalyze(ctx, id)
}
