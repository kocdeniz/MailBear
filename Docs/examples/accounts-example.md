# Example: Bulk Migration Accounts File

This file contains example account pairs for bulk migration.

## Format Options

### Option 1: Simple Format (Terminal UI)

For TUI bulk mode where servers are set globally:

```
src_user,src_pass,dst_user,dst_pass
```

**Example:**
```csv
# Example accounts - replace with your own
alice@oldsrv.com,Password123,alice@newsrv.com,NewPassword456
bob@oldsrv.com,Password123,bob@newsrv.com,NewPassword456
carol@oldsrv.com,Password123,carol@newsrv.com,NewPassword456
```

### Option 2: Advanced Format (Web Dashboard)

For Web Dashboard where each account can have different servers:

```
src_host,src_port,src_user,src_pass,dst_host,dst_port,dst_user,dst_pass
```

**Example:**
```csv
# Company migration to Office 365
imap.gmail.com,993,alice@gmail.com,Pass123,outlook.office365.com,993,alice@company.com,NewPass456
imap.yandex.com,993,bob@yandex.com,Pass123,outlook.office365.com,993,bob@company.com,NewPass456
mail.oldcompany.com,993,carol@oldcompany.com,Pass123,outlook.office365.com,993,carol@company.com,NewPass456
```

## Web Dashboard: Download Example

In the Web Dashboard, click **"📄 Example CSV"** button to download a pre-filled template:

```csv
src_host,src_port,src_user,src_pass,dst_host,dst_port,dst_user,dst_pass
imap.gmail.com,993,eren@oldgmail.com,mypassword123,outlook.office365.com,993,eren@company.com,mypassword456
imap.yandex.com,993,ahmet@yandex.com,pass123,imap.gmail.com,993,ahmet@newgmail.com,pass456
mail.eski.com,993,mehmet@eski.com,secret1,mail.yeni.com,993,mehmet@yeni.com,secret2
```

## Rules

### Simple Format (4 columns)
- One account pair per line
- Four comma-separated fields
- Lines starting with `#` are comments
- Empty lines are ignored
- No spaces between commas

### Advanced Format (8 columns)
- One account per line
- Eight comma-separated fields
- Each account can have different servers
- Lines starting with `#` are comments
- Empty lines are ignored

## Usage

### Terminal UI
1. Create your accounts file (simple format)
2. In MailMole TUI, select "Bulk Migration"
3. Enter global server details
4. Enter the path to your file

### Web Dashboard
1. Click **"📄 Example CSV"** to download template, OR
2. Create your own CSV file (advanced format)
3. Click **"📥 Import Accounts"**
4. Select your CSV file
5. Review imported accounts
6. Test connections and start migration

## Common Scenarios

### Scenario 1: Gmail to Office 365 (Same for all)

**Simple format (TUI):**
```csv
user1@gmail.com,pass1,user1@company.com,pass2
user2@gmail.com,pass1,user2@company.com,pass2
```

Set servers globally in TUI:
- Source: imap.gmail.com
- Dest: outlook.office365.com

### Scenario 2: Mixed Sources to Office 365

**Advanced format (Web):**
```csv
imap.gmail.com,993,user1@gmail.com,pass1,outlook.office365.com,993,user1@company.com,pass2
imap.yandex.com,993,user2@yandex.com,pass1,outlook.office365.com,993,user2@company.com,pass2
imap.gmail.com,993,user3@gmail.com,pass1,outlook.office365.com,993,user3@company.com,pass2
```

### Scenario 3: Different Sources to Different Destinations

**Advanced format (Web):**
```csv
imap.gmail.com,993,alice@gmail.com,pass1,imap.gmail.com,993,alice@newgmail.com,pass2
outlook.office365.com,993,bob@oldcompany.com,pass1,outlook.office365.com,993,bob@newcompany.com,pass2
imap.yandex.com,993,carol@yandex.com,pass1,imap.gmail.com,993,carol@gmail.com,pass2
```

## Security Tips

⚠️ **Important:**
- Keep CSV files secure (they contain passwords)
- Delete after migration
- Don't commit to version control
- Use file permissions to restrict access
- Consider using environment variables for passwords in scripts

## Download Template

Use the Web Dashboard's **"📄 Example CSV"** button to get the latest template with correct formatting.
