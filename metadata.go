package jwt

import "github.com/project-flogo/core/data/coerce"

//Input data structure
type Input struct {
	Header string `md:"header,required"`
	Payload string `md:"payload,required"`
	Secret string `md:"secret,required"`
	Mode string `md:"mode,required"`
	Algorithm string `"md:algorithm,required"`
}

//ToMap Input mapper
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"header":    i.Header,
		"payload":   i.Payload,
		"secret":    i.Secret,
		"mode":      i.Mode,
		"algorithm": i.Algorithm,
	}
}

//FromMap Input from map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error

	i.Header, err = coerce.ToString(values["header"])
	if err != nil {
		return err
	}
	i.Payload, err = coerce.ToString(values["payload"])
	if err != nil {
		return err
	}
	i.Secret, err = coerce.ToString(values["secret"])
	if err != nil {
		return err
	}
	i.Mode, err = coerce.ToString(values["mode"])
	if err != nil {
		return err
	}
	i.Algorithm, err = coerce.ToString(values["algorithm"])
	if err != nil {
		return err
	}
	return nil
}

//Output data structure
type Output struct {
	Token  string `md:"token"`
}

//ToMap Output mapper
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"token":  o.Token,
	}
}

//FromMap Output  from map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Token, err = coerce.ToString(values["token"])
	if err != nil {
		return err
	}

	return nil
}
