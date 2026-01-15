package tui

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mistakeknot/praude/internal/project"
	"github.com/mistakeknot/praude/internal/research"
	"github.com/mistakeknot/praude/internal/scan"
	"github.com/mistakeknot/praude/internal/specs"
	"gopkg.in/yaml.v3"
)

type interviewStep int

const (
	stepScanPrompt interviewStep = iota
	stepDraftConfirm
	stepVision
	stepUsers
	stepProblem
	stepRequirements
	stepResearchPrompt
)

type interviewState struct {
	step         interviewStep
	root         string
	draft        specs.Spec
	scanSummary  string
	vision       string
	users        string
	problem      string
	requirements string
	warnings     []string
	specID       string
	specPath     string
}

func startInterview(root string) interviewState {
	return interviewState{step: stepScanPrompt, root: root}
}

func (m *Model) handleInterviewInput(key string) {
	switch m.interview.step {
	case stepScanPrompt:
		m.handleScanPrompt(key)
	case stepDraftConfirm:
		m.handleDraftConfirm(key)
	case stepVision:
		m.handleTextStep(key, func(input string) {
			m.interview.vision = input
			m.interview.step = stepUsers
		})
	case stepUsers:
		m.handleTextStep(key, func(input string) {
			m.interview.users = input
			m.interview.step = stepProblem
		})
	case stepProblem:
		m.handleTextStep(key, func(input string) {
			m.interview.problem = input
			m.interview.step = stepRequirements
		})
	case stepRequirements:
		m.handleTextStep(key, func(input string) {
			m.interview.requirements = input
			m.finalizeInterview()
			m.interview.step = stepResearchPrompt
		})
	case stepResearchPrompt:
		if key == "y" {
			m.runResearch()
			m.exitInterview()
			return
		}
		if key == "n" {
			m.exitInterview()
		}
	}
}

func (m *Model) handleScanPrompt(key string) {
	if key != "y" && key != "n" {
		return
	}
	if key == "y" {
		res, _ := scan.ScanRepo(m.interview.root, scan.Options{})
		m.interview.scanSummary = renderScanSummary(res)
	}
	m.interview.draft = buildDraftSpec(m.interview.scanSummary)
	m.interview.step = stepDraftConfirm
}

func (m *Model) handleDraftConfirm(key string) {
	if key == "y" {
		m.interview.step = stepVision
		return
	}
	if key == "n" {
		m.exitInterview()
	}
}

func (m *Model) handleTextStep(key string, onDone func(string)) {
	switch key {
	case "enter":
		onDone(strings.TrimSpace(m.input))
		m.input = ""
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
		}
	default:
		m.input += key
	}
}

func (m *Model) finalizeInterview() {
	spec := buildSpecFromInterview(m.interview)
	path, id, warnings := writeSpec(m.interview.root, spec)
	m.interview.specPath = path
	m.interview.specID = id
	m.interview.warnings = warnings
	m.reloadSummaries()
}

func (m *Model) runResearch() {
	if m.interview.specID == "" {
		return
	}
	researchDir := project.ResearchDir(m.interview.root)
	_, _ = research.Create(researchDir, m.interview.specID, time.Now())
}

func (m *Model) exitInterview() {
	m.mode = "list"
	m.input = ""
	m.interview = interviewState{}
}

func renderScanSummary(res scan.Result) string {
	return "Scan summary: " + itoa(len(res.Entries)) + " files, " + itoa(int(res.TotalBytes)) + " bytes"
}

func (m Model) renderInterview() []string {
	lines := []string{
		"Guided interview",
		"PM-focused agent: Codex CLI / Claude Code",
	}
	switch m.interview.step {
	case stepScanPrompt:
		lines = append(lines, "Scan repo now? (y/n)")
	case stepDraftConfirm:
		lines = append(lines, "Draft PRD ready.")
		if m.interview.scanSummary != "" {
			lines = append(lines, m.interview.scanSummary)
		}
		lines = append(lines, "Confirm draft? (y/n)")
	case stepVision:
		lines = append(lines, "Vision:", m.input)
	case stepUsers:
		lines = append(lines, "Users:", m.input)
	case stepProblem:
		lines = append(lines, "Problem:", m.input)
	case stepRequirements:
		lines = append(lines, "Requirements (comma or newline separated):", m.input)
	case stepResearchPrompt:
		lines = append(lines, "Run research now? (y/n)")
	}
	return lines
}

func buildDraftSpec(summary string) specs.Spec {
	text := summary
	if text == "" {
		text = "Draft from scan"
	}
	return specs.Spec{Title: "Draft PRD", Summary: text}
}

func buildSpecFromInterview(state interviewState) specs.Spec {
	reqList := parseRequirements(state.requirements)
	if len(reqList) == 0 {
		reqList = []string{"REQ-001: TBD"}
	}
	firstReq := extractReqID(reqList[0])
	title := firstNonEmpty(state.vision, state.problem, "New PRD")
	summary := firstNonEmpty(state.problem, state.vision, "Summary pending")
	return specs.Spec{
		Title:        title,
		Summary:      summary,
		Requirements: reqList,
		StrategicContext: specs.StrategicContext{
			CUJID:       "CUJ-001",
			CUJName:     "Primary Journey",
			FeatureID:   "",
			MVPIncluded: true,
		},
		UserStory: specs.UserStory{
			Text: "As a user, " + firstNonEmpty(state.users, "I need", "I need") + ", " + summary,
			Hash: "pending",
		},
		CriticalUserJourneys: []specs.CriticalUserJourney{
			{
				ID:                 "CUJ-001",
				Title:              "Primary Journey",
				Priority:           "high",
				Steps:              []string{"Start", "Finish"},
				SuccessCriteria:    []string{"Goal achieved"},
				LinkedRequirements: []string{firstReq},
			},
			{
				ID:                 "CUJ-002",
				Title:              "Maintenance",
				Priority:           "low",
				Steps:              []string{"Routine upkeep"},
				SuccessCriteria:    []string{"System remains stable"},
				LinkedRequirements: []string{firstReq},
			},
		},
	}
}

func writeSpec(root string, spec specs.Spec) (string, string, []string) {
	specDir := project.SpecsDir(root)
	id, err := specs.NextID(specDir)
	if err != nil {
		return "", "", nil
	}
	spec.ID = id
	if spec.CreatedAt == "" {
		spec.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	}
	raw, err := yaml.Marshal(spec)
	if err != nil {
		return "", id, nil
	}
	path := filepath.Join(specDir, id+".yaml")
	if err := osWriteFile(path, raw, 0o644); err != nil {
		return path, id, nil
	}
	res, err := specs.Validate(raw, specs.ValidationOptions{Mode: specs.ValidationSoft, Root: root})
	if err != nil {
		return path, id, nil
	}
	if len(res.Warnings) > 0 {
		_ = specs.StoreValidationWarnings(path, res.Warnings)
	}
	return path, id, res.Warnings
}

func parseRequirements(input string) []string {
	parts := splitInput(input)
	var out []string
	for i, part := range parts {
		id := formatReqID(i + 1)
		out = append(out, id+": "+part)
	}
	return out
}

func splitInput(input string) []string {
	input = strings.ReplaceAll(input, "\n", ",")
	parts := strings.Split(input, ",")
	var out []string
	for _, part := range parts {
		trim := strings.TrimSpace(part)
		if trim != "" {
			out = append(out, trim)
		}
	}
	return out
}

func formatReqID(n int) string {
	return "REQ-" + pad3(n)
}

func pad3(n int) string {
	if n < 10 {
		return "00" + itoa(n)
	}
	if n < 100 {
		return "0" + itoa(n)
	}
	return itoa(n)
}

func extractReqID(req string) string {
	fields := strings.Fields(req)
	if len(fields) == 0 {
		return "REQ-001"
	}
	id := strings.TrimSuffix(fields[0], ":")
	if strings.HasPrefix(id, "REQ-") {
		return id
	}
	return "REQ-001"
}

func firstNonEmpty(values ...string) string {
	for _, val := range values {
		if strings.TrimSpace(val) != "" {
			return val
		}
	}
	return ""
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

var osWriteFile = os.WriteFile
