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

func (r *BabySitterRepository) GetBabySitters() {
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

}

func (r *BabySitterRepository) CreateBabySitter(babySitter request.BabysitterRequest) (int64, error) {
	result, err := r.db.DB.Exec("INSERT INTO babysitters (name, email, phone) VALUES ($1, $2, $3)", babySitter.Name, babySitter.Email, babySitter.Phone)

	if err != nil {
		log.Fatal("Error inserting babysitter:", err)
	}

	id, err := result.RowsAffected()

	if err != nil {
		log.Fatal("Error getting last insert id:", err)
	}

	fmt.Println("Inserted babysitter with ID:", id)
	return id, nil
}
