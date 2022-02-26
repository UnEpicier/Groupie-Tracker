function toggleActive(event) {
    const target = event.target
    switch (target.classList.contains('thActive')) {
        case true:
            target.classList.remove('thActive')
            setCookie('theme', 'light')
            document.body.setAttribute('data-theme', 'light')
            break
        case false:
            target.classList.add('thActive')
            setCookie('theme', 'dark')
            document.body.setAttribute('data-theme', 'dark')
            break
    }
}

function getThemeNode() {
    let th = document.createElement('div')
    th.classList.add('thContainer')

    let toggler = document.createElement('div')
    toggler.classList.add('thToggler')
    th.appendChild(toggler)
    toggler.onclick = (e) => toggleActive(e)

    let moon = document.createElement('div')
    moon.classList.add('thIcon')
    let sun = document.createElement('div')
    sun.classList.add('thIcon')

    th.insertAdjacentElement('afterbegin', moon)
    th.insertAdjacentElement('beforeend', sun)

    return th
}

function initThToggler(active) {
    let css = document.createElement('link')
    css.setAttribute('rel', 'stylesheet')
    css.setAttribute('href', 'css/thToggler.css')

    document.getElementsByTagName('head')[0].appendChild(css)

    let toggler = getThemeNode()
    if (active) {
        toggler.children[1].classList.add('thActive')
    }

    document.getElementById('toggler').parentElement.appendChild(toggler)
    document.getElementById('toggler').remove()
}

function initTheme() {
    let def_active = false
    if (checkCookie('theme')) {
        document.body.setAttribute('data-theme', getCookie('theme'))
        if (getCookie('theme') == "dark") {
            def_active = true
        }
    } else {
        setCookie('theme', 'dark')
        document.body.setAttribute('data-theme', 'dark')
        def_active = true
    }

    initThToggler(def_active)
}

initTheme()