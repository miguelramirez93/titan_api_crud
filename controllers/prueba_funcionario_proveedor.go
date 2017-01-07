package controllers

import (
//	"encoding/json"
//	"errors"
	//"strconv"
	//"strings"
	"github.com/miguelramirez93/titan_api_crud/models"

	"github.com/astaxie/beego"
  "fmt"
)

// ConceptoController operations for Concepto
type PruebaFuncionarioProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *PruebaFuncionarioProveedorController) URLMapping() {
	c.Mapping("FuncionarioProveedor", c.ConsultarIDProveedor)

}

func (c *PruebaFuncionarioProveedorController) ConsultarIDProveedor() {

fmt.Println("aqui etoy")
l := models.GetIdProveedorXFuncionario()

	c.Data["json"] = l
	c.ServeJSON()
}
