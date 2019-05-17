import { ESGClient } from "./request"

export class User {
  static login(handle, password) {
    return new ESGClient()
      .post("/api/auth/session")
      .body({
        username: handle,
        password: password
      })
      .send()
      .then(user => {
        localStorage.setItem(
          "currentUser",
          JSON.stringify({
            username: handle,
            token: user.token,
            roles: user.roles
          })
        )
        return new Promise(resolve => resolve())
      })
  }

  static logout() {
    localStorage.removeItem('currentUser')
  }

  constructor(params) {
    this._token = typeof params === "object" ? params.token : null
    this._roleMask = 0
    this._roles = params.roles

    this._roleMapping = {
      SYS: {
        mask: 0b1000,
        name: "超级管理员"
      },
      COM: {
        mask: 0b0100,
        name: "公司管理员"
      },
      AUD: {
        mask: 0b0010,
        name: "审核员"
      },
      REC: {
        mask: 0b0001,
        name: "数据录入员"
      }
    }

    params.roles.forEach(role => {
      this._roleMask |= this._roleMapping[role].mask
    })
  }

  get roleName() {
    return this._roleMapping[this._roles[0]].name
  }

  get isSuperAdmin() {
    return Boolean(this._roleMask & this._roleMapping.SYS.mask)
  }

  get isCompanyAdmin() {
    return Boolean(this._roleMask & this._roleMapping.COM.mask)
  }

  get isAuditor() {
    return Boolean(this._roleMask & this._roleMapping.AUD.mask)
  }

  get isRecorder() {
    return Boolean(this._roleMask & this._roleMapping.REC.mask)
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
