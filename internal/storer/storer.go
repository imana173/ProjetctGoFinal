package storer

import "contact-cli/internal/models"

type Storer interface {
    Create(models.Contact) error
    List() ([]models.Contact, error)
    Update(models.Contact) error
    Delete(uint) error
}
