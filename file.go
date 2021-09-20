package env

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const configRegex = `^([A-Z][A-Z0-9_]+)([=]{1})([[\S ]*]?)$`

var rx = regexp.MustCompile(configRegex)

// LoadFromFile loads environment variables values from a given text file
// valid lines must comply with regex ^([A-Z][A-Z0-9_]+)([=]{1})([[\S ]*]?)$
// Examples of valid lines:
// ABC=prd
// XYZ=
// ABC="42378462%&&3 178964@"
// mnoPQR=42378462%&&3 ###
//
// Commented/ignored: #XYZ=4334343434 ( starts with # )
// invalid/ignored: opt=ler0ler0 ( has to be all caps/uppercase )
// Invalid/Ignored: _LETTERS=4334343434 ( has to start with a letter )
// Invalid/Ignored: X=4334343434 ( should contain 2 or more chars )
// Environment variables reference for curious: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap08.html
func LoadFromFile(configFile string, debugPrint bool) error {
	f, err := os.Open(configFile)

	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()

		if strings.TrimSpace(s) == "" || s[0] == '#' { // line is commented or empty
			continue
		}

		matches := rx.FindStringSubmatch(s)

		if len(matches) > 0 {
			k := matches[1]
			v := ""

			if len(matches) > 3 {
				v = matches[3]
			}

			if debugPrint {
				fmt.Println("setting", k, "env var")
			}

			if errx := os.Setenv(k, v); errx != nil {
				return errx
			}
		}
	}

	return nil
}
