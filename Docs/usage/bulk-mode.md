# Bulk Migration

Bulk mode allows migrating multiple account pairs from a CSV file.

## Available In

- ✅ Terminal UI (TUI) - CSV file based
- ✅ Web Dashboard - CSV or manual entry

---

## Web Dashboard Method (Recommended)

### Step 1: Open Bulk Migration Tab

```bash
./mailmole -web :8080
```

Go to `http://localhost:8080` → **Bulk Migration** tab

### Step 2: Add Accounts

**Option A: Manual Entry**

Click **+ Add Account** for each account pair:

```
Account #1:
  Source Host: imap.gmail.com
  Source Port: 993
  Source User: user1@gmail.com
  Source Pass: password1
  
  Dest Host: outlook.office365.com
  Dest Port: 993
  Dest User: user1@company.com
  Dest Pass: password2
```

Each account can have different source and destination servers!

**Option B: Import from CSV**

1. Click **Example CSV** button
2. Download `mailmole-example-import.csv`
3. Edit with your accounts
4. Click **Import CSV** and select file

### Step 3: Use Templates (Optional)

Click **Templates** button for quick configuration:

| Template | Description |
|----------|-------------|
| Gmail → Gmail | Both servers Gmail |
| Outlook → Outlook | Both servers Outlook |
| Gmail → Outlook | Gmail to Outlook migration |
| Outlook → Gmail | Outlook to Gmail migration |
| And 15 more... | Various combinations |

### Step 4: Test All Connections

Click **Test All Connections** or **Detailed Test**

Results shown in:
- Toast notification (top-right)
- Validation page (detailed table)

### Step 5: Preview

Click **Preview All Accounts** to see:
- All accounts and their folders
- Total message count
- Estimated size

### Step 6: Start Migration

Click **Start Migration**

Monitor progress for all accounts simultaneously.

---

## CSV File Format

### For Terminal UI (Simple Format)

Create a `.csv` or `.txt` file:

```text
src_user,src_pass,dst_user,dst_pass
user1@source.com,password1,user1@dest.com,password1
user2@source.com,password2,user2@dest.com,password2
user3@source.com,password3,user3@dest.com,password3
```

**Note:** Source and destination servers are set globally in TUI.

### For Web Dashboard (Advanced Format)

Full control with custom servers per account:

```text
src_host,src_port,src_user,src_pass,dst_host,dst_port,dst_user,dst_pass
imap.gmail.com,993,user1@gmail.com,pass1,outlook.office365.com,993,user1@company.com,pass2
imap.yandex.com,993,user2@yandex.com,pass2,imap.gmail.com,993,user2@newgmail.com,pass3
mail.custom.com,993,user3@old.com,pass3,mail.new.com,993,user3@new.com,pass4
```

### File Rules

- One account pair per line
- Comma-separated fields
- Lines starting with `#` are comments
- Empty lines are ignored
- Each account can have different servers (Web only)

### Example CSV

```text
# Company migration to Office 365
alice@oldcompany.com,Pass123,alice@newcompany.com,NewPass456
bob@oldcompany.com,Pass123,bob@newcompany.com,NewPass456
carol@oldcompany.com,Pass123,carol@newcompany.com,NewPass456

# Personal accounts mixed
eren@gmail.com,oldpass,outlook.com,yenişifre
mehmet@yandex.com,eski,mehmet@gmail.com,yeni
```

---

## Terminal UI Bulk Mode

### How It Works

1. Enter global source and destination server settings
2. Provide path to CSV file containing account pairs
3. MailMole tests all connections before migration
4. Migration runs automatically for all accounts
5. Failed accounts are skipped, successful ones continue

### Form Fields

| Field | Description | Example |
|-------|-------------|---------|
| Global Source Host | Source mail server address | `mail.source.com` |
| Global Dest Host | Destination mail server address | `mail.dest.com` |
| Accounts File | Path to CSV/TXT file | `/path/to/accounts.csv` |

**Note:** In TUI, all accounts use the same source and destination servers.

### Progress Tracking

The dashboard shows:
- Overall progress (messages migrated)
- Current account / Total accounts
- Transfer speed (mails/s and KB/s)
- Estimated time remaining

---

## State File

Bulk migrations create a `migration_state.json` file to track completed accounts and enable resume functionality.

## Comparison: Web vs Terminal

| Feature | Web Dashboard | Terminal UI |
|---------|---------------|-------------|
| Account Entry | Form + CSV import | CSV only |
| Per-Account Servers | ✅ Yes | ❌ Global only |
| Templates | ✅ 19 templates | ❌ No |
| Connection Testing | ✅ Detailed table | Basic log |
| Visual Preview | ✅ Yes | ❌ No |
| Ease of Use | ⭐⭐⭐ | ⭐⭐ |

---

## Tips

### Web Dashboard
- Use Templates for common migration scenarios
- Download Example CSV as starting point
- Test connections before starting migration
- Use Preview to verify folder selection

### Terminal UI
- Use absolute paths for CSV file
- Test with 1-2 accounts first
- Check server rate limits before large migrations
- Keep CSV file secure (contains passwords)

### Both
- Keep CSV file secure (contains passwords)
- Use app passwords for Gmail/Outlook
- Monitor logs for any issues
- Check destination storage limits
