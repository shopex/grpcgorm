package model

import (
	"github.com/jinzhu/gorm"
)

//grom model http://doc.gorm.io/models.html#model-definition
type Invoice struct {
	gorm.Model
	Mode               int32  `json:"Mode" name:"发票类型" db:"mode" dbtype:TINYINT(4)`
	Type               int32  `json:"Type" name:"抬头类型" db:"type" dbtype:TINYINT(4)`
	Content            string `json:"Content" sql:"DEFAULT:null" name:"发票抬头" db:"content" dbtype:VARCHAR(128)`
	CompanyName        string `json:"CompanyName" sql:"DEFAULT:null" name:"公司名称" db:"company_name" dbtype:VARCHAR(128)`
	CompanyAddress     string `json:"CompanyAddress" sql:"DEFAULT:null" name:"公司地址" db:"company_address" dbtype:VARCHAR(128)`
	RegistrationNumber string `json:"RegistrationNumber" sql:"DEFAULT:null" name:"纳税人识别号" db:"registration_number" dbtype:VARCHAR(128)`
	Bankaccount        string `json:"Bankaccount" sql:"DEFAULT:null" name:"开户银行帐号" db:"bankaccount" dbtype:VARCHAR(128)`
	CompanyPhone       string `json:"CompanyPhone" sql:"DEFAULT:null" name:"公司电话" db:"company_phone" dbtype:VARCHAR(128)`
	TradeId            int64  `json:"TradeId" name:"NULL" db:"trade_id" dbtype:INT gorm:"index"`
}
