package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/titan_api_crud/controllers:PruebaFuncionarioProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/titan_api_crud/controllers:PruebaFuncionarioProveedorController"],
		beego.ControllerComments{
			Method: "ConsultarIDProveedor",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
