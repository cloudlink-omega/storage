package types

import "time"

const (
	LogGeneric uint8 = 0
	LogDebug   uint8 = 1
	LogInfo    uint8 = 2
	LogWarn    uint8 = 3
	LogError   uint8 = 4
	LogFatal   uint8 = 5
	LogTrace   uint8 = 6
)

type Event struct {
	ID          string `gorm:"primaryKey;type:varchar(100);unique;not null"`
	Description string `gorm:"type:varchar(255);not null;default:''"`
	LogLevel    uint8  `gorm:"not null,default:0"`

	UserEvent      *UserEvent
	DeveloperEvent *DeveloperEvent
	SystemEvent    *SystemEvent
}

type SystemEvent struct {
	ID         string `gorm:"primaryKey;type:varchar(26);unique;not null"`
	EventID    string `gorm:"foreignKey:EventID;not null;index:idx_user_event_event_id"`
	Details    string `gorm:"type:varchar(255);not null"`
	Successful bool   `gorm:"not null;default:false"`
	CreatedAt  time.Time

	Event Event `gorm:"constraint:OnDelete:CASCADE;"`
}

var SystemEvents map[string][]any = map[string][]any{
	"secret_gen_error": {"Failed to generate user secret", LogError},
	"get_user_error":   {"Failed to get user", LogError},
	"email_off":        {"Email services are nonfunctional", LogWarn},
	"totp_error":       {"TOTP validator failure", LogError},
}

type UserEvent struct {
	ID         string `gorm:"primaryKey;type:varchar(26);unique;not null"`
	UserID     string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_event_user_id"`
	EventID    string `gorm:"foreignKey:EventID;not null;index:idx_user_event_event_id"`
	Details    string `gorm:"type:varchar(255);not null"`
	Successful bool   `gorm:"not null;default:false"`
	CreatedAt  time.Time

	User  User  `gorm:"constraint:OnDelete:NO ACTION;"`
	Event Event `gorm:"constraint:OnDelete:CASCADE;"`
}

// Define the events used for logging user activity
var UserEvents map[string][]any = map[string][]any{
	"user_created": {"User was successfully created", LogInfo},
	"user_deleted": {"User was successfully deleted", LogInfo},
	"user_error":   {"User error", LogError},

	"user_login":  {"User was successfully logged in", LogInfo},
	"user_logout": {"User was successfully logged out", LogInfo},

	"user_session_created": {"User session was successfully created", LogInfo},
	"user_session_deleted": {"User session was successfully deleted", LogInfo},
	"user_session_error":   {"User session error", LogError},

	"developer_member_created": {"Developer member was successfully created", LogInfo},
	"developer_member_deleted": {"Developer member was successfully deleted", LogInfo},

	"game_save_created": {"Game save was successfully created", LogInfo},
	"game_save_deleted": {"Game save was successfully deleted", LogInfo},
	"game_save_error":   {"Game save error", LogError},

	"user_auth_password_error": {"Password authentication error", LogError},

	"user_totp_enroll_started": {"User started TOTP enrollment", LogInfo},
	"user_totp_enroll_success": {"User successfully enrolled in TOTP", LogInfo},
	"user_totp_enroll_failure": {"User failed to enroll in TOTP", LogError},
	"user_auth_totp_error":     {"TOTP authentication error", LogError},

	"user_verify_sent":              {"User verification code was sent", LogInfo},
	"user_verify_bypassed_test":     {"User verification was bypassed; testing mode enabled", LogWarn},
	"user_verify_bypassed_disabled": {"User verification was bypassed; email verification is disabled", LogWarn},
	"user_verify_success":           {"User successfully verified", LogInfo},
	"user_verify_failure":           {"User failed to verify", LogError},

	"user_recovery_set":         {"Recovery codes were generated", LogInfo},
	"user_recovery_set_failure": {"Recovery code setup error", LogError},
	"user_recovery_success":     {"User used recovery code", LogInfo},
	"user_recovery_failure":     {"Error using recovery code", LogError},

	"user_password_reset_sent":     {"User password reset code was sent", LogInfo},
	"user_password_reset_verified": {"User password reset code was verified", LogInfo},
	"user_password_reset_success":  {"User successfully reset password", LogInfo},
	"user_password_reset_failure":  {"Error while resetting password", LogError},

	"recovery_code_retrieval_error": {"Failed to retrieve user recovery codes", LogError},
	"recovery_code_store_error":     {"Failed to store user recovery codes", LogError},
}

type DeveloperEvent struct {
	ID          string `gorm:"primaryKey;type:varchar(26);unique;not null"`
	DeveloperID string `gorm:"foreignKey:DeveloperID;type:char(26);not null;index:idx_developer_event_developer_id"`
	EventID     string `gorm:"foreignKey:EventID;not null;index:idx_user_event_event_id"`
	Details     string `gorm:"type:varchar(255);not null"`
	Successful  bool   `gorm:"not null;default:false"`
	CreatedAt   time.Time

	Developer Developer `gorm:"constraint:OnDelete:NO ACTION;"`
	Event     Event     `gorm:"constraint:OnDelete:CASCADE;"`
}

// Define the events used for logging developer activity
var DeveloperEvents map[string][]any = map[string][]any{
	"developer_created": {"Developer account was created", LogInfo},
	"developer_deleted": {"Developer account was deleted", LogInfo},

	"developer_owner_change": {"Owner of developer account was changed", LogInfo},

	"developer_approval_start":   {"Developer account approval was started", LogInfo},
	"developer_approval_success": {"Developer account was approved", LogInfo},
	"developer_approval_deny":    {"Developer account was denied", LogInfo},
	"developer_approval_failure": {"Developer account approval failed", LogError},
}
