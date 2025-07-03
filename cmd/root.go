/*
Copyright © 2025 João Ferreira <jferreira.ba@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vex",
	Short: "Controle de versão semântico simples e automatizado para aplicações",
	Long: `Vex é uma ferramenta de linha de comando escrita em Go que gerencia versões de forma leve, 
sem depender de sistemas de versionamento externos como Git, mas com integração opcional a ele.

Através de um arquivo JSON (por padrão: ./application.json), é possível incrementar automaticamente 
versões major, minor e patch, ou definir manualmente uma versão específica.

Também é possível realizar commits automáticos com mensagens personalizadas e gerar tags Git
para facilitar o controle de releases e a integração com pipelines CI/CD.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vex.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
