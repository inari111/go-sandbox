package models

import (
	"context"
)

func (u *User) GetByName(ctx context.Context, db DB, name string) ([]*User, error) {
	const sqlstr = `SELECT * FROM public.users WHERE name = $1`
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := User{
			_exists: true,
		}
		if err := rows.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, &u)
	}
	rerr := rows.Close()
	if rerr != nil {
		return nil, rerr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
