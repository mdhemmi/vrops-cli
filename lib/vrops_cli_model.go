package vrops_cli

type VROPsData struct {
	Action   string
	FQDN     string
	Auth     string
	Search   string
	Username string
	Password string
	Token    string
	Insecure bool
	Debug    bool
	Extended bool
}

type StatsOfResources struct {
	Values []struct {
		ResourceID string `json:"resourceId"`
		StatList   struct {
			Stat []struct {
				Timestamps []int64 `json:"timestamps"`
				StatKey    struct {
					Key string `json:"key"`
				} `json:"statKey"`
				IntervalUnit struct {
					Quantifier int `json:"quantifier"`
				} `json:"intervalUnit"`
				Data []int `json:"data"`
			} `json:"stat"`
		} `json:"stat-list"`
	} `json:"values"`
}

type PropertiesOfResources struct {
	ResourcePropertiesList []struct {
		ResourceID string `json:"resourceId"`
		Property   []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"property"`
	} `json:"resourcePropertiesList"`
}
