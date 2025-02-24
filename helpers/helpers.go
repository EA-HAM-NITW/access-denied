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

func GetCmdOutput(expr string) string {
	cmd := exec.Command("bash", "-c", expr)

	outputBytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := strings.TrimSpace(string(outputBytes))
	return output
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

func CheckForCmdKeywords(expr string, keywords ...string) bool {
	for _, k := range keywords {
		cmd := strings.Split(expr, " ")[0]
		if !strings.Contains(cmd, k) {
			return false
		}
	}

	return true
}

func SliceToFormattedString(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	if len(slice) == 1 {
		return "`" + slice[0] + "`"
	}

	var result strings.Builder

	for i := 0; i < len(slice)-1; i++ {
		result.WriteString("`" + slice[i] + "`")

		if i < len(slice)-2 {
			result.WriteString(", ")
		} else {
			result.WriteString(" or ")
		}
	}

	result.WriteString("`" + slice[len(slice)-1] + "`")
	return result.String()
}
