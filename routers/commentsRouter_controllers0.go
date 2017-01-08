package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["titan_api_crud/controllers:FuncionarioProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/titan_api_crud/controllers:FuncionarioProveedorController"],
		beego.ControllerComments{
			Method: "ConsultarIDProveedor",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})
		beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:PreliquidacionController"],
			beego.ControllerComments{
				Method: "Resumen",
				Router: `/resumen`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})
			beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["titan_api_crud/controllers:ConceptoPorPersonaController"],
				beego.ControllerComments{
					Method: "NovedadesActivas",
					Router: `novedades_activas/:id`,
					AllowHTTPMethods: []string{"post"},
					Params: nil})
}
