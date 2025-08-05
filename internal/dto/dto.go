package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auth DTOs
type BasicRegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type CompleteRegisterRequest struct {
	Email                        string `json:"email" validate:"required,email"`
	FirstName                    string `json:"fname" validate:"required,min=2,max=50"`
	LastName                     string `json:"lname" validate:"required,min=2,max=50"`
	Bio                          string `json:"bio"`
	DateOfBirth                  string `json:"date_of_birth"`
	Gender                       string `json:"gender" validate:"oneof=Male Female"`
	Member                       bool   `json:"member"`
	Visitor                      bool   `json:"visitor"`
	Usher                        bool   `json:"usher"`
	UserWorkDepartment           string `json:"user_work_unit"`
	DateJoinedChurch             string `json:"date_joined_church"`
	FamilyHead                   bool   `json:"family_head"`
	UserCampus                   string `json:"user_campus"`
	InstagramHandle              string `json:"instagram_handle"`
	FamilyMembers                []int  `json:"family_members"`
	PhoneNumber                  string `json:"phone_number"`
	Profession                   string `json:"profession"`
	UserHouseAddress             string `json:"user_house_address"`
	CampusState                  string `json:"campus_state"`
	CampusCountry                string `json:"campus_country"`
	EmergencyContactName         string `json:"emergency_contact_name"`
	EmergencyContactPhone        string `json:"emergency_contact_phone"`
	EmergencyContactEmail        string `json:"emergency_contact_email" validate:"omitempty,email"`
	EmergencyContactRelationship string `json:"emergency_contact_relationship"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type BasicRegisterResponse struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CompleteRegisterResponse struct {
	UserID                       string              `json:"user_id"`
	Email                        string              `json:"email"`
	FirstName                    string              `json:"FirstName"`
	LastName                     string              `json:"LastName"`
	Bio                          string              `json:"bio"`
	DateOfBirth                  time.Time           `json:"date_of_birth"`
	Gender                       string              `json:"gender"`
	Member                       bool                `json:"member"`
	Visitor                      bool                `json:"visitor"`
	Usher                        bool                `json:"usher"`
	QRCodeToken                  string              `json:"qr_code_token"`
	QRCodeImage                  string              `json:"qr_code_image"`
	Admin                        bool                `json:"admin"`
	UserWorkDepartment           string              `json:"user_work_unit"`
	DateJoinedChurch             time.Time           `json:"date_joined_church"`
	FamilyHead                   bool                `json:"family_head"`
	UserCampus                   string              `json:"user_campus"`
	InstagramHandle              string              `json:"instagram_handle"`
	FamilyMembers                []int               `json:"family_members"`
	PhoneNumber                  string              `json:"phone_number"`
	Profession                   string              `json:"profession"`
	UserHouseAddress             string              `json:"user_house_address"`
	CampusState                  string              `json:"campus_state"`
	CampusCountry                string              `json:"campus_country"`
	EmergencyContactName         string              `json:"emergency_contact_name"`
	EmergencyContactPhone        string              `json:"emergency_contact_phone"`
	EmergencyContactEmail        string              `json:"emergency_contact_email"`
	EmergencyContactRelationship string              `json:"emergency_contact_relationship"`
	Role                         *primitive.ObjectID `json:"role"`
	CreatedAt                    time.Time           `json:"date_joined"`
	UpdatedAt                    time.Time           `json:"date_updated"`
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         UserSummary `json:"user"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserSummary struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

// Attendance DTOs
type CreateAttendanceRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type QRCheckinRequest struct {
	QRCodeToken string `json:"qr_code_token" validate:"required"`
}

type AttendanceResponse struct {
	ID                   string    `json:"id"`
	UserID               string    `json:"user_id"`
	DateTimeOfAttendance time.Time `json:"date_time_of_attendance"`
	QRCodeBasedCheckin   bool      `json:"qrcode_based_checkin"`
	Late                 bool      `json:"late"`
	ManualCheckin        bool      `json:"manual_checkin"`
	Visitor              bool      `json:"visitor"`
	Member               bool      `json:"member"`
}

type AttendanceHistoryItem struct {
	Date            string `json:"date"`
	Members         int    `json:"members"`
	Visitors        int    `json:"visitors"`
	TotalAttendance int    `json:"total_attendance"`
}

type AttendanceAnalytics struct {
	TotalActiveUsersAllTime int `json:"total_active_users_all_time"`
	TotalAttendanceForDate  int `json:"total_attendance_for_date"`
	MembersForMonth         int `json:"members_for_month"`
	VisitorsCount           int `json:"visitors_count"`
}

// QR Code DTOs
type GenerateQRRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type QRCodeResponse struct {
	QRCodeToken string `json:"qr_code_token"`
	QRCodeImage string `json:"qr_code_image"`
}

// Role DTOs
type CreateRoleRequest struct {
	RoleName        string   `json:"role_name" validate:"required,min=2,max=50"`
	RoleDescription string   `json:"role_description" validate:"required,min=5,max=200"`
	Permissions     []string `json:"permissions" validate:"required"`
}

type UpdateRoleRequest struct {
	RoleName        string   `json:"role_name" validate:"omitempty,min=2,max=50"`
	RoleDescription string   `json:"role_description" validate:"omitempty,min=5,max=200"`
	Permissions     []string `json:"permissions" validate:"omitempty"`
}

type RoleResponse struct {
	ID              string    `json:"id"`
	RoleName        string    `json:"role_name"`
	RoleDescription string    `json:"role_description"`
	Permissions     []string  `json:"permissions"`
	TotalMembers    int       `json:"total_members"`
	DateAdded       time.Time `json:"date_added"`
	DateUpdated     time.Time `json:"date_updated"`
}

type RolesAndPermissionsResponse struct {
	Roles                []RoleResponse `json:"roles"`
	AvailablePermissions []string       `json:"available_permissions"`
}

type PaginatedRolesResponse struct {
	Data       []*RoleResponse `json:"data"`
	Pagination Pagination      `json:"pagination"`
}

// Sermon DTOs
type CreateSermonRequest struct {
	Title     string   `json:"title" validate:"required,min=5,max=200"`
	Speaker   string   `json:"speaker" validate:"required,min=2,max=100"`
	Date      string   `json:"date"`
	VideoURL  string   `json:"video_url"`
	AudioURL  string   `json:"audio_url"`
	Notes     string   `json:"notes"`
	Scripture string   `json:"scripture"`
	Series    string   `json:"series"`
	Tags      []string `json:"tags"`
}

type UpdateSermonRequest struct {
	Title     string   `json:"title" validate:"omitempty,min=5,max=200"`
	Speaker   string   `json:"speaker" validate:"omitempty,min=2,max=100"`
	Date      string   `json:"date"`
	VideoURL  string   `json:"video_url"`
	AudioURL  string   `json:"audio_url"`
	Notes     string   `json:"notes"`
	Scripture string   `json:"scripture"`
	Series    string   `json:"series"`
	Tags      []string `json:"tags"`
}

type SermonResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Speaker     string    `json:"speaker"`
	Date        string    `json:"date"`
	VideoURL    string    `json:"video_url"`
	AudioURL    string    `json:"audio_url"`
	Notes       string    `json:"notes"`
	Scripture   string    `json:"scripture"`
	Series      string    `json:"series"`
	Tags        []string  `json:"tags"`
	DateAdded   time.Time `json:"date_added"`
	DateUpdated time.Time `json:"date_updated"`
}

type PaginatedSermonsResponse struct {
	Data       []*SermonResponse `json:"data"`
	Pagination Pagination        `json:"pagination"`
}

// Announcement DTOs
type CreateAnnouncementRequest struct {
	Title               string   `json:"title" validate:"required,min=5,max=200"`
	AnnouncementContent string   `json:"content" validate:"required,min=10,max=2000"`
	AnnouncementDueDate string   `json:"announcement_due_date"`
	StartDate           string   `json:"start_date"`
	EndDate             string   `json:"end_date"`
	AnnouncementType    string   `json:"type" validate:"required,oneof=general event prayer urgent"`
	Priority            string   `json:"priority" validate:"required,oneof=low medium high"`
	TargetUsers         []string `json:"target_users"`
	ImageUrl            string   `json:"image_url"`
}

type UpdateAnnouncementRequest struct {
	Title               string   `json:"title" validate:"required,min=5,max=200"`
	AnnouncementContent string   `json:"content" validate:"required,min=10,max=2000"`
	AnnouncementDueDate string   `json:"announcement_due_date"`
	StartDate           string   `json:"start_date"`
	EndDate             string   `json:"end_date"`
	AnnouncementType    string   `json:"type" validate:"required,oneof=general event prayer urgent"`
	Priority            string   `json:"priority" validate:"required,oneof=low medium high"`
	TargetUsers         []string `json:"target_users"`
	ImageUrl            string   `json:"image_url"`
}

type AnnouncementResponse struct {
	ID                      string             `json:"id"`
	Title                   string             `json:"title"`
	AnnouncementContent     string             `json:"content"`
	AnnouncementType        string             `json:"type"`
	AnnouncementDueDate     time.Time          `json:"announcement_due_date"`
	StartDate               time.Time          `json:"start_date"`
	EndDate                 time.Time          `json:"end_date"`
	Priority                string             `json:"priority"`
	TargetUsers             []string           `json:"target_users"`
	ImageURL                string             `json:"image_url"`
	Status                  string             `json:"status"`
	DateAdded               time.Time          `json:"date_added"`
	DateUpdated             time.Time          `json:"date_updated"`
	AnnouncementEntryMadeBy primitive.ObjectID `json:"entry_made_by"`
}

type PaginatedAnnouncementsResponse struct {
	Data       []*AnnouncementResponse `json:"data"`
	Pagination Pagination              `json:"pagination"`
}

// Family Member DTOs
type CreateFamilyMemberRequest struct {
	FamilyMemberName         string `json:"name" validate:"required,min=2,max=100"`
	FamilyMemberEmail        string `json:"email" validate:"omitempty,email"`
	FamilyMemberRelationship string `json:"relationship" validate:"required"`
	FamilyMemberPhone        string `json:"phone_number" validate:"omitempty,e164"`
	FamilyMemberDateOfBirth  string `json:"date_of_birth"`
	FamilyMemberGender       string `json:"gender" validate:"omitempty,oneof=Male Female Other"`
	FamilyMemberOccupation   string `json:"occupation"`
}

type UpdateFamilyMemberRequest struct {
	FamilyMemberName         string `json:"name" validate:"required,min=2,max=100"`
	FamilyMemberEmail        string `json:"email" validate:"omitempty,email"`
	FamilyMemberRelationship string `json:"relationship" validate:"required"`
	FamilyMemberPhone        string `json:"phone_number" validate:"omitempty,e164"`
	FamilyMemberDateOfBirth  string `json:"date_of_birth"`
	FamilyMemberGender       string `json:"gender" validate:"omitempty,oneof=Male Female Other"`
	FamilyMemberOccupation   string `json:"occupation"`
}

type FamilyMemberResponse struct {
	ID                       string    `json:"id"`
	FamilyMemberName         string    `json:"family_member_name" validate:"required,min=2,max=100"`
	FamilyMemberEmail        string    `json:"family_member__email" validate:"omitempty,email"`
	FamilyMemberRelationship string    `json:"family_member__relationship" validate:"required"`
	FamilyMemberPhone        string    `json:"family_member_phone_number" validate:"omitempty,e164"`
	FamilyMemberDateOfBirth  time.Time `json:"family_member_date_of_birth"`
	FamilyMemberGender       string    `json:"family_member_gender" validate:"omitempty,oneof=Male Female Other"`
	FamilyMemberOccupation   string    `json:"family_member_occupation"`
	FamilyMemberHead         string    `json:"family_head"`
	DateAdded                time.Time `json:"date_added"`
}

type PaginatedFamilyMembersResponse struct {
	Data       []*FamilyMemberResponse `json:"data"`
	Pagination Pagination              `json:"pagination"`
}

// Local Church DTOs
type CreateLocalChurchRequest struct {
	ChurchName         string `json:"church_name" validate:"required,min=5,max=100"`
	ChurchPhone        string `json:"church_phone"`
	ChurchEmail        string `json:"church_email" validate:"required,email"`
	ChurchAddress      string `json:"church_address" validate:"required,min=4,max=200"`
	StateCounty        string `json:"state_county" validate:"required,min=2,max=50"`
	Country            string `json:"country" validate:"required,min=2,max=50"`
	SundayMeetingTime  int    `json:"sunday_meeting_time" validate:"required,min=0,max=23"`
	MidweekMeetingDay  string `json:"midweek_meeting_day" validate:"required,oneof=Monday Tuesday Wednesday Thursday Friday"`
	MidweekMeetingTime int    `json:"midweek_meeting_time" validate:"required,min=0,max=23"`
	Website            string `json:"website" validate:"omitempty,url"`
	SocialMedia        string `json:"social_media"`
	PastorName         string `json:"pastor_name" validate:"required,min=2,max=100"`
	PastorPhone        string `json:"pastor_phone"`
	PastorEmail        string `json:"pastor_email" validate:"required,email"`
	FoundedYear        int    `json:"founded_year" validate:"omitempty,min=1800,max=2500"`
	Description        string `json:"description"`
}

type UpdateLocalChurchRequest struct {
	ChurchName         string `json:"church_name" validate:"omitempty,min=5,max=100"`
	ChurchPhone        string `json:"church_phone"`
	ChurchEmail        string `json:"church_email" validate:"omitempty,email"`
	ChurchAddress      string `json:"church_address" validate:"omitempty,min=10,max=200"`
	StateCounty        string `json:"state_county" validate:"omitempty,min=2,max=50"`
	Country            string `json:"country" validate:"omitempty,min=2,max=50"`
	SundayMeetingTime  int    `json:"sunday_meeting_time" validate:"omitempty,min=0,max=23"`
	MidweekMeetingDay  string `json:"midweek_meeting_day" validate:"omitempty,oneof=Monday Tuesday Wednesday Thursday Friday"`
	MidweekMeetingTime int    `json:"midweek_meeting_time" validate:"omitempty,min=0,max=23"`
	Website            string `json:"website" validate:"omitempty,url"`
	SocialMedia        string `json:"social_media"`
	PastorName         string `json:"pastor_name" validate:"required,min=2,max=100"`
	PastorPhone        string `json:"pastor_phone"`
	PastorEmail        string `json:"pastor_email" validate:"required,email"`
	FoundedYear        int    `json:"founded_year" validate:"omitempty,min=1800,max=2500"`
	Description        string `json:"description"`
}

type LocalChurchResponse struct {
	ID                 string    `json:"id"`
	ChurchName         string    `json:"church_name"`
	ChurchPhone        string    `json:"church_phone"`
	ChurchEmail        string    `json:"church_email"`
	ChurchAddress      string    `json:"church_address"`
	StateCounty        string    `json:"state_county"`
	Country            string    `json:"country"`
	SundayMeetingTime  int       `json:"sunday_meeting_time"`
	MidweekMeetingDay  string    `json:"midweek_meeting_day"`
	MidweekMeetingTime int       `json:"midweek_meeting_time"`
	Website            string    `json:"website"`
	SocialMedia        string    `json:"social_media"`
	PastorName         string    `json:"pastor_name"`
	PastorPhone        string    `json:"pastor_phone"`
	PastorEmail        string    `json:"pastor_email"`
	FoundedYear        int       `json:"founded_year"`
	Description        string    `json:"description"`
	DateAdded          time.Time `json:"date_added"`
	DateUpdated        time.Time `json:"date_updated"`
}

type PaginatedLocalChurchesResponse struct {
	Data       []*LocalChurchResponse `json:"data"`
	Pagination Pagination             `json:"pagination"`
}

// The one being used in the set password endpoint
type SetPasswordRequest struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreatePasswordResponse struct {
	UserID      string    `json:"user_id"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_joined"`
	DateUpdated time.Time `json:"date_updated"`
}

type ResetPasswordRequest struct {
	Token           string `json:"token"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

// Generic Response DTOs
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type SuccessResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
type WorkUnit struct {
	Worship   string `json:"worship_team"`
	Ushering  string `json:"ushering"`
	Protocol  string `json:"protocol"`
	Media     string `json:"media"`
	Pastorate string `json:"pastorate"`
}
