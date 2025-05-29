package database

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/kahbum/eda_balance/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table accounts (id varchar(255), balance int, updated_at date)")

	s.accountDB = NewAccountDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP table accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(uuid.New().String())
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	account := entity.NewAccount(uuid.New().String())
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
}

func (s *AccountDBTestSuite) TestUpdateBalance() {
	account := entity.NewAccount(uuid.New().String())
	err := s.accountDB.Save(account)
	s.Nil(err)

	account.UpdateBalance(100)
	err = s.accountDB.UpdateBalance(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Balance, accountDB.Balance)
}
