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

func (svc *Service) GetTotalGames(ctx context.Context) ([]int64, error) {
	totalGames, err := svc.queries.GetTotalGames(ctx)
	if err != nil {
		return nil, err
	}
	return totalGames, nil
}

func (svc *Service) GetTotalGamesByConsole(ctx context.Context) ([]postgres.GetTotalGamesByConsoleRow, error) {
	totalGamesByConsole, err := svc.queries.GetTotalGamesByConsole(ctx)
	if err != nil {
		return nil, err
	}
	return totalGamesByConsole, nil
}

func (svc *Service) GetThreeRandomGamesRow(ctx context.Context) ([]postgres.GetThreeRandomGamesRow, error) {
	threeRandomGames, err := svc.queries.GetThreeRandomGames(ctx)
	if err != nil {
		return nil, err
	}
	return threeRandomGames, nil
}

func (svc *Service) UpdateSortedGames(ctx context.Context, id []int32) error {
	for _, v := range id {
		err := svc.queries.UpdateSortedGame(ctx, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *Service) FakeSelect() error {
	ctx := context.Background()
	_, err := svc.queries.GetTotalGames(ctx)
	if err != nil {
		return err
	}
	return nil
}
