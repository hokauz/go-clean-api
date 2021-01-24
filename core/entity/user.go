package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	// Credentials -
	Credentials struct {
		Password     string
		Applications []string
	}
	// Access -
	Access struct {
		UserID       string
		AccessToken  string
		RefreshToken string
	}
	// Login -
	Login struct {
		Email    string
		Password string
	}
	// Regiter -
	Regiter struct {
		Email         string
		Password      string
		Name          string
		ApplicationID string
	}
	// Terms -
	Terms struct {
		AcceptedAt string
		CreatedAt  string
		Browser    string
		Device     string
		TermsID    string
		Text       string
		Version    string
		IP         string
	}
	// User -
	User struct {
		ID          primitive.ObjectID
		Email       string
		Name        string
		Credentials *Credentials
		Terms       []*Terms
	}
	// SimpleUser -
	SimpleUser struct {
		ID    primitive.ObjectID
		Email string
		Name  string
	}
)
