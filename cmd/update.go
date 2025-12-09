package cmd

import (
    "fmt"
    "contact-cli/internal/models"

    "github.com/spf13/cobra"
)

var updateID uint
var updateName string
var updateEmail string

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Modifie un contact",
    Run: func(cmd *cobra.Command, args []string) {
        if updateID == 0 {
            fmt.Println("⚠️  --id obligatoire")
            return
        }

        updated := models.Contact{
            ID:    updateID,
            Name:  updateName,
            Email: updateEmail,
        }

        Store.Update(updated)
        fmt.Println("Contact mis à jour.")
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)

    updateCmd.Flags().UintVar(&updateID, "id", 0, "ID du contact (obligatoire)")
    updateCmd.Flags().StringVar(&updateName, "name", "", "Nouveau nom")
    updateCmd.Flags().StringVar(&updateEmail, "email", "", "Nouvel email")
}
