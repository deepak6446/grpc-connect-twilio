package database

import (
	"errors"
	"sync"

	Authv1 "auth-service/gen/auth/v1" // generated by protoc-gen-go
)

// SignupData represents signup data along with OTP
type SignupData struct {
	PhoneNumber string
	Email       string
	FirstName   string
	LastName    string
	OTP         string
	IsVerified  bool
}

var (
	dataMap  = make(map[string]SignupData)
	dataLock sync.Mutex
)

// SaveSignupData saves signup data along with OTP
func SaveSignupData(signupRequest *Authv1.SignupRequest, otp string) error {
	if signupRequest.PhoneNumber == "" {
		return errors.New("phone number is mandatory")
	}
	dataLock.Lock()
	defer dataLock.Unlock()
	dataMap[signupRequest.PhoneNumber] = SignupData{
		PhoneNumber: signupRequest.PhoneNumber,
		Email:       signupRequest.Email,
		FirstName:   signupRequest.FirstName,
		LastName:    signupRequest.LastName,
		OTP:         otp,
		IsVerified:  false,
	}
	return nil
}

// GetProfileByPhoneNumber retrieves signup data by phone number
func GetProfileByPhoneNumber(phoneNumber string) (*SignupData, error) {
	dataLock.Lock()
	defer dataLock.Unlock()
	signupData, ok := dataMap[phoneNumber]
	if !ok {
		return nil, errors.New("profile data not found")
	}
	return &signupData, nil
}

// ValidateOTP validates OTP for a given phone number
func ValidateOTP(phoneNumber, otp string) (bool, error) {
	dataLock.Lock()
	defer dataLock.Unlock()
	storedSignupData, ok := dataMap[phoneNumber]
	if !ok || storedSignupData.OTP != otp {
		return false, errors.New("profile data not found")
	}

	return true, nil
}

// UpdateAccountVerification marks user account as verified
func UpdateAccountVerification(phoneNumber string) error {
	dataLock.Lock()
	defer dataLock.Unlock()
	user, ok := dataMap[phoneNumber]
	if !ok {
		return errors.New("user account not found")
	}
	user.IsVerified = true
	dataMap[phoneNumber] = user
	return nil
}