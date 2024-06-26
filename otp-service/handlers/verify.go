package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	Otpv1 "otp-service/gen/otp/v1" // generated by protoc-gen-go

	"connectrpc.com/connect"
	"github.com/sfreiberg/gotwilio" // Import Twilio package
)

// VerifyPhoneNumber verifies user account using phone number and sends OTP via Twilio
func (s *OtpServer) VerifyPhoneNumber(
	ctx context.Context,
	req *connect.Request[Otpv1.VerifyPhoneNumberRequest],
) (*connect.Response[Otpv1.VerifyPhoneNumberResponse], error) {
	// Log request received
	fmt.Println("Received VerifyPhoneNumber request")

	// Load environment variables
	LoadEnv()

	phoneNumber := req.Msg.PhoneNumber

	// Generate OTP
	otp := generateOTP()
	fmt.Println("Generated OTP:" + otp)

	// Get Twilio API credentials
	twilioConfig := GetTwilioAPIConfig()

	// Send OTP via Twilio
	err := sendOTP(twilioConfig, phoneNumber, otp)
	if err != nil {
		fmt.Println("Error sending OTP:", err)
		return nil, err
	}

	// Log OTP sent
	fmt.Println("OTP sent to:" + phoneNumber)

	// Return success response
	res := connect.NewResponse(&Otpv1.VerifyPhoneNumberResponse{
		Message: "OTP sent successfully",
		Otp:     otp,
	})
	// Log response sent
	fmt.Println("Sent VerifyPhoneNumber response")
	return res, nil
}

// generateOTP generates a random OTP
func generateOTP() string {
	// For simplicity, generating a 4-digit OTP
	otp := strconv.Itoa(rand.Intn(10000))
	return otp
}

// sendOTP sends OTP via Twilio
func sendOTP(twilioConfig *TwilioAPIConfig, phoneNumber, otp string) error {
	// Initialize Twilio client
	twilio := gotwilio.NewTwilioClient(twilioConfig.AccountSID, twilioConfig.AuthToken)

	// Send SMS message with OTP
	_, _, err := twilio.SendSMS(twilioConfig.FromNumber, phoneNumber, "Your OTP: "+otp, "", "")
	if err != nil {
		return err
	}

	return nil
}
