local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')
local shuffleComponent = require('components.shuffle')

return function(posX, posY)
    local e = entity()
    image = love.graphics.newImage('sprites/jon.jpeg')
    e.addComponent(appearanceComponent(image))
    e.addComponent(shuffleComponent())
    e.addComponent(positionComponent(posX or 200, posY or 200))

    return e
end