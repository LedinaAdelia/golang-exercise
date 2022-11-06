package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	// Select target token and delete from listSessions
	// TODO: answer here
	for i := 0; i < len(listSessions); i++ {
		if listSessions[i].Token == tokenTarget {
			listSessions[i] = listSessions[len(listSessions)-1]
			listSessions = listSessions[:len(listSessions)-1]
		}
	}

	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}
	listSessions = append(listSessions, session)
	var jsonData, _ = json.Marshal(listSessions)
	u.db.Save("sessions", jsonData)
	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}

	for _, v := range listSessions {
		if u.TokenExpired(v) {
			u.DeleteSessions(v.Token)
			return model.Session{}, fmt.Errorf("Token is Expired!")
		} else {
			return v, nil
		}
	}
	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
