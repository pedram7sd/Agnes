package main

import ("fmt"
        "time"
        "os"
        "os/exec"
        "math"
        "image"
        "image/jpeg"
        "image/color"
)

// Agnes logo in small case
const AGNES_LOGO string =
`
    ##    #######   ##  ###            ##   ##        ###   ##  ########
    #   ##       ## ##    ##         ##   ####        ##     ###       ###
       ##         ###      ##       #        ##       #      ##         ##
       ##         ###       ##    #           ##     #       ##         ##
       ##         ###        ##  #             ##   #        ##         ##
        ##       ## ##        ###     AGNES     ## #         ##         ##
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
type NNetwork struct {
    numInputs       int             // number of inputs
    numOutputs      int             // number of outputs
    numHNeurons     int             // number of neurons in each hidden layer
    numHLayers      int             // number of hiddne layers
    layers          []*NLayer       // slice of layers
}

func NewNNetwork(nIns, nOuts, nHlNs, nHls int) *NNetwork {
    var ls []*NLayer
    // if there are hidden layers,
    if nHls > 0 {
        ls = make([]*NLayer, nHls)
        // first hidden layer
        ls[0] = NewNLayer(nHlNs, nIns)
        // mid hidden layers
        for i := 1; i < nHls - 1; i++ { ls[i] = NewNLayer(nHlNs, nHlNs) }
        // output layer
        ls[nHls - 1] = NewNLayer(nOuts, nHlNs)
    } else {
        ls = make([]*NLayer, 1)
        ls[0] = NewNLayer(nOuts, nIns)
    }
    return &NNetwork{nIns, nOuts, nHlNs, nHls, ls}
}

func (n *NNetwork) GetWeights() *list.List {
    weights := list.New()
    for i := range n.layers {
    for j := range n.layers[i].neurons {
    for k := range n.layers[i].neurons[j].weights {
        weights.PushBack(n.layers[i].neurons[j].weights[k])
    } } }
    return weights
}

func (n *NNetwork) SetWeights(newWeights []float64) {
    cWeight := 0
    for i := range n.layers {
    for j := range n.layers[i].neurons {
    for k := range n.layers[i].neurons[j].weights {
        n.layers[i].neurons[j].weights[k] = newWeights[cWeight]
        cWeight++
    } } }
}

func (n * NNetwork) Update() {

}

func Sigmoid(x float64) float64 {
    return 1.0 / (1.0 + math.Pow(math.E, -x))
}

// define a Neuron Layer
type NLayer struct {
    numNeurons      int
    neurons         []*Neuron
}

func NewNLayer(nNrns, nIns int) *NLayer {
    ns := make([]*Neuron, nNrns)
    for i := range ns { ns[i] = NewNeuron(nIns) }
    return &NLayer{nNrns, ns}
}

// define a Neuron
type Neuron struct {
    numInputs       int
    weights         []float64
}

// **** change the pararmeter to DNA! ****
func NewNeuron(nIns int) *Neuron {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    w := make([]float64, nIns + 1)
    for i := range w { w[i] = r.Float64() }
    return &Neuron{nIns, w}
}

func (n *Neuron) output(inputs []float64) float64 {
    var sum float64 = 0
    for i := range inputs { sum += inputs[i] * n.weights[i] }
    return sum
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
    intro()                 // print intro
    testImage()
    // feedforwardNeuralNet()
}
