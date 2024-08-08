package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/curtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CurtirPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/descurtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.DescurtirPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/atualizar",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPaginaDeAtualizacaoPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		RequestAutenticacao: true,
	},
}
