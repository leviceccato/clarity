local component = require('library.component')

return function(content, align, padding, colour)
    local c = component('text')
    c.content = {colour or {1, 1, 1, 1}, content or ''}
    c.align = align or 'left'
    c.padding = padding or 0
    return c
end