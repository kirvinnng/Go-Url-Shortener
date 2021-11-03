const generateCard = (res) => {
    const card = document.getElementById('result')

    if (res.isValidURL) {

        card.innerHTML = `<div class="result-card"><p>Here is your shortened URL</p>
        <a target="_blank" id="shortened-url" href=${res.shortenedURL}> ${res.shortenedURL} </a></div>
        <button id="copy">Copy</button>
        `
        const copyButton = document.getElementById('copy')
        copyButton.addEventListener('click', () => {
            const shortenedUrl = document.getElementById('shortened-url')
            navigator.clipboard.writeText(shortenedUrl.innerHTML)
        })

    } else {
        card.innerHTML = `<p> Invalid URL </p>`
    }
}


const submitUrl = async (e) => {
    e.preventDefault()
    const url = document.getElementById('url')
    const bodyContent = { "url": url.value }

    const post = await fetch('/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(bodyContent),
    });
    const content = await post.json()
    generateCard(content)
}

const button = document.getElementById('button')
const input = document.getElementById('url')

button.addEventListener('click', submitUrl, false)


input.addEventListener('keyup', (e) => {
    const key = e.key
    if (key === 'Enter') {
        e.preventDefault();
        button.click();
    }
})

