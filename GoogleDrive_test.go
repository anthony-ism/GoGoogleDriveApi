package GoGoogleDriveApi

import ("os"
	"testing")

func TestGetAccessToken(t *testing.T) {
	gd := new(GoogleDrive)

	gd.Options.ClientId = os.Getenv("clientId")
	gd.Options.ClientSecret = os.Getenv("clientSecret")
	gd.Options.RefreshToken = os.Getenv("refreshToken")

	gd.getAccessToken();
	if (len(gd.AccessToken.Token) == 0) {
		t.Errorf("resolveToken().Token should have a length greater than 0")
	}
	if (gd.AccessToken.Type != "Bearer") {
		t.Errorf("resolveToken().token_type should be equal to Bearer")
	}
	if (gd.AccessToken.ExpiresIn != 3600) {
		t.Errorf("resolveToken().expires_in should be equal to 3600")
	}
}