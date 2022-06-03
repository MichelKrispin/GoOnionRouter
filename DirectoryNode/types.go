package main

type connection struct {
	Address string `json:"address" binding:"required"`
	Success bool   `json:"success" binding:"required"`
}

type node struct {
	Address   string `json:"address"`
	PublicKey string `json:"publickey"`
}
