local component = require('library.component')

return function(mouseEnter, mouseLeave)
    local c = component('hover')
    c.isHovered = false
    c.mouseEnter = mouseEnter or function() end
    c.mouseLeave = mouseLeave or function() end
    return c
end