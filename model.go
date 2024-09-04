package main

type ProductReq struct {
	Title string `json:"title"`
}

type ProductRes struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
