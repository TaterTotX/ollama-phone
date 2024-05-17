package main

import (
	"context"
	"fmt"
)

var globalSettings map[string]string

func init() {
	initSettings() // 初始化数据库
	fmt.Println("初始化数据库成功")

}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	initSettings() // 初始化数据库
	fmt.Println("初始化数据库成功")
}

// Greet returns a greeting for the given name
func (a *App) Greet(input, role string) string {
	response, _ := ChatWithBot(input, role)
	return response
}

func (a *App) Get_settings() map[string]string {
	globalSettings = GetSetting()
	return globalSettings
}

func (a *App) UpdateSetting(data map[string]string) {
	err := UpdateSettings(data)

	if err != nil {
		fmt.Println("保存出错")
		return
	}
}
