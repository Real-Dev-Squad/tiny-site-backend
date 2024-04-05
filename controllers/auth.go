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

    user := new(models.User)
    googleAccountInfo, getInfoError := getUserInfoFromCode(code, getGoogleOAuthConfig(), ctx)

    if getInfoError != nil {
        log.Fatal(getInfoError)
        ctx.JSON(500, gin.H{
            "message": "error",
        })
        return
    }

    count, err := db.NewSelect().Model(user).Where("email = ?", googleAccountInfo.Email).ScanAndCount(ctx)

    if err != nil {
        log.Fatal(err)
        ctx.JSON(500, gin.H{
            "message": "error",
        })
        return
    }

    if count == 0 {
        newUser := &models.User{
            UserName: googleAccountInfo.Name,
            Email:    googleAccountInfo.Email,
        }

        _, err := db.NewInsert().Model(newUser).Exec(ctx)

        if err != nil {
            log.Fatal(err)
            ctx.JSON(500, gin.H{
                "message": "error",
            })
            return
        }
    }

    token, err := utils.GenerateToken(user)

    if err != nil {
        log.Fatal(err)
        ctx.JSON(500, gin.H{
            "message": "error",
        })
        return
    }

    ctx.SetCookie("token", token, tokenExpiration, "/", domain, true, true)
    ctx.Redirect(302, authRedirectUrl)
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
