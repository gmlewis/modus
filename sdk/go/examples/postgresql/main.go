/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"fmt"

	"github.com/gmlewis/modus/sdk/go/pkg/postgresql"
)

// The name of the PostgreSQL connection, as specified in the modus.json manifest
const connection = "my-database"

type Person struct {
	Id   int                  `json:"id"`
	Name string               `json:"name"`
	Age  int                  `json:"age"`
	Home *postgresql.Location `json:"home"`
}

func GetAllPeople() ([]Person, error) {
	const query = "select * from people order by id"
	response, err := postgresql.Query[Person](connection, query)
	return response.Rows, err
}

func GetPeopleByName(name string) ([]Person, error) {
	const query = "select * from people where name = $1"
	response, err := postgresql.Query[Person](connection, query, name)
	return response.Rows, err
}

func GetPerson(id int) (*Person, error) {
	const query = "select * from people where id = $1"
	response, err := postgresql.Query[Person](connection, query, id)
	if err != nil {
		return nil, err
	}

	if len(response.Rows) == 0 {
		return nil, nil // Person not found
	}

	return &response.Rows[0], nil
}

func AddPerson(name string, age int) (*Person, error) {
	const query = "insert into people (name, age) values ($1, $2) RETURNING id"

	response, err := postgresql.QueryScalar[int](connection, query, name, age)
	if err != nil {
		return nil, fmt.Errorf("failed to add person to database: %v", err)
	}

	p := Person{Id: response.Value, Name: name, Age: age}
	return &p, nil
}

func UpdatePersonHome(id int, longitude, latitude float64) (*Person, error) {
	const query = "update people set home = point($1, $2) where id = $3"

	response, err := postgresql.Execute(connection, query, longitude, latitude, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update person in database: %v", err)
	}

	if response.RowsAffected != 1 {
		return nil, fmt.Errorf("failed to update person with id %d - the record may not exist", id)
	}

	return GetPerson(id)
}

func DeletePerson(id int) (string, error) {
	const query = "delete from people where id = $1"

	response, err := postgresql.Execute(connection, query, id)
	if err != nil {
		return "", fmt.Errorf("failed to delete person from database: %v", err)
	}

	if response.RowsAffected != 1 {
		return "", fmt.Errorf("failed to delete person with id %d - the record may not exist", id)
	}

	return "success", nil
}
