package models

import (
  "fmt"

	"github.com/astaxie/beego/orm"
)

type Funcionario_x_Proveedor struct {
	Id                      int                `orm:"column(id_proveedor)"`
  NombreProveedor         string              `orm:"column(nom_proveedor)"`
  NumDocumento            float64            `orm:"column(contratista)"`
  NumeroContrato         string           `orm:"column(numero_contrato)"`

}


func init() {
	orm.RegisterModel(new(Funcionario_x_Proveedor))
}

// last inserted Id on success.
func GetIdProveedorXFuncionario() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()

  var temp []Funcionario_x_Proveedor;
	_, err := o.Raw("SELECT informacionproveedor.id_proveedor, informacionproveedor.nom_proveedor,contratos.contratista,contratos.numero_contrato FROM agora.informacion_proveedor AS informacionproveedor, argo.contrato_general AS contratos where contratos.objeto_contrato = 'Funcionario de planta' AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
  if err == nil{
      fmt.Println("Consulta exitosa")
  }

  return temp;
}

func GetIdProveedorXDocente() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()

	var temp []Funcionario_x_Proveedor
	_, err := o.Raw("SELECT informacionproveedor.id_proveedor, contratos.contratista,contratos.numero_contrato,informacionproveedor.nom_proveedor FROM agora.informacion_proveedor AS informacionproveedor, argo.contrato_general AS contratos where contratos.objeto_contrato = 'Docente de planta' AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp
}

func ListaContratos (v *Preliquidacion) ( datos []Funcionario_x_Proveedor , err error){
	o := orm.NewOrm()
  consulta := `select c.id_proveedor ,
								      c.nom_proveedor ,
								      b.contratista ,
								      b.numero_contrato
								      from argo.acta_inicio as a inner join argo.contrato_general as b on a.numero_contrato = b.numero_contrato
														   inner join agora.informacion_proveedor as c on b.contratista = c.num_documento
														   inner join argo.tipo_contrato on argo.tipo_contrato.id = b.tipo_contrato
														   where (argo.tipo_contrato.tipo_contrato = ?)  and ((? between a.fecha_inicio and a.fecha_fin) or ( ? between a.fecha_inicio and a.fecha_fin) ); `

	_,err = o.Raw(consulta, v.Nomina.TipoNomina.Nombre , v.FechaInicio, v.FechaFin).QueryRows(&datos)
	return
}
