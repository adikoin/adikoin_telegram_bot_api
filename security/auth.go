package security

import (
	"telegram_bot_api/repository"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

// func (authValidator *AuthValidator) ValidateCredentials(email, password string) (*model.User, bool) {
// 	user, err := authValidator.userRepository.FindByEmail(email)
// 	if err != nil || util.VerifyPassword(user.Password, password) != nil {
// 		return nil, false
// 	}
// 	return user, true
// }
