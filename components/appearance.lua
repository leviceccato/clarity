local component = require('library.component')

return function(image, width, height)
    local c = component('appearance')
    if image then
        width = image:getWidth()
        height = image:getHeight()
        c.image = image
        c.frame = love.graphics.newQuad(
            0, 0,
            width, height,
            width, height
        )
    end
    c.width = width
    c.height = height
    return c
end