package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
        // "log"
        // "net/smtp"
        // "bytes"
        // "text/template"
        // "strconv"
)

// Agnes logo in small case
const AGNES_LOGO string =
`
    ##    #######   ##  ###            ##   ##        ###   ##  ########
    #   ##       ## ##    ##         ##   ####        ##   # ###       ###
       ##         ###      ##       #        ##       #      ##         ##
       ##         ###       ##    #           ##     #       ##         ##
       ##         ###        ##  #             ##   #        ##         ##
        ##       ## ##        ###               ## #         ##         ##
          #######    ###      ##      AGNES      ##          ##         ##
                             ###                                        ##
                             ##                                         ##
`

// global constants
I_LAYER = 784           // INPUT_LAYER: size of a box (784 = 28 * 28)
H_LAYER_1 = 100         // HIDDEN LAYER 1
H_LAYER_2 = 100         // HIDDEN LAYER 2
H_LAYER_3 = 100         // HIDDEN LAYER 3
O_LAYER = 10            // OUTPUT_LAYER:

// define a Neural Network
type NeuralNetwork []Layer
// define a neuron layer
type Layer []Neuron
// define a Neuron
type Neuron struct {
    input float64
    output float64
}

// generate and initialize a neural network
func genNeuralNet() *NeuralNetwork {

}

// generate and initialize a neuron layer
func genLayer() *Layer {

}

// generate and initialize a neuron
func genNeuron() *Neuron {

}

// Sigmoid function: choose between 0 and 1 based on input value
func sigmoid(x float64) float64 {
    // sigmoid(0.0) = 0.5
    return 1.0 / (1.0 + math.Pow(math.E, -x))
}

// feedforward neural network where information flows only one way
func feedforwardNeuralNet() {
    // input file
    // read data
    brain := genNeuralNet()
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

// Agnes main function
func main() {
    intro()             // print intro

}
