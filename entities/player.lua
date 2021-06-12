local json = require('utilities.json')
local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')
local controlsComponent = require('components.controls')

return function()
    local e = entity()
    local image = love.graphics.newImage('sprites/player.png')
    local sheet = json('sprites/player.json')
    e.addComponent(appearanceComponent(image, sheet))
    e.addComponent(positionComponent(50, 50))
    e.addComponent(controlsComponent())

    return e
end