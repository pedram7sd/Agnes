package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
)

// Agnes logo in small case
const AGNES_LOGO string =
`
    ##    #######   ##  ###            ##   ##        ###   ##  ########
    #   ##       ## ##    ##         ##   ####        ##   # ###       ###
       ##         ###      ##       #        ##       #      ##         ##
       ##         ###       ##    #           ##     #       ##         ##
       ##         ###        ##  #             ##   #        ##         ##
        ##       ## ##        ###   AGNES I/O   ## #         ##         ##
          #######    ###      ##     v.0.1.0     ##          ##         ##
                             ###                                        ##
                             ##                                         ##
`

// global constants
const (
    I_LAYER = 784   // INPUT_LAYER: 784 = 28 * 28
    H_LAYER = 15    // HIDDEN LAYER: 15 neurons
    O_LAYER = 10    // OUTPUT_LAYER: from 0 to 9
)

// define a Neural Network
type NeuralNetwork struct {
    inputLayer []*Neuron    // input layer
    hiddenLayer []*Neuron   // hidden layer
    outputLayer []*Neuron   // output layer
}

// define a Neuron
type Neuron struct {
    input float64   // sum of previous layer's outputs
    output float64  // result of sigmoid function
}

// genrate a neural network
func genNeuralNetwork() *NeuralNetwork {
    var i []*Neuron = genLayer(I_LAYER)     // input layer
    var h []*Neuron = genLayer(H_LAYER)     // hidden layer
    var o []*Neuron = genLayer(O_LAYER)     // output layer
    return &NeuralNetwork{i, h, o}
}

// generate a neural layer
func genLayer(numNeuron int) []*Neuron {
    var layer []*Neuron = make([]*Neuron, numNeuron)
    for i := 0; i < numNeuron; i++ { layer[i] = genNeuron() }
    return layer
}

// generate and initialize a neuron
func genNeuron() *Neuron { return &Neuron{0.0, 0.0} }

// Sigmoid function: choose between 0 and 1 based on input value
func sigmoid(x float64) float64 { return 1.0 / (1.0 + math.Pow(math.E, -x)) }

// read and analyze a scanned image of hand-written numbers
func getImageData() {
    // input file
    // create data
    // return data
}

// feedforward neural network where information flows only one way
func feedforwardNeuralNet() {
    getImageData()                  // input file and get data
    brain := genNeuralNetwork()     // generate a brain

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
