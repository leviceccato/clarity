local component = require('library.component')

return function(type)
    local c = component('event')
    c.type = type
    return c
end