local component = require('library.component')

return function(handler)
    local c = component('click')
    c.data = nil
    c.handler = handler
    return c
end