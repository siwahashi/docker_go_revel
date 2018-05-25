package controllers

import (
    "github.com/revel/revel"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {

    // コンソールログ出力例
	revel.AppLog.Debug(" DEBUG ")
	revel.AppLog.Info(" INFO ")
	revel.AppLog.Warn(" WARN ")
	revel.AppLog.Error(" ERROR ")

    return c.Render()
}