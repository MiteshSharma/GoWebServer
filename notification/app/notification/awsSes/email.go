package awsSes

import (
	"github.com/MiteshSharma/project/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AwsSes struct {
}

func New() *AwsSes {
	awsSes := &AwsSes{}
	return awsSes
}

func (s AwsSes) Send(to string, message model.NotificationMessage) error {
	aSes, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	creds := credentials.NewStaticCredentials("", "", "")

	sesClient := ses.New(aSes, &aws.Config{Credentials: creds})

	messageContent := &ses.Content{
		Charset: aws.String("UTF-8"),
		Data:    aws.String(message.Message),
	}
	var body *ses.Body
	if message.Type == "html" {
		body = &ses.Body{
			Html: messageContent,
		}
	} else {
		body = &ses.Body{
			Text: messageContent,
		}
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: body,
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(message.Title),
			},
		},
		Source: aws.String("abc@abc.com"),
	}
	_, err = sesClient.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				break
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				break
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				break
			default:
				break
			}
		} else {
		}

		return err
	}
	return err
}
