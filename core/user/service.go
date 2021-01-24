package user

import (
	"fmt"

	"github.com/hokauz/go-clean-api/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Login -
func Login() {}

// Regiter -
func Regiter() {}

// RecoverPassword -
func RecoverPassword() {}

// AddTerms -
func AddTerms() {
	// NOTE used when terms are update or when the user entry a new application
}

// Create -
func Create() {
	// 1. Verify existence
	// 2. Generete hash && credentials
	// 3. Create User
	// 4. Return
}

// Read -
func Read(id string) (u *entity.SimpleUser) {
	ref, _ := primitive.ObjectIDFromHex(id)
	fmt.Sprintln(ref)
	return nil
}

// Update -
func Update() {}

// Delete -
func Delete() {}

func handlerError(err error) error {
	return nil
}
