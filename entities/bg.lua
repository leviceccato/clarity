local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')

return function()
    local e = entity()
    image = love.graphics.newImage('sprites/bg.jpeg')
    e.addComponent(appearanceComponent(image))
    e.addComponent(positionComponent(0, 0))

    return e
end