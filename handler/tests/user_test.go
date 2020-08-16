package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/model"
	"gitlab.com/pplayground/pet_tracking/user-service/handler/repository"
	"reflect"
	"testing"
)

func TestShouldCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO users").WithArgs(
		"test_username",
		"test_password",
		"test_firstName",
		"test_lastName",
		"test@mail.com",
		"2006-01-02 15:04:05").WillReturnResult(sqlmock.NewResult(1, 1))

	user := model.User{
		Username: "test_username",
		Password: "test_password",
		FirstName: "test_firstName",
		LastName: "test_lastName",
		Email: "test@mail.com",
		Role: "user",
		CreatedAt: "2006-01-02 15:04:05",
	}

	// now we execute our method
	if err = repository.CreateUser(db, &user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldShowAllUserInformation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "username", "firstName", "lastName", "email", "role"}).
		AddRow(1, "test_username1", "test_firstName1", "test_lastName1", "test@mail.com", "user").
		AddRow(1, "test_username2", "test_firstName2", "test_lastName2", "test@mail.com", "user")
	mock.ExpectQuery("SELECT id, username, firstName, lastName, email, role FROM users").WillReturnRows(rows)

	// now we execute our method
	if _, err = repository.GetUsers(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldShowUserInformation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "username", "firstName", "lastName", "email", "role"}).
		AddRow(1, "test_username", "test_firstName", "test_lastName", "test@mail.com", "user")
	mock.ExpectQuery("SELECT id, username, firstName, lastName, email, role FROM users").WillReturnRows(rows)

	// now we execute our method
	results, err := repository.GetUser(db, 1)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}


	expectedResult := make(map[string]interface{})
	expectedResult["id"] = 1
	expectedResult["username"] = "test_username"
	expectedResult["firstName"] = "test_firstName"
	expectedResult["lastName"] = "test_lastName"
	expectedResult["email"] = "test@mail.com"
	expectedResult["role"] = "user"

	// check test result is match with expectedResult
	if !reflect.DeepEqual(results, expectedResult) {
		t.Errorf("query not match")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE users").WithArgs(
		"test_username",
		"test_password",
		"test_firstName",
		"test_lastName",
		"test@mail.com",
		"2006-01-02 15:04:05",
		1).WillReturnResult(sqlmock.NewResult(1, 1))

	user := model.User{
		Id: 1,
		Username: "test_username",
		Password: "test_password",
		FirstName: "test_firstName",
		LastName: "test_lastName",
		Email: "test@mail.com",
		Role: "user",
		UpdatedAt: "2006-01-02 15:04:05",
	}

	// now we execute our method
	if err = repository.UpdateUser(db, &user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM users").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err = repository.DeleteUser(db, 1); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}