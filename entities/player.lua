local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')

return function()
    local e = entity()
    image = love.graphics.newImage('sprites/shine.png')
    e.addComponent(appearanceComponent(image))
    e.addComponent(positionComponent(50, 50))

    return e
end