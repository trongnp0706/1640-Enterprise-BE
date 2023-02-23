package otp

import (
    "github.com/twilio/twilio-go"
    twilioApi "github.com/twilio/twilio-go/rest/verify/v2"

	"os"
)



var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
    Username: os.Getenv("TWILIO_ACCOUNT_SID"),
    Password: os.Getenv("TWILIO_AUTHTOKEN"),
})

func TwilioSendOTP(email string) (string, error) {
    params := &twilioApi.CreateVerificationParams{}
    params.SetTo(email)
    params.SetChannel("email")

    resp, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICES_ID"), params)
    if err != nil {
        return "", err
    }

    return *resp.Sid, nil
}

func TwilioVerifyOTP(email string, code string) error {
    params := &twilioApi.CreateVerificationCheckParams{}
    params.SetTo(email)
    params.SetCode(code)

    resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICES_ID"), params)
    if err != nil {
        return err
    } else if *resp.Status == "approved" {
        return nil
    }

    return nil
}