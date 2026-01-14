# AGENTS.md

## Ground Rules
- Run the superpowers bootstrap first:
  - `~/.codex/superpowers/.codex/superpowers-codex bootstrap`
- Use relevant superpowers skills for every task (TDD, writing-plans, executing-plans).
- Do not use git worktrees for this repo unless explicitly requested.
- Keep changes small, testable, and committed frequently.

## Repo Overview
Praude is a TUI-first PM/specs CLI that generates, validates, and governs PRDs.
It orchestrates external agents (Codex, Claude Code, opencode, droid) and stores
specs as YAML files in `.praude/specs/`.

## Commands
- `praude` launches the TUI.
- `praude init` initializes `.praude/`.
- `praude run <id> --agent=<profile>` spawns an agent with a generated brief.

## Testing
- Use TDD for all behavior changes.
- Run targeted tests (`go test ./internal/<package> -v`) while iterating.
