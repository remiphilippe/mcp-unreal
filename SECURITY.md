# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in mcp-unreal, please report it responsibly through [GitHub Security Advisories](https://github.com/remiphilippe/mcp-unreal/security/advisories/new).

**Do not open a public issue for security vulnerabilities.**

We will acknowledge your report within 48 hours and aim to provide a fix or mitigation within 7 days for critical issues.

## Scope

The following are considered security vulnerabilities:

- **Remote code execution** via MCP tool inputs (e.g., command injection through tool parameters)
- **Path traversal** allowing access to files outside the UE project root
- **stdout pollution** from the Go binary that could corrupt the JSON-RPC transport and potentially inject malicious tool responses
- **Unauthorized network exposure** — the Go binary opening listening sockets, or the UE plugin binding to non-loopback addresses
- **Credential leaks** — secrets appearing in logs, tool responses, or error messages

## Out of Scope

The following are **not** considered vulnerabilities in this project:

- Issues in Unreal Engine itself (report to Epic Games)
- Issues requiring local access to the machine running the MCP server (the entire system is designed for local-only use)
- The `execute_script` and `run_console_command` tools executing arbitrary code — this is by design, and access is controlled by Claude Code's built-in permission system
- Denial of service via resource exhaustion (large mesh data, many tool calls) — the system is designed for single-user local use

## Architecture Security Notes

- The MCP server communicates exclusively via **stdio** (JSON-RPC 2.0). It does not open any listening network sockets.
- The UE editor plugin HTTP server binds to **127.0.0.1 only** (ports 30010 for RC API, 8090 for the plugin).
- All subprocess execution uses explicit argument arrays — no shell expansion.
- See `CLAUDE.md` Security section for the full security model.
