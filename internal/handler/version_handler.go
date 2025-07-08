package handler

import (
	"fmt"
	"vex/internal/core"
	"vex/internal/infra"
	"vex/internal/service"
	"vex/internal/utils"
)

func HandleVersion(level, filePath, commitMessage string, noGit bool, manualVersion string) {
	data, err := infra.LoadJSON(filePath)
	if err != nil {
		utils.Error("Estamos tendo dificuldade ao abrir o arquivo: " + err.Error())
	}

	rawVersion, ok := data["version"].(string)
	if !ok {
		utils.Error("Campo 'version' não encontrado ou inválido no JSON.")
	}

	currentVersion, err := service.ParseVersion(rawVersion)
	if err != nil {
		utils.Error("Versão inválida: " + err.Error())
	}

	if level == "" && manualVersion == "" {
		fmt.Println(currentVersion.String())
		return
	}

	var newVersion *core.Version
	if level == "manual" {
		newVersion, err = service.ParseVersion(manualVersion)
		if err != nil {
			utils.Error("Versão manual inválida: " + err.Error())
		}
	} else {
		newVersion = service.BumpVersion(currentVersion, level)
		if err != nil {
			utils.Error("Erro ao incrementar a versão: " + err.Error())
		}
	}

	data["version"] = newVersion.String()

	// Salvar a nova versão no arquivo
	if err := infra.SaveJSON(filePath, data); err != nil {
		utils.Error("Erro ao salvar o JSON dentro do arquivo: " + err.Error())
	}

	utils.Info("Versão atualizada com sucesso para: " + newVersion.String())

	if !noGit {
		if commitMessage == "" {
			commitMessage = fmt.Sprintf("chore(version): new version %s - %s", level, newVersion.String())
		}

		if err := infra.GitCommitAndTag(filePath, newVersion.String(), commitMessage); err != nil {
			utils.Error("Erro ao realizar commit e tag no Git: " + err.Error())
		}
		utils.Info("Commit e tag realizados com sucesso no Git.")
	}
}
