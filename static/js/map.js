/**
 * @name initMap
 * @description Initiate a map
 * @param {string.<[lat: float, lon: float]>} coords - Latitude and longitude
 * @param {Array.<{city: Array.<{lat: float, lon: float}>}>} markers - An array of markers to place on map
 */
const initMap = (coords, markers) => {
    /* Create a map in the element that have 'map' ID
    And set the view on a place with the coords
    */
    let marks = []
    let map = L.map('map').setView([coords[0] || 0, coords[1] || 0], 11)
    let markerClusters = L.markerClusterGroup()

    L.tileLayer('https://{s}.tile.openstreetmap.fr/osmfr/{z}/{x}/{y}.png', {
        minZoom: 1,
        maxZoom: 20
    }).addTo(map)

    if (Object.keys(markers).length > 0) {
        Object.entries(markers).forEach(([key, value]) => {
            // Create a marker and add it to the map
            let m = L.marker([value.lat, value.lon]).addTo(map)
            // Add bubble text
            m.bindPopup(key)
        })
    }
}

/************/
/*   INIT   */
/************/

let coords = []
let markers = {}

/*
Add every coords
*/
let co = document.getElementById('coords').innerText.replaceAll('        ', '').replaceAll('\n', '').replaceAll('[', ',[').substring(1).split(',')
co.forEach(e => {
    coords.push(JSON.parse(e.replace(' ', ', ')))
});

/*
Add every markers
city: {lat, long}
*/
let table = document.getElementsByClassName('tablecols')[1].children
let cities = []

Array.from(table).forEach(element => {
    cities.push(element.innerText)
});
cities[0] = null
cities = cities.filter(Boolean)

for (const [key, value] of cities.entries()) {
    markers[value] = { "lat": coords[key][0], "lon": coords[key][1] }
};

console.log(coords, markers)

initMap(coords[0], markers)