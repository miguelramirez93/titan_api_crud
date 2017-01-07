package models


import (
	"time"
	"github.com/astaxie/beego/orm"

)

type InformacionContrato struct {
	Id             int       `orm:"column(id);pk"`
	NumeroContrato string    `orm:column(numero_contrato);null"`
	Nombre				 string    `orm:column(nombre);null"`
	NumDocumento	 float64    `orm:column(num_documento);null"`
	FechaInicio    time.Time `orm:"column(fecha_inicio);type(date);null"`
	FechaFin       time.Time `orm:"column(fecha_fin);type(date);null"`
  ValorContrato                float64             `orm:"column(valor_contrato)"`
}





func ListaContratos (v *Preliquidacion) ( datos []InformacionContrato , err error){
	o := orm.NewOrm()
  consulta := `select c.id_proveedor as id ,
								      c.nom_proveedor as nombre,
								      c.num_documento ,
								      b.numero_contrato,
								      a.fecha_inicio ,
								      a.fecha_fin ,
								       b.valor_contrato from argo.acta_inicio as a inner join argo.contrato_general as b on a.numero_contrato = b.numero_contrato
														   inner join agora.informacion_proveedor as c on b.contratista = c.num_documento
														   inner join argo.tipo_contrato on argo.tipo_contrato.id = b.tipo_contrato
														   where (argo.tipo_contrato.tipo_contrato = ?)  and ((? between a.fecha_inicio and a.fecha_fin) or ( ? between a.fecha_inicio and a.fecha_fin) ); `

	_,err = o.Raw(consulta, v.Nomina.TipoNomina.Nombre , v.FechaInicio, v.FechaFin).QueryRows(&datos)
	return
}
