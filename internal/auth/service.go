package auth

type AuthModel struct{}

type AuthService struct {
	repository *AuthRepository
}

func NewService(authRepository *AuthRepository) *AuthService {
	return &AuthService{
		repository: authRepository,
	}
}
