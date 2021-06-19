local component = require('library.component')

return function(text)
    local c = component('text')
    c.text = text or ''
    return c
end