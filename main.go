package main

import (
    "fmt"
    "github.com/google/go-containerregistry/pkg/crane"
)

func main() {
    // Get the image reference.
    imageRef := "gcr.io/gae-runtimes/buildpacks/google-gae-22/java/builder"

    // Pull the image.
    image, err := crane.Pull(imageRef)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get the layers of the image.
    layers, err := image.Layers()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print information about the layers.
    for _, layer := range layers {
        fmt.Println("Layer:")
        fmt.Println("  Digest:", layer.Digest)
        fmt.Println("  Size:", layer.Size)
    }
}
