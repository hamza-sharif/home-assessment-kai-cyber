package sqlite

import (
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"github.com/hamza-sharif/home-assessment-kai-cyber/database"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

func NewMockClient() (database.Client, error) {
	// 1. Set up sqlmock
	conn, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err = conn.AutoMigrate(&models.Vulnerability{}); err != nil {
		return nil, errors.Wrap(err, "failed to create tables")
	}

	return &client{conn: conn}, err
}

var cli, _ = NewMockClient()

func getData(severity string) ([]*models.Vulnerability, error) {
	if severity == "HIGH" {
		listOFVuln := []*models.Vulnerability{}
		err := json.Unmarshal([]byte(vulJson), &listOFVuln)
		return listOFVuln, err
	} else if severity == "LOW" {
		listOFVulnLow := []*models.Vulnerability{}
		err := json.Unmarshal([]byte(vulLowJson), &listOFVulnLow)
		return listOFVulnLow, err
	}
	return nil, errors.New("invalid severity")
}
func Test_client_AddVulnerabilities(t *testing.T) {
	highSerData, err := getData("HIGH")
	if err != nil {
		t.Fatal(err)
	}
	lowSerData, err := getData("LOW")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		vuls []*models.Vulnerability
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{vuls: highSerData},
			wantErr: false,
		},
		{
			name:    "success",
			args:    args{vuls: lowSerData},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.AddVulnerabilities(tt.args.vuls); (err != nil) != tt.wantErr {
				t.Errorf("AddVulnerabilities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_FindVulnerabilities(t *testing.T) {
	highSerData, err := getData("HIGH")
	if err != nil {
		t.Fatal(err)
	}
	lowSerData, err := getData("LOW")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		filter interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Vulnerability
		wantErr bool
	}{
		{
			name:    "success with severity HIGH",
			args:    args{filter: map[string]interface{}{"severity": "HIGH"}},
			want:    highSerData,
			wantErr: false,
		},
		{
			name:    "success with severity LOW",
			args:    args{filter: map[string]interface{}{"severity": "LOW"}},
			want:    lowSerData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cli.FindVulnerabilities(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindVulnerabilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindVulnerabilities() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var vulJson = `[
    {
        "id": "CVE-2024-5555",
        "severity": "HIGH",
        "cvss": 8.5,
        "status": "active",
        "package_name": "tensorflow",
        "current_version": "2.7.0",
        "fixed_version": "2.7.1",
        "description": "Remote code execution in TensorFlow model loading",
        "published_date": "2025-01-24T00:00:00Z",
        "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-5555",
        "risk_factors": [
            "Remote Code Execution",
            "High CVSS Score",
            "Public Exploit Available",
            "Exploit in Wild"
        ],
        "Metadata": {
            "scan_id": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "scan_status": "",
            "resource_type": "",
            "resource_name": "",
            "file_name": "vulnscan1213.json"
        }
    },
    {
        "id": "CVE-2024-8802",
        "severity": "HIGH",
        "cvss": 8.4,
        "status": "active",
        "package_name": "envoy-proxy",
        "current_version": "1.22.0",
        "fixed_version": "1.22.1",
        "description": "Buffer overflow in Envoy proxy header processing",
        "published_date": "2025-01-27T00:00:00Z",
        "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-8802",
        "risk_factors": [
            "Buffer Overflow",
            "High CVSS Score",
            "Proof of Concept Available"
        ],
        "Metadata": {
            "scan_id": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "scan_status": "",
            "resource_type": "",
            "resource_name": "",
            "file_name": "vulnscan1213.json"
        }
    },
    {
        "id": "CVE-2024-9902",
        "severity": "HIGH",
        "cvss": 8.8,
        "status": "active",
        "package_name": "mysql-server",
        "current_version": "8.0.31",
        "fixed_version": "8.0.32",
        "description": "SQL injection vulnerability in MySQL stored procedure handling",
        "published_date": "2025-01-27T15:30:00Z",
        "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-9902",
        "risk_factors": [
            "SQL Injection",
            "High CVSS Score",
            "Proof of Concept Available"
        ],
        "Metadata": {
            "scan_id": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "scan_status": "",
            "resource_type": "",
            "resource_name": "",
            "file_name": "vulnscan1213.json"
        }
    }
]`

var vulLowJson = `[
    {
        "id": "CVE-2024-5557",
        "severity": "LOW",
        "cvss": 3.5,
        "status": "active",
        "package_name": "pandas",
        "current_version": "1.3.0",
        "fixed_version": "1.3.1",
        "description": "Information disclosure in pandas data frame handling",
        "published_date": "2025-01-26T00:00:00Z",
        "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-5557",
        "risk_factors": [
            "Information Disclosure",
            "Low CVSS Score"
        ],
        "Metadata": {
            "scan_id": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "scan_status": "",
            "resource_type": "",
            "resource_name": "",
            "file_name": "vulnscan1213.json"
        }
    },
    {
        "id": "CVE-2024-8804",
        "severity": "LOW",
        "cvss": 3.2,
        "status": "active",
        "package_name": "nginx",
        "current_version": "1.20.1",
        "fixed_version": "1.20.2",
        "description": "Information disclosure in NGINX error logs",
        "published_date": "2025-01-25T00:00:00Z",
        "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-8804",
        "risk_factors": [
            "Information Disclosure",
            "Low CVSS Score"
        ],
        "Metadata": {
            "scan_id": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "scan_status": "",
            "resource_type": "",
            "resource_name": "",
            "file_name": "vulnscan1213.json"
        }
    }
]`
