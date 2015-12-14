package main

type Params struct {
	Application string `json:"app"`
	Deployment  string `json:"deployment"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Force       bool   `json:"force"`
	Commit      bool   `json:"commit"`
}
