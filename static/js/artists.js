const reqSuggests = (event) => {
    let input = event.value != "" || event.value !== 'undefined' ? (event.value).toLowerCase() : ""

    // Example slice
    const names = [
        "queen",
        "soja",
        "pink Floyd",
        "scorpions"
    ]

    for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
        document.getElementsByClassName('suggestList')[0].children[i].remove()
    }
    document.getElementsByClassName('suggestList')[0].classList.remove('suggestList_Active')

    if (input == "") {
        for (let i = 0; i < document.getElementsByClassName('suggestList')[0].children.length; i++) {
            document.getElementsByClassName('suggestList')[0].children[i].remove()
        }
        document.getElementsByClassName('suggestList')[0].cclassList.remove('suggestList_Active')
    }

    if (input != "") {
        names.forEach(el => {
            if (el.includes(input)) {
                let n = document.createElement('li');
                let a = document.createElement('a')
                a.href = "#"
                a.appendChild(document.createTextNode(el))
                n.appendChild(a)

                document.getElementsByClassName('suggestList')[0].insertAdjacentElement('afterbegin', n)
            }
        });
    }
}