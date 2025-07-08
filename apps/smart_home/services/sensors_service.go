package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SensorsService handles fetching and mutating sensor data from the external sensors service
type SensorsService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// Sensor represents a smart home sensor (should match your sensors_service model)
type Sensor struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Location    string    `json:"location"`
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewSensorsService creates a new sensors service
func NewSensorsService(baseURL string) *SensorsService {
	return &SensorsService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetSensors fetches all sensors
func (s *SensorsService) GetSensors(ctx context.Context) ([]Sensor, error) {
	url := fmt.Sprintf("%s/sensors", s.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var sensors []Sensor
	if err := json.NewDecoder(resp.Body).Decode(&sensors); err != nil {
		return nil, err
	}
	return sensors, nil
}

// GetSensorByID fetches a sensor by ID
func (s *SensorsService) GetSensorByID(ctx context.Context, id int) (*Sensor, error) {
	url := fmt.Sprintf("%s/sensors/%d", s.BaseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var sensor Sensor
	if err := json.NewDecoder(resp.Body).Decode(&sensor); err != nil {
		return nil, err
	}
	return &sensor, nil
}

// CreateSensor creates a new sensor
func (s *SensorsService) CreateSensor(ctx context.Context, sensor interface{}) (*Sensor, error) {
	url := fmt.Sprintf("%s/sensors", s.BaseURL)
	body, err := json.Marshal(sensor)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var created Sensor
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}
	return &created, nil
}

// UpdateSensor updates a sensor by ID
func (s *SensorsService) UpdateSensor(ctx context.Context, id int, sensor interface{}) (*Sensor, error) {
	url := fmt.Sprintf("%s/sensors/%d", s.BaseURL, id)
	body, err := json.Marshal(sensor)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var updated Sensor
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteSensor deletes a sensor by ID
func (s *SensorsService) DeleteSensor(ctx context.Context, id int) error {
	url := fmt.Sprintf("%s/sensors/%d", s.BaseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// UpdateSensorValue updates only the value and status of a sensor
func (s *SensorsService) UpdateSensorValue(ctx context.Context, id int, value float64, status string) error {
	url := fmt.Sprintf("%s/sensors/%d/value", s.BaseURL, id)
	payload := map[string]interface{}{
		"value":  value,
		"status": status,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
} 