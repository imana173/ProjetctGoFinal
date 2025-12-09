package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "contact-cli/internal/models"

    "github.com/spf13/cobra"
)

var flagName string
var flagEmail string

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Ajoute un contact",
    Run: func(cmd *cobra.Command, args []string) {

        // Si flags fournis → création rapide
        if flagName != "" && flagEmail != "" {
            Store.Create(models.Contact{Name: flagName, Email: flagEmail})
            fmt.Println("Contact ajouté via flags !")
            return
        }

        // Mode interactif
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("Nom : ")
        name, _ := reader.ReadString('\n')

        fmt.Print("Email : ")
        email, _ := reader.ReadString('\n')

        Store.Create(models.Contact{
            Name:  strings.TrimSpace(name),
            Email: strings.TrimSpace(email),
        })

        fmt.Println("Contact ajouté !")
    },
}

func init() {
    rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringVarP(&flagName, "name", "n", "", "Nom du contact")
    addCmd.Flags().StringVarP(&flagEmail, "email", "e", "", "Email du contact")
}
