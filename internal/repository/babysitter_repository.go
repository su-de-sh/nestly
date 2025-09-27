package repository

import (
	"fmt"
	"log"

	"github.com/su-de-sh/nestly/internal/api/request"
	"github.com/su-de-sh/nestly/internal/api/response"
	"github.com/su-de-sh/nestly/internal/database"
)

type BabySitterRepository struct {
	db *database.PostgresDB
}

func NewBabySitterRepository(db *database.PostgresDB) *BabySitterRepository {
	return &BabySitterRepository{db: db}
}

func (r *BabySitterRepository) GetBabySitters() []response.BabysitterResponse {
	rows, err := r.db.DB.Query("SELECT id, name, email, phone FROM babysitters")

	if err != nil {
		log.Fatal("Error fetching babysitters:", err)
	}
	defer rows.Close()

	var babySitters []response.BabysitterResponse

	for rows.Next() {
		var res response.BabysitterResponse
		err = rows.Scan(&res.ID, &res.Name, &res.Email, &res.Phone)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}

		babySitters = append(babySitters, res)

	}
	fmt.Println("Babysitters:", babySitters)

	return babySitters

}

func (r *BabySitterRepository) CreateBabySitter(babySitter request.BabysitterRequest) (int64, error) {
	result, err := r.db.DB.Exec("INSERT INTO babysitters (name, email, phone) VALUES ($1, $2, $3)", babySitter.Name, babySitter.Email, babySitter.Phone)

	if err != nil {

		fmt.Println("Error inserting babysitter:", err)
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Error getting last insert id:", err)
	}

	fmt.Println("Inserted babysitter with ID:", id)
	return id, nil
}

func (r *BabySitterRepository) UpdateById(id string, babysitter request.BabysitterRequest) error {
	_, err := r.db.DB.Exec("UPDATE babysitters SET name=$1, email=$2, phone=$3 WHERE id=$4", babysitter.Name, babysitter.Email, babysitter.Phone, id)

	if err != nil {
		log.Fatal("Error updating babysitter:", err)
		return err
	}
	return nil
}

func (r *BabySitterRepository) DeleteById(id string) error {
	_, err := r.db.DB.Exec("DELETE FROM babysitters WHERE id=$1", id)

	if err != nil {
		log.Fatal("Error deleting babysitter:", err)
		return err
	}
	return nil
}
