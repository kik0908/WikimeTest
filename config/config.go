package config

type vkAuth struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scope        []string
}

var VkAuth *vkAuth

func init() {
	VkAuth = &vkAuth{
		"51421836",
		"UqsuMNV6bHRycTr1ZrEe",
		"localhost:80/auth",
		[]string{"account", "wall"},
	}
}
