package rotas

import "webapp/src/controllers"

var rotaPaginaPrincipal = Rota{
	URI:                 "/home",
	Metodo:              "GET",
	Funcao:              controllers.CarregarPaginaPrincipal,
	RequestAutenticacao: true,
}
