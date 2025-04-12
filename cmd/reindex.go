package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trinhminhtriet/postcage/config"
	"github.com/trinhminhtriet/postcage/internal/logger"
	"github.com/trinhminhtriet/postcage/internal/storage"
)

// reindexCmd represents the reindex command
var reindexCmd = &cobra.Command{
	Use:   "reindex <database>",
	Short: "Reindex the database",
	Long: `This will reindex all messages in the entire database.

If you have several thousand messages in your mailbox, then it is advised to shut down
PostCage while you reindex as this process will likely result in database locking issues.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.Database = args[0]
		config.MaxMessages = 0

		if err := storage.InitDB(); err != nil {
			logger.Log().Error(err)
			os.Exit(1)
		}

		storage.ReindexAll()
	},
}

func init() {
	rootCmd.AddCommand(reindexCmd)
}
