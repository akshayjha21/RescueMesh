package sqlite

import (
	"log"
	"regexp"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"` // Fixed JSON tag
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
}

type Sqlite struct {
	Db *gorm.DB
}

func InitDB() (*Sqlite, error) {
	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection has been established")
	err = db.AutoMigrate(&Message{})
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}
func MigrateFromText(db *gorm.DB, rawData string) error {
	// Updated regex to capture your specific log layout
	re := regexp.MustCompile(`Received message at (.+?) from (.+?): (.+)`)
	lines := re.FindAllStringSubmatch(rawData, -1)

	// Use a transaction for crash recovery
	return db.Transaction(func(tx *gorm.DB) error {
		for _, match := range lines {
			// Parse the custom timestamp format
			// Example: 2026-01-15 00:31:34.2480398 +0530 IST
			const layout = "2006-01-14 15:04:05.9999999 -0700 MST"
			parsedTime, _ := time.Parse(layout, match[1])

			msg := Message{
				CreatedAt: parsedTime,
				Sender:    match[2],
				Content:   match[3],
			}

			// Save to SQLite
			if err := tx.Create(&msg).Error; err != nil {
				return err // Rollback if any write fails
			}
		}
		return nil
	})
}
func (db *Sqlite) GetRecords() ([]Message, error) {
	var message []Message
	results := db.Db.Order("id DESC").
		Limit(10).
		Find(&message)
	if results.Error != nil {
		return nil, results.Error
	}
	return message, nil
}
