/*  This is a practice code for learning image processing in Go.
 *
 */

package main

import ("os"
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

// create and export test
func test() {
    const D = 12
	img := image.NewGray(image.Rect(0, 0, D, D))
	for x := 0; x < D; x++ { img.Set(x, x, color.Gray{byte(2 * x)}) }
    toimg, _ := os.Create("black_box.jpg")
    defer toimg.Close()
    jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}

// main function
func main() {
    test()
}
