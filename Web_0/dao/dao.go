package dao

var database = make(map[string]string)

func AddUser(username, password string) {
	database[username] = password
}

func FindUser(username, password string) bool {
	if pwd, ok := database[username]; ok {
		return pwd == password
	}
	return false
}

func RevisePassword(username, old, new string) bool {
	if pwd, ok := database[username]; ok && pwd == old {
		database[username] = new
		return true
	}
	return false
}
