package models

type Saldo struct {
    UsuarioID int     `json:"usuario_id"`
    Saldo     float64 `json:"saldo"`
}