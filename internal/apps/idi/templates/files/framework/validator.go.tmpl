package lib

type Validator struct {
	errors ErrInvalidData
}

// New is a helper which creates a new Validator instance with an empty errors map.
func NewValidator() *Validator {
	return &Validator{errors: make(ErrInvalidData)}
}

// Valid returns true if the errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for
// the given key).
func (v *Validator) AddError(key, message string) {
	v.errors[key] = append(v.errors[key], message)
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(key string, notOK bool, message string) {
	if notOK {
		v.AddError(key, message)
	}
}

func (v *Validator) Errors() error {
	return v.errors
}
