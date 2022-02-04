let names = [];

const l = document.getElementsByClassName('container')[0].children
for (let i = 0; i < l.length; i++) {
    names.push(l[i].children[1].innerText)
}

const reqSuggests = (event) => {
    let input = event.value != "" || event.value !== 'undefined' ? (event.value).toLowerCase() : ""

    console.log('test')

    for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
        document.getElementsByClassName('suggestList')[0].children[i].remove()
    }
    document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')

    if (input == "") {
        for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
            document.getElementsByClassName('suggestList')[0].children[i].remove()
        }
        document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')
    }

    if (input != "") {
        names.forEach(el => {
            if (el.toLowerCase().includes(input)) {
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