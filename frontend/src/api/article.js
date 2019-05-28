import { ESGClient } from "./request"
import { User } from "./user"

export class Company {
  static list() {
    let client = new ESGClient(User.currentUser.authToken)
    return client
      .get("/api/companies")
      .auth()
      .send()
      .then(result => {
        // 如果请求成功则返回Company对象数组，否则返回错误状态码
        if(typeof result === "object") {
          return result.map(r => new Company(r))
        }
        else {
          return result
        }
      })
  }

  constructor(params) {
    this._id = params._id
    this.status = params.status
    this.name = params.name
    this.level = params.level
    this.parent = params.parent
    this.region = params.region
    this.industry = params.industry
    this.subscription = params.subscription
  }

  static createCompany(params) {
    return new ESGClient(User.currentUser.authToken)
      .post("/api/companies")
      .auth()
      .body(params)
      .send()
      .then(result => {
        // 如果请求成功则返回Company对象，否则返回错误状态码
        if(typeof result === "object") {
          return new Company(result)
        }
        else {
          return result
        }
      })
  }

  static deleteCompany(id) {
    return new ESGClient(User.currentUser.authToken)
      .delete(`/api/companies/${id}`)
      .auth()
      .send()
      .then(resStatus => {
          return resStatus
      })
  }

  static getCompany(id) {
    return new ESGClient(User.currentUser.authToken)
      .get(`/api/companies/${id}`)
      .auth()
      .send()
      .then(result => {
        // 如果请求成功则返回Company对象，否则返回错误状态码
        if(typeof result === "object") {
          return new Company(result)
        }
        else {
          return result
        }
      })
  }

  static reports(id) {
    return new ESGClient(User.currentUser.authToken)
      .get(`/api/reports`)
      .auth()
      .send()
      .then(result => {
        // 如果请求成功则返回Reports数组，否则返回错误状态码
        if(typeof result === "object") {
          let r = []
          for (let item of result) {
            if (item.company == id) {
              r.push(item)
            }
          }
          return r
        }
        else {
          return result
        }
      })
  }
}
