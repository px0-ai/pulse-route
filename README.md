# PulseRoute

PulseRoute is a demo repository designed to evaluate and demonstrate the full review lifecycle, navigation speed, and AI harness integration of px0.

It models an edge API gateway handling authentication, token-bucket rate limiting, and reverse proxying across microservice endpoints.

- < 1 ms cold start inspection
- Multi-format verification: Go source code, GFM markdown, CSV latency data, and raw log streams
- Embedded concurrency issue in `internal/limiter/limiter.go` for testing inline agent edits
- Scoped for zero-config LSP navigation (`gopls`), diffing, and git lifecycle actions

## Quick Start

Inspect this repository using px0:

```bash
px0 .
```

To inspect specific files or line anchors:

```bash
px0 cmd/gateway/main.go:14
```

## Walkthrough Tasks

### 1. Startup and Memory Footprint

- Check the startup time printed in the console (< 1 ms).
- Click the CPU/RAM metrics badge in the bottom-right status bar.
- Compare the daemon resident memory (~20-30 MB RSS) with traditional Electron-based IDEs (~1,440 MB).

### 2. File and Workspace Navigation

- Fuzzy file search: Press `Ctrl+P` and type `arch` to jump to `docs/ARCHITECTURE.md`.
- File jump: Press `Ctrl+P` and type `lim` to open `internal/limiter/limiter.go`.
- Workspace search: Press `Ctrl+Shift+F` and search for `TokenBucket` across the repository.

### 3. Markdown and Document Preview

- Open `docs/ARCHITECTURE.md`.
- Press `Alt+M` (or click Preview in the status bar) to toggle the rendered GFM preview.
- Inspect rendered GitHub alert blocks (`[!NOTE]`, `[!WARNING]`), table layout, and linked architecture assets.

### 4. CSV Table View

- Open `data/benchmark_latencies.csv`.
- Press `Alt+M` to switch to table view mode.
- Scroll through the records to verify the pinned header row and source-line gutter indicators.

### 5. Editor DOM Virtualization

- Open `logs/traffic_stress.log` (100,000 lines).
- Scroll continuously through the file using the mouse wheel or scrollbar.
- Verify smooth rendering performance: px0 maintains only ~60 active DOM rows regardless of document length.

### 6. Code Intelligence and LSP

- Open `cmd/gateway/main.go`.
- Hover over `TokenBucket` to view the inferred type signature and documentation card.
- Place the cursor on `NewTokenBucket` and press `F12` to jump to its definition in `internal/limiter/limiter.go`.
- On `Allow`, press `Shift+F12` to view all usage sites across the workspace in the inspector drawer.
- Press `Alt+Shift+H` to inspect the call trail tree.

### 7. Inline AI Code Editing

- Open `internal/limiter/limiter.go`.
- Locate lines 35-43 inside `Allow()`: note the unprotected access to `tb.tokens` and `tb.lastRefill`.
- Select lines 34-45 and press `Alt+E` (or right-click and choose Edit Inline).
- Check the harness and model selectors in the composer box to ensure your preferred tool (Claude Code, Antigravity, Gemini CLI, or OpenCode) is active.
- Enter the instruction:
  ```text
  Wrap the refill calculation and token deduction in tb.mu.Lock() and tb.mu.Unlock() to prevent data races.
  ```
- Press `Ctrl+Enter` (Apply now).
- Monitor execution progress in the terminal; px0 reloads the file and updates git indicators on disk write.

### 8. Multi-Turn Threads

- Press `Alt+T` to open the Threads drawer in the right inspector.
- Select `cmd/gateway/main.go` and start a thread with your active coding harness:
  ```text
  Expose an HTTP handler at /metrics that returns the request and blocked counters from internal/proxy/metrics.go.
  ```
- Track tool steps streaming via Server-Sent Events and review the files modified during the turn.

### 9. Git Review, Unpushed Commits, and AI Commit

- Press `Ctrl+D` to toggle between the file explorer and uncommitted git changes.
- Review the split and unified diffs generated from previous agent edits.
- In the sidebar Git panel docked below the tree, click Stage All.
- Click Commit with AI (`#git-generate-msg`) to generate a commit message from the staged diff.
- Click Commit to record the changes locally.
- In the sidebar, notice the Unpushed Commits section appear above the Git panel: hover over the commit to view the floating commit card (author, SHA, diffstat), and click it to inspect individual commit diffs before pushing.

### 10. Pull Request Review Mode

- Review any active GitHub pull request directly without branch switching or manual checkout:
  ```bash
  px0 https://github.com/<owner>/<repo>/pull/<id>
  ```
  Or from within an active px0 session, press `Ctrl+K` and select "Git: Open Pull Request...".
- Inspect scoped merge-base diffs within an isolated temporary worktree.
- Select diff lines and press `Alt+R` to draft inline review comments.
- Click Batch Apply in the review header to run your coding agent against all drafted comments simultaneously.
- Use the sidebar Git panel to commit fixes and push directly back to the pull request branch.

## Keyboard Shortcuts

| Shortcut | Context | Action |
| :--- | :--- | :--- |
| `Ctrl+P` / `Cmd+P` | Global | Quick Open fuzzy file finder |
| `Ctrl+Shift+F` / `Cmd+Shift+F` | Global | Workspace text and regex search |
| `Ctrl+D` / `Cmd+D` | Editor | Toggle split/unified diff view |
| `Alt+M` | Markdown / CSV | Toggle rendered preview or table mode |
| `F12` | Code symbol | Go to definition |
| `Shift+F12` | Code symbol | Find all references |
| `Alt+Shift+H` | Function symbol | Open call trail tree |
| `Alt+E` | Code selection | Open inline agent edit composer |
| `Alt+T` | Global / Selection | Toggle Threads drawer |
| `Alt+R` | Diff selection (PR) | Draft inline review comment |
