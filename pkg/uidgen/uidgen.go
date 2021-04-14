package uidgen

import "github.com/golang-plus/uuid"

type UIDGen interface {
	New() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() string {
	uid, _ := uuid.NewV4()
	return uid.Format(uuid.StyleWithoutDash)
}
