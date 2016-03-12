package GoGoogleDriveApi

import ("net/http"
	"fmt"
	"io/ioutil"
 	"encoding/json"
	"net/url")

const endpoint string = "https://www.googleapis.com/drive/v2"
const tokenEndpoint string = "https://accounts.google.com/o/oauth2/token"


type Options struct {
	ClientId string
	ClientSecret string
	RefreshToken string
}

type AccessToken struct {
	Token string `json:"access_token"`
	Type string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}
type GoogleDrive struct {
	RootFolderId  string
	AccessToken
	Options
}

func (gd *GoogleDrive) getAccessToken() string {
	if (len(gd.AccessToken.Token) == 0) {
		return gd.resolveToken().Token;
	}
	return gd.AccessToken.Token
}

func (gd *GoogleDrive) resolveToken() AccessToken {



	resp, err := http.PostForm(tokenEndpoint,
		url.Values{"refresh_token": {gd.Options.RefreshToken},
			"client_id": {gd.Options.ClientId},
			"client_secret": {gd.Options.ClientSecret},
			"grant_type": {"refresh_token"}})

	if nil != err {
		fmt.Println("errorination happened getting the response", err)
		return gd.AccessToken
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return gd.AccessToken
	}


	json.Unmarshal(body, &gd.AccessToken)
	return gd.AccessToken
}