# MailMole Documentation

Welcome to MailMole documentation. Here you'll find everything you need to know about using MailMole for email migration.

## Quick Links

### Getting Started
- [Installation Guide](installation.md) - How to install and run MailMole
- [Getting Started](getting-started.md) - Quick start tutorial
- [Web Dashboard](web-dashboard.md) - Browser-based migration interface (NEW)

### Usage Guides
- [Manual Migration](usage/manual-mode.md) - Migrate single account pairs
- [Bulk Migration](usage/bulk-mode.md) - Migrate multiple accounts from CSV file
- [Resume Migration](usage/resume-migration.md) - Continue interrupted migrations
- [Web Interface](usage/web-interface.md) - Using the web dashboard

### Troubleshooting & Reference
- [Troubleshooting](troubleshooting.md) - Common errors and solutions
- [FAQ](faq.md) - Frequently asked questions
- [Provider Setup](providers/general.md) - Email provider configurations

## What is MailMole?

MailMole is a terminal-based IMAP-to-IMAP mail migration tool that allows you to:

- Migrate emails between any IMAP-compatible mail servers
- Support for Gmail, Outlook, Yahoo, Yandex, and custom mail servers
- Bulk migration with CSV file support
- Real-time progress tracking
- Resume interrupted migrations
- Preserve message flags (read, unread, starred, etc.)

## Interface Options

MailMole provides two interfaces:

### 1. Terminal UI (TUI) - Default
The classic terminal interface with keyboard navigation.

```bash
./mailmole
```

### 2. Web Dashboard (NEW)
Modern browser-based interface with additional features:
- Visual folder preview and selection
- Migration templates (Gmail, Outlook, Yandex, etc.)
- Import/Export functionality
- Connection validation with detailed reports
- Toast notifications
- Multi-language support (English, Turkish)
- Real-time monitoring

```bash
# TUI + Web Dashboard
./mailmole -web :8080

# Web only
./mailmole -web :8080 -web-only
```

Access at: `http://localhost:8080`

## Which Interface Should I Use?

| Feature | TUI | Web Dashboard |
|---------|-----|---------------|
| Terminal/SSH | ✅ | ❌ |
| Visual preview | ❌ | ✅ |
| Bulk templates | ❌ | ✅ |
| Import/Export CSV | Manual | ✅ Built-in |
| Connection testing | Basic | ✅ Detailed |
| Mobile friendly | ❌ | ✅ |
| Keyboard shortcuts | ✅ | Partial |

**Recommendation:** Use Web Dashboard for ease of use, TUI for headless servers.
