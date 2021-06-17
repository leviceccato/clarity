local component = require('library.component')

return function(callback)
    local c = component('click')
    c.callback = callback or function() end
    return c
end