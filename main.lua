Log = require('utilities.log')
local viewport = require('utilities.viewport')
local state = require('library.state')

local mainWorld = require('worlds.main')
local titleWorld = require('worlds.title')

local mainState

love.load = function(arg)
    local font = love.graphics.newFont('assets/dogica.ttf', 8, 'mono')
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

love.exit = function()
    mainState.exit()
end

love.mousepressed = function(_, _, button, isTouch, pressCount)
    local x, y = viewport.getMousePosition()
    mainState.mousepressed(x, y, button, isTouch, pressCount)
end

love.mousereleased = function(_, _, button)
    mainState.mousereleased(button)
end

love.keypressed = function(key, _, isRepeat)
    mainState.keypressed(key, isRepeat)
end

love.keyreleased = function(key)
    mainState.keyreleased(key)
end

love.resize = function(width, height)
    mainState.resize(width, height)
end