package enderecos

import "strings"

func TypeAddress(address string) string {

	typesValid := []string{"rua", "avenida", "estrada", "rodovia"}
	lower := strings.ToLower(address)
	firstWordAddress := strings.Split(lower, " ")[0]

	valid := false

	for _, tipo := range typesValid {
		if tipo == firstWordAddress {
			valid = true
		}
	}

	if valid {
		return firstWordAddress
	}
	return "Endereco invalido"
}
