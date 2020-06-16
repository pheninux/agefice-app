package utils

type Utils struct {
}

func (u *Utils) CheckAuthentification(login string, pwd string) bool {
	return login == "adil" &&
		pwd == "a"

}
