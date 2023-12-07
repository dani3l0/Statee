package utils

import "os/user"

func IsSuperuser() bool {
	name, _ := user.Current()
	return name.Username == "root"
}
