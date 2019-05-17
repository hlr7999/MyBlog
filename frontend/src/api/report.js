export class Report {
  static list() {
    return new Promise((resolve, reject) => {
      let data = [
        {
          name: "环境",
          level: 1,
          children: [
            {
              name: "排放",
              level: 2,
              children: [
                {
                  name: "气体",
                  level: 3,
                  children: [
                    {
                      name: "一个四级指标",
                      level: 4,
                      children: [
                        {
                          name: "二氧化碳",
                          level: 5,
                          type: 1,
                          data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                          unit: "m^3"
                        },
                        {
                          name: "二氧化硫",
                          level: 5,
                          type: 1,
                          data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                          unit: "m^3"
                        },
                        {
                          name: "定性指标1",
                          level: 5,
                          type: 2,
                          description: "this is description",
                          data: "this is data of a qualitative Indicator"
                        }
                      ]
                    },
                    {
                      name: "Test2",
                      level: 4,
                      description: "Thisis a t",
                      children: [
                        {
                          name: "二氧化A",
                          level: 5,
                          type: 1,
                          data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                          unit: "m^3"
                        },
                        {
                          name: "二氧化B",
                          level: 5,
                          type: 1,
                          data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                          unit: "m^3"
                        },
                        {
                          name: "定性指标2",
                          level: 5,
                          type: 2,
                          description: "this is description",
                          data: "2this is data of a qualitative Indicator"
                        }
                      ]
                    }
                  ]
                },
                {
                  name: "液体",
                  level: 3,
                  children: []
                }
              ]
            },
            {
              name: "能源",
              level: 2,
              children: [
                { name: "组织内部的能源消耗量", level: 3, children: [] }
              ]
            }
          ]
        },
        {
          name: "社会",
          level: 1,
          children: [
            {
              name: "法律法规",
              level: 2,
              children: [{ name: "遵守情况", level: 3, children: [] }]
            }
          ]
        }
      ]
      const addNumber = (parentNumber, node, index) => {
        if (parentNumber != null) {
          node.number = `${parentNumber}.${index}`
        } else {
          node.number = index
        }
        node.numberedName = `${node.number} ${node.name}`
        if (node.children) {
          node.children.forEach((n, i) => {
            addNumber(node.number, n, i + 1)
          })
        }
      }
      data.forEach((node, index) => {
        addNumber(null, node, index + 1)
      })
      resolve(data)
    })
  }

  static cutTree(node, maxLevel) {
    if (node.level >= maxLevel) {
      node.mChildren = node.children
      delete node.children
    } else if (Array.isArray(node)) {
      node.forEach(n => Report.cutTree(n, maxLevel))
    } else {
      if (node.children) {
        node.children.forEach(n => Report.cutTree(n, maxLevel))
      }
    }
    return node
  }
}
