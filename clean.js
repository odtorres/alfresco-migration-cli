const fs = require("fs")

if (process.argv[2]) {
    let data = require("./"+process.argv[2])

    data.forEach(element => {
        element.tasks = element.tasks.map
    });
    console.log(data)
}
