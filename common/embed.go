package common

import "time"

var EmbedTypes []string = []string{
	"rich",
	"image",
	"video",
	"gifv",
	"article",
	"link",
	"poll_request",
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-structure
type Embed struct {
	Title       string          `json:"title"`               // title of embed
	Type        string          `json:"type"`                // type of embed (always "rich" for webhook embeds)
	Description string          `json:"description"`         // description of embed
	Url         string          `json:"url"`                 // url of embed
	Timestamp   *time.Time      `json:"timestamp,omitempty"` // timestamp of embed content
	Color       int             `json:"color"`               // color code of the embed
	Footer      *EmbedFooter    `json:"footer,omitempty"`    // footer information
	Image       *EmbedImage     `json:"image,omitempty"`     // image information
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"` // thumbnail information
	Video       *EmbedVideo     `json:"video,omitempty"`     // video information
	Provider    *EmbedProvider  `json:"provider,omitempty"`  // provider information
	Author      *EmbedAuthor    `json:"author,omitempty"`    // author information
	Fields      *[]EmbedField   `json:"fields,omitempty"`    // fields information, max of 25
}

type PollResultsEmbed struct {
	Embed
	PollQuestionText string
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         *string `json:"text"`                     // footer text
	IconUrl      *string `json:"icon_url,omitempty"`       // url of footer icon (only supports http(s) and attachments)
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"` // a proxied url of footer icon
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-image-structure
type EmbedImage struct {
	Url      string  `json:"url"`                 // source url of image (only supports http(s) and attachments)
	ProxyUrl *string `json:"proxy_url,omitempty"` // a proxied url of the image
	Height   *int    `json:"height,omitempty"`    // height of image
	Width    *int    `json:"width,omitempty"`     // width of image
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	Url      string  `json:"url"`                 // source url of thumbnail (only supports http(s) and attachments)
	ProxyUrl *string `json:"proxy_url,omitempty"` // a proxied url of the thumbnail
	Height   *int    `json:"height,omitempty"`    // height of thumbnail
	Width    *int    `json:"width,omitempty"`     // width of thumbnail
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-video-structure
type EmbedVideo struct {
	Url      *string `json:"url,omitempty"`       // source url of video
	ProxyUrl *string `json:"proxy_url,omitempty"` // a proxied url of the video
	Height   *int    `json:"height,omitempty"`    // height of video
	Width    *int    `json:"width,omitempty"`     // width of video

}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name *string `json:"name,omitempty"` // name of provider
	Url  *string `json:"url,omitempty"`  // url of provider
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string  `json:"name"`                     // name of author
	Url          *string `json:"url,omitempty"`            // url of author (only supports http(s))
	IconUrl      *string `json:"icon_url,omitempty"`       // url of author icon (only supports http(s) and attachments)
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"` // a proxied url of author icon
}

// Reference: https://discord.com/developers/docs/resources/message#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`             // name of the field
	Value  string `json:"value"`            // value of the field
	Inline *bool  `json:"inline,omitempty"` // whether or not this field should display inline
}
