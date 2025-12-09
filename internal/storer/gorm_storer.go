package storer

import (
    "contact-cli/internal/models"
    "github.com/glebarez/sqlite"  // driver compatible Windows
    "gorm.io/gorm"
)


type GORMStorer struct {
    DB *gorm.DB
}

func NewGORMStorer(path string) *GORMStorer {
   db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
if err != nil {
    panic("❌ Impossible d'ouvrir la base SQLite : " + err.Error())
}

    db.AutoMigrate(&models.Contact{})
    return &GORMStorer{DB: db}
}

func (s *GORMStorer) Create(c models.Contact) error {
    return s.DB.Create(&c).Error
}

func (s *GORMStorer) List() ([]models.Contact, error) {
    var list []models.Contact
    s.DB.Find(&list)
    return list, nil
}

func (s *GORMStorer) Update(c models.Contact) error {
    return s.DB.Save(&c).Error
}

func (s *GORMStorer) Delete(id uint) error {
    return s.DB.Delete(&models.Contact{}, id).Error
}
