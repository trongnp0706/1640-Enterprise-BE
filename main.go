package main

import (
	db "GDN-delivery-management/db/sql"
	handle "GDN-delivery-management/delivery/http"
	mail "GDN-delivery-management/mail"
	mdw "GDN-delivery-management/middleware"
	"GDN-delivery-management/repository"
	"GDN-delivery-management/router"
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	err := godotenv.Load("./app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	psqlInfo := os.Getenv("DBSOURCE")
	driver, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return
	}
	Migrate(driver)
	_, err = driver.Exec(`INSERT INTO roles (role_name, ticker) 
								VALUES ('System Admin', 'SAD') ON CONFLICT DO NOTHING`)
	if err != nil {
		log.Println(err)
		return
	}
	queries := db.New(driver)
	userRepo := repository.NewUserRepo(queries)
	sessionRepo := repository.NewSessionRepo(queries)
	roleRepo := repository.NewRoleRepo(queries)
	gmail := mail.NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"))
	userHandle := handle.UserHandler{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
		Email:       gmail,
	}
	roleHandler := handle.RoleHandler{
		RoleRepo: roleRepo,
	}

	authMiddleware := mdw.NewAuthMiddleware(roleRepo, userRepo, accessibleRoles())

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	router := router.Router{
		Echo:           e,
		UserHandler:    userHandle,
		RoleHandler:    roleHandler,
		AuthMiddleware: authMiddleware,
	}
	router.SetupRouter()
	e.Logger.Fatal(e.Start(":1313"))
}

func Migrate(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	d, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Println(err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Applied %d & %d migrations!\n", d, n)
}

func accessibleRoles() map[string][]string {
	return map[string][]string{
		"user/add-user":        {"SAD", "GL"},
		"user/update-user":     {"SAD", "GL"},
		"user/delete-user":     {"SAD", "GL"},
		"user/add-avatar":      {"SAD", "GL"},
		"user/all-user":        {"SAD", "GL"},
		"user/profile":         {"SAD", "GL", "PM"},
		"user/get-me":          {"SAD", "GL", "PM"},
		"role/add-role":        {"SAD", "GL"},
		"role/all-role":        {"SAD", "GL"},
		"role/update-role":     {"SAD", "GL"},
		"role/delete-role":     {"SAD", "GL"},
		"client/add-client":    {"SAD", "GL"},
		"client/all-client":    {"SAD", "GL"},
		"client/update-client": {"SAD", "GL"},
		"client/close-client":  {"SAD", "GL"},
	}
}
