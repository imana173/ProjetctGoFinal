package storer

import (
    "encoding/json"
    "os"

    "contact-cli/internal/models"
)

type JSONStorer struct {
    Path     string
    Contacts []models.Contact
    nextID   uint
}

func NewJSONStorer(path string) *JSONStorer {
    s := &JSONStorer{Path: path, nextID: 1}

    file, err := os.ReadFile(path)
    if err == nil {
        json.Unmarshal(file, &s.Contacts)

        for _, c := range s.Contacts {
            if c.ID >= s.nextID {
                s.nextID = c.ID + 1
            }
        }
    }

    return s
}

func (s *JSONStorer) save() {
    data, _ := json.MarshalIndent(s.Contacts, "", "  ")
    os.WriteFile(s.Path, data, 0644)
}

func (s *JSONStorer) Create(c models.Contact) error {
    c.ID = s.nextID
    s.nextID++
    s.Contacts = append(s.Contacts, c)
    s.save()
    return nil
}

func (s *JSONStorer) List() ([]models.Contact, error) {
    return s.Contacts, nil
}

func (s *JSONStorer) Update(c models.Contact) error {
    for i, v := range s.Contacts {
        if v.ID == c.ID {
            s.Contacts[i] = c
            s.save()
            return nil
        }
    }
    return nil
}

func (s *JSONStorer) Delete(id uint) error {
    for i, v := range s.Contacts {
        if v.ID == id {
            s.Contacts = append(s.Contacts[:i], s.Contacts[i+1:]...)
            s.save()
            return nil
        }
    }
    return nil
}
