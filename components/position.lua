local component = require('library.component')

return function(x, y)
    return component('position', {
        x = x or 0,
        y = y or 0
    })
end