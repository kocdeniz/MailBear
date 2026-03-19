# Frequently Asked Questions

Common questions about MailMole.

## General

### What is MailMole?

MailMole is an IMAP-to-IMAP email migration tool with both terminal (TUI) and web dashboard interfaces. It allows you to transfer emails between any two mail servers that support IMAP protocol.

### Which interface should I use?

**Use Web Dashboard if:**
- You prefer visual interfaces
- Want folder preview before migration
- Need migration templates
- Will import accounts from CSV
- Want detailed connection testing

**Use Terminal UI if:**
- Working over SSH
- Running on headless servers
- Prefer keyboard navigation
- Want minimal resource usage

### What does MailMole support?

**Both Interfaces:**
- Migrating between any IMAP-compatible mail servers
- Gmail, Outlook, Yahoo, Yandex, iCloud, and custom servers
- Bulk migration from CSV file
- Preserving message flags (read/unread, starred, etc.)
- Resume interrupted migrations
- Real-time progress tracking

**Web Dashboard Only:**
- Visual folder preview and selection
- 19 migration templates
- Import/Export functionality
- Detailed connection validation
- Toast notifications
- Multi-language support
- Real-time monitoring with charts

### What doesn't MailMole support?

- POP3 mailboxes
- Migrating contacts or calendars
- Migrating filters or rules
- OAuth2 authentication (planned)

## Web Dashboard

### How do I access the web dashboard?

```bash
./mailmole -web :8080
```

Then open `http://localhost:8080` in your browser.

### Can I use web dashboard remotely?

**Yes**, but with caution:
```bash
./mailmole -web 0.0.0.0:8080
```

⚠️ **Security Warning:** The web dashboard has no built-in authentication. Only use on trusted networks or behind a firewall/VPN.

### How do I change the language?

Click the language selector in the **top-right corner** of the page:
- 🇬🇧 English
- 🇹🇷 Turkish

Your preference is saved in browser localStorage.

### What are Migration Templates?

Pre-configured server settings for common providers:
- Click a template button to auto-fill host, port, and SSL settings
- Available for: Gmail, Outlook, Yandex, iCloud, Yahoo
- 19 different migration combinations

### How do I import accounts from CSV?

**Option 1: Download Example**
1. Click **"📄 Example CSV"** button
2. Edit the downloaded file with your accounts
3. Click **"📥 Import Accounts"**
4. Select your edited CSV

**Option 2: Create Your Own**
Create a CSV file with this format:
```csv
src_host,src_port,src_user,src_pass,dst_host,dst_port,dst_user,dst_pass
imap.gmail.com,993,user1@gmail.com,pass1,outlook.office365.com,993,user1@company.com,pass2
```

### What's the difference between Test Connection and Detailed Test?

**Test Connection:**
- Quick connectivity check
- Shows toast notification with result

**Detailed Test:**
- Comprehensive validation
- Shows per-account results in a table
- Lists specific errors
- Exportable results

### Can I select specific folders to migrate?

**Yes!** In the Web Dashboard:
1. Click **Preview Migration**
2. See all folders with checkboxes
3. Uncheck folders you don't want
4. Click **Start Migration**

## Security

### Are my passwords saved?

**No.** Passwords are:
- Never saved to disk
- Never sent to any external server
- Only used for the current IMAP connection
- Cleared from memory after use

### Is the connection secure?

Yes. MailMole:
- Supports SSL/TLS encryption
- Uses hostname verification for TLS
- Falls back to certificate verification for IP addresses

### What about the log file?

The log file (`mailmole.log`) contains:
- Connection timestamps
- Migration progress
- Error messages

It does NOT contain:
- Passwords
- Email content
- Attachments

## Migration

### Will duplicate emails be created?

**No.** MailMole uses Message-ID to detect duplicates:
- Before transferring, it checks if the message exists
- Existing messages are skipped
- Safe to run multiple times

### What happens if migration is interrupted?

MailMole supports resume:
- Completed accounts are saved to `migration_state.json`
- Select "Resume Previous" (TUI) or check state file (Web)
- Already completed accounts are skipped automatically

### Can I migrate only specific folders?

**Yes, in Web Dashboard:**
1. Go to Preview page
2. Uncheck folders you don't want
3. Only selected folders will be migrated

**In Terminal UI:** Currently all folders are migrated.

### Will my read/unread status be preserved?

**Yes.** MailMole preserves:
- `\Seen` - Read/Unread status
- `\Answered` - Answered status
- `\Flagged` - Starred/Flagged status
- `\Draft` - Draft status

### How fast is migration?

Speed depends on:
- Network bandwidth
- Server response time
- Message size
- Number of concurrent workers

Typical speeds: 30-100 emails per second

### Can I run migration on a schedule?

**Yes (Web Dashboard Beta):**
1. Enable scheduling in Connection Setup
2. Set start date/time
3. Choose repeat option (Daily/Weekly/Monthly)
4. Migration will start automatically

## Technical

### What are the system requirements?

- Linux, macOS, or Windows
- Go 1.24+ (to build from source)
- IMAP access to both mail servers
- Network connectivity
- For Web Dashboard: Modern web browser

### Can I run this on a server?

**Yes.** MailMole:
- Has both CLI and Web interfaces
- Creates minimal resource usage
- Can run in background
- Supports resume for long migrations

### Does it use a database?

**No.** MailMole uses simple files:
- `migration_state.json` - tracks completed migrations
- No database installation required

### Can I run multiple instances?

Not recommended. Use separate directories for different migrations to avoid state file conflicts.

## Troubleshooting

### Gmail says "Application-specific password required"

Gmail requires app passwords for accounts with 2-Step Verification. See [Gmail Setup Guide](providers/gmail.md).

### Outlook authentication fails

Microsoft may require app passwords. See [Outlook Setup Guide](providers/outlook.md).

### Connection times out

- Check internet connection
- Try again in a few minutes
- Server might be rate-limiting
- Check firewall settings

### Migration is slow

- Normal for large attachments
- Server may have rate limits
- Consider running during off-peak hours
- Check concurrent worker settings

### Web dashboard won't load

- Check if port is available: `lsof -i :8080`
- Try different port: `./mailmole -web :3000`
- Check firewall settings
- Clear browser cache

### Language not changing

- Clear browser cache (Ctrl+Shift+R or Cmd+Shift+R)
- Check browser console for errors
- Ensure cookies/localStorage enabled

## Contributing

### How can I contribute?

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

### Where is the source code?

[github.com/kocdeniz/MailMole](https://github.com/kocdeniz/MailMole)

### How do I report bugs?

Create an issue on GitHub with:
- Error message
- Steps to reproduce
- Mail server types
- Interface used (TUI or Web)
