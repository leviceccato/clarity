local component = require('library.component')

return function(image)
    local width = image:getWidth()
    local height = image:getHeight()
    local c = component('appearance')
    c.image = image
    c.frame = love.graphics.newQuad(
        0, 0,
        width, height,
        width, height
    )
    return c
end