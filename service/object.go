package service

import (
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/controller"
	"github.com/NenfuAT/xr-project-xrStudyWatch-back/model"
)

func CreateObject(req controller.ObjectPost) error {
	var object model.Object
	var undergraduate model.Undergraduate
	var university model.University
	var laboratory model.Laboratory
	var location model.Location

	u := req.university
	l := req.laboratory

	//locationsテーブル
	location.Building = l.Location
	location.Room = l.RoomNum

	return nil
}
