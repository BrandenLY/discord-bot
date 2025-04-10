package common

// External Reference: https://discord.com/developers/docs/resources/user#user-object
type User struct {
	Id                   string                `json:"id"`                               // The user's id
	Username             string                `json:"username"`                         // The user's username, not unique across the platform
	Discriminator        string                `json:"discriminator"`                    // The user's Discord-tag
	GlobalName           *string               `json:"global_name"`                      // The user's display name, if it is set. For bots, this is the application name
	Avatar               *string               `json:"avatar"`                           // The user's avatar hash
	Bot                  *bool                 `json:"bot,omitempty"`                    // Whether the user belongs to an OAuth2 application
	System               *bool                 `json:"system,omitempty"`                 // Whether the user is an Official Discord System user (part of the urgent message system)
	MfaEnabled           *bool                 `json:"mfa_enabled,omitempty"`            // Whether the user has two factor enabled on their account
	Banner               *string               `json:"banner,omitempty"`                 // The user's banner hash
	AccentColor          *uint                 `json:"accent_color,omitempty"`           // The user's banner color encoded as an integer representation of hexadecimal color code
	Locale               *string               `json:"locale,omitempty"`                 // The user's chosen language option
	Verified             *bool                 `json:"verified,omitempty"`               // Whether the email on this account has been verified
	Email                *string               `json:"email,omitempty"`                  // The user's email
	Flags                *uint                 `json:"flags,omitempty"`                  // The flags on a user's account
	PremiumType          *uint                 `json:"premium_type,omitempty"`           // The type of Nitro subscription on a user's account
	PublicFlags          *uint                 `json:"public_flags,omitempty"`           // The public flags on a user's account
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"` // Data for the user's avatar decoration
}
