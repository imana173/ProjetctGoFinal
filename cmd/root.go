package cmd

import (
    "fmt"
    "contact-cli/internal/storer"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var Store storer.Storer

var rootCmd = &cobra.Command{
    Use:   "contact",
    Short: "Gestionnaire de contacts en CLI",
    Long:  "Outil CLI pour gérer des contacts en Go (JSON, GORM, Mémoire).",
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    // Chargement Viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("⚠️  Aucun fichier config trouvé, utilisation du storer mémoire.")
    }

    // Sélection du storer
    switch viper.GetString("storer") {
    case "json":
        Store = storer.NewJSONStorer(viper.GetString("json_path"))
    case "gorm":
        Store = storer.NewGORMStorer(viper.GetString("db_path"))
    default:
        fmt.Println("Mode mémoire activé.")
        Store = storer.NewMemoryStorer()
    }
}
