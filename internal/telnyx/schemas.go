package telnyx

import "fmt"

type Verification struct {
	Id              string  `json:"id"`
	Type            string  `json:"type"`
	RecordType      string  `json:"record_type"`
	PhoneNumber     string  `json:"phone_number"`
	VerifyProfileId string  `json:"verify_profile_id"`
	CustomCode      *string `json:"custom_code"`
	TimeoutSecs     int     `json:"timeout_secs"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type VerificationResponse struct {
	Data Verification `json:"data"`
}

type VerificationProfile struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	WebhookUrl         string `json:"webhook_url"`
	WebhookFailoverUrl string `json:"webhook_failover_url"`
	RecordType         string `json:"record_type"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Sms                struct {
		MessagingTemplateId            string   `json:"messaging_template_id"`
		AppName                        string   `json:"app_name"`
		AlphaSender                    string   `json:"alpha_sender"`
		CodeLength                     int      `json:"code_length"`
		WhitelistedDestinations        []string `json:"white_listed_destinations"`
		DefaultTimeoutVerificationSecs int      `json:"default_timeout_verification_secs"`
	} `json:"sms"`
	Call struct {
		MessagingTemplateId            string   `json:"messaging_template_id"`
		AppName                        string   `json:"app_name"`
		CodeLength                     int      `json:"code_length"`
		WhitelistedDestinations        []string `json:"white_listed_destinations"`
		DefaultTimeoutVerificationSecs int      `json:"default_timeout_verification_secs"`
	} `json:"call"`
	FlashCall struct {
		DefaultTimeoutVerificationSecs int `json:"default_timeout_verification_secs"`
	} `json:"flash_call"`
	Language string `json:"language"`
}

type VerificationProfileResponse struct {
	Data VerificationProfile `json:"data"`
}

type VerificationProfileListResponse struct {
	Data []VerificationProfile `json:"data"`
	Meta ListMeta
}

type ListMeta struct {
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
}

type TelnyxError struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source struct {
		Pointer   string `json:"pointer"`
		Parameter string `json:"parameter"`
	}
	Meta struct {
		Url string `json:"url"`
	} `json:"meta"`
}

func (e *TelnyxError) Error() string {
	return fmt.Sprintf("Telnyx Error: %d '%s' '%s'", e.Code, e.Title, e.Detail)
}

type TelnyxErrorResponse struct {
	Errors []TelnyxError `json:"errors"`
}
