package testpackage

import "fmt"

func PackageRemote(name string) string {
	return fmt.Sprintf("Hello %s", name)
}