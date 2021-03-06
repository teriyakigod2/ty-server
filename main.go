package main

import (
    "fmt"
    "os"
    "gopkg.in/gomail.v2" //go get gopkg.in/gomail.v2
)

const (
    // Replace sender@example.com with your "From" address. 
    // This address must be verified with Amazon SES.
    Sender = "shulhindvst@gmail.com"
    SenderName = "Sender Name"
    
    // Replace recipient@example.com with a "To" address. If your account 
    // is still in the sandbox, this address must be verified.
    Recipient = "shulhindvst@gmail.com"
    
    // The name of the configuration set to use for this message.
    // If you comment out or remove this variable, you will also need to
    // comment out or remove the header below.
    //ConfigSet = "ConfigSet"
    
    // If you're using Amazon SES in an AWS Region other than US West (Oregon), 
    // replace email-smtp.us-west-2.amazonaws.com with the Amazon SES SMTP  
    // endpoint in the appropriate region.
    Host = "email-smtp.eu-west-1.amazonaws.com"
    Port = 587 
    
    // The subject line for the email.
    Subject = "Amazon SES Test (Gomail)"
    
    // The HTML body for the email.
    HtmlBody =  "<html><head><title>SES Sample Email</title></head><body>" +
                "<h1>Amazon SES Test Email (Gomail)</h1>" +
                "<p>This email was sent with " +
                "<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using " +
                "the <a href='https://github.com/go-gomail/gomail/'>Gomail " +
                "package</a> for <a href='https://golang.org/'>Go</a>.</p>" +
                "</body></html>"
    
    //The email body for recipients with non-HTML email clients.
    TextBody = "This email was sent with Amazon SES using the Gomail package."
    
    // The tags to apply to this message. Separate multiple key-value pairs
    // with commas.
    // If you comment out or remove this variable, you will also need to
    // comment out or remove the header on line 80.
    Tags = "genre=test,genre2=test2"
    
    // The character encoding for the email.
    CharSet = "UTF-8"

)

func main() {
    
    // Create a new message.
    m := gomail.NewMessage()
    
    // Set the main email part to use HTML.
    m.SetBody("text/html", HtmlBody)
    
    // Set the alternative part to plain text.
    m.AddAlternative("text/plain", TextBody)

    // Construct the message headers, including a Configuration Set and a Tag.
    m.SetHeaders(map[string][]string{
        "From": {m.FormatAddress(Sender,SenderName)},
        "To": {Recipient},
        "Subject": {Subject},
        // Comment or remove the next line if you are not using a configuration set
    //    "X-SES-CONFIGURATION-SET": {ConfigSet},
        // Comment or remove the next line if you are not using custom tags
        "X-SES-MESSAGE-TAGS": {Tags},
    })

    // Send the email.
    d := gomail.NewPlainDialer(Host, Port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
    
    // Display an error message if something goes wrong; otherwise, 
    // display a message confirming that the message was sent.
    if err := d.DialAndSend(m); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Email sent!")
    }
}