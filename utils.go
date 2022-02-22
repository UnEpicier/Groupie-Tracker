package groupie

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func GetLocalIP() string {
	conn, err := net.Dial("ip:icmp", "google.com")
	if err != nil {
		fmt.Println(err)
		return "localhost"
	}

	if strings.Count(conn.LocalAddr().String(), ":") < 2 {
		return conn.LocalAddr().String()
	}

	return "localhost"
}

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
	return slice[:s+copy(slice[s:], slice[s+1:])]
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

func getLocation(id int) []string {
	res := []string{}

	GetRelations(id)
	for k := range Rels.DatesLocations {
		res = append(res, k)
	}

	res = removeDuplicateStr(res)

	return res
}

func formatFilter(str string) string {
	str = strings.ToLower(str)
	str = strings.Replace(str, ", ", "-", 1)
	str = strings.ReplaceAll(str, " ", "_")

	return str
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

func processFilters(slice []Artists, index int, filters map[string][]string) []Artists {
	minDateC, _ := strconv.Atoi(fmt.Sprint(filters["Creation"][0]))
	maxDateC, _ := strconv.Atoi(fmt.Sprint(filters["Creation"][1]))

	minDateA, _ := strconv.Atoi(fmt.Sprint(filters["Album"][0]))
	maxDateA, _ := strconv.Atoi(fmt.Sprint(filters["Album"][1]))

	dA, _ := strconv.Atoi(strings.Split(slice[index].FirstAlbum, "-")[2])

	if !((slice[index].CreationDate >= minDateC && slice[index].CreationDate <= maxDateC) && (dA >= minDateA && dA <= maxDateA) && (contains(getLocation(slice[index].Id), formatFilter(filters["Location"][0])))) {
		slice = removeArtist(slice, index)
		index--
	}

	if index >= len(slice)-1 {
		return slice
	}

	return processFilters(slice, index+1, filters)
}
