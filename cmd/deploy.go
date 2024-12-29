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

func runDeploy(cmd *cobra.Command, args []string) error {
	// Get current directory name for repo name
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	repoName := filepath.Base(currentDir)

	// Check if git is initialized and has a remote origin
	gitCmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := gitCmd.CombinedOutput()

	// If there's no error, a remote origin exists
	if err == nil && len(output) > 0 {
		fmt.Println("Git repository already has a remote origin, skipping GitHub repository creation")
	} else {
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
	}

	to, err := cmd.Flags().GetString("to")
	if err != nil {
		return fmt.Errorf("failed to get 'to' flag: %w", err)
	}

	if to == "netlify" {
		siteName := strings.ReplaceAll(repoName, "_", "-")
		// Run netlify init command
		netlifyCmd := exec.Command("netlify", "sites:create", "--name", siteName)
		netlifyCmd.Stdout = os.Stdout
		netlifyCmd.Stderr = os.Stderr
		netlifyCmd.Stdin = os.Stdin
		if err := netlifyCmd.Run(); err != nil {
			return fmt.Errorf("failed to run netlify sites:create: %w", err)
		}

		netlifyCmd = exec.Command("netlify", "init")
		netlifyCmd.Stdout = os.Stdout
		netlifyCmd.Stderr = os.Stderr
		netlifyCmd.Stdin = os.Stdin
		if err := netlifyCmd.Run(); err != nil {
			return fmt.Errorf("failed to run netlify init: %w", err)
		}
	}
	if to == "github" {
		// Run build command
		buildCmd := exec.Command("bin/build")
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			return fmt.Errorf("failed to run build command: %w", err)
		}

		// Create docs directory and copy _site contents
		if err := os.RemoveAll("docs"); err != nil {
			return fmt.Errorf("failed to remove existing docs directory: %w", err)
		}
		if err := os.MkdirAll("docs", 0755); err != nil {
			return fmt.Errorf("failed to create docs directory: %w", err)
		}

		// Copy _site contents to docs
		cpCmd := exec.Command("cp", "-r", "_site/.", "docs/")
		if err := cpCmd.Run(); err != nil {
			return fmt.Errorf("failed to copy _site contents to docs: %w", err)
		}

		// Git add and commit
		gitAddCmd := exec.Command("git", "add", "docs")
		if err := gitAddCmd.Run(); err != nil {
			return fmt.Errorf("failed to git add docs directory: %w", err)
		}

		gitCommitCmd := exec.Command("git", "commit", "-m", "Add docs directory for GitHub Pages")
		if err := gitCommitCmd.Run(); err != nil {
			return fmt.Errorf("failed to commit docs directory: %w", err)
		}

		// Push changes
		gitPushCmd := exec.Command("git", "push", "origin", "main")
		if err := gitPushCmd.Run(); err != nil {
			return fmt.Errorf("failed to push changes: %w", err)
		}

		// Get GitHub token for enabling Pages
		githubClientID := "Ov23likuN3kpIirlvEv0"
		deviceCode, err := getGitHubDeviceCode(githubClientID)
		if err != nil {
			return fmt.Errorf("failed to get device code: %w", err)
		}

		fmt.Printf("\nPlease visit: %s\n", deviceCode.VerificationURI)
		fmt.Printf("And enter code: %s\n", deviceCode.UserCode)

		githubToken, err := pollForGitHubToken(githubClientID, deviceCode)
		if err != nil {
			return fmt.Errorf("failed to get GitHub token: %w", err)
		}

		// Get repo owner and name from remote URL
		gitRemoteCmd := exec.Command("git", "remote", "get-url", "origin")
		output, err := gitRemoteCmd.Output()
		if err != nil {
			return fmt.Errorf("failed to get remote URL: %w", err)
		}

		remoteURL := strings.TrimSpace(string(output))
		parts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(remoteURL, "git@github.com:"), ".git"), "/")
		owner := parts[0]
		repoName := parts[1]

		// Enable GitHub Pages
		ctx := context.Background()
		client := github.NewTokenClient(ctx, githubToken)
		pages := &github.Pages{
			BuildType: github.Ptr("legacy"),
			Source: &github.PagesSource{
				Branch: github.Ptr("main"),
				Path:   github.Ptr("/docs"),
			},
		}

		_, _, err = client.Repositories.EnablePages(ctx, owner, repoName, pages)
		if err != nil {
			return fmt.Errorf("failed to enable GitHub Pages: %w", err)
		}

		fmt.Printf("\nEnabled GitHub Pages for repository. Site will be available at https://%s.github.io/%s\n", owner, repoName)
	}
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
	Long:  "Deploy your site to a hosting provider (currently supports Netlify and GitHub Pages)",
	Run: func(cmd *cobra.Command, args []string) {
		runDeploy(cmd, args)
	},
}

func init() {
	deployCmd.Flags().String("to", "netlify", "Provider to deploy to (netlify or github)")
	deployCmd.MarkFlagRequired("to")

	// Add validation for the 'to' flag
	deployCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		provider, _ := cmd.Flags().GetString("to")
		if provider != "netlify" && provider != "github" {
			return fmt.Errorf("invalid provider %q - must be either 'netlify' or 'github'", provider)
		}
		return nil
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
