let names = [];

const l = document.getElementsByClassName('container')[0].children
for (let i = 0; i < l.length; i++) {
    names.push(l[i].children[1].innerText)
}

const reqSuggests = (event) => {
    let input = event.value != "" || event.value !== 'undefined' ? (event.value).toLowerCase() : ""

    for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
        document.getElementsByClassName('suggestList')[0].children[i].remove()
    }
    document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')

    if (!(input.length > 0)) {
        for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
            document.getElementsByClassName('suggestList')[0].children[i].remove()
        }
        document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')
    }

    if (input.length > 0) {
        for (let el of document.getElementsByClassName('suggestList')[0].children) {
            el.remove()
        }
        names.forEach(el => {
            if (input.substring(0, input.length) === el.substring(0, input.length).toLowerCase()) {
                let n = document.createElement('li');
                let a = document.createElement('a')
                a.setAttribute('href', "artist?a=" + el.replace(" ", "%20"))
                a.appendChild(document.createTextNode(el))
                n.appendChild(a)

                document.getElementsByClassName('suggestList')[0].insertAdjacentElement('afterbegin', n)
            }
        });
        if (document.getElementsByClassName('suggestList')[0].children.length > 0) {
            document.getElementsByClassName('suggestList')[0].classList.add('suggestList_Active')
        } else {
            document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')
        }
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