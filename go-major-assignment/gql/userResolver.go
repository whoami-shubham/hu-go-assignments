package gql

import (
	"log"

	"github.com/graphql-go/graphql"
)

//GetAllUsers ...
func (resolver *QueryResolver) GetAllUsers(p graphql.ResolveParams) (interface{}, error) {

	data := []DbUser{}
	rows, err := db.Query("select * from users;")
	if err != nil {
		log.Println(err)
		return []DbUser{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user DbUser
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Role, &user.CreatedOn, &user.UpdatedOn)
		if err != nil {
			log.Println(err)
			return []DbUser{}, err
		}
		data = append(data, user)

	}

	return data, nil

}
