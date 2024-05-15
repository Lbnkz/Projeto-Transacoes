package models

type Notificacao struct {
    UsuarioID int    `json:"usuario_id"`
    Mensagem  string `json:"mensagem"`
}