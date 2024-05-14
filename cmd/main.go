package main

import (
	"GoAhead/configs"
	"GoAhead/pkg/api"
	"GoAhead/pkg/data"
	"GoAhead/pkg/repository"
	"GoAhead/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {

	config := configs.Config{}
	if err := config.InitConfig(); err != nil {
		logrus.Errorf("Error during init configs, %s", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error during init config.yml, %s", err)
	}

	cfg := configs.Config{
		Username:   viper.GetString("db.username"),
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		DBName:     viper.GetString("db.dbname"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:    viper.GetString("db.sslmode"),
	}

	db := data.NewDB(&cfg)

	dataData, err := data.NewData(&cfg, db)
	if err != nil {

	}
	repo := repository.NewCentralBankRepo(dataData)
	useCase := service.NewCentralBankUseCase(repo)
	cbRoute := api.NewCentralBankRoute(*useCase)

	router := gin.Default()

	cbRoute.Register(router)
	cronDaily(cbRoute)

	err = router.Run("0.0.0.0:8000")
	if err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}

}

func cronDaily(route *api.CentralBankRoute) {
	c := cron.New()

	_, err := c.AddFunc("* 10 * * *", func() {
		route.GetDailyRate()
	})

	if err != nil {
		logrus.Errorf("Error get daily rate: %s", err)
		return
	}

	c.Start()
}
