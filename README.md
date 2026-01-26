# Bingusboard

[![License: AGPL-3.0](https://img.shields.io/badge/License-AGPL%203.0-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

Bingusboard is a work-in-progress fork of the original Focalboard project. The goal of this fork is to keep the core experience lean and familiar to fans of Trello-style task boards while staying lightweight enough to run comfortably on Linux with the built-in SQLite storage backend.

At this stage only the essential web client and Go personal server paths are actively maintained. Other Focalboard features—particularly the enterprise integrations, alternative database layers, and desktop or mobile bundles—may or may not continue to function and could be removed over time.

## ✨ Features

- **Board Views**: Organize tasks with familiar Trello-style Kanban boards, tables, galleries, and calendar views
- **Self-Hosted**: Run on your own infrastructure with full control over your data
- **Lightweight**: SQLite backend for simple, file-based storage (PostgreSQL/MySQL also supported)
- **Multilingual**: Built-in internationalization support
- **Collaborative**: Multi-user support with real-time updates
- **Templates**: Pre-built templates to get started quickly
- **Import/Export**: Board archive format for backup and migration

## 🚀 Quick Start

### Prerequisites

Before you begin, ensure you have the following installed:

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go](https://golang.org/doc/install) 1.21 or higher (see `server/go.mod`)
- [Node.js](https://nodejs.org/en/download/) 20.19 or higher (see `webapp/.nvmrc`)
- npm (comes with Node.js)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/SomethingGeneric/bingusboard.git
   cd bingusboard
   ```

2. **(Optional) Create an `.env` file** in the project root if excluding enterprise features:
   ```bash
   echo EXCLUDE_ENTERPRISE=1 > .env
   ```

3. **Install dependencies and build**
   ```bash
   make prebuild  # Install webapp dependencies
   make           # Build server and webapp
   ```

4. **Run the server**
   ```bash
   ./bin/focalboard-server
   ```

5. **Open your browser** and navigate to:
   ```
   http://localhost:8000
   ```

That's it! You can now create your first board and start organizing your tasks.

## 🛠️ Development

### Development Workflow

For continuous development with automatic rebuilding:

```bash
# Terminal 1: Run the server
./bin/focalboard-server

# Terminal 2: Watch and rebuild webapp on changes
cd webapp && npm run watchdev
```

Then reload your browser to see changes.

### Building Components

```bash
make webapp          # Build webapp only
make server          # Build server only
make                 # Build everything
make clean           # Clean build artifacts
```

### Configuration

The server can be configured via `config.json` in the project root. Key settings include:

- **`port`**: Server port (default: `8000`)
- **`dbtype`**: Database type - `sqlite3` (default), `postgres`, or `mysql`
- **`dbconfig`**: Database connection string
- **`webpath`**: Path to webapp bundle (default: `./webapp/pack`)
- **`filespath`**: File upload storage location (default: `./files`)

For a full list of configuration options, see the example `config.json` in the project root. Note: `server-config.json` uses a relative `webpath` (`./pack`) for running the server directly from the `server/` directory during development, while `config.json` uses `./webapp/pack` for running from the project root.

## 🧪 Testing

### Run All Tests (CI simulation)

```bash
make ci
```

This runs the complete test suite including:
- Server unit tests (SQLite)
- Webapp linting (ESLint + Stylelint)
- Webapp unit tests (Jest)
- Webapp E2E tests (Cypress)

### Individual Test Suites

**Server Tests:**
```bash
make server-test-sqlite      # SQLite tests (fast)
make server-test-postgres    # PostgreSQL tests
make server-test-mysql       # MySQL tests
make server-test-mariadb     # MariaDB tests
make server-test             # All database tests
```

**Webapp Tests:**
```bash
cd webapp
npm run check                # Lint TypeScript and SCSS
npm run fix                  # Auto-fix lint issues
npm run test                 # Run Jest unit tests
npm run updatesnapshot       # Update Jest snapshots
npm run cypress:ci           # Run Cypress E2E tests
```

### Code Quality

**Linting:**
```bash
cd webapp && npm run check                    # Check for issues
cd webapp && npm run fix                      # Auto-fix JS/TS issues
cd webapp && npm run fix:scss                 # Auto-fix SCSS issues
cd server && golangci-lint run ./...          # Lint Go code (requires golangci-lint: https://golangci-lint.run/usage/install/)
```

## 📁 Project Structure

```
bingusboard/
├── server/              # Go backend server
│   ├── main/           # Server entrypoint
│   ├── app/            # Core application logic
│   ├── api/            # HTTP API handlers
│   ├── services/       # Business services
│   ├── model/          # Data models
│   ├── auth/           # Authentication
│   └── integrationtests/ # Integration test suites
├── webapp/             # React/TypeScript web client
│   ├── src/           # UI source code
│   ├── cypress/       # E2E tests
│   └── i18n/          # Internationalization
├── docs/              # Documentation
├── config.json        # Server configuration
└── Makefile           # Build automation
```

## 🤝 Contributing

We welcome contributions! This is a community-maintained fork focused on keeping the project lean and functional.

Before contributing:

1. Check existing issues or create a new one to discuss your idea
2. Read [`CONTRIBUTING.md`](CONTRIBUTING.md) for guidelines
3. Follow the coding style enforced by `.editorconfig`
4. Run tests before submitting: `make ci`
5. Write clear commit messages in imperative mood

See [`docs/`](docs/) for additional development guides.

## 📋 Roadmap & Status

This is a **work-in-progress fork** focused on:

- ✅ Maintaining the core personal server experience
- ✅ SQLite-first storage approach
- ✅ Web client functionality
- ⚠️ Other features (enterprise integrations, mobile/desktop apps) may be removed

For legacy documentation and the broader Focalboard feature set, see [`HISTORIC-README.md`](HISTORIC-README.md).

## 📜 License

This project is licensed under the GNU AGPL v3.0 - see [`LICENSE.txt`](LICENSE.txt) for details.

Source code in Admin Tools and Configuration Files is licensed under Apache License v2.0.

## 🔗 Links

- **Issues & Discussions**: [GitHub Issues](https://github.com/SomethingGeneric/bingusboard/issues)
- **Upstream Project**: [Focalboard (archived)](https://github.com/mattermost/focalboard)
- **Security Policy**: [`SECURITY.md`](SECURITY.md)
- **Changelog**: [`CHANGELOG.md`](CHANGELOG.md)

---

**Note**: Bingusboard is a fork maintained by the community. We're focused on keeping it simple, lightweight, and self-hostable. Expect rapid changes as the project evolves!
