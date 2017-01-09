package controllers

import (
  	"encoding/json"
//	"errors"
	//"strconv"
	//"strings"
	"titan_api_crud/models"

	"github.com/astaxie/beego"
  "fmt"
)

// ConceptoController operations for Concepto
type FuncionarioProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioProveedorController) URLMapping() {
	c.Mapping("FuncionarioProveedor", c.ConsultarIDProveedor)

}


func (c *FuncionarioProveedorController) ConsultarIDProveedor() {
  var v models.Preliquidacion
  if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if (v.Nomina.TipoNomina.Nombre == "HC" || v.Nomina.TipoNomina.Nombre == "HC-SALARIOS" ){
			if listaContratos, err := models.ListaContratos(&v); err == nil {
	      c.Ctx.Output.SetStatus(201)

	      c.Data["json"] = listaContratos
	    } else {
	      c.Data["json"] = err.Error()
				fmt.Println("error : ", err)
	    }
		}else if (v.Nomina.TipoNomina.Nombre == "FP"){
			listaContratos := models.GetIdProveedorXFuncionario()
	      c.Ctx.Output.SetStatus(201)

	      c.Data["json"] = listaContratos

		}

  } else {
    c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
  }
  c.ServeJSON()
  }
