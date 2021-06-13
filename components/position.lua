local component = require('library.component')

return function(x, y)
    local c = component('position')
    c.x = x or 0
    c.y = y or 0
    return c
end