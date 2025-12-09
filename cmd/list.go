package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "Liste tous les contacts",
    Run: func(cmd *cobra.Command, args []string) {
        contacts, _ := Store.List()

        if len(contacts) == 0 {
            fmt.Println("Aucun contact.")
            return
        }

        for _, c := range contacts {
            fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Email)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
