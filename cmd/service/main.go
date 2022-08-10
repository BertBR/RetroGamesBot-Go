package service

import (
	"context"

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

func (svc *Service) GetTop10Console(ctx context.Context) ([]postgres.GetTotalSortedByConsoleRow, error) {
	top10Console, err := svc.queries.GetTotalSortedByConsole(ctx)
	if err != nil {
		return nil, err
	}
	return top10Console, nil
}

func (svc *Service) GetTop10Genre(ctx context.Context) ([]postgres.GetTotalSortedByGenreRow, error) {
	top10Console, err := svc.queries.GetTotalSortedByGenre(ctx)
	if err != nil {
		return nil, err
	}
	return top10Console, nil
}

func (svc *Service) GetTop10Games(ctx context.Context) ([]postgres.GetTop10GamesRow, error) {
	top10Games, err := svc.queries.GetTop10Games(ctx)
	if err != nil {
		return nil, err
	}
	return top10Games, nil
}
