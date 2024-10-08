package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Printf("Rodando o WebApp!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
