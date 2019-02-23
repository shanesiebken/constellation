package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sirupsen/logrus"
)

var (
	svc *ec2.EC2
)

// EC2InstanceConfig represents the Constellation
// abstraction of an AWS ec2 instance, which will
// be transformed into an AWS SDK configuration object
// and validated at creation time
type EC2InstanceConfig struct {
	AmiId          string   `yaml:"ami"`
	securityGroups []string `yaml:`
}

// InitEC2Service creates an AWS EC2 api client
// for creating and managing EC2 resources
func InitEC2Service(sess *session.Session) {
	svc = ec2.New(sess)
}

func CreateEC2(config EC2InstanceConfig) {
	rii, err := runInstancesInput(config)
	if err != nil {
		logrus.Errorf("Error creating EC2 instance: %s", err)
	}
	svc.RunInstances(rii)
}

func runInstancesInput(c EC2InstanceConfig) (*ec2.RunInstancesInput, error) {
	rii := &ec2.RunInstancesInput{}
	rii = rii.SetImageId(c.amiID)
	rii.SetSecurityGroupIds()
	if err != nil {
		return nil, err
	}
	return rii, nil
}
