package main

import ("fmt"
        "time"
        "os"
        "os/exec"
)

// Agnes Logo
// Ἁγνὴ
const AGNES_LOGO string =
`
      ##     #           #############   ###            ####        ###
      #     ###            ##       ##     ##            #           ##
           #  ##           ##              # ##          #
          #    ##          ##              #   ##        #      ### #######
         #      ##         ##              #     ##      #     # ##       ##
        ###########        ##              #       ##    #       ##       ##
       #          ##       ##              #         ##  #       ##       ##
      #            ##      ##              #           ###       ##       ##
    ####           ####   #####          ####   AGNES    #       ##       ##
                                                                          ##
`

// Define Neural Network
type NeuralNetwork struct {




}

// global constants

// clear the terminal screen
func clear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// intro
func intro() {
    clear()                         // clear the terminal screen
    fmt.Println(AGNES_LOGO)         // print agnes logo
    time.Sleep(time.Second * 2)     // wait for 2 seconds...
    clear()                         // clear the terminal screen
}

// send a message through email
func sendReportMail(subject, body string) {
    clear()
    // indicate that mail is being sent
    fmt.Printf("Sending mail to %s...\n", RCVR_ADD)
    var err error           // error
    var doc bytes.Buffer    // buffer of bytes
    // authorize email
    auth := smtp.PlainAuth("", DIVYA_ADD, DIVYA_PWD, EMAIL_SRV)
    // setup a context for the mail
    context := &SmtpTemplateData{"Divya", DIVYA_ADD, RCVR_ADD, subject, body}
    // set up template
    temp := template.New("Email Template")
    temp, err = temp.Parse(TEMPLATE)
    if err != nil { log.Print("error trying to parse mail template") }
    // write the mail using template and context
    err = temp.Execute(&doc, context)
    if err != nil { log.Print("error trying to execute mail template") }
    // send mail
    err = smtp.SendMail(EMAIL_SRV+":"+strconv.Itoa(PORT), auth, DIVYA_ADD,
                        []string{RCVR_ADD}, doc.Bytes())
    if err != nil {
        log.Print("ERROR: attempting to send a mail ", err)
    } else { fmt.Println("Mail sent successfully!") }
}

// Agnes main function
func main() {
    intro()             // print intro
}
