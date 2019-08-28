package models

type VideoInfo struct {
    Id           int `json:"id"`
    Name         string `json:"name"`
    ImageUrl     string `json:"image_url"`
    PresentPrice float64 `json:"present_price"`
    BuyStatus    string `json:"buy_status"`
}

