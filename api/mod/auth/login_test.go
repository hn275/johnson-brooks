package auth

import (
	"context"
	"errors"
	"jb/database"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var testUser Credentials = Credentials{
	Username: "test_foo",
	Password: "test_bar",
	Session: Session{
		SessionID:      "",
		FailedAttempts: 0,
		Locked:         false,
	},
}

var testCredential Credentials = Credentials{
	Username: "test_foo",
	Password: "test_bar",
}

func TestAuthOK(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	db := newDb()

	user := testUser
	if err := authenticateUser(&testCredential, &user, db); err != nil {
		t.Fatalf("expect `err == nil`, got `%v`", err)
	}
}

func TestAuthFailed(t *testing.T) {
	setup()
	t.Cleanup(cleanup)

	db := newDb()
	defer db.Close()

	failedCred := Credentials{
		Username: "test_for",
		Password: "asldfkdlsjf",
	}
	user := testUser

	for i := 0; i < 5; i++ {
		err := authenticateUser(&failedCred, &user, db)

		// Test failed auth
		if err == nil {
			t.Fatalf("failed authentication, expected `err != nil`")
		}

		if !errors.Is(err, ErrAuthFailed) {
			t.Fatalf("unexpected err returned, expected `ErrAuthFailed`")
		}

		// Test failed attempts
		findTestUser(&user, db)
		attempts := user.Session.FailedAttempts
		if attempts != uint8(i+1) {
			t.Fatalf(
				"failed to increment login attempts, expected `%d`, got `%d`",
				i+1,
				user.Session.FailedAttempts,
			)
		}
	}

	isLocked := user.Session.FailedAttempts >= 5 && user.Session.Locked
	if !isLocked {
		t.Fatalf("failed to lock user account, expected `true`, got `%v`", isLocked)
	}

}

func findTestUser(admin *Credentials, db *authDatabase) {
	ctx := context.TODO()
	filter := bson.D{{Key: "_id", Value: admin.ID}}
	err := db.Collection(database.Admin).FindOne(ctx, filter).Decode(&admin)
	if err != nil {
		panic(err)
	}
}

func setup() {
	hashed, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	testUser.Password = string(hashed)

	db := database.New()
	defer db.Close()

	col := db.Collection(database.Admin)
	result, err := col.InsertOne(context.TODO(), &testUser)
	if err != nil {
		log.Fatal(err)
	}

	testUser.ID = result.InsertedID.(primitive.ObjectID)
}

func cleanup() {
	db := database.New()
	defer db.Close()

	col := db.Collection(database.Admin)
	filter := bson.D{{Key: "_id", Value: testUser.ID}}

	_, err := col.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
