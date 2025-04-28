package browser

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/augment-local-manus-clone/backend/web-browsing-service/domain"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
)

// ChromeDPClient implements the BrowserClient interface using chromedp
type ChromeDPClient struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// NewChromeDPClient creates a new ChromeDPClient
func NewChromeDPClient() (*ChromeDPClient, error) {
	// Create a new browser context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-notifications", true),
		chromedp.Flag("disable-infobars", true),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx)

	return &ChromeDPClient{
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Browse navigates to a URL and returns the page content
func (c *ChromeDPClient) Browse(request *domain.WebBrowsingRequest) (*domain.WebBrowsingResult, error) {
	// Create a timeout context
	ctx, cancel := context.WithTimeout(c.ctx, time.Duration(request.Timeout)*time.Second)
	defer cancel()

	// Create result
	result := &domain.WebBrowsingResult{
		URL:       request.URL,
		Timestamp: time.Now(),
	}

	// Record start time
	startTime := time.Now()

	// Define variables to store page data
	var title, content string
	var buf []byte

	// Create tasks
	tasks := []chromedp.Action{
		chromedp.Navigate(request.URL),
	}

	// Add wait for selector if provided
	if request.WaitForSelector != "" {
		tasks = append(tasks, chromedp.WaitVisible(request.WaitForSelector, chromedp.ByQuery))
	}

	// Add tasks to extract data
	tasks = append(tasks,
		chromedp.Title(&title),
		chromedp.OuterHTML("html", &content),
	)

	// Add tasks to extract specific selectors
	extractedData := make(map[string]string)
	if len(request.ExtractSelectors) > 0 {
		for _, selector := range request.ExtractSelectors {
			var text string
			tasks = append(tasks, chromedp.Text(selector, &text, chromedp.ByQuery))
			extractedData[selector] = text
		}
	}

	// Add screenshot task if requested
	if request.TakeScreenshot {
		tasks = append(tasks, chromedp.CaptureScreenshot(&buf))
	}

	// Run tasks
	err := chromedp.Run(ctx, tasks...)
	if err != nil {
		result.Error = err.Error()
		result.StatusCode = 500
		result.Duration = time.Since(startTime).Seconds()
		return result, nil
	}

	// Update result
	result.Title = title
	result.Content = content
	result.ExtractedData = extractedData
	result.StatusCode = 200
	result.Duration = time.Since(startTime).Seconds()

	// Add screenshot if captured
	if request.TakeScreenshot && len(buf) > 0 {
		result.Screenshot = base64.StdEncoding.EncodeToString(buf)
	}

	return result, nil
}

// Search performs a search query and returns the results
func (c *ChromeDPClient) Search(request *domain.WebSearchRequest) (*domain.WebBrowsingResult, error) {
	// Create a timeout context
	ctx, cancel := context.WithTimeout(c.ctx, 60*time.Second)
	defer cancel()

	// Create result
	result := &domain.WebBrowsingResult{
		Timestamp: time.Now(),
	}

	// Record start time
	startTime := time.Now()

	// Determine search URL based on search engine
	var searchURL string
	switch strings.ToLower(request.SearchEngine) {
	case "google":
		searchURL = fmt.Sprintf("https://www.google.com/search?q=%s", request.Query)
	case "bing":
		searchURL = fmt.Sprintf("https://www.bing.com/search?q=%s", request.Query)
	case "duckduckgo":
		searchURL = fmt.Sprintf("https://duckduckgo.com/?q=%s", request.Query)
	default:
		searchURL = fmt.Sprintf("https://www.google.com/search?q=%s", request.Query)
	}

	result.URL = searchURL

	// Define variables to store page data
	var title, content string
	var buf []byte

	// Create tasks
	tasks := []chromedp.Action{
		chromedp.Navigate(searchURL),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Title(&title),
		chromedp.OuterHTML("html", &content),
	}

	// Add screenshot task if requested
	if request.TakeScreenshot {
		tasks = append(tasks, chromedp.CaptureScreenshot(&buf))
	}

	// Run tasks
	err := chromedp.Run(ctx, tasks...)
	if err != nil {
		result.Error = err.Error()
		result.StatusCode = 500
		result.Duration = time.Since(startTime).Seconds()
		return result, nil
	}

	// Update result
	result.Title = title
	result.Content = content
	result.StatusCode = 200
	result.Duration = time.Since(startTime).Seconds()

	// Add screenshot if captured
	if request.TakeScreenshot && len(buf) > 0 {
		result.Screenshot = base64.StdEncoding.EncodeToString(buf)
	}

	return result, nil
}

// Interact interacts with elements on a web page
func (c *ChromeDPClient) Interact(request *domain.WebInteractionRequest) (*domain.WebBrowsingResult, error) {
	// Create a timeout context
	ctx, cancel := context.WithTimeout(c.ctx, 60*time.Second)
	defer cancel()

	// Create result
	result := &domain.WebBrowsingResult{
		URL:       request.URL,
		Timestamp: time.Now(),
	}

	// Record start time
	startTime := time.Now()

	// Define variables to store page data
	var title, content string
	var buf []byte

	// Create initial tasks to navigate to the page
	tasks := []chromedp.Action{
		chromedp.Navigate(request.URL),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	}

	// Add wait for selector if provided
	if request.WaitForSelector != "" {
		tasks = append(tasks, chromedp.WaitVisible(request.WaitForSelector, chromedp.ByQuery))
	}

	// Add interaction task based on action
	switch request.Action {
	case "click":
		tasks = append(tasks, chromedp.Click(request.Selector, chromedp.ByQuery))
	case "type":
		tasks = append(tasks, chromedp.SendKeys(request.Selector, request.Value, chromedp.ByQuery))
	case "select":
		tasks = append(tasks, chromedp.SendKeys(request.Selector, request.Value, chromedp.ByQuery))
	case "focus":
		tasks = append(tasks, chromedp.Focus(request.Selector, chromedp.ByQuery))
	case "hover":
		tasks = append(tasks, chromedp.MouseOver(request.Selector, chromedp.ByQuery))
	case "scroll":
		tasks = append(tasks, chromedp.ScrollIntoView(request.Selector, chromedp.ByQuery))
	case "screenshot":
		// This will be handled below
	}

	// Add tasks to extract page data
	tasks = append(tasks,
		chromedp.Title(&title),
		chromedp.OuterHTML("html", &content),
	)

	// Add screenshot task if requested
	if request.TakeScreenshot || request.Action == "screenshot" {
		if request.Action == "screenshot" && request.Selector != "" {
			// Screenshot of specific element
			tasks = append(tasks, chromedp.Screenshot(request.Selector, &buf, chromedp.ByQuery))
		} else {
			// Full page screenshot
			tasks = append(tasks, chromedp.CaptureScreenshot(&buf))
		}
	}

	// Run tasks
	err := chromedp.Run(ctx, tasks...)
	if err != nil {
		result.Error = err.Error()
		result.StatusCode = 500
		result.Duration = time.Since(startTime).Seconds()
		return result, nil
	}

	// Update result
	result.Title = title
	result.Content = content
	result.StatusCode = 200
	result.Duration = time.Since(startTime).Seconds()

	// Add screenshot if captured
	if (request.TakeScreenshot || request.Action == "screenshot") && len(buf) > 0 {
		result.Screenshot = base64.StdEncoding.EncodeToString(buf)
	}

	return result, nil
}

// Close closes the browser
func (c *ChromeDPClient) Close() error {
	c.cancel()
	return nil
}
