package controllers

import (
	"encoding/json"
	"github.com/miguelramirez93/titan_api_crud/models"
	"github.com/astaxie/beego"
	"fmt"
)

// ActaInicioController operations for ActaInicio
type InformacionContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionContratoController) URLMapping() {
	c.Mapping("Post", c.Post)
}
func (c *InformacionContratoController) Post() {
  var v models.Preliquidacion
  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
    if listaContratos, err := models.ListaContratos(&v); err == nil {
      c.Ctx.Output.SetStatus(201)
			fmt.Println("json : ", listaContratos[0])
      c.Data["json"] = listaContratos
    } else {
      c.Data["json"] = err.Error()
			fmt.Println("error : ", err)
    }
  } else {
    c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
  }
  c.ServeJSON()
  }
