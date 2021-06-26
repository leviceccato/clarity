local component = require('library.component')

return function(outputs)
    local c = component('controls')
    c.outputs = outputs
    return c
end