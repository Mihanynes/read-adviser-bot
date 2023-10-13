package sql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"read-adviser-bot/storage"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, fmt.Errorf("could not open database %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database %w", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Init(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS notes (note_text TEXT, user_name TEXT)`)
	if err != nil {
		return fmt.Errorf("could not init database %w", err)
	}
	return nil
}

func (s *Storage) Save(ctx context.Context, n *storage.Note) error {
	q := `INSERT INTO notes (note_text, user_name) VALUES (?, ?)`
	_, err := s.db.ExecContext(ctx, q, n.NoteText, n.UserName)
	return err

}

func (s *Storage) PickRandom(ctx context.Context, userName string) (*storage.Note, error) {
	var note storage.Note
	err := s.db.QueryRowContext(ctx, "SELECT note_text, user_name FROM notes WHERE user_name = ? ORDER BY RAND() LIMIT 1", userName).Scan(&note.NoteText, &note.UserName)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	}
	if err != nil {
		return nil, fmt.Errorf("could not get random note %w", err)
	}
	return &note, nil
}

func (s *Storage) Remove(ctx context.Context, n *storage.Note) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM notes WHERE note_text = ? AND user_name = ?", n.NoteText)
	return fmt.Errorf("could not remove note %w", err)
}

func (s *Storage) IsExists(ctx context.Context, n *storage.Note) (bool, error) {
	var note storage.Note
	err := s.db.QueryRowContext(ctx, "SELECT note_text, user_name FROM notes WHERE note_text = ? AND user_name = ?", n.NoteText, n.UserName).Scan(&note.NoteText, &note.UserName)
	if err != nil {
		return false, fmt.Errorf("could not check note %w", err)
	}
	return true, nil
}
