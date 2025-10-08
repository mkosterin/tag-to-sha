package models

type Arguments struct {
	SourceFileName string
	ConfigFileName string
}

func NewArguments() (*Arguments, error) {
	return &Arguments{
		SourceFileName: "",
		ConfigFileName: "",
	}, nil
}
