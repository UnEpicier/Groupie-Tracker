package groupie

import (
	"math/rand"
	"time"
)

func getRandomAlbum(as ArtistsStruct) []string {
	result := []string{}
	images := []string{}

	// Add All Images URLs in a slice
	for _, v := range as.Tab {
		images = append(images, v.Image)
	}

	// Shuffle "images" slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(images), func(i, j int) { images[i], images[j] = images[j], images[i] })

	// If we have more than 7 urls (strings), we only take the 7th ones.
	// Else, we take the entire slice
	if len(images) > 7 {
		for k, v := range images {
			if k > 7 {
				break
			}

			result = append(result, v)
		}
	} else if len(images) <= 7 {
		result = append(result, images...)
	}

	return result
}
