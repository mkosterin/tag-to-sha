package models

type Arguments struct {
	SourceFileName string
}

func NewArguments() (*Arguments, error) {
	return &Arguments{SourceFileName: ""}, nil
}
