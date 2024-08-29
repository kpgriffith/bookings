package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() error {
	// connect
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=kevin user=kevin password=")
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
		return err
	}
	defer db.Close()

	// test
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
		return err
	}
	log.Println("Pinged Database!")

	// get rows
	err = getRows(db, 0)
	if err != nil {
		log.Fatalf("Unable to get all rows from database: %v\n", err)
		return err
	}

	// insert a row
	query := `insert into public.users (first_name, last_name) values ($1, $2)`
	_, err = db.Exec(query, "brandt", "griffith")
	if err != nil {
		log.Fatalf("Unable to insert rows to database: %v\n", err)
		return err
	}

	// get rows
	err = getRows(db, 0)
	if err != nil {
		log.Fatalf("Unable to get all rows from database: %v\n", err)
		return err
	}

	// update a row
	update_query := `update public.users set first_name = $1 where id = $2`
	_, err = db.Exec(update_query, "k-love", 1)
	if err != nil {
		log.Fatalf("Unable to update rows to database: %v\n", err)
		return err
	}

	// get rows
	err = getRows(db, 0)
	if err != nil {
		log.Fatalf("Unable to get all rows from database: %v\n", err)
		return err
	}

	// get one row by id
	err = getRows(db, 3)
	if err != nil {
		log.Fatalf("Unable to get all rows from database: %v\n", err)
		return err
	}

	// delete a row
	delete_query := `delete from public.users where id = $1`
	_, err = db.Exec(delete_query, 2)
	if err != nil {
		log.Fatalf("Unable to delete rows from database: %v\n", err)
		return err
	}

	// get rows
	err = getRows(db, 0)
	if err != nil {
		log.Fatalf("Unable to get all rows from database: %v\n", err)
		return err
	}

	return nil
}

func getRows(conn *sql.DB, input int) error {

	var rows *sql.Rows
	var query string
	var err error

	if input > 0 {
		query = "select id, first_name, last_name from public.users where id=$1 order by id"
		rows, err = conn.Query(query, input)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		query = "select id, first_name, last_name from public.users order by id"
		rows, err = conn.Query(query)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	defer rows.Close()

	var fn, ln string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &fn, &ln)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("id=%d, first_name=%s, last_name=%s", id, fn, ln)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	log.Println("--------------------")

	return nil
}
