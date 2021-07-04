local json = require('utilities.json')
local entity = require('library.entity')
local appearance = require('components.appearance')
local animation = require('components.animation')
local position = require('components.position')
local controls = require('components.controls')
local camera = require('components.camera')
local hover = require('components.hover')

return function()
    local e = entity()
    local image = love.graphics.newImage('assets/sprites/player.png')
    local sheet = json('assets/sprites/player.json')

    e.addComponent(appearance(image))
    e.addComponent(animation(sheet))
    e.addComponent(position(50, 50))
    e.addComponent(controls())
    e.addComponent(camera())
    e.addComponent(hover())

    return e
end