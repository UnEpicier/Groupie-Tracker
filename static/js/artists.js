function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}

const grid = [...document.getElementsByClassName('container')[0].children];

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
        grid.forEach(el => {
            if (input.substring(0, input.length) === (el.children[1].innerText).substring(0, input.length).toLowerCase()) {
                document.getElementsByClassName('container')[0].insertAdjacentElement('beforeend', el)
            }
        });
    }
}

/* Filters bars */
// set fix width to the div that contains the
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

let datesCreated = document.getElementById('data-dates').innerText.replaceAll('[', '').replaceAll(']', '').split(' ')
for (let index = 0; index < datesCreated.length; index++) {
    datesCreated[index] = parseInt(datesCreated[index])
}

let slider = document.getElementById('datesSlider')

slider.setAttribute('se-min', datesCreated[0])
slider.setAttribute('se-min-value', datesCreated[0])

slider.setAttribute('se-max', datesCreated[datesCreated.length - 1])
slider.setAttribute('se-max-value', datesCreated[datesCreated.length - 1])

document.getElementById('result').children[0].innerText = datesCreated[0]
document.getElementById('result').children[1].innerText = datesCreated[datesCreated.length - 1]


let creationDate = new ZBRangeSlider('datesSlider');

creationDate.onChange = function (min, max) {
    document.getElementById('result').children[0].innerText = min
    document.getElementById('result').children[1].innerText = max
}

creationDate.didChanged = function (min, max) {
    document.getElementById('result').children[0].innerText = min
    document.getElementById('result').children[1].innerText = max
}