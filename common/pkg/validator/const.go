package validator

// validator tag
const (
	vPort     = "vPort"     // validate http port
	gtField   = "gtField"   // gt field data
	uppercase = "uppercase" // validate field uppercase
)

// validations custom validation
var validations = []validData{
	{Tag: vPort, ValidFunc: ValidPortFn, EvenIfNull: true},
	{Tag: gtField, ValidFunc: ValidGtFieldFn, EvenIfNull: true},
	{Tag: uppercase, ValidFunc: ValidIsUpperCaseFn, EvenIfNull: true},
}

// translations custom translation
var translations = []translData{
	{Tag: vPort, Msg: "{0}范围为1-65535，请重新输入"},
	{Tag: gtField, Msg: "{0}必须大于{1}", TranslFn: translFiledTestFn},
	{Tag: uppercase, Msg: "{0}必须大于{1}"},

	// recover translate
	{Tag: "requited", Msg: "{0}必填字段不能为空"},
}
