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

	// If we have more than 6 urls (strings), we only take the 6th ones.
	// Else, we take the entire slice
	if len(images) > 6 {
		for k, v := range images {
			if k > 6 {
				break
			}

			result = append(result, v)
		}
	} else if len(images) <= 6 {
		result = append(result, images...)
	}

	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeArtist(slice []Artists, s int) []Artists {
	return append(slice[:s], slice[s+1:]...)
}

func getLocations(as ArtistsStruct) []string {
	res := []string{}

	for _, v := range as.Tab {
		GetRelations(v.Id)
		for k := range Rels.DatesLocations {
			res = append(res, k)
		}
	}

	res = removeDuplicateStr(res)

	return res
}

func removeDuplicateStr(strSlice []string) []string {

	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
