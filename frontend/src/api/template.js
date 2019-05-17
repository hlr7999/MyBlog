import { ESGClient } from "./request"
import { User } from "./user"

export class Template {
  // Properties:
  //
  // id:   String
  // name: String
  // root_indicators: [Indicator]
  // standard: Standard
  // industry: Industry

  static fromObject(obj) {
    let tmp = new Template()
    tmp._id = obj._id
    tmp.name = obj.name
  }

  static list() {
    let client = new ESGClient(User.currentUser.authToken)
    return client
      .get("/api/templates")
      .auth()
      .send()
      .then(result => {
        // 如果请求成功则返回Template对象数组，否则返回错误状态码
        if(typeof result === "object") {
          return result.map(r => new Template(r))
        }
        else {
          return result
        }
      })
  }

  constructor(params) {
    this._id = params._id
    this.name = params.name
    this.standard = params.standard
    this.industry = params.industry
    this.root_indicators = params.root_indicators
  }

  static createTemplate(params) {
    return new ESGClient(User.currentUser.authToken)
      .post("/api/templates")
      .auth()
      .body(params)
      .send(result => {
        // 如果请求成功则返回Template对象，否则返回错误状态码
        if(typeof result === "object") {
          return new Template(result)
        }
        else {
          return result
        }
      })
  }

  static deleteTemplate(id) {
    return new ESGClient(User.currentUser.authToken)
      .delete(`/api/templates/${id}`)
      .auth()
      .send()
      .then(resStatus => {
        return resStatus
      })
  }


  static getTemplate(id) {
    return new ESGClient(User.currentUser.authToken)
      .get(`/api/templates/${id}`)
      .auth()
      .send()
      .then(result => {
        // 如果请求成功则返回Template对象，否则返回错误状态码
        if(typeof result === "object") {
          return new Template(result)
        }
        else {
          return result
        }
      })
  }

}
