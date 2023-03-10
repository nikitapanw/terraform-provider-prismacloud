package accountv2

type AccountAndCredentials struct {
	Account AccountResponse `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type CloudAccountResp struct {
	AccountId               string      `json:"accountId"`
	Name                    string      `json:"name"`
	CloudType               string      `json:"cloudType"`
	Enabled                 bool        `json:"enabled"`
	ParentId                string      `json:"parentId"`
	AccountType             string      `json:"accountType"`
	Deleted                 bool        `json:"deleted"`
	ProtectionMode          string      `json:"protectionMode"`
	DeploymentType          string      `json:"deploymentType"`
	CustomerName            string      `json:"customerName"`
	CreatedEpochMillis      int         `json:"createdEpochMillis"`
	LastModifiedEpochMillis int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy          string      `json:"lastModifiedBy"`
	Features                []Features1 `json:"features"`
}

type AccountResponse struct {
	CloudAccountResp CloudAccountResp `json:"cloudAccount"`
	RoleArn          string           `json:"roleArn"`
	ExternalId       string           `json:"externalId"`
	HasMemberRole    bool             `json:"hasMemberRole"`
	TemplateUrl      string           `json:"templateUrl"`
	GroupIds         []string         `json:"groupIds"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Aws struct {
	AccountId   string     `json:"accountId"`
	AccountType string     `json:"accountType"`
	Enabled     bool       `json:"enabled"`
	Features    []Features `json:"features"`
	GroupIds    []string   `json:"groupIds"`
	Name        string     `json:"name"`
	RoleArn     string     `json:"roleArn"`
}

type AwsV2 struct {
	CloudAccountResp          CloudAccountResp `json:"cloudAccount"`
	Name                      string           `json:"name"`
	RoleArn                   string           `json:"roleArn"`
	ExternalId                string           `json:"externalId"`
	HasMemberRole             bool             `json:"hasMemberRole"`
	TemplateUrl               string           `json:"templateUrl"`
	GroupIds                  []string         `json:"groupIds"`
	EventbridgeRuleNamePrefix string           `json:"eventBridgeRuleNamePrefix"`
}

type Features1 struct {
	Name  string `json:"featureName"`
	State string `json:"featureState"`
}

type Features struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type CloudAccount struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}