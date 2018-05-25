package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
    "errors"
)

type User struct {
    *revel.Controller
}

func (c User) Index() revel.Result {

    // コンソールログ出力例
	revel.AppLog.Debug(" DEBUG ")
	revel.AppLog.Info(" INFO ")
	revel.AppLog.Warn(" WARN ")
	revel.AppLog.Error(" ERROR ")

    users := []models.User{}

    result := DB.Order("id desc").Find(&users);
    err := result.Error
    if err != nil {
        return c.RenderError(errors.New("Record Not Found"))
    }
    return c.Render(users)
}

func (c User) Create() revel.Result {
    user := models.User{
        Name: c.Params.Form.Get("name"),
    }
    ret := DB.Create(&user)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Create failure." + ret.Error.Error()))
    }
    return c.Redirect("/users")    
}

func (c User) Delete() revel.Result {
    id := c.Params.Route.Get("id")
    users := []models.User{}
    ret := DB.Delete(&users, id)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Delete failure." + ret.Error.Error()))
    }
    return c.Redirect("/users")    
}

func (c User) RedirectTousers() revel.Result {
    return c.Redirect("/users")    
}
