package main

import (
	"context"
	"tk-robot/internal"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
}

// Call robot call
func (a *App) Call(name string, args internal.Params, config internal.Config) internal.Result {
	return internal.Call(name, args, config)
}

// GetConfig
func (a *App) GetConfig() map[string]any {
	return internal.GetConfig()
}

// GetConfig
func (a *App) Process() []map[string]any {
	return internal.Process()
}

// OpenFileDialog
func (a *App) OpenFileDialog() string {
	exePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{Title: "请选择程序"})
	if err != nil {
		return ""
	}
	return exePath
}

func (a *App) Alert(message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "提示",
		Message: message,
		Buttons: []string{"确定"},
	})
}
