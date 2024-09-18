package user

var (
	qInsertUser = `
		INSERT INTO users(id, phone, name, password) 
		VALUES ($1, $2, $3, $4);`

	qGetUserByPhone = `
		SELECT
		    id,
		    phone,
		    name,
		    password,
		    created_at,
		    updated_at
		FROM users
		WHERE phone = $1;`
)
