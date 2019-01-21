package acceptance

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imega/avro-learning/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *sql.DB

var _ = BeforeSuite(func() {
	err := WaitForBD()
	Expect(err).NotTo(HaveOccurred())
})

func WaitForBD() error {
	var err error
	mysqlConfig, _ := config.NewMysqlConfigFromEnv("STOCK_")
	maxAttempts := 40

	db, err = sql.Open("mysql", mysqlConfig.String())
	if err != nil {
		return errors.New("failed to open sql connection")
	}

	for {
		if err := db.Ping(); err == nil {
			return nil
		}

		log.Printf("ATTEMPTING TO CONNECT ")
		maxAttempts--
		if maxAttempts == 0 {
			return errors.New("SUT is not ready for tests")
		}
		<-time.After(time.Duration(1 * time.Second))
	}
}

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
	db.Close()
}
