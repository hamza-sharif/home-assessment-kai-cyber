package services

import (
	"encoding/json"
	"github.com/hamza-sharif/home-assessment-kai-cyber/database/sqlite"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
	"reflect"
	"testing"
)

func TestService_ScanFiles(t *testing.T) {
	db, _ := sqlite.NewDefaultClient()
	m := &Service{
		db: db,
	}

	type args struct {
		link      string
		fileNames []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				link:      "https://github.com/velancio/vulnerability_scans",
				fileNames: []string{"vulnscan1011.json"},
			},
			wantErr: false,
		},
		{
			name: "success",
			args: args{
				link:      "https://github.com/velancio/vulnerability_scans",
				fileNames: []string{"dasvulnscan1011.json"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.ScanFiles(tt.args.link, tt.args.fileNames); (err != nil) != tt.wantErr {
				t.Errorf("ScanFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_processFiles(t *testing.T) {
	db, _ := sqlite.NewDefaultClient()
	m := &Service{
		db: db,
	}

	type args struct {
		link     string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				link:     "https://github.com/velancio/vulnerability_scans",
				fileName: "vulnscan1011.json",
			},
			wantErr: false,
		},
		{
			name: "success",
			args: args{
				link:     "https://github.com/velancio/vulnerability_scans",
				fileName: "vulnscdsadan1011.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.processFiles(tt.args.link, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("processFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseFile(t *testing.T) {
	listOfVal := []*models.Vulnerability{}
	err := json.Unmarshal([]byte(stringOfVul), &listOfVal)
	if err != nil {
		t.Error(err)
	}

	type args struct {
		rawData  []byte
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Vulnerability
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				rawData:  []byte(fileString),
				fileName: "testFile.json",
			},
			want:    listOfVal,
			wantErr: false,
		}, {
			name: "error",
			args: args{
				rawData:  []byte("error{}{" + fileString),
				fileName: "testFile.json",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFile(tt.args.rawData, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var fileString = `[
  {
    "scanResults": {
      "scan_id": "VULN_scan_456def",
      "timestamp": "2025-01-29T09:15:00Z",
      "scan_status": "completed",
      "resource_type": "container",
      "resource_name": "auth-service:2.1.0",
      "vulnerabilities": [
        {
          "id": "CVE-2024-2222",
          "severity": "HIGH",
          "cvss": 8.2,
          "status": "active",
          "package_name": "spring-security",
          "current_version": "5.6.0",
          "fixed_version": "5.6.1",
          "description": "Authentication bypass in Spring Security",
          "published_date": "2025-01-27T00:00:00Z",
          "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-2222",
          "risk_factors": [
            "Authentication Bypass",
            "High CVSS Score",
            "Proof of Concept Exploit Available"
          ]
        },
        {
          "id": "CVE-2024-2223",
          "severity": "MEDIUM",
          "cvss": 6.5,
          "status": "active",
          "package_name": "tomcat",
          "current_version": "9.0.50",
          "fixed_version": "9.0.51",
          "description": "Information disclosure in Apache Tomcat",
          "published_date": "2025-01-28T00:00:00Z",
          "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-2223",
          "risk_factors": [
            "Information Disclosure",
            "Medium CVSS Score"
          ]
        }
      ],
      "summary": {
        "total_vulnerabilities": 2,
        "severity_counts": {
          "CRITICAL": 0,
          "HIGH": 1,
          "MEDIUM": 1,
          "LOW": 0
        },
        "fixable_count": 2,
        "compliant": false
      },
      "scan_metadata": {
        "scanner_version": "30.1.51",
        "policies_version": "2025.1.29",
        "scanning_rules": [
          "vulnerability",
          "compliance",
          "malware"
        ],
        "excluded_paths": [
          "/tmp",
          "/var/log"
        ]
      }
    }
  },
  {
    "scanResults": {
      "scan_id": "VULN_scan_123abc",
      "timestamp": "2025-01-29T08:00:00Z",
      "scan_status": "completed",
      "resource_type": "container",
      "resource_name": "payment-processor:1.0.0",
      "vulnerabilities": [
        {
          "id": "CVE-2024-1111",
          "severity": "CRITICAL",
          "cvss": 9.9,
          "status": "active",
          "package_name": "openssl",
          "current_version": "3.0.0",
          "fixed_version": "3.0.1",
          "description": "Critical buffer overflow in OpenSSL TLS handling",
          "published_date": "2025-01-28T00:00:00Z",
          "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-1111",
          "risk_factors": [
            "Buffer Overflow",
            "Critical CVSS Score",
            "Public Exploit Available",
            "Exploit in Wild"
          ]
        }
      ],
      "summary": {
        "total_vulnerabilities": 1,
        "severity_counts": {
          "CRITICAL": 1,
          "HIGH": 0,
          "MEDIUM": 0,
          "LOW": 0
        },
        "fixable_count": 1,
        "compliant": false
      },
      "scan_metadata": {
        "scanner_version": "30.1.51",
        "policies_version": "2025.1.29",
        "scanning_rules": [
          "vulnerability",
          "compliance",
          "malware"
        ],
        "excluded_paths": [
          "/tmp",
          "/var/log"
        ]
      }
    }
  }
]`

var stringOfVul = `[{"id":"CVE-2024-2222","severity":"HIGH","cvss":8.2,"status":"active","package_name":"spring-security","current_version":"5.6.0","fixed_version":"5.6.1","description":"Authentication bypass in Spring Security","published_date":"2025-01-27T00:00:00Z","link":"https://nvd.nist.gov/vuln/detail/CVE-2024-2222","risk_factors":["Authentication Bypass","High CVSS Score","Proof of Concept Exploit Available"],"Metadata":{"scan_id":"","timestamp":"0001-01-01T00:00:00Z","scan_status":"","resource_type":"","resource_name":"","file_name":"testFile.json"}},
{"id":"CVE-2024-2223","severity":"MEDIUM","cvss":6.5,"status":"active","package_name":"tomcat","current_version":"9.0.50","fixed_version":"9.0.51","description":"Information disclosure in Apache Tomcat","published_date":"2025-01-28T00:00:00Z","link":"https://nvd.nist.gov/vuln/detail/CVE-2024-2223","risk_factors":["Information Disclosure","Medium CVSS Score"],"Metadata":{"scan_id":"","timestamp":"0001-01-01T00:00:00Z","scan_status":"","resource_type":"","resource_name":"","file_name":"testFile.json"}},
{"id":"CVE-2024-1111","severity":"CRITICAL","cvss":9.9,"status":"active","package_name":"openssl","current_version":"3.0.0","fixed_version":"3.0.1","description":"Critical buffer overflow in OpenSSL TLS handling","published_date":"2025-01-28T00:00:00Z","link":"https://nvd.nist.gov/vuln/detail/CVE-2024-1111","risk_factors":["Buffer Overflow","Critical CVSS Score","Public Exploit Available","Exploit in Wild"],"Metadata":{"scan_id":"","timestamp":"0001-01-01T00:00:00Z","scan_status":"","resource_type":"","resource_name":"","file_name":"testFile.json"}}
]`
