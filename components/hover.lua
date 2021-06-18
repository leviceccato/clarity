local component = require('library.component')

return function()
    local c = component('hover')
    c.isHovered = false
    return c
end