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

let dates = document.getElementById('data-dates').innerText.replaceAll('[', '').replaceAll(']', '').split(' ')
for (let index = 0; index < dates.length; index++) {
    dates[index] = parseInt(dates[index])
}

let slider = document.getElementById('datesSlider')
slider.setAttribute('se-min', dates[0])
slider.setAttribute('set-min-value', dates[0])
slider.setAttribute('se-max', dates[dates.length - 1])
slider.setAttribute('se-max-value', dates[dates.length - 1])

document.getElementById('min-date').value = dates[0]
document.getElementById('min-date').setAttribute('min', dates[0])
document.getElementById('min-date').setAttribute('max', dates[dates.length - 1] - 1)
document.getElementById('max-date').value = dates[dates.length - 1]
document.getElementById('max-date').setAttribute('min', dates[0] + 1)
document.getElementById('max-date').setAttribute('max', dates[dates.length - 1])


let creationDate = new ZBRangeSlider('datesSlider');

creationDate.onChange = function (min, max) {
    document.getElementById('min-date').value = min
    document.getElementById('max-date').value = max
    document.getElementById('min-date').value = min
    document.getElementById('max-date').value = max
}

creationDate.didChanged = function (min, max) {
    document.getElementById('min-date').value = min
    document.getElementById('max-date').value = max
    document.getElementById('min-date').value = min
    document.getElementById('max-date').value = max
}