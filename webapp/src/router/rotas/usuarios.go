package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                 "/criar-usuario",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPaginaDeCadastroUsuario,
		RequestAutenticacao: false,
	},
	{
		URI:                 "/usuario",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequestAutenticacao: false,
	},
	{
		URI:                 "/buscar-usuarios",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPaginaDeUsuarios,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPerfilDoUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.PararDeSeguir,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.Seguir,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/perfil",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPerfilUsuarioLogado,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/editar-usuario",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregarPaginaDeEdicaoDeUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/editar-usuario",
		Metodo:              http.MethodPut,
		Funcao:              controllers.EditarUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/atualizar-senha",
		Metodo:              http.MethodPost,
		Funcao:              controllers.AtualizarSenha,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/deletar-usuario",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequestAutenticacao: true,
	},
}
