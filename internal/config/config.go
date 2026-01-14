package config

type Config struct {
	Agents map[string]AgentProfile `toml:"agents"`
}

type AgentProfile struct {
	Command string   `toml:"command"`
	Args    []string `toml:"args"`
}

const DefaultConfigToml = `# Praude configuration

[agents.codex]
command = "codex"
args = []
# args = ["--profile", "pm", "--prompt-file", "{brief}"]

[agents.claude]
command = "claude"
args = []
# args = ["--profile", "pm", "--prompt-file", "{brief}"]

[agents.opencode]
command = "opencode"
args = []

[agents.droid]
command = "droid"
args = []
`
