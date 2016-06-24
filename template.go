package mobileconfig

var mailTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>PayloadContent</key>
    <array>
      <dict>
        {{- if .AccountDescription}}
        <key>EmailAccountDescription</key>
        <string>{{.AccountDescription}}</string>
        {{- end}}

        {{- if .Imap}}
        <key>EmailAccountType</key>
        <string>EmailTypeIMAP</string>

        <key>EmailAddress</key>
        <string>{{.EmailAddress}}</string>

        <key>IncomingMailServerAuthentication</key>
        <string>EmailAuthPassword</string>

        <key>IncomingMailServerHostName</key>
        <string>{{.Imap.Hostname}}</string>

        <key>IncomingMailServerPortNumber</key>
        <integer>{{.Imap.Port}}</integer>

        <key>IncomingMailServerUseSSL</key>
        {{- if .Imap.Tls}}
        <true/>
        {{- else}}
        <false/>
        {{- end}}

        <key>IncomingMailServerUsername</key>
        <string>{{.Imap.Username}}</string>

        <key>IncomingPassword</key>
        <string>{{.Imap.Password}}</string>
        {{- end}}

        {{ if .Smtp}}
        <key>OutgoingMailServerAuthentication</key>
        <string>{{if .Smtp.Username}}EmailAuthPassword{{else}}EmailAuthNone{{end}}</string>

        <key>OutgoingMailServerHostName</key>
        <string>{{.Smtp.Hostname}}</string>

        <key>OutgoingMailServerPortNumber</key>
        <integer>{{.Smtp.Port}}</integer>

        <key>OutgoingMailServerUseSSL</key>
        {{- if .Smtp.Tls}}
        <true/>
        {{- else}}
        <false/>
        {{- end}}

        {{- if .Smtp.Username}}
        <key>OutgoingMailServerUsername</key>
        <string>{{.Smtp.Username}}</string>
        {{- end}}

        {{- if .Smtp.Password}}
        <key>OutgoingPassword</key>
        <string>{{.Smtp.Password}}</string>
        {{- else}}
        <key>OutgoingPasswordSameAsIncomingPassword</key>
        <true/>
        {{- end}}
        {{end}}

        <key>PayloadDescription</key>
        <string>Configures email account.</string>

        <key>PayloadDisplayName</key>
        <string>{{.DisplayName}}</string>

        <key>PayloadIdentifier</key>
        <string>{{.Identifier}}</string>

        {{- if .Organization}}
        <key>PayloadOrganization</key>
        <string>{{.Organization}}</string>
        {{- end}}

        <key>PayloadType</key>
        <string>com.apple.mail.managed</string>

        <key>PayloadUUID</key>
        <string>{{.ContentUuid}}</string>

        <key>PayloadVersion</key>
        <integer>1</integer>

        <key>PreventAppSheet</key>
        <false/>

        <key>PreventMove</key>
        <false/>

        <key>SMIMEEnabled</key>
        <false/>
      </dict>
    </array>

    <key>PayloadDescription</key>
    <string>{{if .Description}}{{.Description}}{{else}}Install this profile to auto configure email account for {{.EmailAddress}}.{{- end}}</string>

    <key>PayloadDisplayName</key>
    <string>{{.DisplayName}}</string>

    <key>PayloadIdentifier</key>
    <string>{{.Identifier}}</string>

    {{- if .Organization}}
    <key>PayloadOrganization</key>
    <string>{{.Organization}}</string>
    {{- end}}

    <key>PayloadRemovalDisallowed</key>
    <false/>

    <key>PayloadType</key>
    <string>Configuration</string>

    <key>PayloadUUID</key>
    <string>{{.Uuid}}</string>

    <key>PayloadVersion</key>
    <integer>1</integer>
  </dict>
</plist>`
