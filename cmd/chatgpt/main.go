package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	apiKey := ""
	apiUrl := "https://api.openai.com/v1/completions"

	input := `Baseando nos livros 'habitos atomicos', 'a mente do empreendedor'. me responsa os proximos livros recomendados em formato json {"nome":""}`

	// Crie uma estrutura para a solicitação
	requestBody := map[string]interface{}{
		"model":      "text-davinci-002",
		"prompt":     input,
		"max_tokens": 3000,
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Erro ao criar o corpo da solicitação:", err)
		return
	}

	// Crie uma solicitação HTTP POST
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Erro ao criar a solicitação HTTP:", err)
		return
	}

	// Configure a autenticação com sua chave de API
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Faça a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Leia a resposta
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	// Exiba a resposta JSON
	fmt.Println(string(responseBody))

	var data map[string]interface{}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Fatalf("Erro ao analisar a resposta JSON: %v", err)
	}
	response := data["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
	fmt.Printf("Resposta: %s\n", response)
}
