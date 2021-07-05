local component = require('library.component')

return function(data)
    local c = component('event')
    c.type = data[0]
    c.data = data[1] or {}
    return c
end