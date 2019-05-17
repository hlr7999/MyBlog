import { ESGClient } from "./request"

export class ESGModel {
  constructor() {
    this._client = new ESGClient(this._identity.jwtAuthToken)
  }

  save() {
    // If no id provided, we can see it as a create request.
    // Otherwise, it's an update request.
    if (this.id) {
      return this._client
        .put(this.updateUrl)
        .auth()
        .body(this.objectize)
        .send()
    } else {
      return this._client
        .post(this.createUrl)
        .auth()
        .body(this.objectize)
        .send()
    }
  }

  destroy() {
    return this._client
      .post(this.destroyUrl)
      .auth()
      .send()
  }
}
