package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var provider string

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your site to a hosting provider",
	Long:  "Deploy your site to a hosting provider (currently supports Netlify)",
	Run: func(cmd *cobra.Command, args []string) {
		if provider != "netlify" {
			fmt.Println("Error: Only Netlify provider is currently supported")
			os.Exit(1)
		}

		// Get current directory name for site name
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			os.Exit(1)
		}
		siteName := filepath.Base(currentDir)

		ghToken, err := getGitHubToken()
		if err != nil {
			fmt.Printf("Error authenticating with GitHub: %v\n", err)
			os.Exit(1)
		}

		// Get Netlify token from environment
		netlifyToken := os.Getenv("NETLIFY_AUTH_TOKEN")
		if netlifyToken == "" {
			fmt.Println("Please set NETLIFY_AUTH_TOKEN environment variable")
			fmt.Println("You can create one at https://app.netlify.com/user/applications")
			os.Exit(1)
		}

		ctx := context.Background()

		// Initialize GitHub client
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		ghClient := github.NewClient(tc)

		// Create GitHub repository if it doesn't exist
		repo, err := createGitHubRepo(ctx, ghClient, siteName)
		if err != nil {
			fmt.Printf("Error creating GitHub repository: %v\n", err)
			os.Exit(1)
		}

		// Deploy to Netlify
		fmt.Println("Deploying to Netlify...")
		siteURL, err := createNetlifySite(netlifyToken, *repo.HTMLURL, siteName)
		if err != nil {
			fmt.Printf("Error deploying to Netlify: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("🎉 Deployment complete! Your site will be available at: %s\n", siteURL)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&provider, "provider", "p", "netlify", "Hosting provider (only netlify supported)")
}

func createGitHubRepo(ctx context.Context, client *github.Client, name string) (*github.Repository, error) {
	// Check if repo exists
	repo, _, err := client.Repositories.Get(ctx, "", name)
	if err == nil {
		return repo, nil
	}

	// Create new repo
	repo, _, err = client.Repositories.Create(ctx, "", &github.Repository{
		Name:        github.String(name),
		Private:     github.Bool(false),
		Description: github.String("Created with modular_cli"),
		AutoInit:    github.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %v", err)
	}

	return repo, nil
}

func createNetlifySite(token, repoURL, siteName string) (string, error) {
	type createSiteRequest struct {
		Name string `json:"name"`
		Repo struct {
			Provider string `json:"provider"`
			RepoURL  string `json:"repo_url"`
			Branch   string `json:"branch"`
			Dir      string `json:"dir"`
			Cmd      string `json:"cmd"`
		} `json:"repo"`
	}

	type createSiteResponse struct {
		URL string `json:"url"`
	}

	reqBody := createSiteRequest{
		Name: siteName,
	}
	reqBody.Repo.Provider = "github"
	reqBody.Repo.RepoURL = repoURL
	reqBody.Repo.Branch = "main"
	reqBody.Repo.Dir = "_site"
	reqBody.Repo.Cmd = "yarn build"

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.netlify.com/api/v1/sites", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to create site: %s", string(body))
	}

	var site createSiteResponse
	if err := json.NewDecoder(resp.Body).Decode(&site); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return site.URL, nil
}

func getGitHubToken() (string, error) {
	// These would be your app's client ID - it's okay to hardcode this
	// You'll need to register your app at https://github.com/settings/applications/new
	clientID := "Ov23likuN3kpIirlvEv0"

	// Request device and user codes
	deviceCodeResp, err := requestDeviceCode(clientID)
	if err != nil {
		return "", fmt.Errorf("error requesting device code: %v", err)
	}

	// Show instructions to user
	fmt.Printf("\nFirst copy your one-time code: %s\n", deviceCodeResp.UserCode)
	fmt.Printf("Press Enter to open github.com in your browser...\n")
	fmt.Scanln() // Wait for Enter

	// Open browser
	url := deviceCodeResp.VerificationURI
	if err := openBrowser(url); err != nil {
		fmt.Printf("Failed to open browser. Please visit: %s\n", url)
	}

	// Poll for token
	token, err := pollForToken(clientID, deviceCodeResp.DeviceCode, deviceCodeResp.Interval)
	if err != nil {
		return "", fmt.Errorf("error polling for token: %v", err)
	}

	return token, nil
}

type deviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

func requestDeviceCode(clientID string) (*deviceCodeResponse, error) {
	resp, err := http.PostForm("https://github.com/login/device/code", url.Values{"client_id": {clientID}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	// Parse response body as URL-encoded form values
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, fmt.Errorf("error parsing response body: %v", err)
	}

	// Convert form values to deviceCodeResponse struct
	deviceCodeResponse := &deviceCodeResponse{
		DeviceCode:      values.Get("device_code"),
		UserCode:        values.Get("user_code"),
		VerificationURI: values.Get("verification_uri"),
	}

	// Parse numeric fields
	if expiresIn, err := strconv.Atoi(values.Get("expires_in")); err == nil {
		deviceCodeResponse.ExpiresIn = expiresIn
	}
	if interval, err := strconv.Atoi(values.Get("interval")); err == nil {
		deviceCodeResponse.Interval = interval
	}

	return deviceCodeResponse, nil
}

func pollForToken(clientID, deviceCode string, interval int) (string, error) {
	for {
		resp, err := http.PostForm("https://github.com/login/oauth/access_token",
			url.Values{
				"client_id":   {clientID},
				"device_code": {deviceCode},
				"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
			})
		if err != nil {
			return "", err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("Failed to read response body: %v", err))
		}

		// Parse response body as URL-encoded form values
		values, err := url.ParseQuery(string(body))
		if err != nil {
			return "", fmt.Errorf("error parsing response body: %v", err)
		}

		accessToken := values.Get("access_token")
		if accessToken != "" {
			return accessToken, nil
		}

		error := values.Get("error")
		if error != "authorization_pending" {
			return "", fmt.Errorf("authorization failed: %s", error)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
