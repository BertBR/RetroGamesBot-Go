package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/BertBR/RetroGamesBot-Go/pkg/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	queries *postgres.Queries
}

func New(pool *pgxpool.Pool) *Service {
	return &Service{
		queries: postgres.New(pool),
	}
}

func (svc *Service) GetTop10Console(ctx context.Context) (string, error) {
	sb := strings.Builder{}
	top10Console, err := svc.queries.GetTotalSortedByConsole(ctx)
	if err != nil {
		return "", err
	}
	for _, row := range top10Console {
		sb.WriteString(fmt.Sprintf("%s: %d\n", row.Console, row.Sum))
	}
	return sb.String(), nil
}

func (svc *Service) GetTop10Genre(ctx context.Context) (string, error) {
	sb := strings.Builder{}
	top10Console, err := svc.queries.GetTotalSortedByGenre(ctx)
	if err != nil {
		return "", err
	}
	for _, row := range top10Console {
		sb.WriteString(fmt.Sprintf("%s: %d\n", row.Genre, row.Sum))
	}
	return sb.String(), nil
}
