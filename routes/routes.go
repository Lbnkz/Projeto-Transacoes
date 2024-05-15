package routes

import (
	"github.com/labstack/echo/v4"

)

func Setup(e *echo.Echo) {

	e.GET("/usuarios/:id/saldo", getSaldoUsuario)
	e.GET("/usuarios/:id/transacoes", getTransacoaUsuario)
	e.GET("/transacoes/:id", getTransacaoId)

}
