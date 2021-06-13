local component = require('library.component')

return function(min, max)
    local c = component('collision')
    c.min = min || {x = 0, y = 0}
    c.max = max || {x = 0, y = 0}
    return c
end