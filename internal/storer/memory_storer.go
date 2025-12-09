package storer

import "contact-cli/internal/models"

type MemoryStorer struct {
    data []models.Contact
    next uint
}

func NewMemoryStorer() *MemoryStorer {
    return &MemoryStorer{next: 1}
}

func (s *MemoryStorer) Create(c models.Contact) error {
    c.ID = s.next
    s.next++
    s.data = append(s.data, c)
    return nil
}

func (s *MemoryStorer) List() ([]models.Contact, error) {
    return s.data, nil
}

func (s *MemoryStorer) Update(c models.Contact) error {
    for i, v := range s.data {
        if v.ID == c.ID {
            s.data[i] = c
            return nil
        }
    }
    return nil
}

func (s *MemoryStorer) Delete(id uint) error {
    for i, v := range s.data {
        if v.ID == id {
            s.data = append(s.data[:i], s.data[i+1:]...)
            return nil
        }
    }
    return nil
}
