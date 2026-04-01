package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"
	"training-plan-api/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthOAuthServiceImpl struct {
	userRepo    repository.UserRepository
	oauthConfig *oauth2.Config
}

type googleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func NewAuthOAuthServiceImpl(
	userRepo repository.UserRepository,
	clientID string,
	clientSecret string,
	redirectURL string,
) AuthOAuthService {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &AuthOAuthServiceImpl{
		userRepo:    userRepo,
		oauthConfig: config,
	}
}

func (s *AuthOAuthServiceImpl) GetGoogleLoginURL(state string) string {
	return s.oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *AuthOAuthServiceImpl) HandleGoogleCallback(code string) (string, *model.User, error) {
	ctx := context.Background()

	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return "", nil, helper.Unauthorized("Failed to exchange Google token")
	}

	userInfo, err := s.fetchGoogleUserInfo(ctx, token)
	if err != nil {
		return "", nil, err
	}

	user, err := s.userRepo.FindByEmail(userInfo.Email)
	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok && appErr.StatusCode == http.StatusNotFound {
newUser := &model.User{
	Name:       userInfo.Name,
	Email:      userInfo.Email,
	Password:   "",

	EmployeeID: "google_" + userInfo.ID, 
	DepartmentID: 1,

	Role:      model.RoleStaff,
	Status:    model.UserStatusActive,
	CreatedBy: model.CreatedBySelf,

	GoogleID: userInfo.ID,
	Avatar:   userInfo.Picture,
	Provider: "google",

	IsProfileComplete: false,
}
			if err := s.userRepo.Save(newUser); err != nil {
				return "", nil, err
			}
			user = newUser
		} else {
			return "", nil, err
		}
	}

	if user.Status != model.UserStatusActive {
		return "", nil, helper.Unauthorized("Account is deactivated")
	}

	if strings.ToLower(user.Provider) != "google" || user.GoogleID == "" || user.Avatar == "" {
		nameUpdate := ""
		if user.Name == "" {
			nameUpdate = userInfo.Name
		}
		if err := s.userRepo.UpdateOAuthFields(user.ID, userInfo.ID, userInfo.Picture, "google", nameUpdate); err != nil {
			return "", nil, err
		}
		user.GoogleID = userInfo.ID
		user.Avatar = userInfo.Picture
		user.Provider = "google"
		if nameUpdate != "" {
			user.Name = nameUpdate
		}
	}

	jwtToken, err := utils.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return "", nil, helper.InternalServerError("Failed to generate access token")
	}

	return jwtToken, user, nil
}

func (s *AuthOAuthServiceImpl) fetchGoogleUserInfo(ctx context.Context, token *oauth2.Token) (googleUserInfo, error) {
	client := s.oauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return googleUserInfo{}, helper.InternalServerError("Failed to fetch Google user info")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return googleUserInfo{}, helper.Unauthorized("Failed to fetch Google user info")
	}

	var info googleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return googleUserInfo{}, helper.InternalServerError("Failed to parse Google user info")
	}
	if info.Email == "" {
		return googleUserInfo{}, helper.BadRequest("Google account email is missing")
	}

	return info, nil
}
