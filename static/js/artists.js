if (document.getElementsByClassName('container')[0].children.length == 0) {
    if (document.getElementsByClassName('container')[0].children.length == 0) {
        let nr = document.createElement('p')
        nr.classList.add('noresult')
        nr.innerText = 'No result found!'
        document.getElementsByClassName('container')[0].appendChild(nr)
    }
}

function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}

const fullgrid = [...document.getElementsByClassName('fullcontainer')[0].children]
const grid = [...document.getElementsByClassName('container')[0].children]


const reqSuggests = (event) => {
    let input = event.value != "" || event.value !== 'undefined' ? (event.value).toLowerCase() : ""

    removeAllChildNodes(document.getElementsByClassName('container')[0])

    // If input empty
    if (input.length < 1) {
        grid.forEach(el => {
            document.getElementsByClassName('container')[0].insertAdjacentElement('beforeend', el)
        })
    }

    if (input.length > 0) {
        fullgrid.forEach(el => {
            if (input.substring(0, input.length) === (el.children[1].innerText).substring(0, input.length).toLowerCase()) {
                document.getElementsByClassName('container')[0].insertAdjacentElement('beforeend', el)
            }
        });
    }

    // If any result, print "No result found"
    if (document.getElementsByClassName('container')[0].children.length == 0) {
        let nr = document.createElement('p')
        nr.classList.add('noresult')
        nr.innerText = 'No result found!'
        document.getElementsByClassName('container')[0].appendChild(nr)
    }
}

/* Filters bars */
// Set fix width to the div that contains the
// toggle button to keep loading of the slider(s) correct
let btnOffW = document.getElementsByClassName('menu')[0].offsetWidth
document.getElementById('filtersBtnC').style.width = (btnOffW + 50) + 'px'

// filtersContainer
const toggleFilters = () => {
    const el = document.getElementsByClassName('filtersContainer')[0];
    if (el.classList.contains('filtersActive')) {
        el.classList.remove('filtersActive')
    } else {
        el.classList.add('filtersActive')
    }
}

/* SLIDER */

// Creation Slider
let datesCreated = document.getElementById('data-dates').innerText.replaceAll('[', '').replaceAll(']', '').split(' ')
for (let index = 0; index < datesCreated.length; index++) {
    datesCreated[index] = parseInt(datesCreated[index])
}

let sliderC = document.getElementById('datesCSlider')

sliderC.setAttribute('se-min', datesCreated[0])
sliderC.setAttribute('se-min-value', datesCreated[0])

sliderC.setAttribute('se-max', datesCreated[datesCreated.length - 1])
sliderC.setAttribute('se-max-value', datesCreated[datesCreated.length - 1])

document.getElementById('resultC').children[0].innerText = datesCreated[0]
document.getElementById('resultC').children[2].innerText = datesCreated[datesCreated.length - 1]
document.getElementsByName('minDateC')[0].value = datesCreated[0]
document.getElementsByName('maxDateC')[0].value = datesCreated[datesCreated.length - 1]


let creationDate = new ZBRangeSlider('datesCSlider');

creationDate.onChange = function (min, max) {
    document.getElementById('resultC').children[0].innerText = min
    document.getElementById('resultC').children[2].innerText = max
    document.getElementsByName('minDateC')[0].value = min
    document.getElementsByName('maxDateC')[0].value = max
}

creationDate.didChanged = function (min, max) {
    document.getElementById('resultC').children[0].innerText = min
    document.getElementById('resultC').children[2].innerText = max
    document.getElementsByName('minDateC')[0].value = min
    document.getElementsByName('maxDateC')[0].value = max
}

// Album Slider

let datesAlbum = document.getElementById('data-alb').innerText.replaceAll('[', '').replaceAll(']', '').split(' ')
for (let index = 0; index < datesAlbum.length; index++) {
    datesAlbum[index] = parseInt(datesAlbum[index])
}

let sliderA = document.getElementById('datesASlider')

sliderA.setAttribute('se-min', datesAlbum[0])
sliderA.setAttribute('se-min-value', datesAlbum[0])

sliderA.setAttribute('se-max', datesAlbum[datesAlbum.length - 1])
sliderA.setAttribute('se-max-value', datesAlbum[datesAlbum.length - 1])

document.getElementById('resultA').children[0].innerText = datesAlbum[0]
document.getElementById('resultA').children[2].innerText = datesAlbum[datesAlbum.length - 1]
document.getElementsByName('minDateA')[0].value = datesAlbum[0]
document.getElementsByName('maxDateA')[0].value = datesAlbum[datesAlbum.length - 1]


let albumDate = new ZBRangeSlider('datesASlider');

albumDate.onChange = function (min, max) {
    document.getElementById('resultA').children[0].innerText = min
    document.getElementById('resultA').children[2].innerText = max
    document.getElementsByName('minDateA')[0].value = min
    document.getElementsByName('maxDateA')[0].value = max
}

albumDate.didChanged = function (min, max) {
    document.getElementById('resultA').children[0].innerText = min
    document.getElementById('resultA').children[2].innerText = max
    document.getElementsByName('minDateA')[0].value = min
    document.getElementsByName('maxDateA')[0].value = max
}