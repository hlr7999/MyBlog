export class ESGClient {
  constructor(authToken = null) {
    this._authToken = authToken
  }

  get(url) {
    return new ESGRequest("GET", url, this._authToken)
  }

  post(url) {
    return new ESGRequest("POST", url, this._authToken)
  }

  put(url) {
    return new ESGRequest("PUT", url, this._authToken)
  }

  delete(url) {
    return new ESGRequest("DELETE", url, this._authToken)
  }
}

// ESGRequest is the real class packing all about a request.
export class ESGRequest {
  constructor(method, url, authToken = null) {
    this._method = method
    this._url = url
    this._authToken = authToken
    this._headers = {
      "Content-Type": "application/json"
    }
    this._body = {}
    this._params = {}
  }

  header(arg, value = null) {
    if (typeof arg === "object") {
      Object.assign(this._headers, arg)
    } else if (typeof arg === "string") {
      this._headers[arg] = value
    }
    return this
  }

  // Confirm to add authorization header.
  auth() {
    this._headers["Authorization"] = this._authToken
    return this
  }

  body(arg, value = null) {
    if (typeof arg === "object") {
      Object.assign(this._body, arg)
    } else if (typeof arg === "string") {
      this._body[arg] = value
    }
    return this
  }

  // Set param, using either object in 1 argument or key-value in 2 arguments.
  param(arg, value = null) {
    if (typeof arg === "object") {
      Object.assign(this._params, arg)
    } else if (typeof arg === "string") {
      this._params[arg] = value
    }
    return this
  }

  send() {
    let query = Object.keys(this._params).reduce((acc, cur) => {
      acc +=
        "&" +
        encodeURIComponent(cur) +
        "=" +
        encodeURIComponent(this._params[cur])
    }, "")
    let url = this._url
    if (query.length > 0) {
      url += "?" + query
    }
    let fetchParams = {
      method: this._method,
      headers: new Headers(this._headers)
    }
    if (this._method !== "GET" && this._method !== "HEAD") {
      fetchParams.body = JSON.stringify(this._body)
    }
    return fetch(url, fetchParams).then(res => {
      if(res.status == 200) {
        return res.json()
      }
      else {
        return res.status
      }
    })
  }
}
