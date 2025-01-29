# How to Create an App Password for Custom Email Sending

This guide will walk you through the steps to create an App Password for sending emails using Gmail in a custom application. An App Password is a special 16-digit password used for signing into your Google account in third-party applications that don’t support 2-Step Verification.

## Prerequisites

1. A Google account with **2-Step Verification** enabled.
2. Access to the [App Passwords page](https://myaccount.google.com/apppasswords?pli=1&rapt=AEjHL4PsOQgHuaajaPZxH96T_ZL-QWYrvkLZd6n-IFIpsiwguk5q7p4jlYsVCLHYhZFRFOouhKW87L9zUbuz2FxqBOzlD6Tbm1_j8B9yM8bEn6xtzfMXWAA).

## Steps to Create an App Password

### 1. Enable 2-Step Verification
If you haven't already enabled 2-Step Verification for your Google account, you need to do so:

- Go to [Google 2-Step Verification](https://myaccount.google.com/security-checkup).
- Follow the steps to enable 2-Step Verification for your account.

### 2. Go to the App Passwords Page
- Open the [App Passwords page](https://myaccount.google.com/apppasswords?pli=1&rapt=AEjHL4PsOQgHuaajaPZxH96T_ZL-QWYrvkLZd6n-IFIpsiwguk5q7p4jlYsVCLHYhZFRFOouhKW87L9zUbuz2FxqBOzlD6Tbm1_j8B9yM8bEn6xtzfMXWAA).

### 3. Select the App and Device
- From the "Select App" dropdown menu, choose the app you will be using to send emails (e.g., "Mail").
- From the "Select Device" dropdown, choose the device you are using (e.g., "Windows Computer").

### 4. Generate the App Password
- Click on **Generate** to create the app password. A 16-digit password will appear on the screen. Copy this password.

### 5. Use the App Password in Your Application
- In your custom email-sending application, replace your usual Google account password with the newly generated 16-digit App Password.
- For example, if you are using a library like **SMTP** or any email-sending tool, use this password for authentication.

### 6. Save the App Password
For security reasons, make sure to save this password securely. You will not be able to view it again.

### 7. Test Email Sending
Now you can use your application to send emails with your Google account using the generated app password. Make sure the SMTP settings in your application are configured correctly (SMTP server: `smtp.gmail.com`, Port: `587`).

## Conclusion

By following these steps, you have successfully created an App Password that allows your custom application to send emails via Gmail while keeping your account secure with 2-Step Verification enabled.

**Note**: If you ever need to revoke the app password or generate a new one, you can do so from the App Passwords page in your Google account.

## Example Code for Sending Email using Go

Here’s an example of how you can use the App Password in a Go application to send an email:

```go
package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// Sender's email address and the App Password
	email := "myemail@gmail.com"     // Replace with your Gmail address
	pass := "your_app_password"      // Replace with the generated App Password

	// Recipient email address
	to := "recipient@example.com"    // Replace with the recipient's email address

	// Email subject and content
	subject := "Test Email"
	content := "This is a test email sent from Go."

	// SMTP server and port settings for Gmail
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Prepare the email message
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", email, to, subject, content)

	// Authentication using the App Password
	auth := smtp.PlainAuth("", email, pass, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	// Success message
	fmt.Println("Email sent successfully!")
}
