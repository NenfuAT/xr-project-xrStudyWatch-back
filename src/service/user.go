package service

import (
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/common"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/util"
)

func CreateUser(req common.UserPost) error {
	var user model.User
	user.ID = util.CreateUserID()
	user.UserName = req.UserName
	user.Age = req.Age
	user.Gender = req.Gender
	user.Mail = req.Mail
	user.Occupation = req.Occupation
	user.Password = req.Password

	model.InsertUser(user)
	return nil
}
