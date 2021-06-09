local state = require('library.state')
local mainWorld = require('worlds.main')

local mainState

love.load = function(arg)
    love.graphics.setDefaultFilter('nearest', 'nearest')

    mainState = state()
    mainState.addWorld(mainWorld())
    mainState.activateWorld('main')
    mainState.load(arg)
end

love.update = function(dt)
    mainState.update(dt)
end

love.draw = function()
    mainState.draw()
end

love.keyreleased = function(...)
    mainState.keyreleased(...)
end

love.resize = function(...)
    mainState.resize(...)
end