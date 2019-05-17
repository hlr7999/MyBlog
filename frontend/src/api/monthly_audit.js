export class MonthlyAudit {
  static list() {
    let data = [
      {
        id: 'abcd',
        name: '审核报告1',
        records: [
          {
            id: 'aaaa',
            name: '氧气',
            data: '111'
          }
        ]
      }
    ]
    return new Promise(resolve => {
      resolve(data.map(audit => new MonthlyAudit(audit)))
    })
  }

  constructor(params) {
    this.id = params.id
    this.name = params.name
    this.records = params.records
  }

  finish() {
  }
}
