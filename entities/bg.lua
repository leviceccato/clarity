local entity = require('library.entity')
local appearance = require('components.appearance')
local position = require('components.position')

return function()
    local e = entity()
    local image = love.graphics.newImage('sprites/bg.jpeg')

    e.addComponent(appearance(image))
    e.addComponent(position(0, 0))

    return e
end