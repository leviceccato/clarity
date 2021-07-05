local entity = require('library.entity')

local event = require('components.event')

return function(data)
    local e = entity()

    e.addComponent(event(data))

    return e
end