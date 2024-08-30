// Copyright (c) 2021 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	gitpod "github.com/nxpkg/nxpod/gitpod-protocol"
	"github.com/nxpkg/local-app/pkg/prettyprint"
	"github.com/skratchdot/open-golang/open"
	keyring "github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

const keychainServiceName = "nxpkg"

var authScopesLocalCompanion = []string{
	"function:getNxpodTokenScopes",
	"function:getWorkspace",
	"function:getWorkspaces",
	"function:listenForWorkspaceInstanceUpdates",
	"resource:default",
}

func fetchValidCLIScopes(ctx context.Context, serviceURL string) ([]string, error) {
	const clientId = "gitpod-cli"

	endpoint := serviceURL + "/api/oauth/inspect?client=" + clientId

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&authScopesLocalCompanion)
		if err != nil {
			return nil, err
		}
		return authScopesLocalCompanion, nil
	}

	return nil, prettyprint.MarkExceptional(errors.New(serviceURL + " did not provide valid scopes"))
}

type ErrInvalidNxpodToken struct {
	cause error
}

func (e *ErrInvalidNxpodToken) Error() string {
	return "invalid gitpod token: " + e.cause.Error()
}

// ValidateToken validates the given tkn against the given gitpod service
func ValidateToken(client gitpod.APIInterface, tkn string) error {
	hash := sha256.Sum256([]byte(tkn))
	tokenHash := hex.EncodeToString(hash[:])
	tknScopes, err := client.GetNxpodTokenScopes(context.Background(), tokenHash)
	if e, ok := err.(*gitpod.ErrBadHandshake); ok && e.Resp.StatusCode == 401 {
		return &ErrInvalidNxpodToken{err}
	}
	if err != nil && strings.Contains(err.Error(), "jsonrpc2: code 403") {
		return &ErrInvalidNxpodToken{err}
	}
	if err != nil {
		return err
	}
	tknScopesMap := make(map[string]struct{}, len(tknScopes))
	for _, scope := range tknScopes {
		tknScopesMap[scope] = struct{}{}
	}
	for _, scope := range authScopesLocalCompanion {
		_, ok := tknScopesMap[scope]
		if !ok {
			return &ErrInvalidNxpodToken{fmt.Errorf("%v scope is missing in %v", scope, tknScopes)}
		}
	}
	return nil
}

// SetToken returns the persisted Nxpod token
func SetToken(host, token string) error {
	return keyring.Set(keychainServiceName, host, token)
}

// GetToken returns the persisted Nxpod token
func GetToken(host string) (token string, err error) {
	tkn, err := keyring.Get(keychainServiceName, host)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", nil
	}
	return tkn, err
}

// DeleteToken deletes the persisted Nxpod token
func DeleteToken(host string) error {
	return keyring.Delete(keychainServiceName, host)
}

// LoginOpts configure the login process
type LoginOpts struct {
	NxpodURL   string
	RedirectURL string
	AuthTimeout time.Duration

	ExtendScopes bool
}

const html = `
<html>
	<head>
    	<meta charset="utf-8">
    	<title>Done</title>
		<script>
			if (window.opener) {
				const message = new URLSearchParams(window.location.search).get("message");
				window.opener.postMessage(message, "https://${window.location.hostname}");
			} else {
				console.log("This page was not opened by Nxpod.")
				setTimeout("window.close();", 1000);
			}
		</script>
	</head>
	<body>
		If this tab is not closed automatically, feel free to close it and proceed. <button type="button" onclick="window.open('', '_self', ''); window.close();">Close</button>
	</body>
</html>`

// NOTE: the port ranges all need to be valid redirect URI's in the backend
const (
	StartingPortNum = 63110
	EndingPortNum   = 63120
)

// Login walks through the login flow for obtaining a Nxpod token
func Login(ctx context.Context, opts LoginOpts) (token string, err error) {
	// Try a range of ports for local redirect server
	rl, port, err := findOpenPortInRange(StartingPortNum, EndingPortNum)
	if err != nil {
		return "", err
	}

	defer func() {
		closeErr := rl.Close()
		if closeErr != nil {
			slog.Debug("Failed to close listener", "port", port, "err", closeErr)
		}
	}()

	var (
		errChan   = make(chan error, 1)
		queryChan = make(chan url.Values, 1)
	)

	returnHandler := func(rw http.ResponseWriter, req *http.Request) {
		queryChan <- req.URL.Query()
		if opts.RedirectURL != "" {
			http.Redirect(rw, req, opts.RedirectURL, http.StatusSeeOther)
		} else {
			_, _ = io.WriteString(rw, html)
		}
	}

	returnServer := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: http.HandlerFunc(returnHandler),
	}
	go func() {
		err := returnServer.Serve(rl)
		if err != nil {
			errChan <- err
		}
	}()
	defer returnServer.Shutdown(ctx)

	baseURL := opts.NxpodURL
	if baseURL == "" {
		baseURL = "https://gitpod.io"
	}
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	authURL := *reqURL
	authURL.Path = "/api/oauth/authorize"
	tokenURL := *reqURL
	tokenURL.Path = "/api/oauth/token"
	conf := &oauth2.Config{
		ClientID:     "gplctl-1.0",
		ClientSecret: "gplctl-1.0-secret", // Required (even though it is marked as optional?!)
		Scopes:       authScopesLocalCompanion,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL.String(),
			TokenURL: tokenURL.String(),
		},
	}
	if opts.ExtendScopes {
		authScopesLocalCompanion, err = fetchValidCLIScopes(ctx, opts.NxpodURL)
		if err != nil {
			return "", err
		}
		slog.Debug("Using CLI scopes", "scopes", authScopesLocalCompanion)
		conf = &oauth2.Config{
			ClientID:     "gitpod-cli",
			ClientSecret: "gitpod-cli-secret",
			Scopes:       authScopesLocalCompanion,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL.String(),
				TokenURL: tokenURL.String(),
			},
		}
	}
	responseTypeParam := oauth2.SetAuthURLParam("response_type", "code")
	redirectURIParam := oauth2.SetAuthURLParam("redirect_uri", fmt.Sprintf("http://127.0.0.1:%d", rl.Addr().(*net.TCPAddr).Port))
	codeChallengeMethodParam := oauth2.SetAuthURLParam("code_challenge_method", "S256")
	codeVerifier := PKCEVerifier(84)
	codeChallengeParam := oauth2.SetAuthURLParam("code_challenge", PKCEChallenge(codeVerifier))

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	authorizationURL := conf.AuthCodeURL("state", responseTypeParam, redirectURIParam, codeChallengeParam, codeChallengeMethodParam)

	// open a browser window to the authorizationURL
	err = open.Start(authorizationURL)
	if err != nil {
		return "", prettyprint.AddResolution(fmt.Errorf("cannot open browser to URL %s: %s\n", authorizationURL, err),
			"provide a personal access token using --token or the NXPOD_TOKEN environment variable",
		)
	}

	var query url.Values
	var code, approved string
	select {
	case <-ctx.Done():
		return "", errors.New("context cancelled")
	case err = <-errChan:
		return "", err
	case query = <-queryChan:
		code = query.Get("code")
		approved = query.Get("approved")
	case <-time.After(opts.AuthTimeout):
		return "", xerrors.Errorf("auth timeout after %d seconds", uint32(opts.AuthTimeout))
	}

	if approved == "no" {
		return "", errors.New("client approval was not granted")
	}

	// Use the authorization code that is pushed to the redirect URL. Exchange will do the handshake to retrieve the
	// initial access token. We need to add the required PKCE params as well...
	// NOTE: we do not currently support refreshing so using the client from conf.Client will fail when token expires
	codeVerifierParam := oauth2.SetAuthURLParam("code_verifier", codeVerifier)
	tok, err := conf.Exchange(ctx, code, codeVerifierParam, redirectURIParam)
	if err != nil {
		return "", err
	}

	// Extract Nxpod token from OAuth token (JWT)
	// NOTE: we do not verify the token as that requires a shared secret
	//       ... which wouldn't be secret for a publicly accessible app
	claims := jwt.MapClaims{}
	parser := new(jwt.Parser)
	_, _, err = parser.ParseUnverified(tok.AccessToken, &claims)
	if err != nil {
		return "", err
	}

	gitpodToken := claims["jti"].(string)
	return gitpodToken, nil
}

func findOpenPortInRange(start, end int) (net.Listener, int, error) {
	for port := start; port < end; port++ {
		rl, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
		if err != nil {
			slog.Debug("could not open port, trying next port", "port", port, "err", err)
			continue
		}

		return rl, port, nil
	}
	return nil, 0, xerrors.Errorf("could not open any valid port in range %d - %d", start, end)
}

type AuthenticatedTransport struct {
	T     http.RoundTripper
	Token string
}

func (t *AuthenticatedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.Token)
	return t.T.RoundTrip(req)
}
