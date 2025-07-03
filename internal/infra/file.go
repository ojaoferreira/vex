package infra

import (
	"encoding/json"
	"os"
)

func LoadJSON(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		// logger.Error("Estamos tendo dificuldade ao abrir o arquivo: " + err.Error())
		return nil, err
	}
	defer file.Close()

	var data map[string]interface{}
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		// logger.Error("Eita! Estamos com dificuldade em decodificar o arquivo para JSON: " + err.Error())
		return nil, err
	}

	return data, nil
}

func SaveJSON(path string, data map[string]interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		// logger.Error("Ops! Estamos com dificuldade ao abrir o arquivo para escrita: " + err.Error())
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)
	encode.SetIndent("", "  ")
	return encode.Encode(data)
}
