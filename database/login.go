package database

func LoginUser(email string, password string) (userID string, user_type string, err error) {
	loginQuery := "SELECT user_id,user_type from users where email=$1 and password=$2"
	//loginQuery
	err = Sqlite.QueryRow(loginQuery, email, password).Scan(&userID, &user_type)
	if err != nil {
		return
	}

	return
}
