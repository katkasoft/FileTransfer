document.querySelector('form').addEventListener('submit', (e) => {
    e.preventDefault()
    const emailview = document.getElementById('email')
    const passwordview = document.getElementById('password')
    const confirmview = document.getElementById('confirm')
    const error = document.getElementById('error')
    const email = emailview.value
    const password = passwordview.value
    const confirm = confirmview.value
    error.innerHTML = ''
    if (password.length < 8) {
        error.innerHTML = 'Password too short. Minimal 8 characters'
        return
    }
    if (password != confirm) {
        error.innerHTML = 'Passwords don\'t match'
        return
    }
})

function show(el) {
    if (el.checked == true) {
        document.getElementById('password').type = 'text'
        document.getElementById('confirm').type = 'text'
    }
    else if (el.checked == false) {
        document.getElementById('password').type = 'password'
        document.getElementById('confirm').type = 'password'
    }
}