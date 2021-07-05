local component = require('library.component')

return function(data)
    local c = component('event')
    c.name = data[0]
    c.data = data[1] or {}
    return c
end