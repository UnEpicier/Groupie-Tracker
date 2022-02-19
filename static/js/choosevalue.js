const inputField = document.querySelector('.location');
const dropdown = document.querySelector('.value-list');

let locs = document.getElementById('data-locations').innerText.replaceAll('[', '').replaceAll(']', '').split(' ')
locs.forEach((el, index) => {
	// Replace Country separator with a space and uppercase the country
	let l = el.replaceAll('_', ' ')

	l = l.split('-')[0] + ', ' + l.split('-')[1].toUpperCase()

	l = l.trim().replace(/\w\S*/g, (w) => (w.replace(/^\w/, (c) => c.toUpperCase())));

	locs[index] = l
})

locs.forEach((el) => {
	let li = document.createElement('li')
	li.innerText = el
	dropdown.insertAdjacentElement('beforeend', li)
})


const dropdownArray = [...document.querySelectorAll('li')];

let valueArray = [];
dropdownArray.forEach(item => {
	valueArray.push(item.textContent);
});

const closeDropdown = () => {
	dropdown.classList.remove('open');
};

inputField.addEventListener('input', () => {
	dropdown.classList.add('open');
	let inputValue = inputField.value.toLowerCase();

	if (inputValue.length > 0) {
		for (let j = 0; j < valueArray.length; j++) {
			if (!(inputValue.substring(0, inputValue.length) === valueArray[j].substring(0, inputValue.length).toLowerCase())) {
				dropdownArray[j].classList.add('closed');
			} else {
				dropdownArray[j].classList.remove('closed');
			}
		}
	} else {
		for (let i = 0; i < dropdownArray.length; i++) {
			dropdownArray[i].classList.remove('closed');
		}
	}
});

dropdownArray.forEach(item => {
	item.addEventListener('click', evt => {
		inputField.value = item.textContent;
		dropdownArray.forEach(dropdown => {
			dropdown.classList.add('closed');
		});
	});
});

inputField.addEventListener('focus', () => {
	inputField.placeholder = 'Type to filter';
	dropdown.classList.add('open');
	dropdownArray.forEach(dropdown => {
		dropdown.classList.remove('closed');
	});
});

inputField.addEventListener('blur', () => {
	inputField.placeholder = 'Select state';
	dropdown.classList.remove('open');
});

document.addEventListener('click', evt => {
	const isDropdown = dropdown.contains(evt.target);
	const isInput = inputField.contains(evt.target);

	if (!isDropdown && !isInput) {
		dropdown.classList.remove('open');
	}
});