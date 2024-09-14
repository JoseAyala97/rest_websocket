package database

import (
	"context"
	"database/sql"
	"log"
	"rest_websocket/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// constructor - recibe como parametro la url (indica donde se realizara la conexion)
// retorna el repositorio y error en caso de que aplique
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	//se retorna el repositorio y la propiedad que se creo (db)
	return &PostgresRepository{db}, nil
}

// insertar usuario a la base de datos
// reserve func
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	//permite ejectura una sentencia sql
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)

	return err
}

// obtener usuario por id
// reserve func
func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	//procesar la informacion que se esta devolviendo - Go, fuertemente tipado

	//cerrar cuando se haya terminado de ejecutar
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user models.User
	//rows intentara leer la informacion que se esta devolviendo
	for rows.Next() {
		//posterior se estara mapeando la informacion que se esta devolviendo
		if err = rows.Scan(&user.ID, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

// cerrar conexion
// reserve func
func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
