package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func ValidationMsg(err error) []string {
	var errors []string
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, e := range v {
			field := strings.ToLower(e.Field())
			message := fmt.Sprintf("Field '%s' failed validation for tag '%s %s'", field, e.Tag(), e.Param())
			errors = append(errors, message)
		}
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Lower(v any) any {
	switch v := v.(type) {
	case []any:
		lv := make([]any, len(v))
		for i := range v {
			lv[i] = Lower(v[i])
		}
		return lv
	case map[string]any:
		lv := make(map[string]any, len(v))
		for mk, mv := range v {
			lv[strings.ToLower(mk)] = mv
		}
		return lv
	default:
		return v
	}
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func Decode(s string) string {
	data, _ := base64.StdEncoding.DecodeString(s)
	return string(data)
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
