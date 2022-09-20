package vo

type RecordRule struct {
	Record string `yaml:"record" json:"record" dc:"record name"`
	Expr   string `yaml:"expr" json:"expr" dc:"expression for prometheus metrics"`
}
