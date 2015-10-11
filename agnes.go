package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
        "image"
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
    P_FILE = ""     // input file path
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

// read the image data through input layer
func (this *NeuralNetwork) readImageInput(image []float64) {
    // each pixel in a greyscale image is an input for each neuron
    for pixel, neuron := range this.inputLayer { neuron.i = image[pixel] }
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

// convert the image to grayscale

// read and analyze a scanned image of hand-written numbers
func getImageData() []int {
    // input image file
    file, err := os.Open(P_FILE)
    if err != nil { panic(err.String()) }
    defer file.Close()

    // create image from the image file
    img, _, err := image.Decode(file)
    if err != nil { panic(err.String()) }

    // read the input image and organize it into a []int
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
    grayImg := image.NewGray(width, height)
    var data []int = make([]int, I_LAYER)
    index := 0
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            oldPixel := img.At(x, y)
            grayPixel := image.ColorModel().Convert(oldPixel)
            data[index] = grayPixel[0]
            index++
        }
    }
    // return the new color data
    return data
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
    var image []int = getImageData()    // input file and get data
    brain := genNeuralNetwork()         // generate a brain
    // read the image data with input layer of the brain
    brain.readImageInput(image)
}

// Agnes main function
func main() {
    intro()                 // print intro
    feedforwardNeuralNet()  // recognize hand writing using feedforward
}
