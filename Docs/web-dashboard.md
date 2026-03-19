# Web Dashboard

MailMole includes a modern, browser-based web dashboard for managing email migrations.

## Overview

The web dashboard provides an intuitive visual interface for:
- Configuring migration settings
- Managing multiple accounts
- Previewing folders before migration
- Monitoring progress in real-time
- Importing/exporting account lists

## Accessing the Web Dashboard

### Starting the Web Dashboard

```bash
# Start with both TUI and Web Dashboard
./mailmole -web :8080

# Start Web Dashboard only (no terminal interface)
./mailmole -web :8080 -web-only
```

Then open your browser and navigate to: `http://localhost:8080`

### Default Port

The default port is `8080`, but you can use any available port:

```bash
./mailmole -web :3000    # Use port 3000
./mailmole -web :9000    # Use port 9000
```

## Features

### 1. Visual Interface

- **Modern Design**: Clean, dark-themed interface optimized for long-running migrations
- **Responsive Layout**: Works on desktop, tablet, and mobile devices
- **Real-time Updates**: Live progress tracking with Server-Sent Events (SSE)
- **Toast Notifications**: Immediate feedback for actions and status changes

### 2. Migration Modes

#### Single Account Migration
Migrate one email account at a time:
- Enter source and destination server details
- Test connections before migration
- Preview folders and select which to migrate
- Monitor progress in real-time

#### Bulk Migration
Migrate multiple accounts simultaneously:
- Add accounts manually or import from CSV
- Each account can have different source/destination servers
- Test all connections before starting
- Monitor all accounts in parallel

### 3. Quick Templates

Pre-configured settings for popular email providers:

**Available Templates:**
- Gmail (Source/Destination)
- Outlook/Office 365 (Source/Destination)
- Yandex (Source/Destination)
- iCloud (Source/Destination)
- Yahoo (Source/Destination)

**Cross-Provider Templates:**
- Gmail → Outlook
- Outlook → Gmail
- Gmail → Yandex
- Yandex → Gmail
- And more...

Click a template to auto-fill server settings (host, port, SSL).

### 4. Import / Export

#### Import Accounts
Upload CSV files with account information:

**Format:**
```csv
src_host,src_port,src_user,src_pass,dst_host,dst_port,dst_user,dst_pass
imap.gmail.com,993,user1@gmail.com,pass1,outlook.office365.com,993,user1@company.com,pass2
imap.yandex.com,993,user2@yandex.com,pass2,imap.gmail.com,993,user2@newgmail.com,pass3
```

**Steps:**
1. Click "📄 Example CSV" to download a template
2. Edit the CSV file with your accounts
3. Click "📥 Import Accounts (CSV/JSON)"
4. Select your edited file

#### Export Configuration
Save current settings:
- Export as JSON for backup
- Export validation results as CSV
- Export bulk accounts as CSV

### 5. Connection Validation

Two testing modes available:

#### Quick Test
- Tests connection to source and destination
- Shows result in toast notification
- Good for quick verification

#### Detailed Test
- Comprehensive validation report
- Per-account status table
- Success/failure indicators
- Error details and suggestions
- Export results to CSV

### 6. Preview Mode

Before starting migration:
- View all folders from source account
- See message count per folder
- View estimated size
- Select/deselect specific folders
- View estimated duration

### 7. Real-time Monitoring

During migration:
- Live progress bar
- Messages migrated / total count
- Transfer speed (emails/second)
- Estimated time remaining (ETA)
- Per-account progress tracking
- Activity log with timestamps

### 8. Multi-Language Support

Available languages:
- 🇬🇧 English (Default)
- 🇹🇷 Turkish (Türkçe)

**To change language:**
1. Look for the language selector in the top-right corner
2. Select your preferred language
3. Page automatically updates

Language preference is saved in browser storage.

### 9. Scheduling (Beta)

Schedule migrations for later execution:
1. Enable scheduling in Connection Setup
2. Set start date and time
3. Choose repeat option: One Time, Daily, Weekly, or Monthly
4. Migration starts automatically at scheduled time

## Interface Layout

### Sidebar Navigation

```
📧 MIGRATION
  ├── Connection Setup
  ├── Preview
  └── Progress

📜 HISTORY
  └── Activity Logs

🌍 LANGUAGE / DIL
  └── [English ▼]
```

### Main Pages

#### Connection Setup
- Choose between Single Account or Bulk Migration
- Configure source and destination servers
- Use Quick Templates for common providers
- Import accounts from CSV
- Schedule migrations

#### Preview
- Review folders before migration
- Select/deselect folders
- View message counts and sizes
- Start migration

#### Progress
- Real-time migration status
- Progress bars and statistics
- Transfer speed metrics
- Stop migration button

#### Activity Logs
- Complete migration history
- Filter by log level (INFO, WARN, ERROR)
- Timestamped entries
- Export logs

## Usage Examples

### Example 1: Single Account (Gmail to Outlook)

1. Go to **Connection Setup**
2. Select **Single Account** tab
3. Click **Gmail (Source)** template
4. Click **Outlook (Dest)** template
5. Enter your email addresses and passwords
6. Click **Test Connection**
7. Click **Preview Migration**
8. Select folders to migrate
9. Click **Start Migration**

### Example 2: Bulk Migration with CSV

1. Go to **Connection Setup**
2. Select **Bulk Migration** tab
3. Click **Example CSV** button
4. Edit the downloaded file with your accounts
5. Click **Import CSV** and select your file
6. Review imported accounts
7. Click **Test All Connections**
8. Click **Preview All Accounts**
9. Click **Start Migration**

### Example 3: Using Templates

1. In Bulk Migration tab
2. Click **Templates** button
3. Select "Gmail → Outlook" (or any other)
4. Server settings are auto-filled for all accounts
5. Just add usernames and passwords
6. Test and start migration

## Browser Compatibility

Supported browsers:
- ✅ Chrome 90+
- ✅ Firefox 88+
- ✅ Safari 14+
- ✅ Edge 90+

## Security Considerations

⚠️ **Important:**

1. **No Built-in Authentication**: The web dashboard has no login system
2. **Trusted Networks Only**: Run only on trusted networks
3. **Firewall Protection**: Use firewall rules to restrict access if needed
4. **Reverse Proxy**: Consider using nginx/Apache with auth for remote access
5. **HTTPS**: Use HTTPS in production (reverse proxy recommended)

### Securing Remote Access

If you need remote access, use a reverse proxy with authentication:

**Example with nginx:**
```nginx
server {
    listen 443 ssl;
    server_name mailmole.yourdomain.com;
    
    auth_basic "Restricted";
    auth_basic_user_file /etc/nginx/.htpasswd;
    
    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

## Troubleshooting

### Dashboard won't load
- Check if port is available: `lsof -i :8080`
- Try a different port: `./mailmole -web :3000`
- Check firewall settings
- Clear browser cache

### Connection test fails
- Verify IMAP is enabled on both servers
- Check SSL/TLS settings
- Use app passwords for Gmail/Outlook with 2FA
- Verify network connectivity

### Language not changing
- Clear browser cache (Ctrl+Shift+R or Cmd+Shift+R)
- Check browser console for errors
- Ensure localStorage is enabled

### Import fails
- Verify CSV format matches the example
- Check for special characters in passwords
- Ensure no empty lines in CSV
- Use UTF-8 encoding for CSV files

## API Endpoints

The web dashboard exposes these API endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Web dashboard interface |
| `/api/status` | GET | Migration status |
| `/api/logs` | GET | Activity logs |
| `/api/test-connection` | POST | Test connection |
| `/api/preview` | POST | Get folder preview |
| `/api/bulk-preview` | POST | Preview multiple accounts |
| `/api/validate` | POST | Validate connections |
| `/api/schedule` | POST | Schedule migration |
| `/api/schedules` | GET | List scheduled jobs |
| `/api/start` | POST | Start migration |
| `/api/stop` | POST | Stop migration |
| `/ws` | GET | SSE endpoint |
| `/locales/{lang}.json` | GET | Translation files |

## Docker Deployment

Deploy with Docker:

```bash
# Using docker-compose
docker-compose up -d

# Or manually
docker build -t mailmole .
docker run -d -p 8080:8080 mailmole
```

## Tips

- Use **Detailed Test** before large migrations
- Download **Example CSV** as a starting template
- Use **Preview** to verify folder selection
- Monitor **Activity Logs** for any issues
- Export results for reporting and auditing
- Schedule migrations during off-peak hours

## See Also

- [Getting Started](getting-started.md) - Quick start guide
- [Manual Mode](usage/manual-mode.md) - Single account migration
- [Bulk Mode](usage/bulk-mode.md) - Multiple accounts migration
- [FAQ](faq.md) - Frequently asked questions
