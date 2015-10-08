package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
        "log"
        "net/smtp"
        "bytes"
        "text/template"
        "strconv"
)

// Agnes logo in small case
const AGNES_LOGO string =
`
    ##    #######   ##   ###            ##    ##        ###   ##  ########
    #   ##       ## ##     ##         ##    ####        ##   # ###       ###
       ##         ###       ##       #         ##       #      ##         ##
       ##         ###        ##    #            ##     #       ##         ##
       ##         ###         ##  #              ##   #        ##         ##
        ##       ## ##         ###                ## #         ##         ##
          #######    ###       ##       AGNES      ##          ##         ##
                              ###                                         ##
                              ##                                          ##
`

// global constants

// define a Neural Network
type NeuralNetwork []Layer
// define a neuron layer
type Layer []Neuron
// define a Neuron
type Neuron struct {
    input float64       // sum of all the input
    bias float64        // bias for desired results
    weight []float64    // a set of weights that adds up to 1.0
    output float64      // determine 0 or 1 as output
}

// sum up all the inputs from the previous layer
func (this *Neuron) pullInput(prev Layer) {
    // add up outputs from previous layer and assign to input
    for _, neuron := range prev { this.input += sigmoid(neuron.output) }
    this.input += this.bias
}

//

// Sigmoid function: choose between 0 and 1 based on input value
func sigmoid(x float64) float64 {
    // sigmoid(0.0) = 0.5
    return 1.0 / (1.0 + math.Pow(math.E, -x))
}

// generate a new neural network
func genNeuralNet() *NeuralNetwork {

}

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

// **** Send report through email ****
// Agnes's mail address and password and receiver
const (
    AGNES_ADD = ""                      // email address
    AGNES_PWD = ""                      // password
    EMAIL_SRV = "smtp.gmail.com"        // email server
    RCVR_ADD = "jinyeom95@gmail.com"    // receiver's address
    PORT = 587                          // port
)

// email template
const TEMPLATE =
`From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

With love,
{{.Name}}`

// smtp email template data
type SmtpTemplateData struct {
    Name    string      // sender's name
    From    string      // sender address
    To      string      // receiver address
    Subject string      // subject of the mail
    Body    string      // body of the mail
}

// send a message through email
func sendReportMail(subject, body string) {
    clear()
    // indicate that mail is being sent
    fmt.Printf("Sending mail to %s...\n", RCVR_ADD)
    var err error           // error
    var doc bytes.Buffer    // buffer of bytes
    // authorize email
    auth := smtp.PlainAuth("", AGNES_ADD, AGNES_PWD, EMAIL_SRV)
    // setup a context for the mail
    context := &SmtpTemplateData{"Agnes", AGNES_ADD, RCVR_ADD, subject, body}
    // set up template
    temp := template.New("Email Template")
    temp, err = temp.Parse(TEMPLATE)
    if err != nil { log.Print("error trying to parse mail template") }
    // write the mail using template and context
    err = temp.Execute(&doc, context)
    if err != nil { log.Print("error trying to execute mail template") }
    // send mail
    err = smtp.SendMail(EMAIL_SRV+":"+strconv.Itoa(PORT), auth, AGNES_ADD,
                        []string{RCVR_ADD}, doc.Bytes())
    if err != nil {
        log.Print("ERROR: attempting to send a mail ", err)
    } else { fmt.Println("Mail sent successfully!") }
}

// Agnes main function
func main() {
    intro()             // print intro

}
