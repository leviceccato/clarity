local entity = require('library.entity')

local appearance = require('components.appearance')
local position = require('components.position')
local camera = require('components.camera')

return function()
    local e = entity()
    local image = love.graphics.newImage('assets/sprites/bg.jpeg')

    e.addComponent(appearance(image))
    e.addComponent(position(0, 0))
    e.addComponent(camera())

    return e
end