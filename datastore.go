package imapbackend

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// New creates a new datastore with the given database connection string/url
// e.g. postgres://user:pass@localhost/dbname
// e.g. sqlite:/path/to/file.db
func InitDB(dbURL string) error {

	u, err := dburl.Parse(dbURL)
	if err != nil {
		return fmt.Errorf("couldn't parse database connection url: %w", err)
	}

	c := &gorm.Config{}

	switch u.Driver {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(u.DSN), c)

	case "postgres":
		db, err = gorm.Open(postgres.Open(u.DSN), c)

	case "mysql":
		db, err = gorm.Open(mysql.Open(u.DSN), c)

	default:
		return fmt.Errorf("unsupported database driver: %s", u.Driver)
	}

	if err != nil {
		return fmt.Errorf("failed to establish a database connection: %w", err)
	}

	// Migrate
	db.AutoMigrate(&User{}, &Mailbox{}, &Message{})

	return nil

}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("couldn't get sql db: %w", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("couldn't close db connection: %w", err)
	}
	return nil
}

func seedDB() {

	user := &User{Username_: "username", Password: "password", Email: "username@localhost"}

	result := db.Create(&user)
	if result.Error != nil {
		log.Warnf("couldn't seed user: %v", result.Error)
		return
	}

	mailbox := &Mailbox{Name_: "INBOX", User: user}
	result = db.Create(&mailbox)
	if result.Error != nil {
		log.Warnf("couldn't seed mailbox: %v", result.Error)
		return
	}

	body := "From: contact@example.org\r\n" +
		"To: contact@example.org\r\n" +
		"Subject: A little message, just for you\r\n" +
		"Date: Wed, 11 May 2016 14:31:59 +0000\r\n" +
		"Message-ID: <0000000@localhost/>\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi there :)"

	message := &Message{
		UID:   1,
		Date:  time.Now(),
		Flags: []string{"\\Seen"},
		Size:  uint32(len(body)),
		Body:  []byte(body),

		MailboxID: mailbox.ID,
	}
	result = db.Create(&message)
	if result.Error != nil {
		log.Warnf("couldn't seed message: %v", result.Error)
		return
	}

}
