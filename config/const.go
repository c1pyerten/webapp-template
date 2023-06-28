package config


const (
	ApiPathRoot = "/api/v1"
	ApiPathUserGroup = ApiPathRoot + "/user"
	ApiPathUserGetById = ApiPathUserGroup + "/" 

)

const (
	AppConfigPath = "resources/config/application.%s.yml"
	MessagesConfigPath = "resources/config/messages.properties"
)

const (
	ENV_NAME = "WEB_APP_ENV"
)