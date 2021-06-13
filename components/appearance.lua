local component = require('library.component')

return function(image)
    local c = component('appearance')
    c.image = image
    return c
end