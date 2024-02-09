package auth

type AuthEntity struct{}

type AuthRepository struct {
	db string
}

func NewRepository(db string) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
