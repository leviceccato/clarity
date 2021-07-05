local component = require('library.component')

return function(data)
    local c = component('event')
    c.name = data[1]
    c.data = data[2] or {}
    return c
end