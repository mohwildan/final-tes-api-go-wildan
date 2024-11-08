package main

import (
	http_auth "app/app/delivery/http/auth"
	http_blog "app/app/delivery/http/blog"
	http_faq "app/app/delivery/http/faq"
	"app/app/delivery/http/middleware"
	http_user "app/app/delivery/http/user"
	usecase_auth "app/app/usecase/auth"
	usecase_blog "app/app/usecase/blog"
	usecase_faq "app/app/usecase/faq"
	usecase_user "app/app/usecase/user"
	"app/config"
	"app/migrations"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	yureka_redis "github.com/Yureka-Teknologi-Cipta/yureka/services/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = "5"
	}
	timeout, _ := strconv.Atoi(timeoutStr)
	timeoutContext := time.Duration(timeout) * time.Second

	// logger
	writers := make([]io.Writer, 0)
	if logSTDOUT, _ := strconv.ParseBool(os.Getenv("LOG_TO_STDOUT")); logSTDOUT {
		writers = append(writers, os.Stdout)
	}

	if logFILE, _ := strconv.ParseBool(os.Getenv("LOG_TO_FILE")); logFILE {
		logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
		if logMaxSize == 0 {
			logMaxSize = 50 //default 50 megabytes
		}

		logFilename := os.Getenv("LOG_FILENAME")
		if logFilename == "" {
			logFilename = "server.log"
		}

		lg := &lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    logMaxSize,
			MaxBackups: 1,
			LocalTime:  true,
		}

		writers = append(writers, lg)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(io.MultiWriter(writers...))

	// set gin writer to logrus
	gin.DefaultWriter = logrus.StandardLogger().Writer()

	sqlDb := config.ConnectDatabase()

    // Run migrations
    if err := migrations.Migrate(sqlDb); err != nil {
		logrus.Fatalf("Error running migrations: %v", err)
    }

	// init mongo database
	// mongo := yureka_mongodb.Connect(timeoutContext, os.Getenv("MONGO_URL"), "")

	// init redis database
	var redisClient *redis.Client
	if useRedis, err := strconv.ParseBool(os.Getenv("USE_REDIS")); err == nil && useRedis {
		redisClient = yureka_redis.Connect(timeoutContext, os.Getenv("REDIS_URL"))
	}

	// init usecase
    ucAuth := usecase_auth.NewAppUsecase(usecase_auth.RepoInjection{
		SqlDBRepo: sqlDb,
	}, timeoutContext)

	ucUser := usecase_user.NewAppUsecase(usecase_user.RepoInjection{
		SqlDBRepo: sqlDb,
	}, timeoutContext)

	ucFaq := usecase_faq.NewAppUsecase(usecase_faq.RepoInjection{
		SqlDBRepo: sqlDb,
	}, timeoutContext)

	ucBlog := usecase_blog.NewAppUsecase(usecase_blog.RepoInjection{
		SqlDBRepo: sqlDb,
	}, timeoutContext)

	// init middleware
	mdl := middleware.NewMiddleware(redisClient)

	// gin mode realease when go env is production
	if os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// init gin
	ginEngine := gin.New()

	// add exception handler
	ginEngine.Use(mdl.Recovery())

	// add logger
	ginEngine.Use(mdl.Logger(io.MultiWriter(writers...)))

	// cors
	ginEngine.Use(mdl.Cors())

	// default route
	ginEngine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "It works",
		})
	})

	// init route
	http_auth.NewRouteHandler(ginEngine.Group(""), mdl, ucAuth)
	http_user.NewRouteHandler(ginEngine.Group(""), mdl, ucUser)
	http_faq.NewRouteHandler(ginEngine.Group(""), mdl, ucFaq)
	http_blog.NewRouteHandler(ginEngine.Group(""), mdl, ucBlog)

	port := os.Getenv("PORT")

	logrus.Infof("Service running on port %s", port)
	ginEngine.Run(":" + port)
}
