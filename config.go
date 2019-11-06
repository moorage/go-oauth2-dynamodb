package oauth2dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Config dynamodb configuration parameters
type Config struct {
	SESSION *session.Session
	TABLE   *TableConfig
}

// TableConfig internally stores the table names to use for dynamodb
type TableConfig struct {
	BasicCname   string
	AccessCName  string
	RefreshCName string
}

// NewConfig create dynamodb configuration
func NewConfig(awsConfig *aws.Config, basicTableName string, accessTableName string, refreshTablName string) (*Config, error) {
	newSession, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	config := &Config{
		SESSION: newSession,
		TABLE: &TableConfig{
			BasicCname:   basicTableName,
			AccessCName:  accessTableName,
			RefreshCName: refreshTablName,
		},
	}
	return config, nil
}
