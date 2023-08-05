<!-- LTeX: language=en-US -->
# game-gorl (game codebase using Go and raylib)
## Creating a project
First, start by copying the source:
```bash
git clone https://github.com/wintermute-cell/game-gorl <my-project-name>
cd <my-project-name>
rm -rf .git
```
Then, adjust the `makefile`. Finally, run:
```bash
go mod init <my-project-name>
go mod tidy
```

## Structure
- `assets`: Holds game assets.
  - `settings.json`: Settings file.
  - `textures`: Texture files.
- `build`: Destination for build output.
- `main.go`: Application entry point.
- `makefile`: Manages build tasks.
- `pkg`: Go packages.
  - `settings`: Settings handling.
  - `entities`: Game entities.
  - `render`: Rendering logic.
  - `scenes`: Game stages.
- `scripts`: Utility scripts.

## Scripts

### sync_settings.py
Generates Go fallback settings and JSON from GameSettings struct.

Usage:
```bash
python scripts/sync_settings.py [path/to/settings.go path/to/settings.json]
```
If no path is given, the following defaults apply: `./pkg/settings/settings.go`,
`./assets/settings.json`.

> Note: Review changes and test after running.
