package helpers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ExecuteCmd(expr string) {
	parts := strings.Split(expr, " ")
	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func ReadOrCreateFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.WriteFile(path, []byte("{}"), 0644); err != nil {
				return nil, err
			}

			return []byte("{}"), nil
		}

		return nil, err
	}

	return os.ReadFile(path)
}
