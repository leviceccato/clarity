Log = require('utilities.log')
local i18n = require('vendor.i18n.i18n')
local viewport = require('utilities.viewport')
local json = require('utilities.json')
local state = require('library.state')

local mainWorld = require('worlds.main')
local titleWorld = require('worlds.title')

local mainState

love.load = function(arg)
    i18n.load(json('translations/en.json'))
    love.graphics.setFont(love.graphics.newFont('assets/lana-pixel.ttf', 11, 'mono'))
    love.graphics.setDefaultFilter('nearest', 'nearest')

    mainState = state(arg)
    mainState.loadWorld(mainWorld(mainState))
    mainState.loadWorld(titleWorld(mainState))
    mainState.activateWorlds({'title'})
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