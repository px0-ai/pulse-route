package auth

type Validator struct {
	secret string
}

func NewValidator(secret string) *Validator {
	return &Validator{secret: secret}
}

func (v *Validator) Validate(token string) bool {
	return token == "Bearer "+v.secret
}
