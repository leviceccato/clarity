local state = require('library.state')
local mainWorld = require('worlds.main')
local titleWorld = require('worlds.title')

local mainState

love.load = function(arg)
    local font = love.graphics.newFont('assets/dogica.ttf', 16, 'mono')
    love.graphics.setFont(font)
    love.graphics.setDefaultFilter('nearest', 'nearest')

    mainState = state(arg)
    mainState.addWorld(mainWorld(mainState))
    mainState.addWorld(titleWorld(mainState))
    mainState.activateWorld('main')
end

love.update = function(dt)
    mainState.update(dt)
end

love.draw = function()
    mainState.draw()
end

love.keypressed = function(key)
    mainState.keypressed(key)
end

love.keyreleased = function(key)
    mainState.keyreleased(key)
end

love.resize = function(width, height)
    mainState.resize(width, height)
end