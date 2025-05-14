package dbstore

import (
	"goth/internal/store"

	"database/sql"
	"github.com/google/uuid"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type SessionStore struct {
	db *sql.DB
}

type NewSessionStoreParams struct {
	DB *sql.DB
}

func NewSessionStore(params NewSessionStoreParams) *SessionStore {
	return &SessionStore{
		db: params.DB,
	}
}

func (s *SessionStore) CreateSession(session *store.Session) (*store.Session, error) {

	session.SessionID = uuid.New().String()

  // TODO: must fix session for sqlite
	return session, nil
}

func (s *SessionStore) GetUserFromSession(sessionID string, userID string) (*store.User, error) {
	var session store.Session

  //TODO: must fix session for sqlite

	return &session.User, nil
}
