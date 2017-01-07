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
