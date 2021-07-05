local component = require('library.component')

return function(event)
    local c = component('click')
    c.data = nil
    c.event = event
    return c
end