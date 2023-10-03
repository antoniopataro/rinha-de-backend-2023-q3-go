package people

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/antoniopataro/rinha-go/internal/infra/cache"
	"github.com/antoniopataro/rinha-go/internal/infra/database"
	"github.com/google/uuid"
)

func (r *Repository) Create(
	birthdate,
	name,
	nickname string,
	stack []string,
) (*Person, error) {
	person := Person{
		Birthdate: birthdate,
		ID:        uuid.New().String(),
		Name:      name,
		Nickname:  nickname,
		Stack:     stack,
	}

	if _, err := r.database.Client.Exec(context.Background(), `
		INSERT INTO people (birthdate, id, name, nickname, search, stack)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, person.Birthdate, person.ID, person.Name, person.Nickname, fmt.Sprintf("%s %s %s", person.Name, person.Nickname, strings.Join(person.Stack, " ")), strings.Join(person.Stack, ",")); err != nil {
		return nil, err
	}

	r.cache.Set(context.Background(), fmt.Sprintf("person:%s", person.ID), person)

	return &person, nil
}

func (r *Repository) Count() (int, error) {
	var count int

	if err := r.database.Client.QueryRow(context.Background(), `
		SELECT COUNT(*)
		FROM people
	`).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) Find(id string) (*Person, error) {
	var person Person

	var birthdate time.Time
	var stack string

	item, err := r.cache.Get(context.Background(), fmt.Sprintf("person:%s", id))

	if err == nil {
		if err := json.Unmarshal([]byte(item), &person); err != nil {
			return nil, err
		}

		return &person, nil
	}

	if err := r.database.Client.QueryRow(context.Background(), `
		SELECT birthdate, id, name, nickname, stack
		FROM people
		WHERE id = $1
	`, id).Scan(&birthdate, &person.ID, &person.Name, &person.Nickname, &stack); err != nil {
		return nil, err
	}

	if len(stack) != 0 {
		person.Stack = strings.Split(stack, ",")
	}

	person.Birthdate = birthdate.Format("2006-01-02")

	return &person, nil
}

func (r *Repository) Search(t string) ([]Person, error) {
	var people []Person

	// rows, err := r.database.Client.Query(context.Background(), `
	// 	SELECT birthdate, id, name, nickname, stack
	// 	FROM people
	// 	WHERE name LIKE '%' || $1 || '%' OR nickname LIKE '%' || $1 || '%' OR stack LIKE '%' || $1 || '%'
	// `, t)

	rows, err := r.database.Client.Query(context.Background(), `
		SELECT birthdate, id, name, nickname, stack
		FROM people
		WHERE search LIKE '%' || $1 || '%'
		LIMIT 50
	`, t)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person Person

		var birthdate time.Time
		var stack string

		if err := rows.Scan(&birthdate, &person.ID, &person.Name, &person.Nickname, &stack); err != nil {
			return nil, err
		}

		person.Birthdate = birthdate.Format("2006-01-02")

		if len(stack) != 0 {
			person.Stack = strings.Split(stack, ",")
		}

		people = append(people, person)
	}

	return people, nil
}

type Repository struct {
	cache    *cache.Cache
	database *database.Database
}

func MakeReposirory(cache *cache.Cache, database *database.Database) *Repository {
	return &Repository{
		cache:    cache,
		database: database,
	}
}
