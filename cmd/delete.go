package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var deleteID uint

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Supprime un contact",
    Run: func(cmd *cobra.Command, args []string) {
        if deleteID == 0 {
            fmt.Println("⚠️  --id obligatoire")
            return
        }

        Store.Delete(deleteID)
        fmt.Println("Contact supprimé.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
    deleteCmd.Flags().UintVar(&deleteID, "id", 0, "ID du contact")
}
