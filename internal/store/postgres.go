package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // here
	"github.com/montagao/monplan/models"
)

const (
	INSTANCE_CONNECTION_NAME = "10.96.32.3"
	DATABASE_NAME            = "plan"
	DATABASE_USER            = "postgres"
	PASSWORD                 = "montamonta"
)

type PlanStore struct {
	db *sql.DB
}

func New() (*PlanStore, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		INSTANCE_CONNECTION_NAME,
		DATABASE_NAME,
		DATABASE_USER,
		PASSWORD)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	planStore := &PlanStore{
		db: db,
	}
	log.Printf("Initialized postgres DB: %s", DATABASE_NAME)
	return planStore, nil
}

func (s *PlanStore) GetByID(id int64) (*models.Plan, error) {
	rows, err := s.db.Query("select * from plans where id = $1;", id)
	if err != nil {
		return nil, err
	}
	if rows == nil || !rows.Next() {
		return nil, nil
	}
	var (
		nullID     int
		isComplete bool
		list1      sql.NullString
		list2      sql.NullString
		name1      sql.NullString
		name2      sql.NullString
		timestamp  sql.NullString
	)
	err = rows.Scan(&nullID, &isComplete, &list1, &list2, &name1, &name2, &timestamp)
	if err != nil {
		return nil, err
	}
	e := &models.Plan{
		ID:         &id,
		IsComplete: &isComplete,
		List1:      list1.String,
		List2:      list2.String,
		Name1:      name1.String,
		Name2:      name2.String,
		Timestamp:  timestamp.String,
	}
	return e, nil
}

func (s *PlanStore) GetAll(limit int) ([]*models.Plan, error) {
	log.Printf("Getting all plans, limit: %d\n", limit)
	result := []*models.Plan{}
	rows, err := s.db.Query("select * from plans LIMIT $1 ", limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id         int64
			isComplete bool
			list1      string
			list2      string
			name1      string
			name2      string
			timestamp  string
		)
		err := rows.Scan(&id, &isComplete, &list1, &list2, &name1, &name2, &timestamp)
		if err != nil {
			return nil, err
		}
		e := &models.Plan{
			ID:         &id,
			IsComplete: &isComplete,
			List1:      list1,
			List2:      list1,
			Name1:      list1,
			Name2:      list2,
			Timestamp:  timestamp,
		}
		result = append(result, e)
	}

	return result, nil
}

func (s *PlanStore) Put(plan *models.Plan) error {
	var err error
	fmt.Printf("adding new plan with id: %d", *plan.ID)
	// first insertion, insert person 1 info
	if len(plan.List1) != 0 {
		err = s.updateList1(plan)
	} else {
		err = s.updateList2(plan)
	}
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *PlanStore) Update(plan *models.Plan, id int) error {
	// TODO: not a prioity
	return nil
}

func (s *PlanStore) Delete(id int64) error {
	stmt, err := s.db.Prepare("DELETE from plans where id = $1;")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *PlanStore) updateList1(plan *models.Plan) error {
	stmt, err := s.db.Prepare("INSERT INTO plans(id, list1, name1, timestamp, is_complete) VALUES( $1, $2, $3, $4, $5 );")
	// Prepared statements take up server resources and should be closed after use.
	// defer stmt.Close()
	fmt.Printf("got plan %v", plan)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(plan.ID, plan.List1, plan.Name1, plan.Timestamp, plan.IsComplete)
	return err
}

func (s *PlanStore) updateList2(plan *models.Plan) error {
	stmt, err := s.db.Prepare("UPDATE plans SET list2 = $1, name2 = $2, timestamp = $3, is_complete = $4 VALUES( $1, $2, $3, $4 ) WHERE id = $5;")
	// Prepared statements take up server resources and should be closed after use.
	if err != nil {
		return err
	}
	_, err = stmt.Exec(plan.List2, plan.Name2, plan.Timestamp, plan.IsComplete, plan.ID)
	return err
}
