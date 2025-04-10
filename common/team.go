package common

// External reference: https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        *string      `json:"icon"`          // Hash of the image of the team's icon
	Id          string       `json:"id"`            // Unique ID of the team
	Members     []TeamMember `json:"members"`       // Members of the team
	Name        string       `json:"name"`          // Name of the team
	OwnerUserId string       `json:"owner_user_id"` // User ID of the current team owner
}

// External reference: https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState int    `json:"membership_state"` // User's membership state on the team
	TeamId          string `json:"team_id"`          // ID of the parent team of which they are a member
	User            User   `json:"user"`             // Avatar, discriminator, ID, and username of the user
	Role            string `json:"role"`             // Role of the team member
}

// External reference: https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
var MembershipStates map[string]int = map[string]int{
	"INVITED":  1,
	"ACCEPTED": 2,
}
