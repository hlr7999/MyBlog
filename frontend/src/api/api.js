import axios from 'axios'

var baseUrl = "http://localhost:2333/"

// login
const UserLogin = function(data) {
    return axios.post(baseUrl + "login", {
        username: data.username,
        password: data.pass
    })
}

// register
const UserRegister = function(data) {
    return axios.post(baseUrl + "register", {
        username: data.username,
        password: data.pass,
        email: data.email
    })
}

// aboutMe
const AboutMe = function() {
    return axios.get(baseUrl + "aboutMe")
}


// articles
const GetHomeArticles = function() {
    return axios.get(baseUrl + "articles")
}

// like or collect
const GetLCArticles = function(r) {
    r.toLowerCase()
    if (r == '/like') {

    } else if (r == '/collect') {

    }
    return axios.get(baseUrl + "articles")
}

export {
    UserLogin,
    UserRegister,
    AboutMe,
    GetHomeArticles,
    GetLCArticles
}
