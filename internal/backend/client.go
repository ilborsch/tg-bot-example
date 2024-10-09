package backend

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type Client struct {
	chatBotClient
	fileClient
	authClient
	userClient
}

func NewClient(baseURL string) *Client {
	return &Client{
		ChatBotClient{baseURL: baseURL},
		FileClient{baseURL: baseURL},
		AuthClient{baseURL: baseURL},
		UserClient{baseURL: baseURL},
	}
}

func sendGetRequest(url, token string) ([]byte, error) {
	// Create a new HTTP GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// Add the Authorization cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: fmt.Sprintf("Bearer %s", token),
	})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	// Ensure the response body is closed after reading it
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
func sendPostRequest(url string, requestBody []byte, token string) ([]byte, error) {
	// Create a new HTTP POST request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Add the Authorization cookie
	if token != "" {
		req.AddCookie(&http.Cookie{
			Name:  "Authorization",
			Value: fmt.Sprintf("Bearer %s", token),
		})
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	// Ensure the response body is closed after reading it
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("received non-200/201 response status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
func sendPutRequest(url string, requestBody []byte, token string) ([]byte, error) {
	// Create a new HTTP PUT request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Add the Authorization cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: fmt.Sprintf("Bearer %s", token),
	})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send PUT request: %w", err)
	}
	// Ensure the response body is closed after reading it
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("received non-200/204 response status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
func sendDeleteRequest(url, token string) ([]byte, error) {
	// Create a new HTTP DELETE request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create DELETE request: %w", err)
	}

	// Add the Authorization cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: fmt.Sprintf("Bearer %s", token),
	})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send DELETE request: %w", err)
	}
	// Ensure the response body is closed after reading it
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("received non-200/204 response status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func sendPostWithFile(url string, requestBody []byte, token string, filename string, filedata []byte) ([]byte, error) {
	// Create a buffer to write our multipart form data
	var requestBuff bytes.Buffer
	writer := multipart.NewWriter(&requestBuff)

	// Create a form file field
	part, err := writer.CreateFormFile("file", filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	// Write the file data to the form file field
	_, err = io.Copy(part, bytes.NewReader(filedata))
	if err != nil {
		return nil, err
	}

	// Create a form field for the JSON payload
	jsonPart, err := writer.CreateFormField("json")
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the form field
	_, err = io.Copy(jsonPart, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	// Close the writer to finalize the form
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new HTTP POST request with the multipart form data
	req, err := http.NewRequest("POST", url, &requestBuff)
	if err != nil {
		return nil, err
	}

	// Add the Authorization cookie
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: fmt.Sprintf("Bearer %s", token),
	})

	// Set the content type to multipart/form-data with the boundary parameter
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
