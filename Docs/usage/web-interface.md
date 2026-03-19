# Web Dashboard Interface

The web dashboard provides a modern, browser-based interface for MailMole with enhanced features.

## Launching Web Dashboard

```bash
# Start with both TUI and Web
./mailmole -web :8080

# Start Web only (no terminal interface)
./mailmole -web :8080 -web-only
```

Access at: `http://localhost:8080`

## Features Overview

### 1. Visual Interface

- Clean, modern dark theme
- Sidebar navigation
- Responsive design (works on mobile)
- Toast notifications

### 2. Migration Templates

19 pre-configured templates for common scenarios:

**Same Provider:**
- Gmail → Gmail
- Outlook → Outlook
- Yandex → Yandex
- iCloud → iCloud
- Yahoo → Yahoo

**Cross Provider:**
- Gmail → Outlook
- Outlook → Gmail
- Gmail → Yandex
- Yandex → Gmail
- And 10 more combinations...

### 3. Import / Export

**Import:**
- Upload CSV files with account data
- JSON configuration import
- Download Example CSV template

**Export:**
- Save current configuration as JSON
- Export validation results as CSV
- Export bulk accounts as CSV

### 4. Connection Validation

Two testing modes:

**Quick Test:**
- Tests connection to both servers
- Shows pass/fail in toast notification

**Detailed Test:**
- Comprehensive validation report
- Per-account status table
- Error details and suggestions
- Exportable results

### 5. Preview Mode

Before migration:
- View all folders from source
- See message counts
- View estimated sizes
- Select/deselect folders
- View estimated duration

### 6. Real-time Monitoring

During migration:
- Live progress bar
- Messages migrated counter
- Transfer speed (mails/s)
- ETA calculation
- Per-account progress
- Activity logs

### 7. Multi-Language Support

Currently supported:
- 🇬🇧 English (Default)
- 🇹🇷 Turkish (Türkçe)

Language selector in top-right corner.

### 8. Scheduling (Beta)

Schedule migrations for later:
- Set start date/time
- Repeat options: Daily, Weekly, Monthly
- View scheduled jobs

## Navigation

### Sidebar Menu

```
📧 MIGRATION
  Connection Setup
  Preview
  Progress

📜 HISTORY
  Activity Logs

🌍 LANGUAGE / DIL
  [English ▼]
```

### Pages

1. **Connection Setup**
   - Single Account mode
   - Bulk Migration mode
   - Templates
   - Import/Export

2. **Preview**
   - Folder list with checkboxes
   - Message counts
   - Size estimates
   - Start migration button

3. **Progress**
   - Live progress tracking
   - Speed metrics
   - Account status
   - Stop button

4. **Activity Logs**
   - Migration history
   - Log levels (INFO, WARN, ERROR)
   - Timestamped entries

## Usage Examples

### Example 1: Single Account Migration

1. Go to **Connection Setup**
2. Select **Single Account** tab
3. Click **Gmail (Source)** template
4. Click **Outlook (Dest)** template
5. Fill in usernames and passwords
6. Click **Test Connection**
7. Click **Preview Migration**
8. Select folders to migrate
9. Click **Start Migration**

### Example 2: Bulk Migration with CSV

1. Go to **Connection Setup**
2. Select **Bulk Migration** tab
3. Click **Example CSV** button
4. Edit downloaded CSV with your accounts
5. Click **Import CSV** and select file
6. Click **Test All Connections**
7. Click **Preview All Accounts**
8. Click **Start Migration**

### Example 3: Using Templates

1. In Bulk Migration tab
2. Click **Templates** button
3. Select "Gmail → Outlook"
4. All account rows are auto-configured
5. Just add usernames and passwords
6. Test and start migration

## Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `Tab` | Navigate between fields |
| `Enter` | Submit/Activate button |
| `Esc` | Close modal/popup |

## Browser Compatibility

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## Security Notes

⚠️ **Important:**
- No built-in authentication
- Only run on trusted networks
- Use firewall to restrict access if needed
- Credentials are transmitted to IMAP servers only
- No data is sent to third parties

## Troubleshooting

### Dashboard won't load
- Check if port 8080 is available
- Try different port: `./mailmole -web :3000`
- Check firewall settings

### Cannot connect to servers
- Verify IMAP is enabled on both servers
- Check SSL/TLS settings
- Use app passwords for Gmail/Outlook with 2FA

### Language not changing
- Clear browser cache
- Refresh page after language change
- Check browser console for errors

## Tips

- Use **Detailed Test** before large migrations
- Download **Example CSV** as template
- Use **Preview** to verify folder selection
- Monitor **Activity Logs** for issues
- Export results for reporting
