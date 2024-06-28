package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"muzz-service/pkg/cryptography"
	"muzz-service/pkg/dao"
	"muzz-service/pkg/types"
	"muzz-service/pkg/types/dummies"
	"net/http"
)

func CreateUser(c *gin.Context) {
	user := types.User{
		Email:    dummies.GenerateUniqueEmail(),
		Password: dummies.GeneratePassword(),
		Name:     dummies.GenerateName(),
		Gender:   dummies.GenerateGender(),
		Dob:      dummies.GenerateDateOfBirth(),
	}

	// hash the password but keep the value of original password
	password := user.Password
	hashedPassword, err := cryptography.HashPassword(user.Password)
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error hashing password", nil)
		return
	}

	user.Password = hashedPassword
	persistedUser, err := dao.CreateUser(user)
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error creating user", nil)
		return
	}

	log.Println("New user created with id:", persistedUser.ID)

	// important: return the original password
	persistedUser.Password = password
	types.OkResp(c, http.StatusCreated, types.UserCreatedResponse{Result: types.UserCreatedInnerResponse{User: persistedUser, Age: persistedUser.GetAge()}})
}
