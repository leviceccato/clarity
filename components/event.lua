local component = require('library.component')

return function(event)
    local c = component('event')
    c.type = event[0]
    c.data = event[1] or {}
    return c
end