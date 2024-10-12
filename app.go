package main

import (
	"context"
	"fmt"
	"github.com/lxn/win"
	"syscall"
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

	initSettings() // 初始化数据库
	fmt.Println("初始化数据库成功")
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("ollama"))
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)

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
