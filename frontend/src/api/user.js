import { Client } from "./request"

export class User {
  static login(username, password) {
    return new Client()
      .post("/api/auth/session")
      .body({
        username: username,
        password: password
      })
      .send()
      .then(user => {
        localStorage.setItem(
          "currentUser",
          JSON.stringify({
            username: username,
            token: user.token,
          })
        )
        return new Promise(resolve => resolve())
      })
  }

  static register(username, password) {
    
  }

  static logout() {
    localStorage.removeItem('currentUser')
  }

  constructor(params) {
    this._token = typeof params === "object" ? params.token : null
  }

  get signedIn() {
    return Boolean(this._token)
  }

  get authToken() {
    if (!this._token) return null
    else return ["Bearer", this._token].join(" ")
  }
}

// TODO: Always reading local storage is costy.
// Try singleton with listener of local storage's writes.
Object.defineProperty(User, "currentUser", {
  get: function() {
    return new User(JSON.parse(localStorage.getItem("currentUser") || "{}"))
  }
})
