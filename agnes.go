package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
        "image"
        "image/jpeg"
)

// Agnes logo in small case
const AGNES_LOGO string =
`
    ##    #######   ##  ###            ##   ##        ###   ##  ########
    #   ##       ## ##    ##         ##   ####        ##     ###       ###
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
    P_FILE = ""     // input file path
    S_IMG = 500     // size of an image
    I_LAYER = 784   // size of input layer (784 = 28 * 28)
    H_LAYER = 15    // size of hidden layer that can be re-adjusted
    O_LAYER = 10    // size of output layer (from 0 to 9)
)

// define a Neural Network
type NeuralNetwork struct {
    inputLayer []*Neuron    // input layer
    hiddenLayer []*Neuron   // hidden layer
    outputLayer []*Neuron   // output layer
}

// define a Neuron
type Neuron struct {
    i float64   // sum of previous layer's outputs
    o float64   // result of sigmoid function
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

// create an image
func test() {

	img := image.NewRGBA(image.Rect(1, 1, S_IMG, S_IMG))

    newImage, _ := os.Create("new.jpg")
    defer newImage.Close()

    jpeg.Encode(newImage, img, &jpeg.Options{jpeg.DefaultQuality})
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

// feedforward neural network where information flows only one way
func feedforwardNeuralNet() {
    // brain := genNeuralNetwork()             // generate a brain
}

// Agnes main function
func main() {
    intro()                 // print intro
    test()
    // feedforwardNeuralNet()
}
