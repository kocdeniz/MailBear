# Manual Mode

Manual mode is used for migrating a single account pair (source → destination).

## Available In

- ✅ Terminal UI (TUI)
- ✅ Web Dashboard

---

## Web Dashboard Method (Recommended)

### Step 1: Open Web Interface

```bash
./mailmole -web :8080
```

Go to `http://localhost:8080`

### Step 2: Select Single Account Tab

Click **Single Account** tab in the Connection Setup page.

### Step 3: Use Templates (Optional)

Click template buttons for quick configuration:

| Template | Source | Destination |
|----------|--------|-------------|
| Gmail (Source) | imap.gmail.com | - |
| Gmail (Dest) | - | imap.gmail.com |
| Outlook (Source) | outlook.office365.com | - |
| Outlook (Dest) | - | outlook.office365.com |
| Yandex, iCloud | Available | Available |

### Step 4: Fill Server Details

**Source Server:**
- Host: IMAP server address (e.g., `imap.gmail.com`)
- Port: Usually `993` for SSL
- Username: Full email address
- Password: Account password or app password
- SSL/TLS: Checked (recommended)

**Destination Server:**
Same fields for destination.

### Step 5: Test Connection

Click **Test Connection** for quick test, or **Detailed Test** for comprehensive validation.

A toast notification will appear top-right with results.

### Step 6: Preview

Click **Preview Migration** to see:
- All folders from source account
- Message count per folder
- Estimated total size
- Folders selected for migration

Use checkboxes to select/deselect folders.

### Step 7: Start Migration

Click **Start Migration** button.

Watch progress in real-time:
- Progress bar
- Messages migrated / total
- Transfer speed
- ETA

---

## Terminal UI Method

### When to Use

- Migrating one user's email account
- Testing migration before bulk operation
- One-time migrations
- SSH/headless environments

### How It Works

1. Enter source IMAP server credentials
2. Enter destination IMAP server credentials
3. MailMole connects and lists all folders
4. Review folders and press `s` to start migration
5. Watch real-time progress

### Form Fields

#### Source Server
| Field | Description | Example |
|-------|-------------|---------|
| Host | IMAP server address | `imap.gmail.com` |
| Port | IMAP port (993 for SSL) | `993` |
| SSL/TLS | Enable secure connection | Check/uncheck |
| Username | Full email address | `user@gmail.com` |
| Password | Account password or app password | `xxxx xxxx xxxx xxxx` |

#### Destination Server
Same fields as source server.

### Using Provider Presets

Press `F1` on the Host field to see preset providers:

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

Selecting a provider auto-fills:
- Host address
- Correct port (993 for SSL, 143 for TLS)
- SSL/TLS checkbox

---

## What Gets Migrated

- All folders from source to destination
- Message content (full email)
- Message flags:
  - `\Seen` - Read/Unread status
  - `\Answered` - Answered status
  - `\Flagged` - Starred/Flagged status
  - `\Draft` - Draft status

## What Doesn't Get Migrated

- Contacts
- Calendars
- Filters/Rules
- Signatures

---

## Comparison: Web vs Terminal

| Feature | Web Dashboard | Terminal UI |
|---------|---------------|-------------|
| Setup Speed | ⭐⭐⭐ (Templates) | ⭐⭐ |
| Visual Preview | ✅ Yes | ❌ No |
| Folder Selection | ✅ Checkbox | ❌ All folders |
| Connection Test | ✅ Detailed | Basic |
| SSH/Headless | ❌ No | ✅ Yes |
| Resource Usage | Higher | Minimal |

---

## Tips

- **Web**: Test with a small folder first using Preview feature
- **Terminal**: Use `F1` for quick provider selection
- Both: Ensure both servers allow IMAP access
- Both: Use app passwords for Gmail/Outlook if 2FA is enabled
- Both: Check destination server's storage limits before migration
