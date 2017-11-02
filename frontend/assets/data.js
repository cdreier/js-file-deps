var nodeIDs = []

var g = {
  nodes: rawData.map(n => {
    nodeIDs.push(n.id)
    n.color = "#333"
    n.size = 0
    n.x = 0
    n.y = 0
    return n
  }),
  edges: rawData.reduce((all, curr) => {

    if (!curr.imports) {
      return all
    }

    var imports = curr.imports
      .filter(i => nodeIDs.indexOf(i.id) > -1)
      .map(i => {
        return {
          id: curr.id + i.id,
          source: curr.id,
          target: i.id,
        }
      })

    return [
      ...all,
      ...imports,
    ]

  }, [])
}

console.log(g)