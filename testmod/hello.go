package testmod

import (
	"errors"
	"fmt"
)

func Hi(name string)string{
	return fmt.Sprintf("hi,%s",name)
}

func testV2(name string)string{
	return fmt.Sprintf("hi,%s",name)
}

// Hi returns a friendly greeting in language lang
func Hiv2(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}