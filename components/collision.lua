local component = require('library.component')

return function(min, max)
    local c = component('collision')
    c.min = min or {x = 0, y = 0}
    c.max = max or {x = 0, y = 0}
    return c
end