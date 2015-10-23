/*  This is a practice code for learning image processing in Go.
 *
 */

package main

import ("fmt"
        "os"
        "image"
        "image/color"
        "image/jpeg"
)

const INPUT_PATH string = ""    // input file path
const EXPORT_PATH = "./"        // output file path

// convert an image to grayscale only by manipulating pixels individually
func grayscale() {

}

// read a grayscale image and print each pixel's brightness value
func toString() {

}

// // read and analyze a scanned image of hand-written numbers
// func getImageData() []float64 {
//     // input image file
//     file, err := os.Open(P_FILE)
//     if err != nil { panic(err) }
//     defer file.Close()
//
//     // create image from the image file
//     img, _, err := image.Decode(file)
//     if err != nil { panic(err) }
//
//     // read the input image and organize it into a []float64
//     bounds := img.Bounds()
//     w, h := bounds.Max.X, bounds.Max.Y
//     rgbaImg := image.NewRGBA(image.Rect(0, 0, w, h))
//     // re-draw the original image in an RGBA image
//     for x := 0; x < w; x++ { for y := 0; y < h; y++ {
//         rgbaImg.Set(x, y, img.At(x, y))
//     } }
//
//     var data []float64 = make([]float64, I_LAYER)
//
//     // return the new color data
//     return data
// }

// create and export test
func test() {
    const D = 100
	img := image.NewGray(image.Rect(1, 1, D, D))
	for x := 1; x <= D; x++ {
		img.Set(x, x, color.Gray{byte(2 * x)})
	}
	for x := 1; x < D; x++ {
		fmt.Printf("[%2d, %2d]: %v\n", x, x, img.At(x, x))
	}
    toimg, _ := os.Create("new.jpg")
    defer toimg.Close()

    jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}

// main function
func main() {
    test()
}
