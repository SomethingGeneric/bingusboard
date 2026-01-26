# Repository Guidelines

## Product Scope
We are building **Bingusboard**, a lighter-weight fork of the historical Focalboard project. Optimizations or workflows that only matter for the legacy product are out of scope unless they clearly benefit Bingusboard.

Docker is only used for local testing in Bingusboard; production releases rely on the GitHub Actions job that produces a server tarball for non-Docker deployments.

## Project Structure & Module Organization
Focalboard pairs a Go personal server with a React/TypeScript web client. Go source sits in `server/`: entrypoint code in `server/main`, domain logic in `server/app` and `server/services`, HTTP handlers in `server/api`, and data contracts in `server/model`, with integration suites under `server/integrationtests`. The web client lives in `webapp/`; UI code resides in `webapp/src`, Cypress specs in `webapp/cypress`, and locale files in `webapp/i18n`. Templates and static bundles ship from `server/assets` and `webapp/pack`.

## Build, Test, and Development Commands
- `make prebuild` installs webapp dependencies.
- `make` builds the server binary (`bin/focalboard-server`) and production web bundle.
- `cd webapp && npm run watchdev` rebuilds the UI on file changes; keep the server running separately.
- `make server-test-sqlite` runs fast Go tests; add `-mysql`, `-mariadb`, or `-postgres` targets for migration coverage.
- `cd webapp && npm run test` runs Jest, and `npm run cypress:ci` executes headless UI flows against the packaged server.
- `make ci` reproduces the GitHub Actions matrix locally.

## Coding Style & Naming Conventions
`.editorconfig` enforces LF endings, trimmed whitespace, Go tabs, and 4-space indents for TS/JS/SCSS. Follow Go norms: package folders in lowercase, exported identifiers in PascalCase, helpers unexported, and `_test.go` files colocated. React components use PascalCase filenames with matching `*.test.tsx`. Run `golangci-lint run ./...` in `server/` (install manually) and `cd webapp && npm run check` before pushing; `npm run fix` and `npm run fix:scss` tidy most lint violations.

## Testing Guidelines
Use `make server-test-sqlite` for quick regression checks and the full `make server-test` matrix ahead of releases to cover MySQL, MariaDB, and Postgres. Go tests should follow `TestXxx` conventions, prefer table-driven cases, and keep integration flows in `server/integrationtests`. Frontend unit tests rely on Jest with Testing Library and coverage by default; update snapshots with `npm run updatesnapshot`. Cypress suites in `webapp/cypress` should reuse the provided fixtures and hit the default `http://localhost:8080` endpoint.

## Commit & Pull Request Guidelines
Write commit subjects in the imperative mood (“Add board icon badges”) and keep them under ~72 characters; split backend and frontend changes into separate commits when feasible, mirroring recent history. Reference issues with `Fixes #123` or `Refs #123` in the body. PRs should call out behavior changes, include validation steps (commands run, databases tested), and add screenshots or GIFs for UI updates. If you touch configuration, highlight default ports (`8000`) or `.env` requirements so reviewers can reproduce quickly.
