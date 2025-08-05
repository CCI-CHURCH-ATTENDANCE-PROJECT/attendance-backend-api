package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user model
type User struct {
	ID                           primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID                       string              `bson:"user_id" json:"user_id" validate:"required"`
	FirstName                    string              `bson:"fname" json:"fname" validate:"required,min=2,max=50"`
	LastName                     string              `bson:"lname" json:"lname" validate:"required,min=2,max=50"`
	Email                        string              `bson:"email" json:"email" validate:"required,email"`
	Password                     string              `bson:"user_password" json:"-" validate:"required,min=8"`
	Bio                          string              `bson:"bio" json:"bio"`
	DateOfBirth                  time.Time           `bson:"date_of_birth" json:"date_of_birth"`
	Gender                       string              `bson:"gender" json:"gender" validate:"oneof=Male Female Other"`
	Member                       bool                `bson:"member" json:"member"`
	Visitor                      bool                `bson:"visitor" json:"visitor"`
	Usher                        bool                `bson:"usher" json:"usher"`
	Admin                        bool                `bson:"admin" json:"admin"`
	UserWorkDepartment           string              `bson:"user_work_department,omitempty" json:"user_work_department"`
	DateJoinedChurch             time.Time           `bson:"date_joined_church" json:"date_joined_church"`
	QRCodeToken                  string              `bson:"qr_code_token" json:"qr_code_token"`
	QRCodeImage                  string              `bson:"qr_code_image" json:"qr_code_image"`
	FamilyHead                   bool                `bson:"family_head" json:"family_head"`
	UserCampus                   string              `bson:"user_campus" json:"user_campus"`
	CampusState                  string              `bson:"campus_state" json:"campus_state"`
	CampusCountry                string              `bson:"campus_country" json:"campus_country"`
	Profession                   string              `bson:"profession" json:"profession"`
	UserHouseAddress             string              `bson:"user_house_address" json:"user_house_address"`
	PhoneNumber                  string              `bson:"phone_number" json:"phone_number" validate:"omitempty,e164"`
	InstagramHandle              string              `bson:"instagram_handle" json:"instagram_handle"`
	FamilyMembers                []string            `bson:"family_member_id" json:"family_member_id"`
	DateJoined                   time.Time           `bson:"date_joined" json:"date_joined"`
	DateUpdated                  time.Time           `bson:"date_updated" json:"date_updated"`
	Role                         *primitive.ObjectID `bson:"role,omitempty" json:"role"`
	EmergencyContactName         string              `bson:"emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactPhone        string              `bson:"emergency_contact_phone" json:"emergency_contact_phone"`
	EmergencyContactEmail        string              `bson:"emergency_contact_email" json:"emergency_contact_email" validate:"omitempty,email"`
	EmergencyContactRelationship string              `bson:"emergency_contact_relationship" json:"emergency_contact_relationship"`
	PasswordResetToken           string              `bson:"password_reset_token,omitempty" json:"-"`
	PasswordResetExpires         time.Time           `bson:"password_reset_expires,omitempty" json:"-"`
}

// UserResponse represents user data for API responses (without sensitive data)
type UserResponse struct {
	ID                           primitive.ObjectID  `json:"id"`
	UserID                       string              `json:"user_id"`
	FirstName                    string              `json:"fname"`
	LastName                     string              `json:"lname"`
	Email                        string              `json:"email"`
	Bio                          string              `json:"bio"`
	DateOfBirth                  time.Time           `json:"date_of_birth"`
	Gender                       string              `json:"gender"`
	Member                       bool                `json:"member"`
	Visitor                      bool                `json:"visitor"`
	Usher                        bool                `json:"usher"`
	UserWorkDepartment           string              `json:"user_work_department"`
	DateJoinedChurch             time.Time           `json:"date_joined_church"`
	QRCodeToken                  string              `json:"qr_code_token"`
	QRCodeImage                  string              `json:"qr_code_image"`
	FamilyHead                   bool                `json:"family_head"`
	UserCampus                   string              `json:"user_campus"`
	CampusState                  string              `json:"campus_state"`
	CampusCountry                string              `json:"campus_country"`
	Profession                   string              `json:"profession"`
	UserHouseAddress             string              `json:"user_house_address"`
	PhoneNumber                  string              `json:"phone_number"`
	InstagramHandle              string              `json:"instagram_handle"`
	FamilyMembers                []string            `json:"family_member_id"`
	DateJoined                   time.Time           `json:"date_joined"`
	DateUpdated                  time.Time           `json:"date_updated"`
	Role                         *primitive.ObjectID `json:"role"`
	EmergencyContactName         string              `json:"emergency_contact_name"`
	EmergencyContactPhone        string              `json:"emergency_contact_phone"`
	EmergencyContactEmail        string              `json:"emergency_contact_email"`
	EmergencyContactRelationship string              `json:"emergency_contact_relationship"`
}

// Attendance represents the attendance model
type Attendance struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User                 primitive.ObjectID `bson:"user" json:"user" validate:"required"`
	DateTimeOfAttendance time.Time          `bson:"date_time_of_attendance" json:"date_time_of_attendance"`
	QRCodeBasedCheckin   bool               `bson:"qrcode_based_checkin" json:"qrcode_based_checkin"`
	Late                 bool               `bson:"late" json:"late"`
	ManualCheckin        bool               `bson:"manual_checkin" json:"manual_checkin"`
	Visitor              bool               `bson:"visitor,omitempty" json:"visitor"`
	Member               bool               `bson:"member,omitempty" json:"member"`
}

// FamilyMember represents the family member model
type FamilyMember struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FamilyHead               string             `bson:"family_head" json:"family_head" validate:"required"`
	FamilyMemberName         string             `bson:"family_members" json:"family_members" validate:"required"`
	DateJoined               time.Time          `bson:"date_joined" json:"date_joined"`
	FamilyMemberPhone        string             `bson:"phone" json:"phone" validate:"required,e164"`
	FamilyMemberEmail        string             `bson:"email" json:"email" validate:"required,email"`
	FamilyMemberRelationship string             `bson:"relationship" json:"relationship" validate:"required"`
	FamilyMemberDateOfBirth  time.Time          `bson:"date_of_birth" json:"date_of_birth" validate:"required"`
	FamilyMemberGender       string             `bson:"gender" json:"gender" validate:"oneof=Male Female Other"`
	FamilyMemberOccupation   string             `bson:"occupation" json:"occupation"`
	DateAdded                time.Time          `bson:"date_added" json:"date_added"`
}

// Sermon represents the sermon model
type Sermon struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Preacher      string             `bson:"preacher" json:"preacher" validate:"required,min=2,max=100"`
	DateOfMeeting string             `bson:"date_of_meeting" json:"date_of_meeting"`
	SermonTopic   string             `bson:"sermon_topic" json:"sermon_topic" validate:"required,min=5,max=200"`
	SermonNote    string             `bson:"sermon_note" json:"sermon_note" validate:"required,min=10"`
	EntryMadeBy   primitive.ObjectID `bson:"entry_made_by" json:"entry_made_by" validate:"required"`
	VideoUrl      string             `bson:"video_url" json:"video_url"`
	AudioUrl      string             `bson:"audio_url" json:"audio_url"`
	Scripture     string             `bson:"scripture" json:"scripture"`
	Series        string             `bson:"series" json:"series"`
	Tags          []string           `bson:"tags" json:"tags"`
	DateAdded     time.Time          `bson:"date_added" json:"date_added"`
	DateUpdated   time.Time          `bson:"date_updated" json:"date_updated"`
}

// Announcement represents the announcement model
type Announcement struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title                   string             `bson:"title" json:"title" validate:"required,min=5,max=200"`
	AnnouncementContent     string             `bson:"announcement_content" json:"announcement_content" validate:"required,min=10"`
	AnnouncementDueDate     time.Time          `bson:"announcement_due_date" json:"announcement_due_date"`
	StartDate               time.Time          `bson:"start_date" json:"start_date"`
	EndDate                 time.Time          `bson:"end_date" json:"end_date"`
	AnnouncementEntryMadeBy primitive.ObjectID `bson:"announcement_entry_made_by" json:"announcement_entry_made_by" validate:"required"`
	AnnouncementType        string             `bson:"type" json:"type"`
	Priority                string             `bson:"priority" json:"priority"`
	TargetUsers             []string           `bson:"target_users" json:"target_users"`
	ImageUrl                string             `bson:"image_url" json:"image_url"`
	Status                  string             `bson:"status" json:"status" validate:"oneof=Pending Done"`
	DateAdded               time.Time          `bson:"date_added" json:"date_added"`
	DateUpdated             time.Time          `bson:"date_updated" json:"date_updated"`
}

// Department represents the department model
type Department struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string             `bson:"name" json:"name" validate:"required,min=2,max=100"`
	TotalMembers      int                `bson:"total_members" json:"total_members"`
	ActiveRate        int                `bson:"active_rate" json:"active_rate"`
	MonthlyAttendance int                `bson:"monthly_attendance" json:"monthly_attendance"`
	AverageAttendance int                `bson:"average_attendance" json:"average_attendance"`
	Performance       string             `bson:"performance" json:"performance" validate:"oneof=GOOD POOR EXCELLENT NEEDS_ATTENTION UNDECIDED"`
}

// Role represents the role model
type Role struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RoleName        string             `bson:"role_name" json:"role_name" validate:"required,min=2,max=50"`
	RoleDescription string             `bson:"role_description" json:"role_description" validate:"required,min=5,max=200"`
	TotalMembers    int                `bson:"total_members" json:"total_members"`
	Permissions     []string           `bson:"permissions" json:"permissions" validate:"required"`
	DateAdded       time.Time          `bson:"date_added" json:"date_added"`
	DateUpdated     time.Time          `bson:"date_updated" json:"date_updated"`
}

type RolesAndPermissions struct {
	Roles                []Role   `json:"roles"`
	AvailablePermissions []string `json:"available_permissions"`
}

// LocalChurch represents the local church model
type LocalChurch struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChurchName         string             `bson:"church_name" json:"church_name" validate:"required,min=2,max=100"`
	ChurchPhone        string             `bson:"church_phone" json:"church_phone"`
	ChurchEmail        string             `bson:"church_email" json:"church_email" validate:"required,email"`
	ChurchAddress      string             `bson:"church_address" json:"church_address" validate:"required,min=5,max=200"`
	StateCounty        string             `bson:"state_county" json:"state_county" validate:"required,min=2,max=50"`
	Country            string             `bson:"country" json:"country" validate:"required,min=2,max=50"`
	SundayMeetingTime  int                `bson:"sunday_meeting_time" json:"sunday_meeting_time" validate:"required,min=0,max=23"`
	MidweekMeetingDay  string             `bson:"midweek_meeting_day" json:"midweek_meeting_day" validate:"required,oneof=Monday Tuesday Wednesday Thursday Friday"`
	MidweekMeetingTime int                `bson:"midweek_meeting_time" json:"midweek_meeting_time" validate:"required,min=0,max=23"`
	Website            string             `bson:"website" json:"website" validate:"omitempty,url"`
	SocialMedia        string             `bson:"social_media" json:"social_media"`
	PastorName         string             `bson:"pastor_name" json:"pastor_name" validate:"required,min=2,max=100"`
	PastorPhone        string             `bson:"pastor_phone" json:"pastor_phone"`
	PastorEmail        string             `bson:"pastor_email" json:"pastor_email" validate:"required,email"`
	FoundedYear        int                `bson:"founded_year" json:"founded_year" validate:"omitempty,min=1800,max=2500"`
	Description        string             `bson:"description" json:"description"`
	DateAdded          time.Time          `bson:"date_added" json:"date_added"`
	DateUpdated        time.Time          `bson:"date_updated" json:"date_updated"`
}

// Notification represents the notification model
type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" validate:"required,min=5,max=100"`
	UserCreated primitive.ObjectID `bson:"user_created" json:"user_created" validate:"required"`
	Body        string             `bson:"body" json:"body" validate:"required,min=10"`
	Email       bool               `bson:"email" json:"email"`
	SMS         bool               `bson:"sms" json:"sms"`
	Event       bool               `bson:"event" json:"event"`
	Newsletter  bool               `bson:"newsletter" json:"newsletter"`
	DateCreated time.Time          `bson:"date_created" json:"date_created"`
}

// UserNotification represents the user notification model
type UserNotification struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User          primitive.ObjectID `bson:"user" json:"user" validate:"required"`
	Title         string             `bson:"title" json:"title" validate:"required,min=5,max=100"`
	Body          string             `bson:"body" json:"body" validate:"required,min=10"`
	Email         bool               `bson:"email" json:"email"`
	SMS           bool               `bson:"sms" json:"sms"`
	Event         bool               `bson:"event" json:"event"`
	Newsletter    bool               `bson:"newsletter" json:"newsletter"`
	DateDelivered time.Time          `bson:"date_delivered" json:"date_delivered"`
}

// RefreshToken represents a refresh token
type RefreshToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// Pagination represents pagination information
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type Permissions struct {
	CanViewDashboard string `json:"can_view_dashboard"`
	CanCreateUser    string `json:"can_create_user"`
	CanEditUser      string `json:"can_edit_user"`
	CanDeleteUser    string `json:"can_delete_user"`
	CanViewAnalytics string `json:"can_view_analytics"`
}

type WorkUnit struct {
	Worship   string `json:"worship_team"`
	Ushering  string `json:"ushering"`
	Protocol  string `json:"protocol"`
	Media     string `json:"media"`
	Pastorate string `json:"pastorate"`
}
