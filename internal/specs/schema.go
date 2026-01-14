package specs

type StrategicContext struct {
	CUJID       string `yaml:"cuj_id"`
	CUJName     string `yaml:"cuj_name"`
	FeatureID   string `yaml:"feature_id"`
	MVPIncluded bool   `yaml:"mvp_included"`
}

type UserStory struct {
	Text string `yaml:"text"`
	Hash string `yaml:"hash"`
}

type AcceptanceCriterion struct {
	ID          string `yaml:"id"`
	Description string `yaml:"description"`
}

type FileChange struct {
	Action      string `yaml:"action"`
	Path        string `yaml:"path"`
	Description string `yaml:"description"`
}

type Spec struct {
	ID               string                `yaml:"id"`
	Title            string                `yaml:"title"`
	CreatedAt        string                `yaml:"created_at"`
	StrategicContext StrategicContext      `yaml:"strategic_context"`
	UserStory        UserStory             `yaml:"user_story"`
	Summary          string                `yaml:"summary"`
	Requirements     []string              `yaml:"requirements"`
	Acceptance       []AcceptanceCriterion `yaml:"acceptance_criteria"`
	FilesToModify    []FileChange          `yaml:"files_to_modify"`
	Complexity       string                `yaml:"complexity"`
	EstimatedMinutes int                   `yaml:"estimated_minutes"`
	Priority         int                   `yaml:"priority"`
}
