# Getting Started

This guide will walk you through your first email migration with MailMole.

## Choose Your Interface

MailMole offers two interfaces:

### Option 1: Web Dashboard (Recommended)

Modern browser interface with visual tools:

```bash
./mailmole -web :8080
```

Then open `http://localhost:8080` in your browser.

**Features:**
- Click-based navigation
- Visual folder preview
- Migration templates
- Import/Export CSV files
- Detailed connection testing
- Toast notifications

### Option 2: Terminal UI (TUI)

Classic terminal interface:

```bash
./mailmole
```

**Features:**
- Keyboard-only navigation
- Works over SSH
- Minimal resource usage
- No browser required

---

## Web Dashboard Quick Start

### Step 1: Launch Web Interface

```bash
./mailmole -web :8080
```

Open `http://localhost:8080` in your browser.

### Step 2: Choose Migration Mode

Click on the sidebar menu:
- **Single Account** - Migrate one account
- **Bulk Migration** - Migrate multiple accounts

### Step 3: Use Templates (Optional)

Click **Quick Templates** buttons to auto-fill server settings:
- Gmail (Source/Dest)
- Outlook (Source/Dest)
- Yandex, iCloud, Yahoo

### Step 4: Enter Credentials

Fill in the form:
```
Source Server:
  Host: imap.gmail.com
  Port: 993
  Username: your@email.com
  Password: ********

Destination Server:
  Host: outlook.office365.com
  Port: 993
  Username: your@email.com
  Password: ********
```

### Step 5: Test Connection

Click **Test Connection** or **Detailed Test** button.

A toast notification will show the result in the top-right corner.

### Step 6: Preview Migration

Click **Preview Migration** to see:
- All folders from source
- Message counts
- Estimated size
- Which folders will be migrated

Select/deselect folders as needed.

### Step 7: Start Migration

Click **Start Migration** button.

Watch real-time progress with:
- Progress bar
- Transfer speed
- Messages migrated count
- ETA

---

## Terminal UI Quick Start

### Step 1: Launch MailMole

```bash
./mailmole
```

Press any key to see the main menu.

### Step 2: Choose Migration Mode

You'll see three options:

| Option | Description |
|--------|-------------|
| `1` Manual Entry | Migrate a single account pair |
| `2` Bulk Migration | Migrate multiple accounts from CSV file |
| `3` Resume Previous | Continue an interrupted migration |

### Step 3: Enter Credentials

#### For Manual Mode

Enter your source and destination server details:

```
SOURCE SERVER
- Host: mail.source.com
- Port: 993
- [x] SSL/TLS
- Username: your@email.com
- Password: ********

DESTINATION SERVER
- Host: mail.dest.com
- Port: 993
- [x] SSL/TLS
- Username: your@email.com
- Password: ********
```

#### Using Provider Presets

When on the Host field, press `F1` to see preset providers:

```
1. Gmail / Google Workspace
2. Outlook / Office 365
3. Yahoo Mail
4. Yandex Mail
5. ProtonMail (Bridge)
6. Zoho Mail
7. iCloud Mail
8. FastMail
9. Custom (manual entry)
```

Selecting a provider automatically fills in the host, port, and SSL settings.

### Step 4: Test Connection

Before migration starts, MailMole tests all account connections:

```
[INFO] Testing connections for 3 accounts...
[TEST] Testing connection for user@example.com (1/3)...
[SUCCESS] All accounts OK.
```

Accounts with connection issues are automatically skipped.

### Step 5: Start Migration

Press `s` to start the migration. You'll see real-time progress:

```
OVERALL PROGRESS  1250 / 5000 messages  Account 2/10
[████████░░░░░░░░░░░░░░░░░░░░░░░░░░░] 25%
Speed: 45.23 mails/s  125.3 KB/s
Est. remaining: 1.4 min
```

---

## Bulk Migration

### Web Dashboard

1. Switch to **Bulk Migration** tab
2. Click **+ Add Account** for each account
3. Or click **Example CSV** to download a template
4. Fill and import the CSV
5. Test all connections
6. Preview and start migration

### Terminal UI

See [Bulk Mode Guide](usage/bulk-mode.md) for CSV format and details.

---

## Keyboard Shortcuts (TUI Only)

| Key | Action |
|-----|--------|
| `Tab` | Next field |
| `Shift+Tab` | Previous field |
| `Space` | Toggle SSL/TLS |
| `Enter` | Submit / Connect |
| `F1` | Provider menu |
| `s` | Start migration |
| `Esc` | Go back |
| `q` | Quit |

---

## Next Steps

- [Web Dashboard Guide](usage/web-interface.md) - Full web interface documentation
- [Manual Mode Details](usage/manual-mode.md) - Single account migration
- [Bulk Mode Details](usage/bulk-mode.md) - Multiple accounts migration
- [Troubleshooting](troubleshooting.md) - If you encounter issues
