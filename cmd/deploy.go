package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"

	"path/filepath"
	"strings"
	"time"

	"runtime"

	"github.com/google/go-github/v68/github"
	"github.com/spf13/cobra"
)

type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type NetlifyDeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	VerificationURI string `json:"verification_uri"`
	UserCode        string `json:"user_code"`
	ExpiresIn       string `json:"expires_in"`
	Interval        string `json:"interval"`
}

func runDeploy(args []string) error {
	// GitHub OAuth Device Flow
	githubClientID := "Ov23likuN3kpIirlvEv0"
	deviceCode, err := getGitHubDeviceCode(githubClientID)
	if err != nil {
		return fmt.Errorf("failed to get device code: %w", err)
	}

	fmt.Printf("\nPlease visit: %s\n", deviceCode.VerificationURI)
	fmt.Printf("And enter code: %s\n", deviceCode.UserCode)

	// Poll for GitHub token
	githubToken, err := pollForGitHubToken(githubClientID, deviceCode)
	if err != nil {
		return fmt.Errorf("failed to get GitHub token: %w", err)
	}

	// Create GitHub repository
	ctx := context.Background()
	client := github.NewTokenClient(ctx, githubToken)

	// Get current directory name for repo name
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	repoName := filepath.Base(currentDir)

	// Create private repository
	repo, _, err := client.Repositories.Create(ctx, "", &github.Repository{
		Name:    &repoName,
		Private: &[]bool{true}[0],
	})
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}

	fmt.Printf("\nCreated GitHub repository: %s\n", *repo.SSHURL)

	// Initialize git and push
	if err := initAndPushToGitHub(*repo.SSHURL); err != nil {
		return fmt.Errorf("failed to initialize and push to GitHub: %w", err)
	}

	// Get Netlify Personal Access Token
	fmt.Printf("\nPlease visit: https://app.netlify.com/user/applications/personal")
	fmt.Printf("\nCreate a Personal Access Token and enter it here: ")
	var netlifyToken string
	if _, err := fmt.Scanln(&netlifyToken); err != nil {
		return fmt.Errorf("failed to read Netlify token: %w", err)
	}

	// Netlify OAuth Device Flow
	// netlifyClientID := "GiCAYjfVFqU4xupFsPzFYYg5Zepxh8_Aqc8rYCx0mmQ"

	// Get device code
	// deviceCode, err := getNetlifyDeviceCode(netlifyClientID)
	// if err != nil {
	// 	panic(err)
	// 	return fmt.Errorf("failed to get Netlify device code: %w", err)
	// }

	// fmt.Printf("\nPlease visit: %s\n", deviceCode.VerificationURI)
	// fmt.Printf("And enter code: %s\n", deviceCode.UserCode)

	// // Poll for Netlify token
	// netlifyToken, err := pollForNetlifyToken(netlifyClientID, deviceCode)
	// if err != nil {
	// 	panic(err)
	// 	return fmt.Errorf("failed to get Netlify token: %w", err)
	// }
	// panic(netlifyToken)

	// netlify init

	// Create Netlify site using their API
	netlifyAPI := "https://api.netlify.com/api/v1"
	siteData := map[string]interface{}{
		"name": "modular_cms_test_5",
		"repo": map[string]interface{}{
			"provider":    "github",
			"repo_url":    repo.HTMLURL,
			"repo_branch": "main",
		},
		"build_settings": map[string]interface{}{
			"cmd":             "yarn run build",
			"dir":             "_site",
			"base_rel_dir":    true,
			"stop_builds":     false,
			"public_repo":     false,
			"provider":        "github",
			"repo_type":       "git",
			"repo_url":        repo.HTMLURL,
			"repo_branch":     "main",
			"repo_owner_type": "Organization",
			// "repo_visual_editor_template": null,
			// "repo_template_required_extensions": null,
			// "repo_template_usage_url": null,
			// "repo_template_created_from_git_host": null,
			// "repo_template_created_from_slug": null,
			// "base": "",
			// "deploy_key_id": null
		},
	}

	siteJSON, _ := json.Marshal(siteData)
	req, _ := http.NewRequest("POST", netlifyAPI+"/sites", strings.NewReader(string(siteJSON)))
	req.Header.Set("Authorization", "Bearer "+netlifyToken)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create Netlify site: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read Netlify response: %w", err)
	}
	fmt.Printf("Netlify response: %s\n", string(body))

	var site struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&site); err != nil {
		return fmt.Errorf("failed to decode Netlify response: %w", err)
	}

	fmt.Printf("\nSuccess! Your site is deployed at: %s\n", site.URL)
	// fmt.Printf("GitHub repository: %s\n", *repo.HTMLURL)

	return nil
}

func getGitHubDeviceCode(clientID string) (*DeviceCodeResponse, error) {
	data := strings.NewReader(fmt.Sprintf("client_id=%s&scope=repo", clientID))
	req, err := http.NewRequest("POST", "https://github.com/login/device/code", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deviceCode DeviceCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&deviceCode); err != nil {
		return nil, err
	}

	return &deviceCode, nil
}

func pollForGitHubToken(clientID string, deviceCode *DeviceCodeResponse) (string, error) {
	expiration := time.Now().Add(time.Duration(deviceCode.ExpiresIn) * time.Second)

	for time.Now().Before(expiration) {
		data := strings.NewReader(fmt.Sprintf(
			"client_id=%s&device_code=%s&grant_type=urn:ietf:params:oauth:grant-type:device_code",
			clientID,
			deviceCode.DeviceCode,
		))

		req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", data)
		if err != nil {
			return "", err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		var tokenResp TokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
			resp.Body.Close()
			return "", err
		}
		resp.Body.Close()

		if tokenResp.AccessToken != "" {
			return tokenResp.AccessToken, nil
		}

		time.Sleep(time.Duration(deviceCode.Interval) * time.Second)
	}

	return "", fmt.Errorf("authentication timeout")
}

func initAndPushToGitHub(cloneURL string) error {
	commands := []struct {
		name string
		args []string
	}{
		{"git", []string{"remote", "add", "origin", cloneURL}},
		{"git", []string{"push", "-u", "origin", "main"}},
	}

	for _, cmd := range commands {
		command := exec.Command(cmd.name, cmd.args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			return fmt.Errorf("failed to run '%s %s': %w", cmd.name, strings.Join(cmd.args, " "), err)
		}
	}

	return nil
}

func getNetlifyDeviceCode(clientID string) (*NetlifyDeviceCodeResponse, error) {
	req, err := http.NewRequest("GET", "https://app.netlify.com/authorize", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("client_id", clientID)
	q.Add("response_type", "code")
	q.Add("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	req.URL.RawQuery = q.Encode()

	fmt.Printf("Opening browser to: %s\n", req.URL.String())
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", req.URL.String())
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", req.URL.String())
	case "darwin":
		cmd = exec.Command("open", req.URL.String())
	default:
		fmt.Printf("Please open this URL in your browser: %s\n", req.URL.String())
	}
	if cmd != nil {
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to open browser, please visit: %s\n", req.URL.String())
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	panic(string(body))
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}
	deviceCode := NetlifyDeviceCodeResponse{
		DeviceCode:      values.Get("device_code"),
		UserCode:        values.Get("user_code"),
		VerificationURI: values.Get("verification_uri"),
		ExpiresIn:       values.Get("expires_in"),
		Interval:        values.Get("interval"),
	}

	return &deviceCode, nil
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func pollForNetlifyToken(clientID string, deviceCode *NetlifyDeviceCodeResponse) (string, error) {
	expiration := time.Now().Add(time.Duration(parseInt(deviceCode.ExpiresIn)) * time.Second)

	for time.Now().Before(expiration) {
		req, err := http.NewRequest("POST", "https://api.netlify.com/oauth/device/token", nil)
		if err != nil {
			return "", err
		}
		q := req.URL.Query()
		q.Add("client_id", clientID)
		q.Add("device_code", deviceCode.DeviceCode)
		q.Add("grant_type", "device_code")
		req.URL.RawQuery = q.Encode()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		var tokenResp TokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
			resp.Body.Close()
			return "", err
		}
		resp.Body.Close()

		if tokenResp.AccessToken != "" {
			return tokenResp.AccessToken, nil
		}

		time.Sleep(time.Duration(parseInt(deviceCode.Interval)) * time.Second)
	}

	return "", fmt.Errorf("authentication timeout")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your site to a hosting provider",
	Long:  "Deploy your site to a hosting provider (currently supports Netlify)",
	Run: func(cmd *cobra.Command, args []string) {
		runDeploy(args)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
