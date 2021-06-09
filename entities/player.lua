local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')
local controlsComponent = require('components.controls')

return function()
    local e = entity()
    image = love.graphics.newImage('sprites/shine.png')
    e.addComponent(appearanceComponent(image))
    e.addComponent(positionComponent(50, 50))
    e.addComponent(controlsComponent())

    return e
end