package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/BertBR/RetroGamesBot-Go/pkg/database"
)

func handleTopGames(sender string) string {
	numbers := [10]string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}

	var sb strings.Builder
	snap := database.LoadGames(context.Background())
	var g database.Game
	for i, v := range snap {
		err := v.DataTo(&g)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(&sb, "%s [%s](%s) - %d\n", numbers[i], g.Title, g.FileURL, g.Sorted)
	}

	return fmt.Sprintf("Olá %s\nAqui está a lista dos TOP 10 games mais sorteados!\n\n%s", sender, sb.String())
}
