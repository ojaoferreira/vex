/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"vex/internal/handler"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version [major|minor|patch|manual <version>]",
	Short: "Atualiza a versão da aplicação baseada em um arquivo JSON",
	Long: `O comando 'version' permite incrementar ou definir manualmente a versão de uma aplicação 
no formato semântico (major.minor.patch), utilizando um arquivo JSON como base.

Você pode especificar o tipo de incremento desejado (major, minor, patch) ou definir 
uma versão estática específica (ex: manual 2.5.1). O arquivo de versão pode ser passado 
via flag ou detectado automaticamente na raiz do projeto (./application.json por padrão).

O comando também permite realizar commit e tag automáticos via Git, com suporte a mensagens 
customizadas e opção para desabilitar essa integração.

Exemplos:

  vex version patch
  vex version major -m "chore(ci): new version major"
  vex version manual 3.2.1 --file ./config/ver.json
  vex version minor --no-git-tag-version`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			handler.HandleVersion("invalid", "", "", false, "")
			return
		}

		level := args[0]
		manualVersion := ""
		if level == "manual" {
			if len(args) < 2 {
				handler.HandleVersion("manual-missing", "", "", false, "")
				return
			}
			manualVersion = args[1]
		}

		filePath, _ := cmd.Flags().GetString("file")
		commitMsg, _ := cmd.Flags().GetString("commit")
		noGit, _ := cmd.Flags().GetBool("no-git-tag-version")

		handler.HandleVersion(level, filePath, commitMsg, noGit, manualVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Define as flags para o comando version
	versionCmd.Flags().StringP("file", "f", "./application.json", "Caminho do arquivo JSON de versão (Padrão: application.json)")
	versionCmd.Flags().StringP("commit", "m", "", "Mensagem do commit.")
	versionCmd.Flags().BoolP("no-git-tag-version", "", false, "Impede que seja feito commit e tag da versão (modo somente local)")
}
