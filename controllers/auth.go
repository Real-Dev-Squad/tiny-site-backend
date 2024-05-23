package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/Real-Dev-Squad/tiny-site-backend/models"
	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2Api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	googleOAuthConfig *oauth2.Config
	googleConfigMu    sync.Mutex
)

var (
    tokenExpiration = 31536000
)

func getGoogleOAuthConfig() *oauth2.Config {
	googleConfigMu.Lock()
	defer googleConfigMu.Unlock()

	if googleOAuthConfig == nil {
		googleOAuthConfig = &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		}
	}
	return googleOAuthConfig
}

func GoogleLogin(ctx *gin.Context) {
	url := getGoogleOAuthConfig().AuthCodeURL("state")
	ctx.Redirect(http.StatusFound, url)
}

func GoogleCallback(ctx *gin.Context, db *bun.DB) {
    code := ctx.Query("code")
    domain := os.Getenv("DOMAIN")
    authRedirectUrl := os.Getenv("AUTH_REDIRECT_URL")

    googleAccountInfo, getInfoError := getUserInfoFromCode(code, getGoogleOAuthConfig(), ctx)
    if getInfoError != nil {
        log.Printf("Failed to get user info: %v", getInfoError)
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
        return
    }

    log.Printf("Google account info: %v", googleAccountInfo)

    var user models.User

    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        log.Printf("Failed to begin transaction: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
        return
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    count, err := tx.NewSelect().Model(&user).Where("email = ?", googleAccountInfo.Email).ScanAndCount(ctx)
    if err != nil && err.Error() != "sql: no rows in result set" {
        log.Printf("Failed to query user: %v", err)
        return
    }

    log.Printf("User count: %d", count)

    if count == 0 {
        newUser := &models.User{
            UserName: googleAccountInfo.Name,
            Email:    googleAccountInfo.Email,
        }

        result, insertErr := tx.NewInsert().Model(newUser).Exec(ctx)
        if insertErr != nil {
            log.Printf("Failed to create new user: %v", insertErr)
            return
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
            log.Printf("Failed to get rows affected: %v", err)
            return
        }

        log.Printf("Rows affected: %d", rowsAffected)

        if rowsAffected == 0 {
            log.Println("No rows affected, user was not created.")
            return
        }

        log.Printf("User created successfully: %v", newUser)

        user = *newUser
    }

    err = tx.NewSelect().Model(&user).Where("email = ?", googleAccountInfo.Email).Scan(ctx)
    if err != nil {
        log.Printf("Failed to re-query user after insertion: %v", err)
        return
    }

    token, err := utils.GenerateToken(&user)
    if err != nil {
        log.Printf("Failed to generate token: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
        return
    }

    log.Printf("Generated token: %s", token)

    ctx.SetCookie("token", token, tokenExpiration, "/", domain, true, true)
    ctx.Redirect(http.StatusFound, authRedirectUrl)
}

func Logout(ctx *gin.Context) {
	domain := os.Getenv("DOMAIN")
	authRedirectUrl := os.Getenv("AUTH_REDIRECT_URL")

	ctx.SetCookie("token", "", -1, "/", domain, false, true)
	ctx.Redirect(http.StatusFound, authRedirectUrl)
}

func getUserInfoFromCode(code string, conf *oauth2.Config, ctx *gin.Context) (*oauth2Api.Userinfo, error) {
	tok, exchangeErr := conf.Exchange(context.TODO(), code)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	oauth2Service, serviceError := oauth2Api.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))

	if serviceError != nil {
		return nil, serviceError
	}

	userInfo, getInfoError := oauth2Service.Userinfo.Get().Context(ctx).Do()

	if getInfoError != nil {
		return nil, getInfoError
	}

	return userInfo, nil
}
