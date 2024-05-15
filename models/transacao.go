package models

type Transacao struct {
    ID           int     `json:"id"`
    UsuarioID    int     `json:"usuario_id"`
    Destinatario int     `json:"destinatario"`
    Valor        float64 `json:"valor"`
    Data         string  `json:"data"`
}