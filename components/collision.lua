local component = require('library.component')

return function(min, max)
    return component('collision', {
        min = min || { x = 0, y = 0 },
        max = max || { x = 0, y = 0 }
    })
end