package common

// Reference: https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	Id            *string `json:"id"`                       // emoji id
	Name          *string `json:"name"`                     // emoji name
	Roles         *[]Role `json:"roles,omitempty"`          // roles allowed to use this emoji
	User          *User   `json:"user,omitempty"`           // user that created this emoji
	RequireColons *bool   `json:"require_colons,omitempty"` // whether this emoji must be wrapped in colons
	Managed       *bool   `json:"managed,omitempty"`        // whether this emoji is managed
	Animated      *bool   `json:"animated,omitempty"`       // whether this emoji is animated
	Available     *bool   `json:"available,omitempty"`      // whether this emoji can be used, may be false due to loss of Server Boosts
}
