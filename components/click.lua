local component = require('library.component')

return function()
    local c = component('click')
    c.isClicked = false
    return c
end