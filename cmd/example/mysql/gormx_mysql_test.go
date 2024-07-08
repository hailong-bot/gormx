package mysql

import (
	"github.com/hailong-bot/gormx/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"testing"
	"time"
)

type ModelXTestSuit struct {
	suite.Suite
	DB *gorm.DB
}

func (m *ModelXTestSuit) SetupTest() {
	gormLoggerMode := logger.Info
	mysqlURI := os.Getenv("MYSQL_URI")
	var err error
	db, err := gorm.Open(
		mysql.New(mysql.Config{DSN: mysqlURI}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 log.Default().LogMode(gormLoggerMode), // 使用 logrus 作为 log
		},
	)
	if err != nil {
		panic(err)
	}
	m.DB = db
}

func TestESignTestSuite(t *testing.T) {
	suite.Run(t, &ModelXTestSuit{})
}

func (m *ModelXTestSuit) TestInsertBatch() {
	as := assert.New(m.T())
	cpeModel := NewCpeModel()
	cpeDos := make([]CpeDO, 0)
	for i := 0; i < 3; i++ {
		cpeDos = append(cpeDos, CpeDO{
			FlowID:     strconv.FormatInt(int64(i), 10),
			ContractID: strconv.FormatInt(int64(i), 10),
			UserID:     int64(i),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			IsAlter:    0,
			IsDeleted:  0,
		})
	}
	err := cpeModel.InsertBatch(m.DB, &cpeDos)
	as.NoError(err)
}

func (m *ModelXTestSuit) TestGetByID() {
	as := assert.New(m.T())
	cpeModel := NewCpeModel()
	do, err := cpeModel.GetByID(m.DB, 3)
	logrus.Infof("%+v", do)
	as.NoError(err)
}
