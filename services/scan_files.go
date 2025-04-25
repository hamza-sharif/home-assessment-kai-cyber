package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/spf13/viper"

	"github.com/hamza-sharif/home-assessment-kai-cyber/config"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

type ScanResults struct {
	details         models.ScanFileInfo
	Vulnerabilities []*models.Vulnerability `json:"vulnerabilities"`
}

type ScanWrapper struct {
	ScanResults ScanResults `json:"scanResults"`
}

func (m *Service) ScanFiles(link string, fileNames []string) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(fileNames))

	cLimit := viper.GetInt(config.ConcurrencyLimit)

	sem := make(chan struct{}, cLimit)
	for _, file := range fileNames {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			sem <- struct{}{}
			err := m.processFiles(link, file)
			<-sem
			if err != nil {
				errChan <- err
			}
		}(file)
	}
	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return errors.New("some files failed to process")
	}
	return nil
}

func (m *Service) processFiles(link, fileName string) error {
	getRequiredPart := strings.Split(link, ".com")
	if len(getRequiredPart) != 2 {
		return errors.New("invalid link")
	}

	rURL := "https://raw.githubusercontent.com/" + getRequiredPart[1] + "/main/" + fileName
	retries := viper.GetInt(config.GitRetries)
	for i := 0; i < retries; i++ {
		resp, err := http.Get(rURL)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			body, errBody := ioutil.ReadAll(resp.Body)
			if errBody != nil {
				continue
			}

			results, errParsing := parseFile(body, fileName)
			if errParsing != nil {
				continue
			}

			if err = m.db.AddVulnerabilities(results); err != nil {
				log().Error("Failed to add vulnerabilities:", err)
				return err
			}
			return nil
		}

	}
	return fmt.Errorf("failed to fetch %s after retries", rURL)
}

func parseFile(rawData []byte, fileName string) ([]*models.Vulnerability, error) {
	var results []ScanWrapper
	vulList := make([]*models.Vulnerability, 0)

	if err := json.Unmarshal(rawData, &results); err != nil {
		log().Error("Failed to parse JSON:", err)
		return nil, err
	}

	for _, item := range results {
		for _, vul := range item.ScanResults.Vulnerabilities {
			vul.Metadata = item.ScanResults.details
			vul.Metadata.FileName = fileName
			vulList = append(vulList, vul)
		}
	}

	return vulList, nil
}
